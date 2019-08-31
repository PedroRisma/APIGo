package domains

import (
	"../utils"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type User struct {
	ID               int    `json:"id"`
	Nickname         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	CountryID        string `json:"country_id"`
	Address          struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"address"`
	UserType         string      `json:"user_type"`
	Tags             []string    `json:"tags"`
	Logo             interface{} `json:"logo"`
	Points           int         `json:"points"`
	SiteID           string      `json:"site_id"`
	Permalink        string      `json:"permalink"`
	SellerReputation struct {
		LevelID           interface{} `json:"level_id"`
		PowerSellerStatus interface{} `json:"power_seller_status"`
		Transactions      struct {
			Canceled  int    `json:"canceled"`
			Completed int    `json:"completed"`
			Period    string `json:"period"`
			Ratings   struct {
				Negative int `json:"negative"`
				Neutral  int `json:"neutral"`
				Positive int `json:"positive"`
			} `json:"ratings"`
			Total int `json:"total"`
		} `json:"transactions"`
	} `json:"seller_reputation"`
	BuyerReputation struct {
		Tags []interface{} `json:"tags"`
	} `json:"buyer_reputation"`
	Status struct {
		SiteStatus string `json:"site_status"`
	} `json:"status"`
}


func (user *User) Get() *utils.ApiError{
	if user.ID <= 0 {
		return &utils.ApiError{
			Message: "User ID invalid",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%d", utils.UrlUser, user.ID)

	response, err := http.Get(url)

	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &user); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}


	return nil

}

func (user *User) GetW(ae *utils.ApiError) {
	if user.ID <= 0 {
		ae = &utils.ApiError{
			Message: "User ID invalid",
			Status: http.StatusBadRequest,
		}
		return
	}

	url := fmt.Sprintf("%s%d", utils.UrlUser, user.ID)

	response, err := http.Get(url)

	if err != nil {
		ae = &utils.ApiError{
			Message: "User ID invalid",
			Status: http.StatusBadRequest,
		}
		return
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		ae = &utils.ApiError{
			Message: "User ID invalid",
			Status: http.StatusBadRequest,
		}
		return
	}

	if err := json.Unmarshal(data, &user); err != nil {
		ae = &utils.ApiError{
			Message: "User ID invalid",
			Status: http.StatusBadRequest,
		}
		return
	}

	ae = nil
	return

}

func (user *User) GetC(cerror chan *utils.ApiError){
	if user.ID <= 0 {
		err := &utils.ApiError{
			Message: "User ID invalid",
			Status: http.StatusBadRequest,
		}
		cerror <- err
		return
	}

	url := fmt.Sprintf("%s%d", utils.UrlUser, user.ID)

	response, err := http.Get(url)

	if err != nil {
		err := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		cerror <- err
		return
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		err := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		cerror <- err
		return
	}

	if err := json.Unmarshal(data, &user); err != nil {
		err := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		cerror <- err
		return
	}

	cerror <- nil
	return

}

package domains

import (
	"../utils"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Country struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Locale             string `json:"locale"`
	CurrencyID         string `json:"currency_id"`
	DecimalSeparator   string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone           string `json:"time_zone"`
	GeoInformation     struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"states"`
}

func (country *Country) Get() *utils.ApiError{
	if country.ID == "" {
		return &utils.ApiError{
			Message: "User ID invalid",
			Status: http.StatusBadRequest,
		}

	}

	url := fmt.Sprintf("%s%s", utils.UrlCountry, country.ID)

	response, err := http.Get(url)

	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}

	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}

	}

	if err := json.Unmarshal(data, &country); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}

	}

	return nil

}

func (country *Country) GetW(ae *utils.ApiError){
	if country.ID == "" {
		ae = &utils.ApiError{
			Message: "Country ID invalid",
			Status: http.StatusBadRequest,
		}
		return
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountry, country.ID)

	response, err := http.Get(url)

	if err != nil {
		ae = &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		return
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		ae = &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		return
	}

	if err := json.Unmarshal(data, &country); err != nil {
		ae = &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		return
	}

	ae = nil
	return

}

func (country *Country) GetC(cerror chan *utils.ApiError){
	if country.ID == "" {
		err := &utils.ApiError{
			Message: "Country ID invalid",
			Status: http.StatusBadRequest,
		}
		cerror <- err
		return
	}

	url := fmt.Sprintf("%s%s", utils.UrlCountry, country.ID)

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

	if err := json.Unmarshal(data, &country); err != nil {
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
package domains

import (
	"../utils"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Site struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	CountryID          string   `json:"country_id"`
	SaleFeesMode       string   `json:"sale_fees_mode"`
	MercadopagoVersion int      `json:"mercadopago_version"`
	DefaultCurrencyID  string   `json:"default_currency_id"`
	ImmediatePayment   string   `json:"immediate_payment"`
	PaymentMethodIds   []string `json:"payment_method_ids"`
	Settings           struct {
		IdentificationTypes      []string `json:"identification_types"`
		TaxpayerTypes            []string `json:"taxpayer_types"`
		IdentificationTypesRules []struct {
			IdentificationType string        `json:"identification_type"`
			Rules              []interface{} `json:"rules"`
		} `json:"identification_types_rules"`
	} `json:"settings"`
	Currencies []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Categories []struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"categories"`
}


func (site *Site) Get() *utils.ApiError{
	if site.ID == "" {
		return &utils.ApiError{
			Message: "Site ID is empty",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

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

	if err := json.Unmarshal(data, &site); err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}


	return nil

}

func (site *Site) GetW(ae *utils.ApiError) {
	if site.ID == "" {
		ae = &utils.ApiError{
			Message: "User ID invalid",
			Status: http.StatusBadRequest,
		}
		return
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

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

	if err := json.Unmarshal(data, &site); err != nil {
		ae = &utils.ApiError{
			Message: "User ID invalid",
			Status: http.StatusBadRequest,
		}
		return
	}

	ae = nil
	return

}

func (site *Site) GetC(cerror chan *utils.ApiError) {
	if site.ID == "" {
		err := &utils.ApiError{
			Message: "Site ID invalid",
			Status: http.StatusBadRequest,
		}
		cerror <- err
		return
	}

	url := fmt.Sprintf("%s%s", utils.UrlSite, site.ID)

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

	if err := json.Unmarshal(data, &site); err != nil {
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


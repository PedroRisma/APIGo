package services


import (
	"../domains"
	"sync"
	"../utils"
)


func GetResult (userId int) (*domains.Result){

	result := &domains.Result{

	}

	user := &domains.User{
		ID:userId,
	}

	if err := user.Get(); err!= nil {
		result.ApiError = err
		return result
	}
	result.User = user

	country := &domains.Country{
		ID:user.CountryID,
	}

	if err := country.Get(); err!= nil {
		result.ApiError = err
		return result
	}

	result.Country = country

	site := &domains.Site{
		ID:user.SiteID,
	}

	if err := site.Get(); err!= nil {
		result.ApiError = err
		return result
	}
	result.Site = site

	return result

}

func GetResultW (userId int, wg sync.WaitGroup) (*domains.Result){

	resultW := &domains.Result{

	}

	var error *utils.ApiError

	user := &domains.User{
		ID:userId,
	}

	if err := user.Get(); err!= nil {
		resultW.ApiError = err
		return resultW
	}
	resultW.User = user

	country := &domains.Country{
		ID:user.CountryID,
	}

	wg.Add(1)
	go country.GetW(error)
	wg.Done()

	if error != nil {
		resultW.ApiError = error
		return resultW
	}

	resultW.Country = country

	site := &domains.Site{
		ID:user.SiteID,
	}

	wg.Add(1)
	go site.GetW(error)
	wg.Done()

	if error != nil {
		resultW.ApiError = error
		return resultW
	}
	resultW.Site = site

	return resultW

}

func GetResultC (userId int) (*domains.Result){

	resultC := &domains.Result{

	}

	cerror := make(chan *utils.ApiError, 2)

	user := &domains.User{
		ID:userId,
	}

	if err := user.Get(); err!= nil {
		resultC.ApiError = err
		return resultC
	}

	resultC.User = user

	country := &domains.Country{
		ID:user.CountryID,
	}
	site := &domains.Site{
		ID:user.SiteID,
	}

	go country.GetC(cerror)
	go site.GetC(cerror)

	for i := 0; i < 2 ;i++  {
		resultC.ApiError = <- cerror
		if resultC.ApiError != nil {
			return resultC
		}
	}

	resultC.Country = country
	resultC.Site = site
	cerror <- nil
	resultC.ApiError = <- cerror

	return resultC

}


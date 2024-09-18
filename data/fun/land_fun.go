package data

import (
	"fitroom-autotests/config"
	"fitroom-autotests/models"
	"fitroom-autotests/utils"
)

var (
	LandFunRequestList = []models.Requester{
		&GetLanding,
		&GetCityList,
		&GetTagsListPublic,
		&GetPackagesList,
		&GetClubsList,
		&GetTrainersList,
		&Get2GisMap,
		&Get2GisCatalog,
		&GetAvailable,
	}

	city = models.City{
		City: 5,
	}

	cities = models.Cities{
		Cities:     []int{2}, // Пушкин
		OnlyGroups: 1,
	}

	GetLanding = models.Request{
		URL:    config.LAND_FUN,
		Method: "GET",
		Token:  &config.TokenUserFun,
	}

	GetCityList = models.Request{
		URL:    config.BACK_FUN + "/public/city/get-list",
		Method: "GET",
		Token:  &config.TokenUserFun,
	}

	GetPackagesList = models.Request{
		URL:     config.BACK_FUN + "/public/package/get-packages",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(city),
		Token:   &config.TokenUserFun,
	}

	GetTagsListPublic = models.Request{
		URL:    config.BACK_FUN + "/public/tag/get-list",
		Method: "POST",
		Token:  &config.TokenUserFun,
	}

	GetClubsList = models.Request{
		URL:      config.BACK_FUN + "/public/club/get-list",
		Method:   "POST",
		Token:    &config.TokenUserFun,
		RespData: &clubList,
	}

	GetTrainersList = models.Request{
		URL:     config.BACK_FUN + "/public/user/get-trainers",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(city),
		Token:   &config.TokenUserFun,
	}

	Get2GisMap = models.Request{
		URL:    "https://maps.api.2gis.ru/2.0/css/?version=v3.7.3",
		Method: "GET",
		Token:  &config.TokenUserFun,
	}

	Get2GisCatalog = models.Request{
		URL:    "https://catalog.api.2gis.ru/2.0/region/list?format=json&key=rubnkm7490&fields=items.bounds,items.zoom_level,items.time_zone,items.code,items.flags,items.country_code,items.domain,items.default_pos",
		Method: "GET",
		Token:  &config.TokenUserFun,
	}

	availableLandFun = models.Available{
		Workplace:           120,
		ServiceProductTypes: []int{1},
		StartDate:           config.P_Date,
		EndDate:             config.P_Date,
		Trainers:            []int{2058, 1625},
	}

	GetAvailable = models.Request{
		URL:     config.BACK_FUN + "/public/appointment/available",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(availableLandFun),
		Token:   &config.TokenUserFun,
	}

	GerSerivceList = models.Request{
		URL:     config.BACK_FUN + "/public/service/get-list",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(cities),
		Token:   &config.TokenUserFun,
	}
)

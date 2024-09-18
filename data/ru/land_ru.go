package data_ru

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
		URL:    config.LAND_RU,
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	GetCityList = models.Request{
		URL:    config.BACK_RU + "/public/city/get-list",
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	GetPackagesList = models.Request{
		URL:     config.BACK_RU + "/public/package/get-packages",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(city),
		Token:   &config.TokenUserRu,
	}

	GetTagsListPublic = models.Request{
		URL:    config.BACK_RU + "/public/tag/get-list",
		Method: "POST",
		Token:  &config.TokenUserRu,
	}

	GetClubsList = models.Request{
		URL:      config.BACK_RU + "/public/club/get-list",
		Method:   "POST",
		Token:    &config.TokenUserRu,
		RespData: &clubList,
	}

	GetTrainersList = models.Request{
		URL:     config.BACK_RU + "/public/user/get-trainers",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(city),
		Token:   &config.TokenUserRu,
	}

	Get2GisMap = models.Request{
		URL:    "https://maps.api.2gis.ru/2.0/css/?version=v3.7.3",
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	Get2GisCatalog = models.Request{
		URL:    "https://catalog.api.2gis.ru/2.0/region/list?format=json&key=rubnkm7490&fields=items.bounds,items.zoom_level,items.time_zone,items.code,items.flags,items.country_code,items.domain,items.default_pos",
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	availableLandRu = models.Available{
		Workplace:           27,
		ServiceProductTypes: []int{1},
		StartDate:           config.P_Date,
		EndDate:             config.P_Date,
		Trainers:            []int{16},
	}

	GetAvailable = models.Request{
		URL:     config.BACK_RU + "/public/appointment/available",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(availableLandRu),
		Token:   &config.TokenUserRu,
	}

	GerSerivceList = models.Request{
		URL:     config.BACK_RU + "/public/service/get-list",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(cities),
		Token:   &config.TokenUserRu,
	}
)

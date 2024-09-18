package data_ru

import (
	"fitroom-autotests/config"
	"fitroom-autotests/models"
	"fitroom-autotests/utils"
	"strconv"
)

var (
	AppFunRequestList = []models.Requester{
		&GetProfile,
		&GetAppointments,
		&GetAppointmentsGroup,
		&GetTerm,
		&GetSettingsList,
		&GetClubByID,
		&GetClubList,
		&GetPackageList,
		&GetAvatar,
		&GetHistory,
		&GetMarketHTML,
		&Rent,
		// &Train,
		// &GroupTrain,
	}

	// Requests
	GetAppointments = models.Request{
		URL:     config.BACK_RU + "/api/lk/v2/appointment/available",
		Method:  "POST",
		Token:   &config.TokenUserRu,
		ReqBody: utils.PrepareReqBody(available_1),
	}

	available_1 = models.Available{
		Workplace:           27, // Матисов
		ServiceProductTypes: []int{1},
		StartDate:           config.C_Date,
		EndDate:             config.F_Date,
	}

	GetAppointmentsGroup = models.Request{
		URL:     config.BACK_RU + "/api/lk/v2/appointment/available",
		Method:  "POST",
		Token:   &config.TokenUserRu,
		ReqBody: utils.PrepareReqBody(available_2),
	}

	available_2 = models.Available{
		Workplace:           27,
		ServiceProductTypes: []int{1, 2},
		StartDate:           config.C_Date,
		EndDate:             config.F_Date,
	}

	GetAvatar = models.Request{
		URL:    config.BACK_RU + "/public/image/get-by-id/2076",
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	GetClubByID = models.Request{
		URL:    config.BACK_RU + "/public/club/get-by-id/21", // Матисов
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	GetClubList = models.Request{
		URL:     config.BACK_RU + "/public/club/get-list",
		Method:  "POST",
		Token:   &config.TokenUserRu,
		ReqBody: utils.PrepareReqBody(cities),
	}

	GetHistory = models.Request{
		URL:     config.BACK_RU + "/api/lk/appointment/history",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(pages),
		Token:   &config.TokenUserRu,
	}

	pageLimit = models.Paging{
		Limit: 20,
		Page:  1,
	}

	pages = models.Page{
		Paging: pageLimit,
	}

	GetMarketHTML = models.Request{
		URL:    "https://app.fitroom.fun/market",
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	GetPackageList = models.Request{
		URL:     config.BACK_RU + "/api/lk/package/get-list",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(club_1),
		Token:   &config.TokenUserRu,
	}

	club_1 = models.Club{
		Club: 21, // Матисов
	}

	GetProfile = models.Request{
		URL:    config.BACK_RU + "/api/lk/user/profile",
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	GetList = models.Request{
		URL:    config.BACK_RU + "/api/lk/setting/get-list",
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	GetSettingsList = models.Request{
		URL:    config.BACK_RU + "/api/lk/setting/get-list",
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	GetTerm = models.Request{
		URL:    config.BACK_RU + "/public/term/get",
		Method: "GET",
		Token:  &config.TokenUserRu,
	}

	Rent = models.Request{
		URL:      config.BACK_RU + "/api/lk/appointment/book",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(bookRent),
		RespData: &rentResp,
		Token:    &config.TokenUserRu,
		Actions:  []func(){Rent_F1},
		Next:     &CancelRentLk,
	}

	bookRent = models.Book{
		User:               6775,
		Club:               21,
		Workplace:          27,
		ServiceProductType: 1,
		StartDate:          config.StartDate1,
		EndDate:            config.EndDate1,
		UserPackage:        33965,
	}

	rentResp models.Resp

	Rent_F1 = func() {
		CancelRentLk.URL = config.BACK_RU + "/api/lk/appointment/cancel/" + strconv.Itoa(rentResp.ID)
	}

	Train = models.Request{
		URL:      config.BACK_RU + "/api/lk/appointment/book",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(bookTrain),
		RespData: &rentResp,
		Token:    &config.TokenUserRu,
		Actions:  []func(){Rent_F1},
		Next:     &CancelRentLk,
	}

	bookTrain = models.Book{
		User:               2058,
		Club:               13,
		Workplace:          13,
		ServiceProductType: 2,
		StartDate:          config.StartDate2,
		EndDate:            config.EndDate2,
		UserPackage:        8012,
	}

	GroupTrain = models.Request{
		URL:      config.BACK_RU + "/api/lk/appointment/book",
		Method:   "POST",
		Token:    &config.TokenUserRu,
		ReqBody:  utils.PrepareReqBody(groupTrain),
		RespData: &rentResp,
		Actions:  []func(){Rent_F1},
		Next:     &CancelRentLk,
	}

	groupTrain = models.GroupTrainBook{
		User:               1625,
		Club:               166,
		Workplace:          120,
		ServiceProductType: 8,
		StartDate:          config.StartDate3,
		EndDate:            config.EndDate3,
		UserPackage:        8057,
		GroupTraining:      16,
	}

	CancelRentLk = models.Request{
		Method: "POST",
		Token:  &config.TokenUserRu,
	}
)

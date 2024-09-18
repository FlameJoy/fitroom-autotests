package data

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
		&Train,
		&GroupTrain,
	}

	// Requests
	GetAppointments = models.Request{
		URL:     config.BACK_FUN + "/api/lk/v2/appointment/available",
		Method:  "POST",
		Token:   &config.TokenUserFun,
		ReqBody: utils.PrepareReqBody(available_1),
	}

	available_1 = models.Available{
		Workplace:           13, // Пушкин
		ServiceProductTypes: []int{1},
		StartDate:           config.C_Date,
		EndDate:             config.F_Date,
	}

	GetAppointmentsGroup = models.Request{
		URL:     config.BACK_FUN + "/api/lk/v2/appointment/available",
		Method:  "POST",
		Token:   &config.TokenUserFun,
		ReqBody: utils.PrepareReqBody(available_2),
	}

	available_2 = models.Available{
		Workplace:           13,
		ServiceProductTypes: []int{10, 11, 8},
		StartDate:           config.C_Date,
		EndDate:             config.F_Date,
	}

	GetAvatar = models.Request{
		URL:    config.BACK_FUN + "/public/image/get-by-id/2076",
		Method: "GET",
		Token:  &config.TokenUserFun,
	}

	GetClubByID = models.Request{
		URL:    config.BACK_FUN + "/public/club/get-by-id/13", // Pushkin
		Method: "GET",
		Token:  &config.TokenUserFun,
	}

	GetClubList = models.Request{
		URL:     config.BACK_FUN + "/public/club/get-list",
		Method:  "POST",
		Token:   &config.TokenUserFun,
		ReqBody: utils.PrepareReqBody(cities),
	}

	GetHistory = models.Request{
		URL:     config.BACK_FUN + "/api/lk/appointment/history",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(pages),
		Token:   &config.TokenUserFun,
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
		Token:  &config.TokenUserFun,
	}

	GetPackageList = models.Request{
		URL:     config.BACK_FUN + "/api/lk/package/get-list",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(club_1),
		Token:   &config.TokenUserFun,
	}

	club_1 = models.Club{
		Club: 13, // Пушкин
	}

	GetProfile = models.Request{
		URL:    config.BACK_FUN + "/api/lk/user/profile",
		Method: "GET",
		Token:  &config.TokenUserFun,
	}

	GetList = models.Request{
		URL:    config.BACK_FUN + "/api/lk/setting/get-list",
		Method: "GET",
		Token:  &config.TokenUserFun,
	}

	GetSettingsList = models.Request{
		URL:    config.BACK_FUN + "/api/lk/setting/get-list",
		Method: "GET",
		Token:  &config.TokenUserFun,
	}

	GetTerm = models.Request{
		URL:    config.BACK_FUN + "/public/term/get",
		Method: "GET",
		Token:  &config.TokenUserFun,
	}

	Rent = models.Request{
		URL:      config.BACK_FUN + "/api/lk/appointment/book",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(bookRent),
		RespData: &rentResp,
		Token:    &config.TokenUserFun,
		Actions:  []func(){Rent_F1},
		Next:     &CancelRentLk,
	}

	bookRent = models.Book{
		User:               2953,
		Club:               13,
		Workplace:          13,
		ServiceProductType: 1,
		StartDate:          config.StartDate1,
		EndDate:            config.EndDate1,
		UserPackage:        7959,
	}

	rentResp models.Resp

	Rent_F1 = func() {
		CancelRentLk.URL = config.BACK_FUN + "/api/lk/appointment/cancel/" + strconv.Itoa(rentResp.ID)
	}

	Train = models.Request{
		URL:      config.BACK_FUN + "/api/lk/appointment/book",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(bookTrain),
		RespData: &rentResp,
		Token:    &config.TokenUserFun,
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
		URL:      config.BACK_FUN + "/api/lk/appointment/book",
		Method:   "POST",
		Token:    &config.TokenUserFun,
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
		Token:  &config.TokenUserFun,
	}
)

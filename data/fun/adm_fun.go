package data

import (
	"fitroom-autotests/config"
	"fitroom-autotests/models"
	"fitroom-autotests/utils"
	"log"
	"strconv"
)

var (
	AdmFunRequestList = []models.Requester{
		&GetUserByID,
		&GetUsersList,
		&GetMenu,
		&GetPermission,
		&GetCityList,
		&GetDistrictList,
		&GetComplexList,
		&GetClubsList,
		&GetClubByID,
		&GetMetroList,
		&GetWorkplaceList,
		&GetBranchList,
		&GetAppointmentsList,
		&GetProductType,
		&GetPackagesList,
		&GetPromotionList,
		&GetPromotionType,
		&GetTagsList,
		&GetconfigList,
		&GetCountryList,
		&GetTerms,
		&GetNotifications,
		&GetOrdersList,
		&GetOrders,
		&GetReview,
		&GetGroupAppintments,
		&GetNotification,
		// Book ---------------------------------------------------------
		&AdminBook,
		// Create / Edit / Delete elements ------------------------------
		&CreateBranch,
		&CreateClub,
		&CreatePackage,
		&CreatePromotion,
		&CreateTag,
		&CreateWorkplace,
		&CreateTerm,
		&ASD,
		&CreateUser,
		&CreateGroupTrain,
		// // Bonus --------------------------------------------------------
		&BonusRefferal,
		&BonusAdd,
		&BonusRemove,
		&BonusDayBeforeDel,
		&BonusMaxProcent,
		&BonusMinProcent,
		&BonusDelProcent,
		// // Reports ------------------------------------------------------
		&ReportFinancial,
		&ReportClub,
		&ReportTrainers,
		&ReportBonus,
		&ReportConversion,
	}

	resp models.Resp

	// Requests

	GetPermission = models.Request{
		URL:    config.ADM_FUN + "/permission/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetProductType = models.Request{
		URL:    config.ADM_FUN + "/service-product-type/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetPromotionList = models.Request{
		URL:    config.ADM_FUN + "/promotion/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetPromotionType = models.Request{
		URL:    config.ADM_FUN + "/promotion-type/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetMenu = models.Request{
		URL:    config.ADM_FUN + "/menu/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetDistrictList = models.Request{
		URL:    config.ADM_FUN + "/district/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetComplexList = models.Request{
		URL:    config.ADM_FUN + "/complex/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetMetroList = models.Request{
		URL:    config.ADM_FUN + "/metro/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetBranchList = models.Request{
		URL:    config.ADM_FUN + "/branch/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetAppointmentsList = models.Request{
		URL:     config.ADM_FUN + "/appointment/get-list",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(appointmentList),
		Token:   &config.TokenAdmFun,
	}

	appointmentList = models.Appointments{
		Status:   "booked",
		FromDate: config.C_Date,
		ToDate:   config.F_Date,
	}

	GetCountryList = models.Request{
		URL:    config.ADM_FUN + "/country/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetNotifications = models.Request{
		URL:    config.ADM_FUN + "/notification/get-list",
		Method: "POST",

		Token: &config.TokenAdmFun,
	}

	GetOrdersList = models.Request{
		URL:     config.ADM_FUN + "/order/get-list",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(pageLimit),
		Token:   &config.TokenAdmFun,
	}

	GetOrders = models.Request{
		URL:     config.ADM_FUN + "/order/get-list",
		Method:  "POST",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(&pageLimit),
	}

	GetGroupAppintments = models.Request{
		URL:     config.ADM_FUN + "/appointment/available",
		Method:  "POST",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(groupAppointmentList),
	}

	GetNotification = models.Request{
		URL:    config.ADM_FUN + "/notification/get-list",
		Method: "POST",
		Token:  &config.TokenAdmFun,
	}

	groupAppointmentList = models.GroupAppointments{
		StartDate:           config.P_Date,
		EndDate:             config.C_Date,
		Workplace:           120,
		ServiceProductTypes: []int{10, 8, 11},
	}

	GetReview = models.Request{
		URL:     config.ADM_FUN + "/review/get-list",
		Method:  "POST",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(pages),
	}

	GetconfigList = models.Request{
		URL:    config.ADM_FUN + "/setting/get-list",
		Method: "POST",
		Token:  &config.TokenAdmFun,
	}

	GetTagsList = models.Request{
		URL:    config.ADM_FUN + "/tag/get-list",
		Method: "POST",
		Token:  &config.TokenAdmFun,
	}

	GetTerms = models.Request{
		URL:    config.ADM_FUN + "/term/get",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetUserByID = models.Request{
		URL:    config.ADM_FUN + "/user/get-by-id/2058",
		Method: "GET",
		Token:  &config.TokenAdmFun,
	}

	GetUsersList = models.Request{
		URL:     config.ADM_FUN + "/user/get-list",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(pageLimit),
		Token:   &config.TokenAdmFun,
	}

	GetWorkplaceList = models.Request{
		URL:    config.ADM_FUN + "/workplace/get-list",
		Method: "POST",
		Token:  &config.TokenAdmFun,
	}

	ReportBonus = models.Request{
		URL:     config.ADM_FUN + "/report/bonus-flow",
		Method:  "POST",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(newRepBonus),
	}

	newRepBonus = models.RepBonus{
		FromDate: config.C_Date,
		ToDate:   config.F_Date,
	}

	ReportClub = models.Request{
		URL:     config.ADM_FUN + "/report/club",
		Method:  "POST",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(newRepClub),
	}

	newRepClub = models.RepClub{
		Clubs:         GetClubIDList(),
		FromDate:      config.P_Date,
		ToDate:        config.C_Date,
		SalaryPercent: 50,
		Paging: models.Paging{
			Page:  1,
			Limit: 300,
		},
	}

	clubList = []struct {
		ID int `json:"id"`
	}{}

	clubIDList = []int{}

	GetClubIDList = func() []int {
		// models.Authorization()
		GetClubsList.POST()
		for _, club := range clubList {
			clubIDList = append(clubIDList, club.ID)
		}
		log.Println(clubIDList)
		return clubIDList
	}

	ReportConversion = models.Request{
		URL:     config.ADM_FUN + "/report/conversion",
		Method:  "POST",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(newRepConversion),
	}

	newRepConversion = models.RepConversion{
		FromDate: "2023-01-01",
		ToDate:   config.C_Date,
	}

	ReportFinancial = models.Request{
		URL:     config.ADM_FUN + "/report/revenue",
		Method:  "POST",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(newRepFinancial),
	}

	newRepFinancial = models.RepFinancial{
		Clubs:    GetClubIDList(),
		FromDate: config.P_Date,
		ToDate:   config.C_Date,
	}

	ReportTrainers = models.Request{
		URL:     config.ADM_FUN + "/report/club-trainers",
		Method:  "POST",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(newRepTrainers),
	}

	newRepTrainers = models.RepTrainers{
		FromDate:      config.P_Date,
		ToDate:        config.C_Date,
		SalaryPercent: 50,
	}

	AdminBook = models.Request{
		URL:      config.ADM_FUN + "/appointment/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(adminBook),
		RespData: &adminBookResp,
		Token:    &config.TokenAdmFun,
		Actions:  []func(){AdminBook_F1},
		Next:     &CancelRentAdm,
	}

	adminBook = models.Book{
		Clients:            []uint{2953},
		Club:               13,
		Workplace:          13,
		ServiceProductType: 1,
		StartDate:          config.StartDate1,
		EndDate:            config.EndDate1,
		Status:             "booked",
		User:               2953,
		UserPackage:        7959,
	}

	adminBookResp models.Resp

	AdminBook_F1 = func() {
		CancelRentAdm.URL = config.ADM_FUN + "/appointment/delete/" + strconv.Itoa(adminBookResp.ID)
	}

	CancelRentAdm = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmFun,
	}

	CreateBranch = models.Request{
		URL:      config.ADM_FUN + "/branch/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newBranch),
		Token:    &config.TokenAdmFun,
		RespData: &branchResp,
		Actions:  []func(){Branch_F1},
		Next:     &EditBranch,
	}

	EditBranch = models.Request{
		Method:  "PUT",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(editBranch),
		Next:    &DeleteBranch,
	}

	DeleteBranch = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmFun,
	}

	newBranch = models.BranchReq{
		Active:      1,
		Title:       "AUTOTEST",
		Description: "AUTOTEST",
	}

	editBranch = models.BranchReq{
		Active:      1,
		Title:       "EDITTED AUTOTEST",
		Description: "EDITTED AUTOTEST",
	}

	branchResp models.BranchResp

	Branch_F1 = func() {
		EditBranch.URL = config.ADM_FUN + "/branch/edit/" + strconv.Itoa(branchResp.ID)
		DeleteBranch.URL = config.ADM_FUN + "/branch/delete/" + strconv.Itoa(branchResp.ID)
	}

	CreateClub = models.Request{
		URL:      config.ADM_FUN + "/club/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newClub),
		Token:    &config.TokenAdmFun,
		RespData: &clubResp,
		Actions:  []func(){Club_F1},
		Next:     &EditClub,
	}

	EditClub = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editClub),
		Token:   &config.TokenAdmFun,
		Next:    &DeleteClub,
	}

	DeleteClub = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmFun,
	}

	newClub = models.ClubReq{
		Title:     "AUTOTEST",
		City:      1,
		District:  1,
		Complex:   1,
		Street:    "AUTOTEST",
		House:     "1",
		Building:  "1",
		Latitude:  "50.00000",
		Longitude: "50.00000",
		OpenDate:  config.F_Date,
	}

	editClub = models.ClubReq{
		Title:     "EDITTED AUTOTEST",
		City:      2,
		District:  2,
		Complex:   2,
		Street:    "EDITTED AUTOTEST",
		House:     "2",
		Building:  "2",
		Latitude:  "50.20000",
		Longitude: "50.20000",
		OpenDate:  config.F_Date,
	}

	clubResp models.ClubResp

	Club_F1 = func() {
		EditClub.URL = config.ADM_FUN + "/club/edit/" + strconv.Itoa(clubResp.ID)
		DeleteClub.URL = config.ADM_FUN + "/club/delete/" + strconv.Itoa(clubResp.ID)
	}

	CreatePackage = models.Request{
		URL:      config.ADM_FUN + "/package/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newPackage),
		Token:    &config.TokenAdmFun,
		RespData: &packageResp,
		Actions:  []func(){Package_F1},
		Next:     &EditPackage,
	}

	EditPackage = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editPackage),
		Token:   &config.TokenAdmFun,
		Next:    &DeletePackage,
	}

	DeletePackage = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmFun,
	}

	newPackage = models.PackageReq{
		Title:               "AUTOTEST",
		ServiceProductType:  1,
		TrainingAmount:      5,
		Cost:                1000,
		Cashback:            5,
		WritingOff:          10,
		DaysAmount:          30,
		CashbackStartDate:   "2023-03-01",
		CashbackEndDate:     "2023-03-03",
		WritingOffStartDate: "2023-03-01",
		WritingOffEndDate:   "2023-03-03",
	}

	editPackage = models.PackageReq{
		Title:               "EDITTED AUTOTEST",
		ServiceProductType:  2,
		TrainingAmount:      10,
		Cost:                2000,
		Cashback:            10,
		WritingOff:          20,
		DaysAmount:          40,
		CashbackStartDate:   "2023-03-01",
		CashbackEndDate:     "2023-03-03",
		WritingOffStartDate: "2023-03-01",
		WritingOffEndDate:   "2023-03-03",
	}

	packageResp models.PackageResp

	Package_F1 = func() {
		EditPackage.URL = config.ADM_FUN + "/package/edit/" + strconv.Itoa(packageResp.ID)
		DeletePackage.URL = config.ADM_FUN + "/package/delete/" + strconv.Itoa(packageResp.ID)
	}

	CreatePromotion = models.Request{
		URL:      config.ADM_FUN + "/promotion/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newPromotion),
		Token:    &config.TokenAdmFun,
		RespData: &promotionResp,
		Actions:  []func(){Promotion_F1},
		Next:     &EditPromotion,
	}

	EditPromotion = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editPromotion),
		Token:   &config.TokenAdmFun,
		Next:    &DeletePromotion,
	}

	DeletePromotion = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmFun,
	}

	newPromotion = models.PromotionReq{
		Clubs:         []int{},
		Count:         1,
		Discount:      10,
		EndDate:       config.EndDate,
		Packages:      []int{},
		PersonalCount: 1,
		PromotionType: 1,
		Segments:      []int{1},
		StartDate:     config.StartDate,
		Title:         "AUTOTEST",
	}

	editPromotion = models.PromotionReq{
		Clubs:         []int{1, 2},
		Count:         20,
		Discount:      20,
		EndDate:       config.EndDate,
		Packages:      []int{},
		PersonalCount: 2,
		PromotionType: 2,
		Segments:      []int{1, 2},
		StartDate:     config.StartDate,
		Title:         "EDITTED AUTOTEST",
	}

	promotionResp models.PromotionResp

	Promotion_F1 = func() {
		EditPromotion.URL = config.ADM_FUN + "/promotion/edit/" + strconv.Itoa(promotionResp.ID)
		DeletePromotion.URL = config.ADM_FUN + "/promotion/delete/" + strconv.Itoa(promotionResp.ID)
	}

	CreateTag = models.Request{
		URL:      config.ADM_FUN + "/tag/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newTag),
		Token:    &config.TokenAdmFun,
		RespData: &tagResp,
		Actions:  []func(){Tag_F1},
		Next:     &EditTag,
	}

	EditTag = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editTag),
		Token:   &config.TokenAdmFun,
		Next:    &DeleteTag,
	}

	DeleteTag = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmFun,
	}

	newTag = models.TagReq{
		Title: "AUTOTEST",
		Color: "yellow",
	}

	editTag = models.TagReq{
		Title: "EDITTED AUTOTEST",
		Color: "red",
	}

	tagResp models.TagResp

	Tag_F1 = func() {
		EditTag.URL = config.ADM_FUN + "/tag/edit/" + strconv.Itoa(tagResp.ID)
		DeleteTag.URL = config.ADM_FUN + "/tag/delete/" + strconv.Itoa(tagResp.ID)
	}

	CreateWorkplace = models.Request{
		URL:      config.ADM_FUN + "/workplace/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newWorkplace), // Pushkin
		Token:    &config.TokenAdmFun,
		RespData: &workplaceResp,
		Actions:  []func(){WP_F1},
		Next:     &EditWorkplace,
	}

	EditWorkplace = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editWorkplace), // Pushkin
		Token:   &config.TokenAdmFun,
		Next:    &DeleteWorkplace,
	}

	DeleteWorkplace = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmFun,
	}

	newWorkplace = models.WorkplaceReq{
		Title: "AUTOTEST",
		Club:  13,
	}

	editWorkplace = models.WorkplaceReq{
		Title: "EDITTED AUTOTEST",
		Club:  13,
	}

	workplaceResp models.WorkplaceResp

	WP_F1 = func() {
		EditWorkplace.URL = config.ADM_FUN + "/workplace/edit/" + strconv.Itoa(workplaceResp.ID)
		DeleteWorkplace.URL = config.ADM_FUN + "/workplace/delete/" + strconv.Itoa(workplaceResp.ID)
	}

	CreateTerm = models.Request{
		URL:      config.ADM_FUN + "/term/add",
		Method:   "POST",
		Token:    &config.TokenAdmFun,
		ReqBody:  utils.PrepareReqBody(newTerm),
		RespData: &resp,
		Actions:  []func(){Term_F1},
		Next:     &EditTerm,
	}

	EditTerm = models.Request{
		Method:  "PUT",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(editTerm),
		Next:    &DeleteTerm,
	}

	DeleteTerm = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmFun,
	}

	newTerm = models.TermReq{
		Title:       "AUTOTEST",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
	}

	editTerm = models.TermReq{
		Title:       "EDITTED AUTOTEST",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
	}

	Term_F1 = func() {
		EditTerm.URL = config.ADM_FUN + "/term/edit/" + strconv.Itoa(resp.ID)
		DeleteTerm.URL = config.ADM_FUN + "/term/delete/" + strconv.Itoa(resp.ID)
	}

	ASD = models.Request{
		URL:     config.ADM_FUN + "/access-system-device/add",
		Method:  "POST",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(&newASD),
	}

	newASD = models.ASDReq{
		SerialNumber:   "010101",
		Type:           "AUTOTEST",
		Code:           "1111",
		OpenAfterTime:  1200,
		OpenBeforeTime: 600,
		Workplace:      3,
	}

	CreateUser = models.Request{
		URL:      config.ADM_FUN + "/user/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newUser),
		Token:    &config.TokenAdmFun,
		RespData: &userResp,
		Actions:  []func(){User_F1},
		Next:     &EditUser,
	}

	EditUser = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editUser),
		Token:   &config.TokenAdmFun,
		Next:    &DeleteUser,
	}

	DeleteUser = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmFun,
	}

	newUser = models.UserReq{
		Phone:              "73456861286",
		Email:              "autotest@gmail.com",
		FirstName:          "TEST",
		SecondName:         "TEST",
		LastName:           "TEST",
		BirthDate:          "1996-01-01",
		WorkExperience:     "1996-01-01",
		Gender:             "male",
		VerificationStatus: "not_verified",
		Tags:               []int{},
	}

	editUser = models.UserReq{
		Phone:              "73456861286",
		Email:              "editted@gmail.com",
		FirstName:          "Editted",
		SecondName:         "Editted",
		LastName:           "Editted",
		BirthDate:          "1996-01-01",
		WorkExperience:     "1996-01-01",
		Gender:             "male",
		VerificationStatus: "not_verified",
		Tags:               []int{},
	}

	userResp models.UserResp

	User_F1 = func() {
		EditUser.URL = config.ADM_FUN + "/user/edit/" + strconv.Itoa(userResp.ID)
		DeleteUser.URL = config.ADM_FUN + "/user/delete/" + strconv.Itoa(userResp.ID)
	}

	CreateGroupTrain = models.Request{
		URL:      config.ADM_FUN + "/group-training/add",
		Method:   "POST",
		Token:    &config.TokenAdmFun,
		ReqBody:  utils.PrepareReqBody(&newGroupTrain),
		RespData: &resp,
		Actions:  []func(){GroupTrain_F1},
		Next:     &EditGroupTrain,
	}

	EditGroupTrain = models.Request{
		Method:  "PUT",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(&editGroupTrain),
		Next:    &DeleteGroupTrain,
	}

	DeleteGroupTrain = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmFun,
	}

	newGroupTrain = models.GroupTrainReq{
		Title:   "AUTOTEST",
		Slots:   6,
		Service: 10,
		Club:    29,
		Trainer: 6289,
	}

	editGroupTrain = models.GroupTrainReq{
		Title:   "EDITTED AUTOTEST",
		Slots:   10,
		Service: 10,
		Club:    29,
		Trainer: 6290,
	}

	GroupTrain_F1 = func() {
		EditGroupTrain.URL = config.ADM_FUN + "/group-training/edit/" + strconv.Itoa(resp.ID)
		DeleteGroupTrain.URL = config.ADM_FUN + "/group-training/delete/" + strconv.Itoa(resp.ID)
	}

	BonusRefferal = models.Request{
		URL:     config.ADM_FUN + "/setting/edit/3",
		Method:  "PUT",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(value1),
	}

	BonusMaxProcent = models.Request{
		URL:     config.ADM_FUN + "/setting/edit/5",
		Method:  "PUT",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(value2),
	}

	BonusMinProcent = models.Request{
		URL:     config.ADM_FUN + "/setting/edit/4",
		Method:  "PUT",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(value3),
	}

	BonusDayBeforeDel = models.Request{
		URL:     config.ADM_FUN + "/setting/edit/7",
		Method:  "PUT",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(value4),
	}

	BonusDelProcent = models.Request{
		URL:     config.ADM_FUN + "/setting/edit/6",
		Method:  "PUT",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(value5),
	}

	BonusAdd = models.Request{
		URL:     config.ADM_FUN + "/setting/edit/1",
		Method:  "PUT",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(value6),
	}

	BonusRemove = models.Request{
		URL:     config.ADM_FUN + "/setting/edit/2",
		Method:  "PUT",
		Token:   &config.TokenAdmFun,
		ReqBody: utils.PrepareReqBody(value7),
	}

	value1 = models.BonusReq{Value: 500}
	value2 = models.BonusReq{Value: 50}
	value3 = models.BonusReq{Value: 20}
	value4 = models.BonusReq{Value: 7}
	value5 = models.BonusReq{Value: 30}
	value6 = models.BonusReq{Value: 5}
	value7 = models.BonusReq{Value: 4}
)

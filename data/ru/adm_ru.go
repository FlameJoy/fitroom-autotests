package data_ru

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
		// &CreateBranch,
		// &CreateClub,
		// &CreatePackage,
		// &CreatePromotion,
		// &CreateTag,
		// &CreateWorkplace,
		// &CreateTerm,
		// &ASD,
		// &CreateUser,
		// &CreateGroupTrain,
		// // Bonus --------------------------------------------------------
		// &BonusRefferal,
		// &BonusAdd,
		// &BonusRemove,
		// &BonusDayBeforeDel,
		// &BonusMaxProcent,
		// &BonusMinProcent,
		// &BonusDelProcent,
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
		URL:    config.ADM_RU + "/permission/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetProductType = models.Request{
		URL:    config.ADM_RU + "/service-product-type/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetPromotionList = models.Request{
		URL:    config.ADM_RU + "/promotion/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetPromotionType = models.Request{
		URL:    config.ADM_RU + "/promotion-type/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetMenu = models.Request{
		URL:    config.ADM_RU + "/menu/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetDistrictList = models.Request{
		URL:    config.ADM_RU + "/district/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetComplexList = models.Request{
		URL:    config.ADM_RU + "/complex/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetMetroList = models.Request{
		URL:    config.ADM_RU + "/metro/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetBranchList = models.Request{
		URL:    config.ADM_RU + "/branch/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetAppointmentsList = models.Request{
		URL:     config.ADM_RU + "/appointment/get-list",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(appointmentList),
		Token:   &config.TokenAdmRu,
	}

	appointmentList = models.Appointments{
		Status:   "booked",
		FromDate: config.C_Date,
		ToDate:   config.F_Date,
	}

	GetCountryList = models.Request{
		URL:    config.ADM_RU + "/country/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetNotifications = models.Request{
		URL:    config.ADM_RU + "/notification/get-list",
		Method: "POST",

		Token: &config.TokenAdmRu,
	}

	GetOrdersList = models.Request{
		URL:     config.ADM_RU + "/order/get-list",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(pageLimit),
		Token:   &config.TokenAdmRu,
	}

	GetOrders = models.Request{
		URL:     config.ADM_RU + "/order/get-list",
		Method:  "POST",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(&pageLimit),
	}

	GetGroupAppintments = models.Request{
		URL:     config.ADM_RU + "/appointment/available",
		Method:  "POST",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(groupAppointmentList),
	}

	GetNotification = models.Request{
		URL:    config.ADM_RU + "/notification/get-list",
		Method: "POST",
		Token:  &config.TokenAdmRu,
	}

	groupAppointmentList = models.GroupAppointments{
		StartDate:           config.P_Date,
		EndDate:             config.C_Date,
		Workplace:           27,
		ServiceProductTypes: []int{1, 2},
	}

	GetReview = models.Request{
		URL:     config.ADM_RU + "/review/get-list",
		Method:  "POST",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(pages),
	}

	GetconfigList = models.Request{
		URL:    config.ADM_RU + "/setting/get-list",
		Method: "POST",
		Token:  &config.TokenAdmRu,
	}

	GetTagsList = models.Request{
		URL:    config.ADM_RU + "/tag/get-list",
		Method: "POST",
		Token:  &config.TokenAdmRu,
	}

	GetTerms = models.Request{
		URL:    config.ADM_RU + "/term/get",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetUserByID = models.Request{
		URL:    config.ADM_RU + "/user/get-by-id/2058",
		Method: "GET",
		Token:  &config.TokenAdmRu,
	}

	GetUsersList = models.Request{
		URL:     config.ADM_RU + "/user/get-list",
		Method:  "POST",
		ReqBody: utils.PrepareReqBody(pageLimit),
		Token:   &config.TokenAdmRu,
	}

	GetWorkplaceList = models.Request{
		URL:    config.ADM_RU + "/workplace/get-list",
		Method: "POST",
		Token:  &config.TokenAdmRu,
	}

	ReportBonus = models.Request{
		URL:     config.ADM_RU + "/report/bonus-flow",
		Method:  "POST",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(newRepBonus),
	}

	newRepBonus = models.RepBonus{
		FromDate: config.C_Date,
		ToDate:   config.F_Date,
	}

	ReportClub = models.Request{
		URL:     config.ADM_RU + "/report/club",
		Method:  "POST",
		Token:   &config.TokenAdmRu,
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
		URL:     config.ADM_RU + "/report/conversion",
		Method:  "POST",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(newRepConversion),
	}

	newRepConversion = models.RepConversion{
		FromDate: config.P_Date,
		ToDate:   config.C_Date,
	}

	ReportFinancial = models.Request{
		URL:     config.ADM_RU + "/report/revenue",
		Method:  "POST",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(newRepFinancial),
	}

	newRepFinancial = models.RepFinancial{
		Clubs:    GetClubIDList(),
		FromDate: config.P_Date,
		ToDate:   config.C_Date,
	}

	ReportTrainers = models.Request{
		URL:     config.ADM_RU + "/report/club-trainers",
		Method:  "POST",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(newRepTrainers),
	}

	newRepTrainers = models.RepTrainers{
		FromDate:      config.P_Date,
		ToDate:        config.C_Date,
		SalaryPercent: 50,
	}

	AdminBook = models.Request{
		URL:      config.ADM_RU + "/appointment/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(adminBook),
		RespData: &adminBookResp,
		Token:    &config.TokenAdmRu,
		Actions:  []func(){AdminBook_F1},
		Next:     &CancelRentAdm,
	}

	adminBook = models.Book{
		Clients:            []uint{6775},
		Club:               21,
		Workplace:          27,
		ServiceProductType: 1,
		StartDate:          config.StartDate1,
		EndDate:            config.EndDate1,
		Status:             "booked",
		User:               6775,
		UserPackage:        33965,
	}

	adminBookResp models.Resp

	AdminBook_F1 = func() {
		CancelRentAdm.URL = config.ADM_RU + "/appointment/delete/" + strconv.Itoa(adminBookResp.ID)
	}

	CancelRentAdm = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmRu,
	}

	CreateBranch = models.Request{
		URL:      config.ADM_RU + "/branch/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newBranch),
		Token:    &config.TokenAdmRu,
		RespData: &branchResp,
		Actions:  []func(){Branch_F1},
		Next:     &EditBranch,
	}

	EditBranch = models.Request{
		Method:  "PUT",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(editBranch),
		Next:    &DeleteBranch,
	}

	DeleteBranch = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmRu,
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
		EditBranch.URL = config.ADM_RU + "/branch/edit/" + strconv.Itoa(branchResp.ID)
		DeleteBranch.URL = config.ADM_RU + "/branch/delete/" + strconv.Itoa(branchResp.ID)
	}

	CreateClub = models.Request{
		URL:      config.ADM_RU + "/club/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newClub),
		Token:    &config.TokenAdmRu,
		RespData: &clubResp,
		Actions:  []func(){Club_F1},
		Next:     &EditClub,
	}

	EditClub = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editClub),
		Token:   &config.TokenAdmRu,
		Next:    &DeleteClub,
	}

	DeleteClub = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmRu,
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
		EditClub.URL = config.ADM_RU + "/club/edit/" + strconv.Itoa(clubResp.ID)
		DeleteClub.URL = config.ADM_RU + "/club/delete/" + strconv.Itoa(clubResp.ID)
	}

	CreatePackage = models.Request{
		URL:      config.ADM_RU + "/package/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newPackage),
		Token:    &config.TokenAdmRu,
		RespData: &packageResp,
		Actions:  []func(){Package_F1},
		Next:     &EditPackage,
	}

	EditPackage = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editPackage),
		Token:   &config.TokenAdmRu,
		Next:    &DeletePackage,
	}

	DeletePackage = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmRu,
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
		EditPackage.URL = config.ADM_RU + "/package/edit/" + strconv.Itoa(packageResp.ID)
		DeletePackage.URL = config.ADM_RU + "/package/delete/" + strconv.Itoa(packageResp.ID)
	}

	CreatePromotion = models.Request{
		URL:      config.ADM_RU + "/promotion/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newPromotion),
		Token:    &config.TokenAdmRu,
		RespData: &promotionResp,
		Actions:  []func(){Promotion_F1},
		Next:     &EditPromotion,
	}

	EditPromotion = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editPromotion),
		Token:   &config.TokenAdmRu,
		Next:    &DeletePromotion,
	}

	DeletePromotion = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmRu,
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
		EditPromotion.URL = config.ADM_RU + "/promotion/edit/" + strconv.Itoa(promotionResp.ID)
		DeletePromotion.URL = config.ADM_RU + "/promotion/delete/" + strconv.Itoa(promotionResp.ID)
	}

	CreateTag = models.Request{
		URL:      config.ADM_RU + "/tag/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newTag),
		Token:    &config.TokenAdmRu,
		RespData: &tagResp,
		Actions:  []func(){Tag_F1},
		Next:     &EditTag,
	}

	EditTag = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editTag),
		Token:   &config.TokenAdmRu,
		Next:    &DeleteTag,
	}

	DeleteTag = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmRu,
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
		EditTag.URL = config.ADM_RU + "/tag/edit/" + strconv.Itoa(tagResp.ID)
		DeleteTag.URL = config.ADM_RU + "/tag/delete/" + strconv.Itoa(tagResp.ID)
	}

	CreateWorkplace = models.Request{
		URL:      config.ADM_RU + "/workplace/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newWorkplace), // Pushkin
		Token:    &config.TokenAdmRu,
		RespData: &workplaceResp,
		Actions:  []func(){WP_F1},
		Next:     &EditWorkplace,
	}

	EditWorkplace = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editWorkplace), // Pushkin
		Token:   &config.TokenAdmRu,
		Next:    &DeleteWorkplace,
	}

	DeleteWorkplace = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmRu,
	}

	newWorkplace = models.WorkplaceReq{
		Title: "AUTOTEST",
		Club:  21,
	}

	editWorkplace = models.WorkplaceReq{
		Title: "EDITTED AUTOTEST",
		Club:  21,
	}

	workplaceResp models.WorkplaceResp

	WP_F1 = func() {
		EditWorkplace.URL = config.ADM_RU + "/workplace/edit/" + strconv.Itoa(workplaceResp.ID)
		DeleteWorkplace.URL = config.ADM_RU + "/workplace/delete/" + strconv.Itoa(workplaceResp.ID)
	}

	CreateTerm = models.Request{
		URL:      config.ADM_RU + "/term/add",
		Method:   "POST",
		Token:    &config.TokenAdmRu,
		ReqBody:  utils.PrepareReqBody(newTerm),
		RespData: &resp,
		Actions:  []func(){Term_F1},
		Next:     &EditTerm,
	}

	EditTerm = models.Request{
		Method:  "PUT",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(editTerm),
		Next:    &DeleteTerm,
	}

	DeleteTerm = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmRu,
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
		EditTerm.URL = config.ADM_RU + "/term/edit/" + strconv.Itoa(resp.ID)
		DeleteTerm.URL = config.ADM_RU + "/term/delete/" + strconv.Itoa(resp.ID)
	}

	// ASD = models.Request{
	// 	URL:     config.ADM_RU + "/access-system-device/add",
	// 	Method:  "POST",
	// 	Token:   &config.TokenAdmRu,
	// 	ReqBody: utils.PrepareReqBody(&newASD),
	// }

	// newASD = models.ASDReq{
	// 	SerialNumber:   "010101",
	// 	Type:           "AUTOTEST",
	// 	Code:           "1111",
	// 	OpenAfterTime:  1200,
	// 	OpenBeforeTime: 600,
	// 	Workplace:      3,
	// }

	CreateUser = models.Request{
		URL:      config.ADM_RU + "/user/add",
		Method:   "POST",
		ReqBody:  utils.PrepareReqBody(newUser),
		Token:    &config.TokenAdmRu,
		RespData: &userResp,
		Actions:  []func(){User_F1},
		Next:     &EditUser,
	}

	EditUser = models.Request{
		Method:  "PUT",
		ReqBody: utils.PrepareReqBody(editUser),
		Token:   &config.TokenAdmRu,
		Next:    &DeleteUser,
	}

	DeleteUser = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmRu,
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
		EditUser.URL = config.ADM_RU + "/user/edit/" + strconv.Itoa(userResp.ID)
		DeleteUser.URL = config.ADM_RU + "/user/delete/" + strconv.Itoa(userResp.ID)
	}

	CreateGroupTrain = models.Request{
		URL:      config.ADM_RU + "/group-training/add",
		Method:   "POST",
		Token:    &config.TokenAdmRu,
		ReqBody:  utils.PrepareReqBody(&newGroupTrain),
		RespData: &resp,
		Actions:  []func(){GroupTrain_F1},
		Next:     &EditGroupTrain,
	}

	EditGroupTrain = models.Request{
		Method:  "PUT",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(&editGroupTrain),
		Next:    &DeleteGroupTrain,
	}

	DeleteGroupTrain = models.Request{
		Method: "DELETE",
		Token:  &config.TokenAdmRu,
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
		EditGroupTrain.URL = config.ADM_RU + "/group-training/edit/" + strconv.Itoa(resp.ID)
		DeleteGroupTrain.URL = config.ADM_RU + "/group-training/delete/" + strconv.Itoa(resp.ID)
	}

	BonusRefferal = models.Request{
		URL:     config.ADM_RU + "/setting/edit/3",
		Method:  "PUT",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(value1),
	}

	BonusMaxProcent = models.Request{
		URL:     config.ADM_RU + "/setting/edit/5",
		Method:  "PUT",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(value2),
	}

	BonusMinProcent = models.Request{
		URL:     config.ADM_RU + "/setting/edit/4",
		Method:  "PUT",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(value3),
	}

	BonusDayBeforeDel = models.Request{
		URL:     config.ADM_RU + "/setting/edit/7",
		Method:  "PUT",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(value4),
	}

	BonusDelProcent = models.Request{
		URL:     config.ADM_RU + "/setting/edit/6",
		Method:  "PUT",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(value5),
	}

	BonusAdd = models.Request{
		URL:     config.ADM_RU + "/setting/edit/1",
		Method:  "PUT",
		Token:   &config.TokenAdmRu,
		ReqBody: utils.PrepareReqBody(value6),
	}

	BonusRemove = models.Request{
		URL:     config.ADM_RU + "/setting/edit/2",
		Method:  "PUT",
		Token:   &config.TokenAdmRu,
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

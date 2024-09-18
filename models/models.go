package models

import "io"

type AuthReq struct {
	Code string `json:"code"`
}

type AuthResp struct {
	Token string `json:"token"`
}

type Requester interface {
	GET() error
	POST() error
	PUT() error
	DELETE() error
	AddToFailed()
}

type Request struct {
	URL      string
	Method   string
	ReqBody  []byte
	RespData any
	Token    *string
	Actions  []func() `json:"-"` // ignore for MarshalIndent
	Next     *Request
}

type FailedRequest struct {
	URL      string
	Method   string
	ReqBody  []byte
	RespBody io.ReadCloser
	RespData any
	Token    string
}

type History struct {
	FromDate   string `json:"fromDate"`
	ToDate     string `json:"toDate"`
	OnlyGroups int    `json:"onlyGroups"`
	Page
}

type Page struct {
	Paging Paging `json:"paging"`
}

type Paging struct {
	Page  uint `json:"page"`
	Limit uint `json:"limit"`
}

type GroupTrainBook struct {
	User               uint   `json:"user"`
	Club               uint   `json:"club"`
	Workplace          uint   `json:"workplace"`
	ServiceProductType uint   `json:"serviceProductType"`
	StartDate          string `json:"startDate"`
	EndDate            string `json:"endDate"`
	UserPackage        uint   `json:"userPackage"`
	GroupTraining      int    `json:"groupTraining"`
}

type Resp struct {
	ID int `json:"id"`
}

type Available struct {
	Workplace           int    `json:"workplace"`
	ServiceProductTypes []int  `json:"serviceProductTypes"`
	Trainers            []int  `json:"trainers"`
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
}

type City struct {
	City uint `json:"city"`
}

type Cities struct {
	Cities     []int `json:"cities"`
	OnlyGroups int   `json:"onlyGroups"`
}

type Book struct {
	Clients            []uint `json:"clients"`
	Status             string `json:"status"`
	User               uint   `json:"user"`
	Club               uint   `json:"club"`
	Workplace          uint   `json:"workplace"`
	ServiceProductType uint   `json:"serviceProductType"`
	StartDate          string `json:"startDate"`
	EndDate            string `json:"endDate"`
	UserPackage        uint   `json:"userPackage"`
}
type Club struct {
	Club uint `json:"club"`
}

type Appointments struct {
	Status   string `json:"status"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

type GroupAppointments struct {
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
	Workplace           int    `json:"workplace"`
	ServiceProductTypes []int  `json:"serviceProductTypes"`
}

type RepBonus struct {
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

type RepClub struct {
	Clubs         []int  `json:"clubs"`
	FromDate      string `json:"fromDate"`
	ToDate        string `json:"toDate"`
	SalaryPercent int    `json:"salaryPercent"`
	Paging        `json:"paging"`
}

type RepConversion struct {
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

type RepFinancial struct {
	Clubs    []int  `json:"clubs"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

type RepTrainers struct {
	SalaryPercent uint   `json:"salaryPercent"`
	FromDate      string `json:"fromDate"`
	ToDate        string `json:"toDate"`
}

type BranchReq struct {
	Active      uint   `json:"active"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BranchResp struct {
	ID          int    `json:"id"`
	Actice      bool   `json:"actice"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Gallery     []struct {
		ID int `json:"id"`
	} `json:"gallery"`
}

type ClubReq struct {
	Title     string `json:"title"`
	City      uint   `json:"city"`
	District  uint   `json:"district"`
	Complex   uint   `json:"complex"`
	Street    string `json:"street"`
	House     string `json:"house"`
	Building  string `json:"building"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Cardio    uint   `json:"cardio"`
	OpenDate  string `json:"openDate"`
}

type ClubResp struct {
	ID        int    `json:"id"`
	Actice    bool   `json:"actice"`
	Title     string `json:"title"`
	Street    string `json:"street"`
	House     string `json:"house"`
	Building  string `json:"building"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	City      struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"city"`
	Complex struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"complex"`
	District struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"district"`
	Users   []int `json:"users"`
	Gallery struct {
		ID int `json:"id"`
	} `json:"gallery"`
	ServiceProductTypes []int  `json:"serviceProductTypes"`
	Cardio              bool   `json:"cardio"`
	OpenDate            string `json:"openDate"`
}

type PackageReq struct {
	Title               string `json:"title"`
	ServiceProductType  uint   `json:"serviceProductType"`
	TrainingAmount      uint   `json:"trainingAmount"`
	Cost                uint   `json:"cost"`
	Cashback            uint   `json:"cashback"`
	WritingOff          uint   `json:"writingOff"`
	DaysAmount          uint   `json:"daysAmount"`
	CashbackStartDate   string `json:"cashbackStartDate"`
	CashbackEndDate     string `json:"cashbackEndDate"`
	WritingOffStartDate string `json:"writingOffStartDate"`
	WritingOffEndDate   string `json:"writingOffEndDate"`
}

type PackageResp struct {
	ID                 int    `json:"id"`
	Title              string `json:"title"`
	DaysAmount         uint   `json:"daysAmount"`
	TrainingAmount     uint   `json:"trainingAmount"`
	Cost               string `json:"cost"`
	Clubs              []int  `json:"clubs"`
	ServiceProductType struct {
		ID                   int    `json:"id"`
		Title                string `json:"title"`
		VerificationRequired bool   `json:"verificationRequired"`
	} `json:"serviceProductType"`
	IsTrial            bool  `json:"isTrial"`
	PackageConstraints []int `json:"packageConstraints"`
}

type PromotionReq struct {
	Title         string `json:"title"`
	PromotionType int    `json:"promotionType"`
	Segments      []int  `json:"segments"`
	Clubs         []int  `json:"clubs"`
	Packages      []int  `json:"packages"`
	Discount      int    `json:"discount"`
	Count         int    `json:"count"`
	PersonalCount int    `json:"personalCount"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
}

type PromotionResp struct {
	ID            int  `json:"id"`
	Actice        bool `json:"actice"`
	PromotionType struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"promotionType"`
	Title         string `json:"title"`
	Discount      string `json:"discount"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	Count         int    `json:"count"`
	PersonalCount int    `json:"personalCount"`
	Packages      []int  `json:"packages"`
	Clubs         []int  `json:"clubs"`
	Segments      []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"segments"`
	RestrictUse bool `json:"restrictUse"`
}

type TagReq struct {
	Title string `jsom:"title"`
	Color string `jsom:"color"`
}

type TagResp struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

type WorkplaceReq struct {
	Title string `json:"title"`
	Club  uint   `json:"club"`
}

type WorkplaceResp struct {
	ID     int  `json:"id"`
	Actice bool `json:"actice"`
	Club   struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Street   string `json:"street"`
		House    string `json:"house"`
		Building string `json:"building"`
		City     struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"club"`
		Complex  []int `json:"complex"`
		District []int `json:"district"`
	} `json:"club"`
	Title   string `json:"title"`
	Gallery struct {
		ID int `json:"id"`
	} `json:"gallery"`
}

type TermReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ASDReq struct {
	SerialNumber   string `json:"serialNumber"`
	Type           string `json:"type"`
	Code           string `json:"code"`
	OpenAfterTime  int    `json:"openAfterTime"`
	OpenBeforeTime int    `json:"openBeforeTime"`
	Workplace      int    `json:"workplace"`
}

type UserReq struct {
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	FirstName          string `json:"firstName"`
	SecondName         string `json:"secondName"`
	LastName           string `json:"lastName"`
	BirthDate          string `json:"birthDate"`
	WorkExperience     string `json:"workExperience"`
	Gender             string `json:"gender"`
	VerificationStatus string `json:"verificationStatus"`
	Tags               []int  `json:"tags"`
}

type UserResp struct {
	ID                 int      `json:"id"`
	Email              string   `json:"email"`
	Phone              string   `json:"phone"`
	Roles              []string `json:"roles"`
	Actice             bool     `json:"actice"`
	FirstName          string   `json:"firstName"`
	SecondName         string   `json:"secondName"`
	LastName           string   `json:"lastName"`
	Gender             string   `json:"gender"`
	VerificationStatus string   `json:"verificationStatus"`
	BirthDate          string   `json:"birthDate"`
	CreatedAt          string   `json:"createdAt"`
	UpdatedAt          string   `json:"updatedAt"`
	Gallery            struct {
		ID int `json:"id"`
	} `json:"gallery"`
	Tags             []int  `json:"tags"`
	BonusBalance     string `json:"bonusBalance  "`
	ReferralNeedGift bool   `json:"referralNeedGift"`
}

type GroupTrainReq struct {
	Title   string `json:"title"`
	Slots   int    `json:"slots"`
	Service int    `json:"service"`
	Club    int    `json:"club"`
	Trainer int    `json:"trainer"`
}

type BonusReq struct {
	Value int `json:"value"`
}

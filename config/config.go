package config

import (
	"fmt"
	"time"
)

const (
	LAND_FUN = "https://fitroom.fun"
	ADM_FUN  = "https://back.fitroom.fun/api/admin"
	APP_FUN  = "https://back.fitroom.fun/api/lk"
	LAND_RU  = "https://fitroom.ru"
	APP_RU   = "https://back.fitroom.ru/api/lk"
	ADM_RU   = "https://back.fitroom.ru/api/admin"
	BACK_RU  = "https://back.fitroom.ru"
	BACK_FUN = "https://back.fitroom.fun"
)

var (
	TokenUserFun string
	TokenUserRu  string
	TokenAdmFun  string
	TokenAdmRu   string
	LogsLength   = 90
)

var Year, Month, Day = time.Now().Date()
var PastDate = time.Date(Year, Month, Day-3, 0, 0, 0, 0, time.Local)
var CurrentDate = time.Date(Year, Month, Day, 0, 0, 0, 0, time.Local)
var FutureDate = time.Date(Year, Month, Day+3, 0, 0, 0, 0, time.Local)
var year, month, day = CurrentDate.Date()
var fYear, fMonth, fDay = FutureDate.Date()
var pYear, pMonth, pDay = PastDate.Date()

// Dates
var C_Date = fmt.Sprintf("%d-%02d-%02d", year, month, day)
var F_Date = fmt.Sprintf("%d-%02d-%02d", fYear, fMonth, fDay)
var P_Date = fmt.Sprintf("%d-%02d-%02d", pYear, pMonth, pDay)

// Rent
var StartDate1 = fmt.Sprintf("%d-%02d-%02d 00:00", fYear, fMonth, fDay)
var EndDate1 = fmt.Sprintf("%d-%02d-%02d 01:00", fYear, fMonth, fDay)

// Training
var StartDate2 = fmt.Sprintf("%d-%02d-%02d 01:00", fYear, fMonth, fDay)
var EndDate2 = fmt.Sprintf("%d-%02d-%02d 02:00", fYear, fMonth, fDay)

// Cardio
var StartDate3 = fmt.Sprintf("%d-%02d-%02d 02:00", fYear, fMonth, fDay)
var EndDate3 = fmt.Sprintf("%d-%02d-%02d 03:00", fYear, fMonth, fDay)

// Default 5 days
var StartDate = fmt.Sprintf("%d-%02d-%02dT00:00:00.760Z", year, month, day)
var EndDate = fmt.Sprintf("%d-%02d-%02dT00:00:00.760Z", fYear, fMonth, fDay)

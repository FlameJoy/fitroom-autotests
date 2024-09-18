package main

import (
	"fitroom-autotests/config"
	"fitroom-autotests/models"
	"fmt"
	"strconv"
)

var GetClubsList = models.Request{
	URL:      config.BACK_FUN + "/public/club/get-list",
	Method:   "POST",
	Token:    &config.TokenUserFun,
	RespData: &clubList,
}

var clubList = []struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}{}

var DeleteClub = models.Request{
	Method: "DELETE",
	Token:  &config.TokenAdmFun,
}

var token = ""

func main() {
	GetClubsList.POST()
	fmt.Println(clubList)
	for _, club := range clubList {
		if club.Title == "Test" {
			req := models.Request{
				URL:    config.ADM_FUN + "/club/delete/" + strconv.Itoa(club.ID),
				Method: "DELETE",
				Token:  &token,
			}
			req.DELETE()
		}
	}
}

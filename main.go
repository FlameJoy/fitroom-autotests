package main

import (
	"fitroom-autotests/config"
	data "fitroom-autotests/data/fun"
	"fitroom-autotests/initializers"
	"fitroom-autotests/models"
	"fitroom-autotests/utils"
	"flag"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	TotalReqCounter int
	landFun         = flag.Bool("landing", false, "Regress landing")
	appFun          = flag.Bool("appfun", false, "Regress app.Stage.fun")
	admFun          = flag.Bool("admfun", false, "Regress adm.stage.fun")
	failed          = flag.Bool("failed", false, "Resend failed requests")
	// special         = flag.Bool("special", false, "Special requests")
)

func main() {
	start := time.Now()
	flag.Parse()
	// Init
	initializers.LoadENV()
	// Auth
	fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("Authorization"))
	// User
	config.TokenUserFun = GetToken("https://back.fitroom.fun", initializers.UserTel, initializers.UserCode)
	config.TokenUserRu = GetToken("https://back.fitroom.ru", initializers.UserTel, initializers.UserCode)
	// Adm
	config.TokenAdmFun = GetToken("https://back.fitroom.fun", initializers.AdmTel, initializers.AdmCode)
	config.TokenAdmRu = GetToken("https://back.fitroom.ru", initializers.AdmTel, initializers.AdmCode)
	// Regress
	switch {
	case *failed:
		models.LoadFailedRequests()
		regressAsync(models.FailedMethodSlice)
		models.SaveFailedRequests()
	case *landFun:
		fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("Landing"))
		regressAsync(data.LandFunRequestList)
	case *appFun:
		fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("AppFun"))
		regressAsync(data.AppFunRequestList)
	case *admFun:
		fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("AdmFun"))
		regressAsync(data.AdmFunRequestList)
	default:
		fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("Landing"))
		regressAsync(data.LandFunRequestList)
		fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("AppFun"))
		regressAsync(data.AppFunRequestList)
		fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("AdmFun"))
		regressAsync(data.AdmFunRequestList)
	}
	models.SaveFailedRequests()
	fmt.Println(time.Since(start))
}

func regress(reqList []models.Requester) {
	TotalReqCounter += len(reqList)
	for _, req := range reqList {
		switch req.(*models.Request).Method {
		case "GET":
			req.GET()
		case "POST":
			req.POST()
		case "PUT":
			req.PUT()
		case "DELETE":
			req.DELETE()
		default:
			log.Printf("\033[96m%s\033[0m method \033[33m%s\033[0m is not supported \n", utils.URLFormat(req.(*models.Request).Method, req.(*models.Request).URL), req.(*models.Request).Method)
		}
	}
}

func regressAsync(reqList []models.Requester) {
	var wg sync.WaitGroup
	wg.Add(len((reqList)))
	for _, req := range reqList {
		go func(req models.Requester) {
			defer wg.Done()
			switch req.(*models.Request).Method {
			case "GET":
				req.GET()
			case "POST":
				req.POST()
			case "PUT":
				req.PUT()
			case "DELETE":
				req.DELETE()
			default:
				log.Printf("\033[96m%s\033[0m method \033[33m%s\033[0m is not supported \n", utils.URLFormat(req.(*models.Request).Method, req.(*models.Request).URL), req.(*models.Request).Method)
			}
		}(req)
	}
	wg.Wait()
}

func GetToken(url, tel, code string) string {
	var r models.Request
	r.URL = url + "/auth/authorize/" + tel
	r.Method = "POST"
	r.ReqBody = utils.PrepareReqBody(models.AuthReq{Code: code})
	var data = models.AuthResp{}
	r.RespData = &data
	r.POST()
	return "Bearer " + data.Token
}

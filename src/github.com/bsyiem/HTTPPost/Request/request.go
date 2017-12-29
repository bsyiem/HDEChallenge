package Request

import (
	"github.com/parnurzeal/gorequest"
)

func SendPOST(userid string,pass string) string{
	request := gorequest.New().SetBasicAuth(userid,pass)
	_, body, errs := request.Post("http://hdechallenge-solve.appspot.com/challenge/003/endpoint").
		Send(`{
			"github_url": "https://gist.github.com/bsyiem/72fb785fa6b8a2abbcc7fd64bc72591f",
			"contact_email": "bsyiem92@gmail.com"
			}`).
		End()
	if errs != nil{
		panic(errs[0])
	}
	return body
}
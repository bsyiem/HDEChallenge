package main

import (
	"fmt"
	"github.com/bsyiem/HTTPPost/OTP"
	"github.com/bsyiem/HTTPPost/Request"

)

func main() {
	reply := Request.SendPOST("bsyiem92@gmail.com",OTP.GetCode())
	fmt.Println(reply)
}

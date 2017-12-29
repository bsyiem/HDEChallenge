/*
code for generating valid secret code for HDE challenge
 */

package OTP

import (
	"github.com/pquerna/otp/totp"
	"github.com/pquerna/otp"
	"time"
	"encoding/base32"
	)

/*
function to return valid opts for HDE challenge
 */
func getOpts() totp.ValidateOpts {
	var opt totp.ValidateOpts
	opt.Algorithm = otp.AlgorithmSHA512
	opt.Period = 30
	opt.Digits = 10
	opt.Skew = 1

	return opt
}

/*
generates the code specific for the HDE challenge
 */
func GetCode() string{
	secret := base32.StdEncoding.EncodeToString([]byte("bsyiem92@gmail.comHDECHALLENGE003"))
	code,err := totp.GenerateCodeCustom(secret,time.Now(),getOpts())
	if err != nil{
		panic(err);
	}
	return code
}
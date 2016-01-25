package middlewares

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-martini/martini"
	"github.com/tcyd/goconf"
	"net/http"
)

func ApiVerify() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, context martini.Context, conf *goconf.ConfigFile) {
		//no verify on dev
		if martini.Env == martini.Dev {
			context.Next()
		}
		signatue := req.Header.Get("X-API-Signature")

		url := req.RequestURI
		apiKey, _ := conf.GetString("default", "apikey")
		timestamp := req.Header.Get("Timestamp")

		hash := md5.New()
		hash.Write([]byte(url + "?apiKey=" + apiKey + "&timestamp=" + timestamp))
		sig := hex.EncodeToString(hash.Sum(nil))
		if sig == signatue {
			context.Next()
		} else {
			res.WriteHeader(http.StatusUnauthorized)
		}
	}
}

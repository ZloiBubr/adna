package redirect

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	_ "runtime"
	"path/filepath"
	"runtime"
)

var (
	https_baseUrl string
//isProd        bool
)

const HTTPSAddr = "HTTPSAddr"
const HTTPSPort = "HTTPSPort"


func init() {
	if beego.BConfig.RunMode == "test" {
		_, file, _, _ := runtime.Caller(1)
		apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
		beego.TestBeegoInit(apppath)
	}

	//isProd = beego.BConfig.RunMode == "prod"
	//if isProd {
	if len(beego.AppConfig.String(HTTPSAddr)) == 0 {
		panic(fmt.Errorf("missing configuration '%s'", HTTPSAddr))
	}
	if len(beego.AppConfig.String(HTTPSPort)) == 0 {
		panic(fmt.Errorf("missing configuration '%s'", HTTPSPort))
	}
	https_baseUrl = "https://" + beego.AppConfig.String(HTTPSAddr) + ":" + beego.AppConfig.String(HTTPSPort)
	//}
}

func Https(ctx *context.Context) {
	//if isProd {
	if ctx.Request.TLS == nil {
		ctx.Redirect(http.StatusTemporaryRedirect, https_baseUrl+ctx.Request.URL.String())
	}
	//}
}


package redirect

import (
	"os"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"runtime"
	"path/filepath"
)

var (
	https_baseUrl string
)

const httpsConfigKey = "https_base"

func Https(ctx *context.Context) {
    if ctx.Input.Context.Request.TLS == nil {
        ctx.Redirect(http.StatusTemporaryRedirect, https_baseUrl+ctx.Input.Context.Request.URL.String())
    }
}

func init() {
    //workaround to let tests pass
    if(os.Getenv("RUNMODE") == "test") {
        _, file, _, _ := runtime.Caller(1)
        apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
        beego.TestBeegoInit(apppath)
    }
    //end workaround

    https_baseUrl = beego.AppConfig.String(httpsConfigKey)
    if len(https_baseUrl) == 0 {
        panic(fmt.Errorf("missing configuration '%s'", httpsConfigKey))
    }
}
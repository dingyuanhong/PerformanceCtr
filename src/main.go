package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"os"
	"os/exec"
	"net/http"
	"strings"
    "fmt"
)

func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func getCurrentPath() string {
    s, _ := exec.LookPath(os.Args[0])
    i := strings.LastIndex(s, "\\")
    path := string(s[0 : i+1])
    return path
}


var ROOT string = "static"
var Indexs []string = []string{};
//var Indexs []string = []string{"index.html","index.htm"};

func MapPath(path string) string{
	return "./" + ROOT + path
}

func TransparentStatic(ctx *context.Context) {
    defInd := 0
    maxInd := len(Indexs) - 1
    orpath := ctx.Request.URL.Path

DefaultIndexPage:
    p := orpath
    if strings.EqualFold(p, "/") {
        p += Indexs[defInd];
        defInd++
    }
    fp := MapPath(p);

    if !PathExists(fp) {
        if defInd > 0 && defInd < maxInd {
            goto DefaultIndexPage
        }
    } else {
        http.ServeFile(ctx.ResponseWriter, ctx.Request, fp)
    }
}

func main() {
    beego.SetStaticPath("/", "static/Extension/dist")
    // beego.DelStaticPath("/static");
    // beego.SetStaticPath("/static", "static/Extension/dist")
    //beego.SetStaticPath("/my", "static/Extension/dist")
    fmt.Println(beego.BConfig.WebConfig);
	beego.Run()
}


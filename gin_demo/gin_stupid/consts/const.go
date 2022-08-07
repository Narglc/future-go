package consts

import "github.com/gin-gonic/gin"

/*
const (
	TMPL_PAHT = "~/Narglc/go/src/future-go/gin_demo/gin_stupid/templates"
)
*/

var TMPL_PAHT string

func GetTmplPath() string {
	if mode := gin.Mode(); mode == gin.TestMode {
		TMPL_PAHT = "./../templates/*"
	} else {
		TMPL_PAHT = "templates/*"
	}
	return TMPL_PAHT
}

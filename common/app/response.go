package app

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/rolexkdev/emvn-music-library-server/common/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
	Code int         `json:"code"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	if data == nil {
		data = struct{}{}
	}

	if reflect.TypeOf(data).Name() == "string" {
		data = map[string]interface{}{
			"message": data,
		}
	}

	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}

func (g *Gin) Response200(data interface{}) {
	g.Response(http.StatusOK, e.SUCCESS, data)
	return
}

func (g *Gin) Response201(data interface{}) {
	g.Response(http.StatusCreated, e.CREATED, data)
	return
}

func (g *Gin) Response400(errCode int, data interface{}) {
	g.Response(http.StatusBadRequest, errCode, data)
	return
}

func (g *Gin) Response401(errCode int, data interface{}) {
	g.Response(http.StatusUnauthorized, errCode, data)
	return
}

func (g *Gin) Response403(errCode int, data interface{}) {
	g.Response(http.StatusForbidden, errCode, data)
	return
}

func (g *Gin) Response404(errCode int, data interface{}) {
	g.Response(http.StatusNotFound, errCode, data)
	return
}

func (g *Gin) Response409(errCode int, data interface{}) {
	g.Response(http.StatusConflict, errCode, data)
	return
}

func (g *Gin) Response500(errCode int, data interface{}) {
	g.Response(http.StatusInternalServerError, errCode, data)
	return
}

package controller

import (
	"awesomeProject/api/service"
	"awesomeProject/pkg/enums"
	"awesomeProject/pkg/models"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

func createtask(e echo.Context) error {
	var reqbody models.Task
	err := json.NewDecoder(e.Request().Body).Decode(&reqbody)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, enums.Faileddecode)
	}
	v := validator.New()
	err = v.Struct(&reqbody)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, enums.Validation)
	}
	err = service.Createtask(reqbody)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusInternalServerError, enums.ServerIssue)
	}
	return e.JSON(http.StatusOK, enums.Statusok)
}
func getalltask(e echo.Context) error {
	res, err := service.Getalltask()
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusInternalServerError, enums.ServerIssue)
	}
	return e.JSON(http.StatusOK, res)
}
func updatetask(e echo.Context) error {
	str, err := strconv.Atoi(e.Param("id"))
	var reqbody models.Task
	err = json.NewDecoder(e.Request().Body).Decode(&reqbody)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, enums.Faileddecode)
	}
	v := validator.New()
	err = v.Struct(&reqbody)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, enums.Validation)
	}
	err = service.Updatetask(reqbody, str)
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusInternalServerError, enums.ServerIssue)
	}
	return e.JSON(http.StatusOK, enums.Statusok)
}
func gettask(e echo.Context) error {
	str, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, enums.Faileddecode)
	}
	res, err1, err2 := service.Gettask(str)
	if err1 != nil {
		return e.JSON(http.StatusAccepted, err1.Error())
	} else if err2 != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusInternalServerError, enums.ServerIssue)
	}
	return e.JSON(http.StatusOK, res)
}
func deletetask(e echo.Context) error {
	str, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		log.Error(err.Error())
		return e.JSON(http.StatusBadRequest, enums.Faileddecode)
	}
	err1, err2 := service.Deletetask(str)
	if err1 != nil && err2 == nil {
		return e.JSON(http.StatusInternalServerError, enums.ServerIssue)
	} else if err1 != nil && err2 != nil {
		return e.JSON(http.StatusInternalServerError, "no such key exist")
	}
	return e.JSON(http.StatusOK, enums.Deletesuccess)

}

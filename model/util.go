package model

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ResponseFormat format structure for response API
type ResponseFormat struct {
	Status   string            `json:"status,omitempty"`
	Data     interface{}       `json:"data,omitempty"`
	Total    int               `json:"total,omitempty"`
	Messages map[string]string `json:"messages,omitempty"`
}

// Serve used to response the data when request is success
func (m *ResponseFormat) Serve(ctx echo.Context, data interface{}, total int) error {
	m.Data = data
	m.Status = "success"
	m.Total = total
	if len(m.Messages) != int(0) {
		m.Status = "not success"
	}
	return ctx.JSON(http.StatusOK, m)
}

// Error used to response an error when request is fail
func (m *ResponseFormat) Error(ctx echo.Context, e error) (err error) {
	if e != nil {
		m.Status = e.Error()
	} else {
		m.Status = "invalid"
	}
	m.Total = 0
	err = ctx.JSON(http.StatusUnprocessableEntity, m)
	return err
}

// Message used to insert a message to response
func (m *ResponseFormat) Message(rs map[string]string) (f bool) {
	m.Messages = make(map[string]string)
	if len(rs) != int(0) {
		m.Messages = rs
	} else {
		f = true
	}
	return
}

// ConvertID for converting ID in parameter
func ConvertID(ID string) (number int64, err error) {
	if ID != "" {
		if number, err = strconv.ParseInt(ID, 10, 64); err != nil {
			err = errors.New("wrong format for id")
		}
	} else {
		err = errors.New("wrong format for id")
	}
	return
}

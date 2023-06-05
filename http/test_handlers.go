package http

import (
	"github.com/labstack/echo"
	"net/http"
	"test/testEntities"
)

type httpTestHandler struct {
	testLogic test_entities.TestLogic
}

func NewTestHandler(rg *echo.Group, testLogic test_entities.TestLogic) {
	h := httpTestHandler{
		testLogic: testLogic,
	}
	rg.POST("/test1", h.Test1Func)
	rg.POST("/test2", h.Test2Func)
}

func (h *httpTestHandler) Test1Func(c echo.Context) error {
	ctx := c.Request().Context()
	test1Struct := test_entities.Test1StructRequest{}
	err := c.Bind(&test1Struct)
	if err != nil {
		return err
	}
	respStruct, err := h.testLogic.Test1Func(ctx, test1Struct)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, respStruct)
}
func (h *httpTestHandler) Test2Func(c echo.Context) error {
	ctx := c.Request().Context()
	test2Struct := test_entities.Test2StructRequest{}
	err := c.Bind(&test2Struct)
	if err != nil {
		return err
	}
	respStruct, err := h.testLogic.Test2Func(ctx, test2Struct)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, respStruct)
}

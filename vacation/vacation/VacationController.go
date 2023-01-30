package vacation
import (
"net/http"
"strconv"
"gorm.io/gorm"
"github.com/labstack/echo"
)

func Cancel(c echo.Context) error  {
    repository := VacationRepository()
    id, _ := strconv.Atoi(c.Param("id"))
    params := make(map[string] string)

    c.Bind(&params)
    entity, err := repository.FindById(id)
    if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, err)
		} else {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
    // TODO 로직
}

import (
"net/http"
"strconv"
"gorm.io/gorm"
"github.com/labstack/echo"
)

func Approve(c echo.Context) error  {
    repository := VacationRepository()
    id, _ := strconv.Atoi(c.Param("id"))
    params := make(map[string] string)

    c.Bind(&params)
    entity, err := repository.FindById(id)
    if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, err)
		} else {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
    // TODO 로직
}

import (
"net/http"
"strconv"
"gorm.io/gorm"
"github.com/labstack/echo"
)

func ConfirmUsed(c echo.Context) error  {
    repository := VacationRepository()
    id, _ := strconv.Atoi(c.Param("id"))
    params := make(map[string] string)

    c.Bind(&params)
    entity, err := repository.FindById(id)
    if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, err)
		} else {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
    // TODO 로직
}

import (
"net/http"
"strconv"
"gorm.io/gorm"
"github.com/labstack/echo"
)

func Update(c echo.Context) error  {
    repository := VacationRepository()
    id, _ := strconv.Atoi(c.Param("id"))
    params := make(map[string] string)

    c.Bind(&params)
    entity, err := repository.FindById(id)
    if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, err)
		} else {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
    // TODO 로직
}


package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const EX1_ANSWER int = 2397483

type Endpoint struct {
	route    string
	endpoint gin.HandlerFunc
	method   string
}

func getDefaultEndpoints() []*Endpoint {
	return []*Endpoint{
		&Endpoint{
			route:    "/ex1",
			endpoint: ex1,
			method:   "POST",
		},
	}
}

type ex1Guess struct {
	Guess int `json:"guess"`
}

type ex1Response struct {
	Message int `json:"message"`
}

func ex1(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	var guess ex1Guess
	err = json.Unmarshal(bodyBytes, &guess)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}
	if guess.Guess > EX1_ANSWER {
		ctx.JSON(http.StatusOK, ex1Response{Message: 1})
		return
	}
	if guess.Guess < EX1_ANSWER {
		ctx.JSON(http.StatusOK, ex1Response{Message: -1})
		return
	}
	ctx.JSON(http.StatusOK, ex1Response{Message: 0})
}

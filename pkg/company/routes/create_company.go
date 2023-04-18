package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s-shaida/xm-go-exercise/pkg/company/pb"
)

type CreateCompanyRequestBody struct {
	Name              string `json:"name" binding:"required"`
	Description       string `json:"description"`
	AmountOfEmployees int64  `json:"amount" binding:"required"`
	Registered        string `json:"registered, bool" binding:"required"`
	Type              string `json:"type" binding:"required"`
}

func CreateCompany(ctx *gin.Context, c pb.CompanyServiceClient) {
	body := CreateCompanyRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateCompany(context.Background(), &pb.CreateCompanyRequest{
		Name:        body.Name,
		Description: body.Description,
		Amount:      body.AmountOfEmployees,
		Registered:  body.Registered,
		Type:        body.Type,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

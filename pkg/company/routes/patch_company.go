package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s-shaida/xm-go-exercise/pkg/company/pb"
)

type PatchCompanyRequestBody struct {
	Id                int64  `json:"id" binding:"required"`
	Name              string `json:"name" binding:"required"`
	Description       string `json:"description"`
	AmountOfEmployees int64  `json:"amount" binding:"required"`
	Registered        string `json:"registered" binding:"required"`
	Type              string `json:"type" binding:"required"`
}

func PatchCompany(ctx *gin.Context, c pb.CompanyServiceClient) {
	body := PatchCompanyRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.PatchCompany(context.Background(), &pb.PatchCompanyRequest{
		Id:          body.Id,
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

package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s-shaida/xm-go-exercise/pkg/company/pb"
)

var ()

type CreateCompanyRequestBody struct {
	Name              string `json:"name" binding:"required"`
	Description       string `json:"description"`
	AmountOfEmployees int64  `json:"amount" binding:"required"`
	Registered        *bool  `json:"registered" binding:"required"`
	Type              string `json:"type" binding:"required"`
}

func CreateCompany(ctx *gin.Context, c pb.CompanyServiceClient) {
	body := CreateCompanyRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	val, ok := pb.CompanyType_value[body.Type]

	if !ok {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var companyType = pb.CompanyType(val)

	res, err := c.CreateCompany(context.Background(), &pb.CreateCompanyRequest{
		Name:        body.Name,
		Description: body.Description,
		Amount:      body.AmountOfEmployees,
		Registered:  *body.Registered,
		Type:        companyType,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}

// curl --request POST \
//   --url http://localhost:3000/company \
//   --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzODkxNzEsImlzcyI6InhtLWdvLWF1dGgtc3ZjIiwiSWQiOjEsIkVtYWlsIjoiZWxvbkBtdXNrLmNvbSJ9.V0CS3AnX_fwyo9ld4Ibo8GVRVOI5Fh9x_-JNR3aq4cg' \
//   --header 'Content-Type: application/json' \
//   --data '{
//  "name": "Company 1",
//  "description": "Small description",
//  "amount": 15,
//  "registered": false,
//  "type": "NonProfit"
// }'

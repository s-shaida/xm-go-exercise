package company

import (
	"github.com/gin-gonic/gin"
	"github.com/s-shaida/xm-go-exercise/pkg/auth"
	"github.com/s-shaida/xm-go-exercise/pkg/company/routes"
	"github.com/s-shaida/xm-go-exercise/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/company")
	routes.GET("/:id", svc.GetOneCompany)
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateCompany)
}

func (svc *ServiceClient) GetOneCompany(ctx *gin.Context) {
	routes.GetOneCompany(ctx, svc.Client)
}

func (svc *ServiceClient) CreateCompany(ctx *gin.Context) {
	routes.CreateCompany(ctx, svc.Client)
}

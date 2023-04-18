package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/s-shaida/xm-go-exercise/pkg/auth"
	"github.com/s-shaida/xm-go-exercise/pkg/company"
	"github.com/s-shaida/xm-go-exercise/pkg/config"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	company.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}

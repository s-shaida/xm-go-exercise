package company

import (
	"fmt"

	"github.com/s-shaida/xm-go-exercise/pkg/company/pb"
	"github.com/s-shaida/xm-go-exercise/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CompanyServiceClient
}

func InitServiceClient(c *config.Config) pb.CompanyServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CompanySvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCompanyServiceClient(cc)
}

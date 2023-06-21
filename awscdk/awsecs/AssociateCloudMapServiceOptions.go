package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
)

// The options for using a cloudmap service.
//
// Example:
//   var cloudMapService service
//   var ecsService fargateService
//
//
//   ecsService.AssociateCloudMapService(&AssociateCloudMapServiceOptions{
//   	Service: cloudMapService,
//   })
//
type AssociateCloudMapServiceOptions struct {
	// The cloudmap service to register with.
	Service awsservicediscovery.IService `field:"required" json:"service" yaml:"service"`
	// The container to point to for a SRV record.
	Container ContainerDefinition `field:"optional" json:"container" yaml:"container"`
	// The port to point to for a SRV record.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
}


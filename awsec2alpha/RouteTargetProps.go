package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The type of endpoint or gateway being targeted by the route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var routeTarget iRouteTarget
//   var vpcEndpoint vpcEndpoint
//
//   routeTargetProps := &RouteTargetProps{
//   	Endpoint: vpcEndpoint,
//   	Gateway: routeTarget,
//   }
//
// Experimental.
type RouteTargetProps struct {
	// The endpoint route target.
	//
	// This is used for targets such as
	// VPC endpoints.
	// Default: none.
	//
	// Experimental.
	Endpoint awsec2.IVpcEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The gateway route target.
	//
	// This is used for targets such as
	// egress-only internet gateway or VPC peering connection.
	// Default: none.
	//
	// Experimental.
	Gateway IRouteTarget `field:"optional" json:"gateway" yaml:"gateway"`
}


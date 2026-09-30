package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for a NAT gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   natGatewayProps := &NatGatewayProps{
//   	EipAllocationIds: []*string{
//   		jsii.String("eipAllocationIds"),
//   	},
//   	MaxDrainDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type NatGatewayProps struct {
	// EIP allocation IDs for the NAT gateways.
	// Default: - No fixed EIPs allocated for the NAT gateways.
	//
	EipAllocationIds *[]*string `field:"optional" json:"eipAllocationIds" yaml:"eipAllocationIds"`
	// Maximum amount of time to wait before forcibly releasing IP addresses if connections are still in progress.
	// Default: Duration.seconds(350)
	//
	MaxDrainDuration awscdk.Duration `field:"optional" json:"maxDrainDuration" yaml:"maxDrainDuration"`
}


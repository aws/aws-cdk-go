package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for adding a NAT gateway to a public subnet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addNatGatewayOptions := &AddNatGatewayOptions{
//   	MaxDrainDuration: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type AddNatGatewayOptions struct {
	// Maximum amount of time to wait before forcibly releasing the IP addresses if connections are still in progress.
	// Default: Duration.seconds(350)
	//
	MaxDrainDuration awscdk.Duration `field:"optional" json:"maxDrainDuration" yaml:"maxDrainDuration"`
}


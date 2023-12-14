package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Attributes when importing a new VpcLink.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpc vpc
//
//   vpcLinkAttributes := &VpcLinkAttributes{
//   	Vpc: vpc,
//   	VpcLinkId: jsii.String("vpcLinkId"),
//   }
//
type VpcLinkAttributes struct {
	// The VPC to which this VPC link is associated with.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The VPC Link id.
	VpcLinkId *string `field:"required" json:"vpcLinkId" yaml:"vpcLinkId"`
}


package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVPNGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPNGatewayProps := &cfnVPNGatewayProps{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	amazonSideAsn: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVPNGatewayProps struct {
	// The type of VPN connection the virtual private gateway supports.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	AmazonSideAsn *float64 `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// Any tags assigned to the virtual private gateway.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCustomerGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomerGatewayProps := &cfnCustomerGatewayProps{
//   	bgpAsn: jsii.Number(123),
//   	ipAddress: jsii.String("ipAddress"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCustomerGatewayProps struct {
	// For devices that support BGP, the customer gateway's BGP ASN.
	//
	// Default: 65000.
	BgpAsn *float64 `field:"required" json:"bgpAsn" yaml:"bgpAsn"`
	// The Internet-routable IP address for the customer gateway's outside interface.
	//
	// The address must be static.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The type of VPN connection that this customer gateway supports ( `ipsec.1` ).
	Type *string `field:"required" json:"type" yaml:"type"`
	// One or more tags for the customer gateway.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCustomerGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomerGatewayProps := &CfnCustomerGatewayProps{
//   	BgpAsn: jsii.Number(123),
//   	IpAddress: jsii.String("ipAddress"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	DeviceName: jsii.String("deviceName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html
//
type CfnCustomerGatewayProps struct {
	// For devices that support BGP, the customer gateway's BGP ASN.
	//
	// Default: 65000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-bgpasn
	//
	// Default: - 65000.
	//
	BgpAsn *float64 `field:"required" json:"bgpAsn" yaml:"bgpAsn"`
	// IPv4 address for the customer gateway device's outside interface.
	//
	// The address must be static.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-ipaddress
	//
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The type of VPN connection that this customer gateway supports ( `ipsec.1` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// The name of customer gateway device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-devicename
	//
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// One or more tags for the customer gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


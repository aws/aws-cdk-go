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
//   	IpAddress: jsii.String("ipAddress"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	BgpAsn: jsii.Number(123),
//   	BgpAsnExtended: jsii.Number(123),
//   	CertificateArn: jsii.String("certificateArn"),
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
	// IPv4 address for the customer gateway device's outside interface.
	//
	// The address must be static. If `OutsideIpAddressType` in your VPN connection options is set to `PrivateIpv4` , you can use an RFC6598 or RFC1918 private IPv4 address. If `OutsideIpAddressType` is set to `PublicIpv4` , you can use a public IPv4 address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-ipaddress
	//
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The type of VPN connection that this customer gateway supports ( `ipsec.1` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// For customer gateway devices that support BGP, specify the device's ASN.
	//
	// You must specify either `BgpAsn` or `BgpAsnExtended` when creating the customer gateway. If the ASN is larger than `2,147,483,647` , you must use `BgpAsnExtended` .
	//
	// Default: 65000
	//
	// Valid values: `1` to `2,147,483,647`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-bgpasn
	//
	// Default: - 65000.
	//
	BgpAsn *float64 `field:"optional" json:"bgpAsn" yaml:"bgpAsn"`
	// For customer gateway devices that support BGP, specify the device's ASN.
	//
	// You must specify either `BgpAsn` or `BgpAsnExtended` when creating the customer gateway. If the ASN is larger than `2,147,483,647` , you must use `BgpAsnExtended` .
	//
	// Valid values: `2,147,483,648` to `4,294,967,295`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-bgpasnextended
	//
	BgpAsnExtended *float64 `field:"optional" json:"bgpAsnExtended" yaml:"bgpAsnExtended"`
	// The Amazon Resource Name (ARN) for the customer gateway certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-certificatearn
	//
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The name of customer gateway device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-devicename
	//
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// One or more tags for the customer gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-customergateway.html#cfn-ec2-customergateway-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


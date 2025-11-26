package previewawsvpclatticemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnResourceGatewayPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceGatewayMixinProps := &CfnResourceGatewayMixinProps{
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Ipv4AddressesPerEni: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcIdentifier: jsii.String("vpcIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourcegateway.html
//
type CfnResourceGatewayMixinProps struct {
	// The type of IP address used by the resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourcegateway.html#cfn-vpclattice-resourcegateway-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The number of IPv4 addresses in each ENI for the resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourcegateway.html#cfn-vpclattice-resourcegateway-ipv4addressespereni
	//
	Ipv4AddressesPerEni *float64 `field:"optional" json:"ipv4AddressesPerEni" yaml:"ipv4AddressesPerEni"`
	// The name of the resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourcegateway.html#cfn-vpclattice-resourcegateway-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IDs of the security groups applied to the resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourcegateway.html#cfn-vpclattice-resourcegateway-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The IDs of the VPC subnets for the resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourcegateway.html#cfn-vpclattice-resourcegateway-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The tags for the resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourcegateway.html#cfn-vpclattice-resourcegateway-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC for the resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourcegateway.html#cfn-vpclattice-resourcegateway-vpcidentifier
	//
	VpcIdentifier *string `field:"optional" json:"vpcIdentifier" yaml:"vpcIdentifier"`
}


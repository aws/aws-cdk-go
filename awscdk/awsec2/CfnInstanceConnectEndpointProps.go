package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInstanceConnectEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInstanceConnectEndpointProps := &CfnInstanceConnectEndpointProps{
//   	SubnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	ClientToken: jsii.String("clientToken"),
//   	PreserveClientIp: jsii.Boolean(false),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceconnectendpoint.html
//
type CfnInstanceConnectEndpointProps struct {
	// The ID of the subnet in which to create the EC2 Instance Connect Endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceconnectendpoint.html#cfn-ec2-instanceconnectendpoint-subnetid
	//
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceconnectendpoint.html#cfn-ec2-instanceconnectendpoint-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Indicates whether the client IP address is preserved as the source. The following are the possible values.
	//
	// - `true` - Use the client IP address as the source.
	// - `false` - Use the network interface IP address as the source.
	//
	// > `PreserveClientIp` is only supported on IPv4 EC2 Instance Connect Endpoints. To use `PreserveClientIp` , the value for `IpAddressType` must be `ipv4` .
	//
	// Default: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceconnectendpoint.html#cfn-ec2-instanceconnectendpoint-preserveclientip
	//
	PreserveClientIp interface{} `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// One or more security groups to associate with the endpoint.
	//
	// If you don't specify a security group, the default security group for your VPC will be associated with the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceconnectendpoint.html#cfn-ec2-instanceconnectendpoint-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The tags to apply to the EC2 Instance Connect Endpoint during creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-instanceconnectendpoint.html#cfn-ec2-instanceconnectendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


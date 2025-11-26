package previewawsapprunnermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVpcConnectorPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVpcConnectorMixinProps := &CfnVpcConnectorMixinProps{
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcConnectorName: jsii.String("vpcConnectorName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-vpcconnector.html
//
type CfnVpcConnectorMixinProps struct {
	// A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets.
	//
	// If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-vpcconnector.html#cfn-apprunner-vpcconnector-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC.
	//
	// Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	//
	// > App Runner only supports subnets of IP address type *IPv4* and *dual stack* (IPv4 and IPv6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-vpcconnector.html#cfn-apprunner-vpcconnector-subnets
	//
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// A list of metadata items that you can associate with your VPC connector resource.
	//
	// A tag is a key-value pair.
	//
	// > A `VpcConnector` is immutable, so you cannot update its tags. To change the tags, replace the resource. To replace a `VpcConnector` , you must provide a new combination of security groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-vpcconnector.html#cfn-apprunner-vpcconnector-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A name for the VPC connector.
	//
	// If you don't specify a name, CloudFormation generates a name for your VPC connector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apprunner-vpcconnector.html#cfn-apprunner-vpcconnector-vpcconnectorname
	//
	VpcConnectorName *string `field:"optional" json:"vpcConnectorName" yaml:"vpcConnectorName"`
}


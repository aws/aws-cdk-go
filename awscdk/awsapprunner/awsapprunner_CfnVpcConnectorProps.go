package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVpcConnector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcConnectorProps := &cfnVpcConnectorProps{
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//
//   	// the properties below are optional
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcConnectorName: jsii.String("vpcConnectorName"),
//   }
//
type CfnVpcConnectorProps struct {
	// A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC.
	//
	// Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.
	//
	// > App Runner currently only provides support for IPv4.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets.
	//
	// If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// A list of metadata items that you can associate with your VPC connector resource.
	//
	// A tag is a key-value pair.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A name for the VPC connector.
	//
	// If you don't specify a name, AWS CloudFormation generates a name for your VPC connector.
	VpcConnectorName *string `field:"optional" json:"vpcConnectorName" yaml:"vpcConnectorName"`
}


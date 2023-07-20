package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVpcLink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcLinkProps := &CfnVpcLinkProps{
//   	Name: jsii.String("name"),
//   	TargetArns: []*string{
//   		jsii.String("targetArns"),
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html
//
type CfnVpcLinkProps struct {
	// The name used to label and identify the VPC link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html#cfn-apigateway-vpclink-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the network load balancer of the VPC targeted by the VPC link.
	//
	// The network load balancer must be owned by the same AWS account of the API owner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html#cfn-apigateway-vpclink-targetarns
	//
	TargetArns *[]*string `field:"required" json:"targetArns" yaml:"targetArns"`
	// The description of the VPC link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html#cfn-apigateway-vpclink-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of arbitrary tags (key-value pairs) to associate with the VPC link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-vpclink.html#cfn-apigateway-vpclink-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


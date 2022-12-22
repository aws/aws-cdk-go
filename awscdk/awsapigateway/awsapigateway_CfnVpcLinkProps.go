package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVpcLink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVpcLinkProps := &cfnVpcLinkProps{
//   	name: jsii.String("name"),
//   	targetArns: []*string{
//   		jsii.String("targetArns"),
//   	},
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVpcLinkProps struct {
	// A name for the VPC link.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of network load balancer of the VPC targeted by the VPC link.
	//
	// The network load balancer must be owned by the same AWS account of the API owner.
	TargetArns *[]*string `field:"required" json:"targetArns" yaml:"targetArns"`
	// A description of the VPC link.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of arbitrary tags (key-value pairs) to associate with the VPC link.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


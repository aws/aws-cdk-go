package awsdevicefarm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTestGridProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTestGridProjectProps := &cfnTestGridProjectProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		vpcId: jsii.String("vpcId"),
//   	},
//   }
//
type CfnTestGridProjectProps struct {
	// A human-readable name for the project.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A human-readable description for the project.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The VPC security groups and subnets that are attached to a project.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}


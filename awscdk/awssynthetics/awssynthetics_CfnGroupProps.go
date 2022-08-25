package awssynthetics

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGroupProps := &cfnGroupProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	resourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnGroupProps struct {
	// `AWS::Synthetics::Group.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::Synthetics::Group.ResourceArns`.
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// `AWS::Synthetics::Group.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


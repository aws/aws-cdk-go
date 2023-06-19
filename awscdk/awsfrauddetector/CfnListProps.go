package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnList`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnListProps := &CfnListProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Elements: []*string{
//   		jsii.String("elements"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VariableType: jsii.String("variableType"),
//   }
//
type CfnListProps struct {
	// The name of the list.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the list.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The elements in the list.
	Elements *[]*string `field:"optional" json:"elements" yaml:"elements"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The variable type of the list.
	//
	// For more information, see [Variable types](https://docs.aws.amazon.com/frauddetector/latest/ug/variables.html#variable-types)
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}


package awsiotevents

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnInput`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInputProps := &cfnInputProps{
//   	inputDefinition: &inputDefinitionProperty{
//   		attributes: []interface{}{
//   			&attributeProperty{
//   				jsonPath: jsii.String("jsonPath"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	inputDescription: jsii.String("inputDescription"),
//   	inputName: jsii.String("inputName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnInputProps struct {
	// The definition of the input.
	InputDefinition interface{} `field:"required" json:"inputDefinition" yaml:"inputDefinition"`
	// A brief description of the input.
	InputDescription *string `field:"optional" json:"inputDescription" yaml:"inputDescription"`
	// The name of the input.
	InputName *string `field:"optional" json:"inputName" yaml:"inputName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


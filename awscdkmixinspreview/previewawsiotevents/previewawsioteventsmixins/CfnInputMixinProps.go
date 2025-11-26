package previewawsioteventsmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnInputPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInputMixinProps := &CfnInputMixinProps{
//   	InputDefinition: &InputDefinitionProperty{
//   		Attributes: []interface{}{
//   			&AttributeProperty{
//   				JsonPath: jsii.String("jsonPath"),
//   			},
//   		},
//   	},
//   	InputDescription: jsii.String("inputDescription"),
//   	InputName: jsii.String("inputName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html
//
type CfnInputMixinProps struct {
	// The definition of the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdefinition
	//
	InputDefinition interface{} `field:"optional" json:"inputDefinition" yaml:"inputDefinition"`
	// A brief description of the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdescription
	//
	InputDescription *string `field:"optional" json:"inputDescription" yaml:"inputDescription"`
	// The name of the input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputname
	//
	InputName *string `field:"optional" json:"inputName" yaml:"inputName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


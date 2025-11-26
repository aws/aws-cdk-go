package previewawsfrauddetectormixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnListPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnListMixinProps := &CfnListMixinProps{
//   	Description: jsii.String("description"),
//   	Elements: []*string{
//   		jsii.String("elements"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VariableType: jsii.String("variableType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-list.html
//
type CfnListMixinProps struct {
	// The description of the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-list.html#cfn-frauddetector-list-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The elements in the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-list.html#cfn-frauddetector-list-elements
	//
	Elements *[]*string `field:"optional" json:"elements" yaml:"elements"`
	// The name of the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-list.html#cfn-frauddetector-list-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-list.html#cfn-frauddetector-list-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The variable type of the list.
	//
	// For more information, see [Variable types](https://docs.aws.amazon.com/frauddetector/latest/ug/variables.html#variable-types)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-frauddetector-list.html#cfn-frauddetector-list-variabletype
	//
	VariableType *string `field:"optional" json:"variableType" yaml:"variableType"`
}


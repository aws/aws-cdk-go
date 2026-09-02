package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnHumanTaskUiPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnHumanTaskUiMixinProps := &CfnHumanTaskUiMixinProps{
//   	HumanTaskUiName: jsii.String("humanTaskUiName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UiTemplate: &UiTemplateProperty{
//   		Content: jsii.String("content"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-humantaskui.html
//
type CfnHumanTaskUiMixinProps struct {
	// The name of the human task user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-humantaskui.html#cfn-sagemaker-humantaskui-humantaskuiname
	//
	HumanTaskUiName *string `field:"optional" json:"humanTaskUiName" yaml:"humanTaskUiName"`
	// An array of key-value pairs that contain metadata to help you categorize and organize a human review workflow user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-humantaskui.html#cfn-sagemaker-humantaskui-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Liquid template for the worker user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-humantaskui.html#cfn-sagemaker-humantaskui-uitemplate
	//
	UiTemplate interface{} `field:"optional" json:"uiTemplate" yaml:"uiTemplate"`
}


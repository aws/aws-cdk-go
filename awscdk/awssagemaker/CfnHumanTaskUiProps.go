package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnHumanTaskUi`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHumanTaskUiProps := &CfnHumanTaskUiProps{
//   	HumanTaskUiName: jsii.String("humanTaskUiName"),
//
//   	// the properties below are optional
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
type CfnHumanTaskUiProps struct {
	// The name of the human task user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-humantaskui.html#cfn-sagemaker-humantaskui-humantaskuiname
	//
	HumanTaskUiName *string `field:"required" json:"humanTaskUiName" yaml:"humanTaskUiName"`
	// An array of key-value pairs that contain metadata to help you categorize and organize a human review workflow user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-humantaskui.html#cfn-sagemaker-humantaskui-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The Liquid template for the worker user interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-humantaskui.html#cfn-sagemaker-humantaskui-uitemplate
	//
	UiTemplate interface{} `field:"optional" json:"uiTemplate" yaml:"uiTemplate"`
}


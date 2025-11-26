package previewawssagemakermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnStudioLifecycleConfigPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnStudioLifecycleConfigMixinProps := &CfnStudioLifecycleConfigMixinProps{
//   	StudioLifecycleConfigAppType: jsii.String("studioLifecycleConfigAppType"),
//   	StudioLifecycleConfigContent: jsii.String("studioLifecycleConfigContent"),
//   	StudioLifecycleConfigName: jsii.String("studioLifecycleConfigName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-studiolifecycleconfig.html
//
type CfnStudioLifecycleConfigMixinProps struct {
	// The App type to which the Lifecycle Configuration is attached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-studiolifecycleconfig.html#cfn-sagemaker-studiolifecycleconfig-studiolifecycleconfigapptype
	//
	StudioLifecycleConfigAppType *string `field:"optional" json:"studioLifecycleConfigAppType" yaml:"studioLifecycleConfigAppType"`
	// The content of your Amazon SageMaker Studio Lifecycle Configuration script.
	//
	// This content must be base64 encoded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-studiolifecycleconfig.html#cfn-sagemaker-studiolifecycleconfig-studiolifecycleconfigcontent
	//
	StudioLifecycleConfigContent *string `field:"optional" json:"studioLifecycleConfigContent" yaml:"studioLifecycleConfigContent"`
	// The name of the Amazon SageMaker AI Studio Lifecycle Configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-studiolifecycleconfig.html#cfn-sagemaker-studiolifecycleconfig-studiolifecycleconfigname
	//
	StudioLifecycleConfigName *string `field:"optional" json:"studioLifecycleConfigName" yaml:"studioLifecycleConfigName"`
	// Tags to be associated with the Lifecycle Configuration.
	//
	// Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-studiolifecycleconfig.html#cfn-sagemaker-studiolifecycleconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


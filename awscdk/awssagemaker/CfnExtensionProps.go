package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnExtension`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnExtensionProps := &CfnExtensionProps{
//   	ClusterArn: jsii.String("clusterArn"),
//   	ExtensionName: jsii.String("extensionName"),
//
//   	// the properties below are optional
//   	ConfigurationValues: jsii.String("configurationValues"),
//   	ExtensionVersion: jsii.String("extensionVersion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-extension.html
//
type CfnExtensionProps struct {
	// The Amazon Resource Name (ARN) of the Amazon SageMaker HyperPod cluster to install the extension on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-extension.html#cfn-sagemaker-extension-clusterarn
	//
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// The name of the extension to install, such as sagemaker-observability.
	//
	// You can't install the same extension twice on a cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-extension.html#cfn-sagemaker-extension-extensionname
	//
	ExtensionName *string `field:"required" json:"extensionName" yaml:"extensionName"`
	// The configuration settings for the extension, as a JSON string.
	//
	// The settings must match the configuration schema for the extension version. For each setting that you don't specify, SageMaker HyperPod uses its default value. When you update the stack, SageMaker HyperPod applies the settings in the string to the current configuration. Settings that you leave out keep their current values. To restore the default value of a setting, set it to null. To write the settings as YAML or JSON in your template, use the Fn::ToJsonString intrinsic function with the AWS::LanguageExtensions transform. The JSON string can be up to 10,240 characters long.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-extension.html#cfn-sagemaker-extension-configurationvalues
	//
	ConfigurationValues *string `field:"optional" json:"configurationValues" yaml:"configurationValues"`
	// The version of the extension to install, such as 1.0.0. When you create the extension without a version, SageMaker HyperPod installs the latest generally available version. If the extension doesn't have a generally available version, you must specify one. If you remove this property from your template during a stack update, the installed version doesn't change. You can't downgrade an extension. To install an earlier version, delete the extension and create it again.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-extension.html#cfn-sagemaker-extension-extensionversion
	//
	ExtensionVersion *string `field:"optional" json:"extensionVersion" yaml:"extensionVersion"`
	// The tags to add to the extension.
	//
	// You can add up to 50 tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-extension.html#cfn-sagemaker-extension-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


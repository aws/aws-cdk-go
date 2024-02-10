package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnAppImageConfig`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAppImageConfigProps := &CfnAppImageConfigProps{
//   	AppImageConfigName: jsii.String("appImageConfigName"),
//
//   	// the properties below are optional
//   	JupyterLabAppImageConfig: &JupyterLabAppImageConfigProperty{
//   		ContainerConfig: &ContainerConfigProperty{
//   			ContainerArguments: []*string{
//   				jsii.String("containerArguments"),
//   			},
//   			ContainerEntrypoint: []*string{
//   				jsii.String("containerEntrypoint"),
//   			},
//   			ContainerEnvironmentVariables: []interface{}{
//   				&CustomImageContainerEnvironmentVariableProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	KernelGatewayImageConfig: &KernelGatewayImageConfigProperty{
//   		KernelSpecs: []interface{}{
//   			&KernelSpecProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				DisplayName: jsii.String("displayName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		FileSystemConfig: &FileSystemConfigProperty{
//   			DefaultGid: jsii.Number(123),
//   			DefaultUid: jsii.Number(123),
//   			MountPath: jsii.String("mountPath"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-appimageconfig.html
//
type CfnAppImageConfigProps struct {
	// The name of the AppImageConfig.
	//
	// Must be unique to your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-appimageconfig.html#cfn-sagemaker-appimageconfig-appimageconfigname
	//
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// The configuration for the file system and the runtime, such as the environment variables and entry point.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-appimageconfig.html#cfn-sagemaker-appimageconfig-jupyterlabappimageconfig
	//
	JupyterLabAppImageConfig interface{} `field:"optional" json:"jupyterLabAppImageConfig" yaml:"jupyterLabAppImageConfig"`
	// The configuration for the file system and kernels in the SageMaker image.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-appimageconfig.html#cfn-sagemaker-appimageconfig-kernelgatewayimageconfig
	//
	KernelGatewayImageConfig interface{} `field:"optional" json:"kernelGatewayImageConfig" yaml:"kernelGatewayImageConfig"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-appimageconfig.html#cfn-sagemaker-appimageconfig-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


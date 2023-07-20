package awssagemaker


// A collection of space settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spaceSettingsProperty := &SpaceSettingsProperty{
//   	JupyterServerAppSettings: &JupyterServerAppSettingsProperty{
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   	KernelGatewayAppSettings: &KernelGatewayAppSettingsProperty{
//   		CustomImages: []interface{}{
//   			&CustomImageProperty{
//   				AppImageConfigName: jsii.String("appImageConfigName"),
//   				ImageName: jsii.String("imageName"),
//
//   				// the properties below are optional
//   				ImageVersionNumber: jsii.Number(123),
//   			},
//   		},
//   		DefaultResourceSpec: &ResourceSpecProperty{
//   			InstanceType: jsii.String("instanceType"),
//   			SageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			SageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html
//
type CfnSpace_SpaceSettingsProperty struct {
	// The JupyterServer app settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-jupyterserverappsettings
	//
	JupyterServerAppSettings interface{} `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The KernelGateway app settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-spacesettings.html#cfn-sagemaker-space-spacesettings-kernelgatewayappsettings
	//
	KernelGatewayAppSettings interface{} `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
}


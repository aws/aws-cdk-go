package awssagemaker


// A collection of space settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spaceSettingsProperty := &spaceSettingsProperty{
//   	jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   	kernelGatewayAppSettings: &kernelGatewayAppSettingsProperty{
//   		customImages: []interface{}{
//   			&customImageProperty{
//   				appImageConfigName: jsii.String("appImageConfigName"),
//   				imageName: jsii.String("imageName"),
//
//   				// the properties below are optional
//   				imageVersionNumber: jsii.Number(123),
//   			},
//   		},
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   }
//
type CfnSpace_SpaceSettingsProperty struct {
	// The JupyterServer app settings.
	JupyterServerAppSettings interface{} `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The KernelGateway app settings.
	KernelGatewayAppSettings interface{} `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
}


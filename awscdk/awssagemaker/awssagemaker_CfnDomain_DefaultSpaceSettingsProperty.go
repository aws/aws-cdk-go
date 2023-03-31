package awssagemaker


// A collection of settings that apply to spaces created in the Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   defaultSpaceSettingsProperty := &defaultSpaceSettingsProperty{
//   	executionRole: jsii.String("executionRole"),
//   	jupyterServerAppSettings: &jupyterServerAppSettingsProperty{
//   		defaultResourceSpec: &resourceSpecProperty{
//   			instanceType: jsii.String("instanceType"),
//   			lifecycleConfigArn: jsii.String("lifecycleConfigArn"),
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
//   			lifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   			sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   			sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   		},
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   }
//
type CfnDomain_DefaultSpaceSettingsProperty struct {
	// The execution role for the space.
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The JupyterServer app settings.
	JupyterServerAppSettings interface{} `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The KernelGateway app settings.
	KernelGatewayAppSettings interface{} `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// The security groups for the Amazon Virtual Private Cloud that the space uses for communication.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}


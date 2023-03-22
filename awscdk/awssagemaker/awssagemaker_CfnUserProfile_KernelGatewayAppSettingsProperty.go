package awssagemaker


// The KernelGateway app settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kernelGatewayAppSettingsProperty := &kernelGatewayAppSettingsProperty{
//   	customImages: []interface{}{
//   		&customImageProperty{
//   			appImageConfigName: jsii.String("appImageConfigName"),
//   			imageName: jsii.String("imageName"),
//
//   			// the properties below are optional
//   			imageVersionNumber: jsii.Number(123),
//   		},
//   	},
//   	defaultResourceSpec: &resourceSpecProperty{
//   		instanceType: jsii.String("instanceType"),
//   		sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   		sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   	},
//   }
//
type CfnUserProfile_KernelGatewayAppSettingsProperty struct {
	// A list of custom SageMaker images that are configured to run as a KernelGateway app.
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.
	//
	// > The Amazon SageMaker Studio UI does not use the default instance type value set here. The default instance type set here is used when Apps are created using the AWS Command Line Interface or AWS CloudFormation and the instance type parameter value is not passed.
	DefaultResourceSpec interface{} `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}


package awssagemaker


// Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceSpecProperty := &resourceSpecProperty{
//   	instanceType: jsii.String("instanceType"),
//   	lifecycleConfigArn: jsii.String("lifecycleConfigArn"),
//   	sageMakerImageArn: jsii.String("sageMakerImageArn"),
//   	sageMakerImageVersionArn: jsii.String("sageMakerImageVersionArn"),
//   }
//
type CfnDomain_ResourceSpecProperty struct {
	// The instance type that the image version runs on.
	//
	// > JupyterServer Apps only support the `system` value. KernelGateway Apps do not support the `system` value, but support all other values for available instance types.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// `CfnDomain.ResourceSpecProperty.LifecycleConfigArn`.
	LifecycleConfigArn *string `field:"optional" json:"lifecycleConfigArn" yaml:"lifecycleConfigArn"`
	// The ARN of the SageMaker image that the image version belongs to.
	SageMakerImageArn *string `field:"optional" json:"sageMakerImageArn" yaml:"sageMakerImageArn"`
	// The ARN of the image version created on the instance.
	SageMakerImageVersionArn *string `field:"optional" json:"sageMakerImageVersionArn" yaml:"sageMakerImageVersionArn"`
}


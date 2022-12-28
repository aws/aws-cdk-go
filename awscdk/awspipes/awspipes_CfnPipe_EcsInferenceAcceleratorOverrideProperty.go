package awspipes


// Details on an Elastic Inference accelerator task override.
//
// This parameter is used to override the Elastic Inference accelerator specified in the task definition. For more information, see [Working with Amazon Elastic Inference on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/userguide/ecs-inference.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ecsInferenceAcceleratorOverrideProperty := &ecsInferenceAcceleratorOverrideProperty{
//   	deviceName: jsii.String("deviceName"),
//   	deviceType: jsii.String("deviceType"),
//   }
//
type CfnPipe_EcsInferenceAcceleratorOverrideProperty struct {
	// The Elastic Inference accelerator device name to override for the task.
	//
	// This parameter must match a `deviceName` specified in the task definition.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// The Elastic Inference accelerator type to use.
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}


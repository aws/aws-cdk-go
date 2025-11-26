package previewawspipesmixins


// The type and amount of a resource to assign to a container.
//
// The supported resource types are GPUs and Elastic Inference accelerators. For more information, see [Working with GPUs on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-gpu.html) or [Working with Amazon Elastic Inference on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-inference.html) in the *Amazon Elastic Container Service Developer Guide*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ecsResourceRequirementProperty := &EcsResourceRequirementProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecsresourcerequirement.html
//
type CfnPipePropsMixin_EcsResourceRequirementProperty struct {
	// The type of resource to assign to a container.
	//
	// The supported values are `GPU` or `InferenceAccelerator` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecsresourcerequirement.html#cfn-pipes-pipe-ecsresourcerequirement-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value for the specified resource type.
	//
	// If the `GPU` type is used, the value is the number of physical `GPUs` the Amazon ECS container agent reserves for the container. The number of GPUs that's reserved for all containers in a task can't exceed the number of available GPUs on the container instance that the task is launched on.
	//
	// If the `InferenceAccelerator` type is used, the `value` matches the `deviceName` for an InferenceAccelerator specified in a task definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-ecsresourcerequirement.html#cfn-pipes-pipe-ecsresourcerequirement-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


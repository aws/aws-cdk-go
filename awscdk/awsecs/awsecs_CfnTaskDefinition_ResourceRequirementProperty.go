package awsecs


// The `ResourceRequirement` property specifies the type and amount of a resource to assign to a container.
//
// The only supported resource is a GPU. For more information, see [Working with GPUs on Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-gpu.html) in the *Amazon Elastic Container Service Developer Guide*
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceRequirementProperty := &resourceRequirementProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnTaskDefinition_ResourceRequirementProperty struct {
	// The type of resource to assign to a container.
	//
	// The supported values are `GPU` or `InferenceAccelerator` .
	Type *string `field:"required" json:"type" yaml:"type"`
	// The value for the specified resource type.
	//
	// If the `GPU` type is used, the value is the number of physical `GPUs` the Amazon ECS container agent will reserve for the container. The number of GPUs reserved for all containers in a task should not exceed the number of available GPUs on the container instance the task is launched on.
	//
	// If the `InferenceAccelerator` type is used, the `value` should match the `DeviceName` for an [InferenceAccelerator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-inferenceaccelerator.html) specified in a task definition.
	Value *string `field:"required" json:"value" yaml:"value"`
}


package awsecs


// Information about the platform for the Amazon ECS service or task.
//
// For more information about `RuntimePlatform` , see [RuntimePlatform](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#runtime-platform) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimePlatformProperty := &RuntimePlatformProperty{
//   	CpuArchitecture: jsii.String("cpuArchitecture"),
//   	OperatingSystemFamily: jsii.String("operatingSystemFamily"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-runtimeplatform.html
//
type CfnTaskDefinition_RuntimePlatformProperty struct {
	// The CPU architecture.
	//
	// You can run your Linux tasks on an ARM-based platform by setting the value to `ARM64` . This option is available for tasks that run on Linux Amazon EC2 instance or Linux containers on Fargate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-runtimeplatform.html#cfn-ecs-taskdefinition-runtimeplatform-cpuarchitecture
	//
	CpuArchitecture *string `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// The operating system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-runtimeplatform.html#cfn-ecs-taskdefinition-runtimeplatform-operatingsystemfamily
	//
	OperatingSystemFamily *string `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}


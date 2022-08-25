package awsecs


// Information about the platform for the Amazon ECS service or task.
//
// For more informataion about `RuntimePlatform` , see [RuntimePlatform](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#runtime-platform) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimePlatformProperty := &runtimePlatformProperty{
//   	cpuArchitecture: jsii.String("cpuArchitecture"),
//   	operatingSystemFamily: jsii.String("operatingSystemFamily"),
//   }
//
type CfnTaskDefinition_RuntimePlatformProperty struct {
	// The CPU architecture.
	//
	// You can run your Linux tasks on an ARM-based platform by setting the value to `ARM64` . This option is avaiable for tasks that run on Linux Amazon EC2 instance or Linux containers on Fargate.
	CpuArchitecture *string `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// The operating system.
	OperatingSystemFamily *string `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}


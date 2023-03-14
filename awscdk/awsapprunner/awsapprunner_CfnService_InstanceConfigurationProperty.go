package awsapprunner


// Describes the runtime configuration of an AWS App Runner service instance (scaling unit).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceConfigurationProperty := &instanceConfigurationProperty{
//   	cpu: jsii.String("cpu"),
//   	instanceRoleArn: jsii.String("instanceRoleArn"),
//   	memory: jsii.String("memory"),
//   }
//
type CfnService_InstanceConfigurationProperty struct {
	// The number of CPU units reserved for each instance of your App Runner service.
	//
	// Default: `1 vCPU`.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// The Amazon Resource Name (ARN) of an IAM role that provides permissions to your App Runner service.
	//
	// These are permissions that your code needs when it calls any AWS APIs.
	InstanceRoleArn *string `field:"optional" json:"instanceRoleArn" yaml:"instanceRoleArn"`
	// The amount of memory, in MB or GB, reserved for each instance of your App Runner service.
	//
	// Default: `2 GB`.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}


package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// A reference to an existing task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   taskDefinitionAttributes := &TaskDefinitionAttributes{
//   	TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	Compatibility: awscdk.Aws_ecs.Compatibility_EC2,
//   	ExecutionRole: role,
//   	NetworkMode: awscdk.*Aws_ecs.NetworkMode_NONE,
//   	TaskRole: role,
//   }
//
type TaskDefinitionAttributes struct {
	// The arn of the task definition.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The IAM role that grants containers and Fargate agents permission to make AWS API calls on your behalf.
	//
	// Some tasks do not have an execution role.
	// Default: - undefined.
	//
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The networking mode to use for the containers in the task.
	// Default: Network mode cannot be provided to the imported task.
	//
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Default: Permissions cannot be granted to the imported task.
	//
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
	// What launch types this task definition should be compatible with.
	// Default: Compatibility.EC2_AND_FARGATE
	//
	Compatibility Compatibility `field:"optional" json:"compatibility" yaml:"compatibility"`
}


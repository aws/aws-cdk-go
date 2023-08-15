package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The common task definition attributes used across all types of task definitions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   commonTaskDefinitionAttributes := &CommonTaskDefinitionAttributes{
//   	TaskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	ExecutionRole: role,
//   	NetworkMode: awscdk.Aws_ecs.NetworkMode_NONE,
//   	TaskRole: role,
//   }
//
type CommonTaskDefinitionAttributes struct {
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
}


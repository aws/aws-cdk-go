package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Attributes used to import an existing External task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   externalTaskDefinitionAttributes := &externalTaskDefinitionAttributes{
//   	taskDefinitionArn: jsii.String("taskDefinitionArn"),
//
//   	// the properties below are optional
//   	networkMode: awscdk.Aws_ecs.networkMode_NONE,
//   	taskRole: role,
//   }
//
// Experimental.
type ExternalTaskDefinitionAttributes struct {
	// The arn of the task definition.
	// Experimental.
	TaskDefinitionArn *string `field:"required" json:"taskDefinitionArn" yaml:"taskDefinitionArn"`
	// The networking mode to use for the containers in the task.
	// Experimental.
	NetworkMode NetworkMode `field:"optional" json:"networkMode" yaml:"networkMode"`
	// The name of the IAM role that grants containers in the task permission to call AWS APIs on your behalf.
	// Experimental.
	TaskRole awsiam.IRole `field:"optional" json:"taskRole" yaml:"taskRole"`
}


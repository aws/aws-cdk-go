package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// The properties for the ScheduledEc2Task using a task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var ec2TaskDefinition Ec2TaskDefinition
//
//   scheduledEc2TaskDefinitionOptions := &ScheduledEc2TaskDefinitionOptions{
//   	TaskDefinition: ec2TaskDefinition,
//   }
//
type ScheduledEc2TaskDefinitionOptions struct {
	// The task definition to use for tasks in the service. One of image or taskDefinition must be specified.
	//
	// [disable-awslint:ref-via-interface].
	// Default: - none.
	//
	TaskDefinition awsecs.Ec2TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
}


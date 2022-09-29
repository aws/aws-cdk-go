package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
)

// The properties for the ScheduledEc2Task using a task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var ec2TaskDefinition ec2TaskDefinition
//
//   scheduledEc2TaskDefinitionOptions := &scheduledEc2TaskDefinitionOptions{
//   	taskDefinition: ec2TaskDefinition,
//   }
//
// Experimental.
type ScheduledEc2TaskDefinitionOptions struct {
	// The task definition to use for tasks in the service. One of image or taskDefinition must be specified.
	//
	// [disable-awslint:ref-via-interface].
	// Experimental.
	TaskDefinition awsecs.Ec2TaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
}


package awsecspatterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// The properties for the ScheduledFargateTask using a task definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var fargateTaskDefinition fargateTaskDefinition
//
//   scheduledFargateTaskDefinitionOptions := &scheduledFargateTaskDefinitionOptions{
//   	taskDefinition: fargateTaskDefinition,
//   }
//
type ScheduledFargateTaskDefinitionOptions struct {
	// The task definition to use for tasks in the service. Image or taskDefinition must be specified, but not both.
	//
	// [disable-awslint:ref-via-interface].
	TaskDefinition awsecs.FargateTaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
}


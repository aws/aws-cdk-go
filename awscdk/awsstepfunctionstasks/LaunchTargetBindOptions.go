package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Options for binding a launch target to an ECS run job task.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cluster cluster
//   var taskDefinition taskDefinition
//
//   launchTargetBindOptions := &LaunchTargetBindOptions{
//   	TaskDefinition: taskDefinition,
//
//   	// the properties below are optional
//   	Cluster: cluster,
//   }
//
type LaunchTargetBindOptions struct {
	// Task definition to run Docker containers in Amazon ECS.
	TaskDefinition awsecs.ITaskDefinition `field:"required" json:"taskDefinition" yaml:"taskDefinition"`
	// A regional grouping of one or more container instances on which you can run tasks and services.
	// Default: - No cluster.
	//
	Cluster awsecs.ICluster `field:"optional" json:"cluster" yaml:"cluster"`
}


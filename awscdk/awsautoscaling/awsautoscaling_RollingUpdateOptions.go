package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for customizing the rolling update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rollingUpdateOptions := &rollingUpdateOptions{
//   	maxBatchSize: jsii.Number(123),
//   	minInstancesInService: jsii.Number(123),
//   	minSuccessPercentage: jsii.Number(123),
//   	pauseTime: cdk.duration.minutes(jsii.Number(30)),
//   	suspendProcesses: []scalingProcess{
//   		awscdk.Aws_autoscaling.*scalingProcess_LAUNCH,
//   	},
//   	waitOnResourceSignals: jsii.Boolean(false),
//   }
//
type RollingUpdateOptions struct {
	// The maximum number of instances that AWS CloudFormation updates at once.
	//
	// This number affects the speed of the replacement.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// The minimum number of instances that must be in service before more instances are replaced.
	//
	// This number affects the speed of the replacement.
	MinInstancesInService *float64 `field:"optional" json:"minInstancesInService" yaml:"minInstancesInService"`
	// The percentage of instances that must signal success for the update to succeed.
	MinSuccessPercentage *float64 `field:"optional" json:"minSuccessPercentage" yaml:"minSuccessPercentage"`
	// The pause time after making a change to a batch of instances.
	PauseTime awscdk.Duration `field:"optional" json:"pauseTime" yaml:"pauseTime"`
	// Specifies the Auto Scaling processes to suspend during a stack update.
	//
	// Suspending processes prevents Auto Scaling from interfering with a stack
	// update.
	SuspendProcesses *[]ScalingProcess `field:"optional" json:"suspendProcesses" yaml:"suspendProcesses"`
	// Specifies whether the Auto Scaling group waits on signals from new instances during an update.
	WaitOnResourceSignals *bool `field:"optional" json:"waitOnResourceSignals" yaml:"waitOnResourceSignals"`
}


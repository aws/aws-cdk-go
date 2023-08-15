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
//   rollingUpdateOptions := &RollingUpdateOptions{
//   	MaxBatchSize: jsii.Number(123),
//   	MinInstancesInService: jsii.Number(123),
//   	MinSuccessPercentage: jsii.Number(123),
//   	PauseTime: cdk.Duration_Minutes(jsii.Number(30)),
//   	SuspendProcesses: []scalingProcess{
//   		awscdk.Aws_autoscaling.*scalingProcess_LAUNCH,
//   	},
//   	WaitOnResourceSignals: jsii.Boolean(false),
//   }
//
type RollingUpdateOptions struct {
	// The maximum number of instances that AWS CloudFormation updates at once.
	//
	// This number affects the speed of the replacement.
	// Default: 1.
	//
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// The minimum number of instances that must be in service before more instances are replaced.
	//
	// This number affects the speed of the replacement.
	// Default: 0.
	//
	MinInstancesInService *float64 `field:"optional" json:"minInstancesInService" yaml:"minInstancesInService"`
	// The percentage of instances that must signal success for the update to succeed.
	// Default: - The `minSuccessPercentage` configured for `signals` on the AutoScalingGroup.
	//
	MinSuccessPercentage *float64 `field:"optional" json:"minSuccessPercentage" yaml:"minSuccessPercentage"`
	// The pause time after making a change to a batch of instances.
	// Default: - The `timeout` configured for `signals` on the AutoScalingGroup.
	//
	PauseTime awscdk.Duration `field:"optional" json:"pauseTime" yaml:"pauseTime"`
	// Specifies the Auto Scaling processes to suspend during a stack update.
	//
	// Suspending processes prevents Auto Scaling from interfering with a stack
	// update.
	// Default: HealthCheck, ReplaceUnhealthy, AZRebalance, AlarmNotification, ScheduledActions.
	//
	SuspendProcesses *[]ScalingProcess `field:"optional" json:"suspendProcesses" yaml:"suspendProcesses"`
	// Specifies whether the Auto Scaling group waits on signals from new instances during an update.
	// Default: true if you configured `signals` on the AutoScalingGroup, false otherwise.
	//
	WaitOnResourceSignals *bool `field:"optional" json:"waitOnResourceSignals" yaml:"waitOnResourceSignals"`
}


package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Additional settings when a rolling update is selected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//
//   rollingUpdateConfiguration := &rollingUpdateConfiguration{
//   	maxBatchSize: jsii.Number(123),
//   	minInstancesInService: jsii.Number(123),
//   	minSuccessfulInstancesPercent: jsii.Number(123),
//   	pauseTime: duration,
//   	suspendProcesses: []scalingProcess{
//   		awscdk.Aws_autoscaling.*scalingProcess_LAUNCH,
//   	},
//   	waitOnResourceSignals: jsii.Boolean(false),
//   }
//
// Deprecated: use `UpdatePolicy.rollingUpdate()`
type RollingUpdateConfiguration struct {
	// The maximum number of instances that AWS CloudFormation updates at once.
	// Deprecated: use `UpdatePolicy.rollingUpdate()`
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// The minimum number of instances that must be in service before more instances are replaced.
	//
	// This number affects the speed of the replacement.
	// Deprecated: use `UpdatePolicy.rollingUpdate()`
	MinInstancesInService *float64 `field:"optional" json:"minInstancesInService" yaml:"minInstancesInService"`
	// The percentage of instances that must signal success for an update to succeed.
	//
	// If an instance doesn't send a signal within the time specified in the
	// pauseTime property, AWS CloudFormation assumes that the instance wasn't
	// updated.
	//
	// This number affects the success of the replacement.
	//
	// If you specify this property, you must also enable the
	// waitOnResourceSignals and pauseTime properties.
	// Deprecated: use `UpdatePolicy.rollingUpdate()`
	MinSuccessfulInstancesPercent *float64 `field:"optional" json:"minSuccessfulInstancesPercent" yaml:"minSuccessfulInstancesPercent"`
	// The pause time after making a change to a batch of instances.
	//
	// This is intended to give those instances time to start software applications.
	//
	// Specify PauseTime in the ISO8601 duration format (in the format
	// PT#H#M#S, where each # is the number of hours, minutes, and seconds,
	// respectively). The maximum PauseTime is one hour (PT1H).
	// Deprecated: use `UpdatePolicy.rollingUpdate()`
	PauseTime awscdk.Duration `field:"optional" json:"pauseTime" yaml:"pauseTime"`
	// Specifies the Auto Scaling processes to suspend during a stack update.
	//
	// Suspending processes prevents Auto Scaling from interfering with a stack
	// update.
	// Deprecated: use `UpdatePolicy.rollingUpdate()`
	SuspendProcesses *[]ScalingProcess `field:"optional" json:"suspendProcesses" yaml:"suspendProcesses"`
	// Specifies whether the Auto Scaling group waits on signals from new instances during an update.
	//
	// AWS CloudFormation must receive a signal from each new instance within
	// the specified PauseTime before continuing the update.
	//
	// To have instances wait for an Elastic Load Balancing health check before
	// they signal success, add a health-check verification by using the
	// cfn-init helper script. For an example, see the verify_instance_health
	// command in the Auto Scaling rolling updates sample template.
	// Deprecated: use `UpdatePolicy.rollingUpdate()`
	WaitOnResourceSignals *bool `field:"optional" json:"waitOnResourceSignals" yaml:"waitOnResourceSignals"`
}


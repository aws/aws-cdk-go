package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Specifies an action that AWS Batch will take after the job has remained at the head of the queue in the specified state for longer than the specified time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobStateTimeLimitAction := &JobStateTimeLimitAction{
//   	MaxTime: cdk.Duration_Minutes(jsii.Number(30)),
//   	Reason: awscdk.Aws_batch.JobStateTimeLimitActionsReason_INSUFFICIENT_INSTANCE_CAPACITY,
//
//   	// the properties below are optional
//   	Action: awscdk.*Aws_batch.JobStateTimeLimitActionsAction_CANCEL,
//   	State: awscdk.*Aws_batch.JobStateTimeLimitActionsState_RUNNABLE,
//   }
//
type JobStateTimeLimitAction struct {
	// The approximate amount of time, that must pass with the job in the specified state before the action is taken.
	//
	// The minimum value is 10 minutes and the maximum value is 24 hours.
	MaxTime awscdk.Duration `field:"required" json:"maxTime" yaml:"maxTime"`
	// The reason to log for the action being taken.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html#job_stuck_in_runnable
	//
	Reason JobStateTimeLimitActionsReason `field:"required" json:"reason" yaml:"reason"`
	// The action to take when a job is at the head of the job queue in the specified state for the specified period of time.
	// Default: JobStateTimeLimitActionsAction.CANCEL
	//
	Action JobStateTimeLimitActionsAction `field:"optional" json:"action" yaml:"action"`
	// The state of the job needed to trigger the action.
	// Default: JobStateTimeLimitActionsState.RUNNABLE
	//
	State JobStateTimeLimitActionsState `field:"optional" json:"state" yaml:"state"`
}


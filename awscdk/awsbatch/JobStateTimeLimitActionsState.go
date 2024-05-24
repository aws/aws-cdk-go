package awsbatch


// The state of the job needed to trigger the action.
//
// Example:
//   batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
//   	JobStateTimeLimitActions: []jobStateTimeLimitAction{
//   		&jobStateTimeLimitAction{
//   			Action: batch.JobStateTimeLimitActionsAction_CANCEL,
//   			MaxTime: cdk.Duration_Minutes(jsii.Number(10)),
//   			Reason: batch.JobStateTimeLimitActionsReason_INSUFFICIENT_INSTANCE_CAPACITY,
//   			State: batch.JobStateTimeLimitActionsState_RUNNABLE,
//   		},
//   	},
//   })
//
type JobStateTimeLimitActionsState string

const (
	// RUNNABLE state triggers the action.
	JobStateTimeLimitActionsState_RUNNABLE JobStateTimeLimitActionsState = "RUNNABLE"
)


package awsbatch


// The action to take when a job is at the head of the job queue in the specified state for the specified period of time.
//
// Example:
//   batch.NewJobQueue(this, jsii.String("JobQueue"), &JobQueueProps{
//   	JobStateTimeLimitActions: []JobStateTimeLimitAction{
//   		&JobStateTimeLimitAction{
//   			Action: batch.JobStateTimeLimitActionsAction_CANCEL,
//   			MaxTime: cdk.Duration_Minutes(jsii.Number(10)),
//   			Reason: batch.JobStateTimeLimitActionsReason_INSUFFICIENT_INSTANCE_CAPACITY,
//   			State: batch.JobStateTimeLimitActionsState_RUNNABLE,
//   		},
//   	},
//   })
//
type JobStateTimeLimitActionsAction string

const (
	// Cancel the job.
	JobStateTimeLimitActionsAction_CANCEL JobStateTimeLimitActionsAction = "CANCEL"
	// Terminate the job.
	JobStateTimeLimitActionsAction_TERMINATE JobStateTimeLimitActionsAction = "TERMINATE"
)


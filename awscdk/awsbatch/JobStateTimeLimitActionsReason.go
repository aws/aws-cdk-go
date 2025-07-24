package awsbatch


// The reason to log for the action being taken.
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
// See: https://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html#job_stuck_in_runnable
//
type JobStateTimeLimitActionsReason string

const (
	// All connected compute environments have insufficient capacity errors.
	JobStateTimeLimitActionsReason_INSUFFICIENT_INSTANCE_CAPACITY JobStateTimeLimitActionsReason = "INSUFFICIENT_INSTANCE_CAPACITY"
	// All compute environments have a maxvCpus parameter that is smaller than the job requirements.
	JobStateTimeLimitActionsReason_COMPUTE_ENVIRONMENT_MAX_RESOURCE JobStateTimeLimitActionsReason = "COMPUTE_ENVIRONMENT_MAX_RESOURCE"
	// None of the compute environments have instances that meet the job requirements.
	JobStateTimeLimitActionsReason_JOB_RESOURCE_REQUIREMENT JobStateTimeLimitActionsReason = "JOB_RESOURCE_REQUIREMENT"
)


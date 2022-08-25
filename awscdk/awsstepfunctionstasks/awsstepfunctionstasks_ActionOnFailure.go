package awsstepfunctionstasks


// The action to take when the cluster step fails.
//
// Example:
//   tasks.NewEmrAddStep(this, jsii.String("Task"), &emrAddStepProps{
//   	clusterId: jsii.String("ClusterId"),
//   	name: jsii.String("StepName"),
//   	jar: jsii.String("Jar"),
//   	actionOnFailure: tasks.actionOnFailure_CONTINUE,
//   })
//
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_StepConfig.html
//
// Here, they are named as TERMINATE_JOB_FLOW, TERMINATE_CLUSTER, CANCEL_AND_WAIT, and CONTINUE respectively.
//
type ActionOnFailure string

const (
	// Terminate the Cluster on Step Failure.
	ActionOnFailure_TERMINATE_CLUSTER ActionOnFailure = "TERMINATE_CLUSTER"
	// Cancel Step execution and enter WAITING state.
	ActionOnFailure_CANCEL_AND_WAIT ActionOnFailure = "CANCEL_AND_WAIT"
	// Continue to the next Step.
	ActionOnFailure_CONTINUE ActionOnFailure = "CONTINUE"
)


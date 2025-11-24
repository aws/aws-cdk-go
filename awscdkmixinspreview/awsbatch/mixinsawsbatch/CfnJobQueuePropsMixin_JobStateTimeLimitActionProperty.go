package mixinsawsbatch


// Specifies an action that AWS Batch will take after the job has remained at the head of the queue in the specified state for longer than the specified time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   jobStateTimeLimitActionProperty := &JobStateTimeLimitActionProperty{
//   	Action: jsii.String("action"),
//   	MaxTimeSeconds: jsii.Number(123),
//   	Reason: jsii.String("reason"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-jobstatetimelimitaction.html
//
type CfnJobQueuePropsMixin_JobStateTimeLimitActionProperty struct {
	// The action to take when a job is at the head of the job queue in the specified state for the specified period of time.
	//
	// The only supported value is `CANCEL` , which will cancel the job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-jobstatetimelimitaction.html#cfn-batch-jobqueue-jobstatetimelimitaction-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// The approximate amount of time, in seconds, that must pass with the job in the specified state before the action is taken.
	//
	// The minimum value is 600 (10 minutes) and the maximum value is 86,400 (24 hours).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-jobstatetimelimitaction.html#cfn-batch-jobqueue-jobstatetimelimitaction-maxtimeseconds
	//
	MaxTimeSeconds *float64 `field:"optional" json:"maxTimeSeconds" yaml:"maxTimeSeconds"`
	// The reason to log for the action being taken.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-jobstatetimelimitaction.html#cfn-batch-jobqueue-jobstatetimelimitaction-reason
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// The state of the job needed to trigger the action.
	//
	// The only supported value is `RUNNABLE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-jobstatetimelimitaction.html#cfn-batch-jobqueue-jobstatetimelimitaction-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}


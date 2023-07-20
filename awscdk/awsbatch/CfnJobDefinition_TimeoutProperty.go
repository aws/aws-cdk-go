package awsbatch


// An object that represents a job timeout configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeoutProperty := &TimeoutProperty{
//   	AttemptDurationSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-timeout.html
//
type CfnJobDefinition_TimeoutProperty struct {
	// The job timeout time (in seconds) that's measured from the job attempt's `startedAt` timestamp.
	//
	// After this time passes, AWS Batch terminates your jobs if they aren't finished. The minimum value for the timeout is 60 seconds.
	//
	// For array jobs, the timeout applies to the child jobs, not to the parent array job.
	//
	// For multi-node parallel (MNP) jobs, the timeout applies to the whole job, not to the individual nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-timeout.html#cfn-batch-jobdefinition-timeout-attemptdurationseconds
	//
	AttemptDurationSeconds *float64 `field:"optional" json:"attemptDurationSeconds" yaml:"attemptDurationSeconds"`
}


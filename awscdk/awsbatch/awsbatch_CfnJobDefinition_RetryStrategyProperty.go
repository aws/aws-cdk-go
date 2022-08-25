package awsbatch


// The retry strategy associated with a job.
//
// For more information, see [Automated job retries](https://docs.aws.amazon.com/batch/latest/userguide/job_retries.html) in the *AWS Batch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retryStrategyProperty := &retryStrategyProperty{
//   	attempts: jsii.Number(123),
//   	evaluateOnExit: []interface{}{
//   		&evaluateOnExitProperty{
//   			action: jsii.String("action"),
//
//   			// the properties below are optional
//   			onExitCode: jsii.String("onExitCode"),
//   			onReason: jsii.String("onReason"),
//   			onStatusReason: jsii.String("onStatusReason"),
//   		},
//   	},
//   }
//
type CfnJobDefinition_RetryStrategyProperty struct {
	// The number of times to move a job to the `RUNNABLE` status.
	//
	// You can specify between 1 and 10 attempts. If the value of `attempts` is greater than one, the job is retried on failure the same number of attempts as the value.
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// Array of up to 5 objects that specify conditions under which the job should be retried or failed.
	//
	// If this parameter is specified, then the `attempts` parameter must also be specified.
	EvaluateOnExit interface{} `field:"optional" json:"evaluateOnExit" yaml:"evaluateOnExit"`
}


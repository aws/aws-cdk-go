package awspipes


// The retry strategy that's associated with a job.
//
// For more information, see [Automated job retries](https://docs.aws.amazon.com/batch/latest/userguide/job_retries.html) in the *AWS Batch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchRetryStrategyProperty := &batchRetryStrategyProperty{
//   	attempts: jsii.Number(123),
//   }
//
type CfnPipe_BatchRetryStrategyProperty struct {
	// The number of times to move a job to the `RUNNABLE` status.
	//
	// If the value of `attempts` is greater than one, the job is retried on failure the same number of attempts as the value.
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
}


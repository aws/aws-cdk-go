package previewawspipesmixins


// The retry strategy that's associated with a job.
//
// For more information, see [Automated job retries](https://docs.aws.amazon.com/batch/latest/userguide/job_retries.html) in the *AWS Batch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   batchRetryStrategyProperty := &BatchRetryStrategyProperty{
//   	Attempts: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchretrystrategy.html
//
type CfnPipePropsMixin_BatchRetryStrategyProperty struct {
	// The number of times to move a job to the `RUNNABLE` status.
	//
	// If the value of `attempts` is greater than one, the job is retried on failure the same number of attempts as the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-batchretrystrategy.html#cfn-pipes-pipe-batchretrystrategy-attempts
	//
	// Default: - 0.
	//
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
}


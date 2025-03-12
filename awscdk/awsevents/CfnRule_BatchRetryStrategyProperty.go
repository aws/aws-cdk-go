package awsevents


// The retry strategy to use for failed jobs, if the target is an AWS Batch job.
//
// If you specify a retry strategy here, it overrides the retry strategy defined in the job definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchRetryStrategyProperty := &BatchRetryStrategyProperty{
//   	Attempts: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-batchretrystrategy.html
//
type CfnRule_BatchRetryStrategyProperty struct {
	// The number of times to attempt to retry, if the job fails.
	//
	// Valid values are 1–10.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-batchretrystrategy.html#cfn-events-rule-batchretrystrategy-attempts
	//
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
}


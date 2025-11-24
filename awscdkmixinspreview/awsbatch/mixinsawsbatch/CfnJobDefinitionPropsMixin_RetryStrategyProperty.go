package mixinsawsbatch


// The retry strategy that's associated with a job.
//
// For more information, see [Automated job retries](https://docs.aws.amazon.com/batch/latest/userguide/job_retries.html) in the *AWS Batch User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   retryStrategyProperty := &RetryStrategyProperty{
//   	Attempts: jsii.Number(123),
//   	EvaluateOnExit: []interface{}{
//   		&EvaluateOnExitProperty{
//   			Action: jsii.String("action"),
//   			OnExitCode: jsii.String("onExitCode"),
//   			OnReason: jsii.String("onReason"),
//   			OnStatusReason: jsii.String("onStatusReason"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-retrystrategy.html
//
type CfnJobDefinitionPropsMixin_RetryStrategyProperty struct {
	// The number of times to move a job to the `RUNNABLE` status.
	//
	// You can specify between 1 and 10 attempts. If the value of `attempts` is greater than one, the job is retried on failure the same number of attempts as the value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-retrystrategy.html#cfn-batch-jobdefinition-retrystrategy-attempts
	//
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// Array of up to 5 objects that specify the conditions where jobs are retried or failed.
	//
	// If this parameter is specified, then the `attempts` parameter must also be specified. If none of the listed conditions match, then the job is retried.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-retrystrategy.html#cfn-batch-jobdefinition-retrystrategy-evaluateonexit
	//
	EvaluateOnExit interface{} `field:"optional" json:"evaluateOnExit" yaml:"evaluateOnExit"`
}


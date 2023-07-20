package awsevents


// The custom parameters to be used when the target is an AWS Batch job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   batchParametersProperty := &BatchParametersProperty{
//   	JobDefinition: jsii.String("jobDefinition"),
//   	JobName: jsii.String("jobName"),
//
//   	// the properties below are optional
//   	ArrayProperties: &BatchArrayPropertiesProperty{
//   		Size: jsii.Number(123),
//   	},
//   	RetryStrategy: &BatchRetryStrategyProperty{
//   		Attempts: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-batchparameters.html
//
type CfnRule_BatchParametersProperty struct {
	// The ARN or name of the job definition to use if the event target is an AWS Batch job.
	//
	// This job definition must already exist.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-batchparameters.html#cfn-events-rule-batchparameters-jobdefinition
	//
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// The name to use for this execution of the job, if the target is an AWS Batch job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-batchparameters.html#cfn-events-rule-batchparameters-jobname
	//
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// The array properties for the submitted job, such as the size of the array.
	//
	// The array size can be between 2 and 10,000. If you specify array properties for a job, it becomes an array job. This parameter is used only if the target is an AWS Batch job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-batchparameters.html#cfn-events-rule-batchparameters-arrayproperties
	//
	ArrayProperties interface{} `field:"optional" json:"arrayProperties" yaml:"arrayProperties"`
	// The retry strategy to use for failed jobs, if the target is an AWS Batch job.
	//
	// The retry strategy is the number of times to retry the failed job execution. Valid values are 1–10. When you specify a retry strategy here, it overrides the retry strategy defined in the job definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-batchparameters.html#cfn-events-rule-batchparameters-retrystrategy
	//
	RetryStrategy interface{} `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}


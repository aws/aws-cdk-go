package mixinsawssagemaker


// Configures conditions under which the processing job should be stopped, such as how long the processing job has been running.
//
// After the condition is met, the processing job is stopped.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stoppingConditionProperty := &StoppingConditionProperty{
//   	MaxRuntimeInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-stoppingcondition.html
//
type CfnProcessingJobPropsMixin_StoppingConditionProperty struct {
	// The maximum length of time, in seconds, that a training or compilation job can run before it is stopped.
	//
	// For compilation jobs, if the job does not complete during this time, a `TimeOut` error is generated. We recommend starting with 900 seconds and increasing as necessary based on your model.
	//
	// For all other jobs, if the job does not complete during this time, SageMaker ends the job. When `RetryStrategy` is specified in the job request, `MaxRuntimeInSeconds` specifies the maximum time for all of the attempts in total, not each individual attempt. The default value is 1 day. The maximum value is 28 days.
	//
	// The maximum time that a `TrainingJob` can run in total, including any time spent publishing metrics or archiving and uploading models after it has been stopped, is 30 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-processingjob-stoppingcondition.html#cfn-sagemaker-processingjob-stoppingcondition-maxruntimeinseconds
	//
	MaxRuntimeInSeconds *float64 `field:"optional" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}


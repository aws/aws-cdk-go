package awsschedulertargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties for a SageMaker Target.
//
// Example:
//   import sagemaker "github.com/aws/aws-cdk-go/awscdk"
//
//   var pipeline iPipeline
//
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(60))),
//   	Target: targets.NewSageMakerStartPipelineExecution(pipeline, &SageMakerStartPipelineExecutionProps{
//   		PipelineParameterList: []sageMakerPipelineParameter{
//   			&sageMakerPipelineParameter{
//   				Name: jsii.String("parameter-name"),
//   				Value: jsii.String("parameter-value"),
//   			},
//   		},
//   	}),
//   })
//
type SageMakerStartPipelineExecutionProps struct {
	// The SQS queue to be used as deadLetterQueue.
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Default: - no dead-letter queue.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Input passed to the target.
	// Default: - no input.
	//
	Input awsscheduler.ScheduleTargetInput `field:"optional" json:"input" yaml:"input"`
	// The maximum age of a request that Scheduler sends to a target for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Default: Duration.hours(24)
	//
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the target returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Default: 185.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// An execution role is an IAM role that EventBridge Scheduler assumes in order to interact with other AWS services on your behalf.
	//
	// If none provided templates target will automatically create an IAM role with all the minimum necessary
	// permissions to interact with the templated target. If you wish you may specify your own IAM role, then the templated targets
	// will grant minimal required permissions.
	// Default: - created by target.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// List of parameter names and values to use when executing the SageMaker Model Building Pipeline.
	//
	// The length must be between 0 and 200.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-sagemakerpipelineparameters.html#cfn-scheduler-schedule-sagemakerpipelineparameters-pipelineparameterlist
	//
	// Default: - no pipeline parameter list.
	//
	PipelineParameterList *[]*SageMakerPipelineParameter `field:"optional" json:"pipelineParameterList" yaml:"pipelineParameterList"`
}


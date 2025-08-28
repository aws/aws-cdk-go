package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Customize the Lambda Event Target.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunc"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("exports.handler = handler.toString()")),
//   })
//
//   rule := events.NewRule(this, jsii.String("rule"), &RuleProps{
//   	EventPattern: &EventPattern{
//   		Source: []*string{
//   			jsii.String("aws.ec2"),
//   		},
//   	},
//   })
//
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   rule.AddTarget(targets.NewLambdaFunction(fn, &LambdaFunctionProps{
//   	DeadLetterQueue: queue,
//   	 // Optional: add a dead letter queue
//   	MaxEventAge: awscdk.Duration_Hours(jsii.Number(2)),
//   	 // Optional: set the maxEventAge retry policy
//   	RetryAttempts: jsii.Number(2),
//   }))
//
type LambdaFunctionProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Default: - no dead-letter queue.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Default: Duration.hours(24)
	//
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Default: 185.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The event to send to the Lambda.
	//
	// This will be the payload sent to the Lambda Function.
	// Default: the entire EventBridge event.
	//
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
}


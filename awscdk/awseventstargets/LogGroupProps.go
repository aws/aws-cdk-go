package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Customize the CloudWatch LogGroup Event Target.
//
// Example:
//   import logs "github.com/aws/aws-cdk-go/awscdk"
//   var logGroup logGroup
//   var rule rule
//
//
//   rule.AddTarget(targets.NewCloudWatchLogGroup(logGroup, &LogGroupProps{
//   	LogEvent: targets.LogGroupTargetInput_FromObject(&LogGroupTargetInputOptions{
//   		Message: jSON.stringify(map[string]*string{
//   			"CustomField": jsii.String("CustomValue"),
//   		}),
//   	}),
//   }))
//
type LogGroupProps struct {
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
	// The event to send to the CloudWatch LogGroup.
	//
	// This will be the event logged into the CloudWatch LogGroup.
	// Default: - the entire EventBridge event.
	//
	// Deprecated: use logEvent instead.
	Event awsevents.RuleTargetInput `field:"optional" json:"event" yaml:"event"`
	// Whether the custom resource created wll default to install latest AWS SDK.
	// Default: - install latest AWS SDK.
	//
	InstallLatestAwsSdk *bool `field:"optional" json:"installLatestAwsSdk" yaml:"installLatestAwsSdk"`
	// The event to send to the CloudWatch LogGroup.
	//
	// This will be the event logged into the CloudWatch LogGroup.
	// Default: - the entire EventBridge event.
	//
	LogEvent LogGroupTargetInput `field:"optional" json:"logEvent" yaml:"logEvent"`
}


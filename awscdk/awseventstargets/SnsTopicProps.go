package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Customize the SNS Topic Event Target.
//
// Example:
//   var onCommitRule rule
//   var topic topic
//
//
//   onCommitRule.AddTarget(targets.NewSnsTopic(topic, &SnsTopicProps{
//   	Message: events.RuleTargetInput_FromText(
//   	fmt.Sprintf("A commit was pushed to the repository %v on branch %v", codecommit.ReferenceEvent_RepositoryName(), codecommit.ReferenceEvent_ReferenceName())),
//   }))
//
type SnsTopicProps struct {
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
	// The message to send to the topic.
	// Default: the entire EventBridge event.
	//
	Message awsevents.RuleTargetInput `field:"optional" json:"message" yaml:"message"`
}


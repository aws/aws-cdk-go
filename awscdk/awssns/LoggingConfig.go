package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// A logging configuration for delivery status of messages sent from SNS topic to subscribed endpoints.
//
// For more information, see https://docs.aws.amazon.com/sns/latest/dg/sns-topic-attributes.html.
//
// Example:
//   var role role
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   topic.AddLoggingConfig(&LoggingConfig{
//   	Protocol: sns.LoggingProtocol_SQS,
//   	FailureFeedbackRole: role,
//   	SuccessFeedbackRole: role,
//   	SuccessFeedbackSampleRate: jsii.Number(50),
//   })
//
type LoggingConfig struct {
	// Indicates one of the supported protocols for the SNS topic.
	Protocol LoggingProtocol `field:"required" json:"protocol" yaml:"protocol"`
	// The IAM role to be used when logging failed message deliveries in Amazon CloudWatch.
	// Default: None.
	//
	FailureFeedbackRole awsiam.IRole `field:"optional" json:"failureFeedbackRole" yaml:"failureFeedbackRole"`
	// The IAM role to be used when logging successful message deliveries in Amazon CloudWatch.
	// Default: None.
	//
	SuccessFeedbackRole awsiam.IRole `field:"optional" json:"successFeedbackRole" yaml:"successFeedbackRole"`
	// The percentage of successful message deliveries to be logged in Amazon CloudWatch.
	//
	// Valid values are integer between 0-100.
	// Default: None.
	//
	SuccessFeedbackSampleRate *float64 `field:"optional" json:"successFeedbackSampleRate" yaml:"successFeedbackSampleRate"`
}


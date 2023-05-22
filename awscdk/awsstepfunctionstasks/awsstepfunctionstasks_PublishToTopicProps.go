package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for PublishTask.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var taskInput taskInput
//
//   publishToTopicProps := &PublishToTopicProps{
//   	Message: taskInput,
//
//   	// the properties below are optional
//   	IntegrationPattern: awscdk.Aws_stepfunctions.ServiceIntegrationPattern_FIRE_AND_FORGET,
//   	MessagePerSubscriptionType: jsii.Boolean(false),
//   	Subject: jsii.String("subject"),
//   }
//
// Deprecated: Use `SnsPublish`.
type PublishToTopicProps struct {
	// The text message to send to the topic.
	// Deprecated: Use `SnsPublish`.
	Message awsstepfunctions.TaskInput `field:"required" json:"message" yaml:"message"`
	// The service integration pattern indicates different ways to call Publish to SNS.
	//
	// The valid value is either FIRE_AND_FORGET or WAIT_FOR_TASK_TOKEN.
	// Deprecated: Use `SnsPublish`.
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// If true, send a different message to every subscription type.
	//
	// If this is set to true, message must be a JSON object with a
	// "default" key and a key for every subscription type (such as "sqs",
	// "email", etc.) The values are strings representing the messages
	// being sent to every subscription type.
	// See: https://docs.aws.amazon.com/sns/latest/api/API_Publish.html#API_Publish_RequestParameters
	//
	// Deprecated: Use `SnsPublish`.
	MessagePerSubscriptionType *bool `field:"optional" json:"messagePerSubscriptionType" yaml:"messagePerSubscriptionType"`
	// Used as the "Subject" line when the message is delivered to email endpoints.
	//
	// Also included, if present, in the standard JSON messages delivered to other endpoints.
	// Deprecated: Use `SnsPublish`.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}


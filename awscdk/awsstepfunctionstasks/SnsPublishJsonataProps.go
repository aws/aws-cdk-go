package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for publishing a message to an SNS topic using JSONata.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assign interface{}
//   var outputs interface{}
//   var taskInput taskInput
//   var taskRole taskRole
//   var timeout timeout
//   var topic topic
//   var value interface{}
//
//   snsPublishJsonataProps := &SnsPublishJsonataProps{
//   	Message: taskInput,
//   	Topic: topic,
//
//   	// the properties below are optional
//   	Assign: map[string]interface{}{
//   		"assignKey": assign,
//   	},
//   	Comment: jsii.String("comment"),
//   	Credentials: &Credentials{
//   		Role: taskRole,
//   	},
//   	Heartbeat: cdk.Duration_Minutes(jsii.Number(30)),
//   	HeartbeatTimeout: timeout,
//   	IntegrationPattern: awscdk.Aws_stepfunctions.IntegrationPattern_REQUEST_RESPONSE,
//   	MessageAttributes: map[string]messageAttribute{
//   		"messageAttributesKey": &messageAttribute{
//   			"value": value,
//
//   			// the properties below are optional
//   			"dataType": awscdk.aws_stepfunctions_tasks.MessageAttributeDataType_STRING,
//   		},
//   	},
//   	MessageDeduplicationId: jsii.String("messageDeduplicationId"),
//   	MessageGroupId: jsii.String("messageGroupId"),
//   	MessagePerSubscriptionType: jsii.Boolean(false),
//   	Outputs: outputs,
//   	QueryLanguage: awscdk.*Aws_stepfunctions.QueryLanguage_JSON_PATH,
//   	StateName: jsii.String("stateName"),
//   	Subject: jsii.String("subject"),
//   	TaskTimeout: timeout,
//   	Timeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   }
//
type SnsPublishJsonataProps struct {
	// A comment describing this state.
	// Default: No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The name of the query language used by the state.
	//
	// If the state does not contain a `queryLanguage` field,
	// then it will use the query language specified in the top-level `queryLanguage` field.
	// Default: - JSONPath.
	//
	QueryLanguage awsstepfunctions.QueryLanguage `field:"optional" json:"queryLanguage" yaml:"queryLanguage"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	// Default: - None (Task is executed using the State Machine's execution role).
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Default: - None.
	//
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	//
	// Depending on the AWS Service, the Service Integration Pattern availability will vary.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-supported-services.html
	//
	// Default: - `IntegrationPattern.REQUEST_RESPONSE` for most tasks.
	// `IntegrationPattern.RUN_JOB` for the following exceptions:
	// `BatchSubmitJob`, `EmrAddStep`, `EmrCreateCluster`, `EmrTerminationCluster`, and `EmrContainersStartJobRun`.
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Default: - None.
	//
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// Workflow variables to store in this step.
	//
	// Using workflow variables, you can store data in a step and retrieve that data in future steps.
	// See: https://docs.aws.amazon.com/ja_jp/step-functions/latest/dg/workflow-variables.html
	//
	// Default: - Not assign variables.
	//
	Assign *map[string]interface{} `field:"optional" json:"assign" yaml:"assign"`
	// Used to specify and transform output from the state.
	//
	// When specified, the value overrides the state output default.
	// The output field accepts any JSON value (object, array, string, number, boolean, null).
	// Any string value, including those inside objects or arrays,
	// will be evaluated as JSONata if surrounded by {% %} characters.
	// Output also accepts a JSONata expression directly.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-input-output-filtering.html
	//
	// Default: - $states.result or $states.errorOutput
	//
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
	// The message you want to send.
	//
	// With the exception of SMS, messages must be UTF-8 encoded strings and
	// at most 256 KB in size.
	// For SMS, each message can contain up to 140 characters.
	Message awsstepfunctions.TaskInput `field:"required" json:"message" yaml:"message"`
	// The SNS topic that the task will publish to.
	Topic awssns.ITopic `field:"required" json:"topic" yaml:"topic"`
	// Add message attributes when publishing.
	//
	// These attributes carry additional metadata about the message and may be used
	// for subscription filters.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html
	//
	// Default: {}.
	//
	MessageAttributes *map[string]*MessageAttribute `field:"optional" json:"messageAttributes" yaml:"messageAttributes"`
	// This parameter applies only to FIFO topics.
	//
	// Every message must have a unique MessageDeduplicationId, which is a token used for deduplication of sent messages.
	// If a message with a particular MessageDeduplicationId is sent successfully, any message sent with the same MessageDeduplicationId
	// during the 5-minute deduplication interval is treated as a duplicate.
	//
	// If the topic has ContentBasedDeduplication set, the system generates a MessageDeduplicationId
	// based on the contents of the message. Your MessageDeduplicationId overrides the generated one.
	// Default: - Not used for standard topics, required for FIFO topics with ContentBasedDeduplication disabled.
	//
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// This parameter applies only to FIFO topics.
	//
	// The MessageGroupId is a tag that specifies that a message belongs to a specific message group.
	// Messages that belong to the same message group are processed in a FIFO manner
	// (however, messages in different message groups might be processed out of order).
	// Every message must include a MessageGroupId.
	// Default: - Not used for standard topics, required for FIFO topics.
	//
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
	// Send different messages for each transport protocol.
	//
	// For example, you might want to send a shorter message to SMS subscribers
	// and a more verbose message to email and SQS subscribers.
	//
	// Your message must be a JSON object with a top-level JSON key of
	// "default" with a value that is a string
	// You can define other top-level keys that define the message you want to
	// send to a specific transport protocol (i.e. "sqs", "email", "http", etc)
	// See: https://docs.aws.amazon.com/sns/latest/api/API_Publish.html#API_Publish_RequestParameters
	//
	// Default: false.
	//
	MessagePerSubscriptionType *bool `field:"optional" json:"messagePerSubscriptionType" yaml:"messagePerSubscriptionType"`
	// Used as the "Subject" line when the message is delivered to email endpoints.
	//
	// This field will also be included, if present, in the standard JSON messages
	// delivered to other endpoints.
	// Default: - No subject.
	//
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
}


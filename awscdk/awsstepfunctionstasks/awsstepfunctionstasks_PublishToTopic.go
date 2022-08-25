package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
)

// A Step Functions Task to publish messages to SNS topic.
//
// A Function can be used directly as a Resource, but this class mirrors
// integration with other AWS services via a specific class instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var taskInput taskInput
//   var topic topic
//
//   publishToTopic := awscdk.Aws_stepfunctions_tasks.NewPublishToTopic(topic, &publishToTopicProps{
//   	message: taskInput,
//
//   	// the properties below are optional
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	messagePerSubscriptionType: jsii.Boolean(false),
//   	subject: jsii.String("subject"),
//   })
//
// Deprecated: Use `SnsPublish`.
type PublishToTopic interface {
	awsstepfunctions.IStepFunctionsTask
	// Called when the task object is used in a workflow.
	// Deprecated: Use `SnsPublish`.
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for PublishToTopic
type jsiiProxy_PublishToTopic struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: Use `SnsPublish`.
func NewPublishToTopic(topic awssns.ITopic, props *PublishToTopicProps) PublishToTopic {
	_init_.Initialize()

	j := jsiiProxy_PublishToTopic{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.PublishToTopic",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Deprecated: Use `SnsPublish`.
func NewPublishToTopic_Override(p PublishToTopic, topic awssns.ITopic, props *PublishToTopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.PublishToTopic",
		[]interface{}{topic, props},
		p,
	)
}

func (p *jsiiProxy_PublishToTopic) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}


package awsiotactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsiot"
	"github.com/aws/aws-cdk-go/awscdk/awsiotactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
)

// The action to send data to Amazon CloudWatch Logs.
// Experimental.
type CloudWatchLogsAction interface {
	awsiot.IAction
	Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for CloudWatchLogsAction
type jsiiProxy_CloudWatchLogsAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewCloudWatchLogsAction(logGroup awslogs.ILogGroup, props *CloudWatchLogsActionProps) CloudWatchLogsAction {
	_init_.Initialize()

	j := jsiiProxy_CloudWatchLogsAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchLogsAction",
		[]interface{}{logGroup, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchLogsAction_Override(c CloudWatchLogsAction, logGroup awslogs.ILogGroup, props *CloudWatchLogsActionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.CloudWatchLogsAction",
		[]interface{}{logGroup, props},
		c,
	)
}

// Returns the topic rule action specification.
// Experimental.
func (c *jsiiProxy_CloudWatchLogsAction) Bind(rule awsiot.ITopicRule) *awsiot.ActionConfig {
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{rule},
		&returns,
	)

	return returns
}

// Configuration properties of an action for CloudWatch Logs.
// Experimental.
type CloudWatchLogsActionProps struct {
	// The IAM role that allows access to the CloudWatch log group.
	// Experimental.
	Role awsiam.IRole `json:"role"`
}

// The action to invoke an AWS Lambda function, passing in an MQTT message.
// Experimental.
type LambdaFunctionAction interface {
	awsiot.IAction
	Bind(topicRule awsiot.ITopicRule) *awsiot.ActionConfig
}

// The jsii proxy struct for LambdaFunctionAction
type jsiiProxy_LambdaFunctionAction struct {
	internal.Type__awsiotIAction
}

// Experimental.
func NewLambdaFunctionAction(func_ awslambda.IFunction) LambdaFunctionAction {
	_init_.Initialize()

	j := jsiiProxy_LambdaFunctionAction{}

	_jsii_.Create(
		"monocdk.aws_iot_actions.LambdaFunctionAction",
		[]interface{}{func_},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunctionAction_Override(l LambdaFunctionAction, func_ awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iot_actions.LambdaFunctionAction",
		[]interface{}{func_},
		l,
	)
}

// Returns the topic rule action specification.
// Experimental.
func (l *jsiiProxy_LambdaFunctionAction) Bind(topicRule awsiot.ITopicRule) *awsiot.ActionConfig {
	var returns *awsiot.ActionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{topicRule},
		&returns,
	)

	return returns
}


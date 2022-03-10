package awsioteventsactions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiotevents"
	"github.com/aws/aws-cdk-go/awscdk/awsioteventsactions/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
)

// The action to write the data to an AWS Lambda function.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaInvokeAction interface {
	awsiotevents.IAction
	Bind(_scope constructs.Construct, options *awsiotevents.ActionBindOptions) *awsiotevents.ActionConfig
}

// The jsii proxy struct for LambdaInvokeAction
type jsiiProxy_LambdaInvokeAction struct {
	internal.Type__awsioteventsIAction
}

// Experimental.
func NewLambdaInvokeAction(func_ awslambda.IFunction) LambdaInvokeAction {
	_init_.Initialize()

	j := jsiiProxy_LambdaInvokeAction{}

	_jsii_.Create(
		"monocdk.aws_iotevents_actions.LambdaInvokeAction",
		[]interface{}{func_},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaInvokeAction_Override(l LambdaInvokeAction, func_ awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_iotevents_actions.LambdaInvokeAction",
		[]interface{}{func_},
		l,
	)
}

// Returns the AWS IoT Events action specification.
// Experimental.
func (l *jsiiProxy_LambdaInvokeAction) Bind(_scope constructs.Construct, options *awsiotevents.ActionBindOptions) *awsiotevents.ActionConfig {
	var returns *awsiotevents.ActionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}


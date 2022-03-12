// Receipt Detector Model actions for AWS IoT Events
package awscdkioteventsactionsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkioteventsactionsalpha/v2/internal"
	"github.com/aws/aws-cdk-go/awscdkioteventsalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// The action to write the data to an AWS Lambda function.
//
// TODO: EXAMPLE
//
// Experimental.
type LambdaInvokeAction interface {
	awscdkioteventsalpha.IAction
	Bind(_scope constructs.Construct, options *awscdkioteventsalpha.ActionBindOptions) *awscdkioteventsalpha.ActionConfig
}

// The jsii proxy struct for LambdaInvokeAction
type jsiiProxy_LambdaInvokeAction struct {
	internal.Type__awscdkioteventsalphaIAction
}

// Experimental.
func NewLambdaInvokeAction(func_ awslambda.IFunction) LambdaInvokeAction {
	_init_.Initialize()

	j := jsiiProxy_LambdaInvokeAction{}

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.LambdaInvokeAction",
		[]interface{}{func_},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaInvokeAction_Override(l LambdaInvokeAction, func_ awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-iotevents-actions-alpha.LambdaInvokeAction",
		[]interface{}{func_},
		l,
	)
}

// Returns the AWS IoT Events action specification.
// Experimental.
func (l *jsiiProxy_LambdaInvokeAction) Bind(_scope constructs.Construct, options *awscdkioteventsalpha.ActionBindOptions) *awscdkioteventsalpha.ActionConfig {
	var returns *awscdkioteventsalpha.ActionConfig

	_jsii_.Invoke(
		l,
		"bind",
		[]interface{}{_scope, options},
		&returns,
	)

	return returns
}


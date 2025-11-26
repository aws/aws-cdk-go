package previewawsstepfunctionsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsstepfunctions/previewawsstepfunctionsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provisions a state machine.
//
// A state machine consists of a collection of states that can do work ( `Task` states), determine to which states to transition next ( `Choice` states), stop an execution with an error ( `Fail` states), and so on. State machines are specified using a JSON-based, structured language.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnStateMachineLogsMixin := awscdkmixinspreview.Mixins.NewCfnStateMachineLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html
//
type CfnStateMachineLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnStateMachineLogsMixin
type jsiiProxy_CfnStateMachineLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnStateMachineLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnStateMachineLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::StepFunctions::StateMachine`.
func NewCfnStateMachineLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnStateMachineLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnStateMachineLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnStateMachineLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::StepFunctions::StateMachine`.
func NewCfnStateMachineLogsMixin_Override(c CfnStateMachineLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnStateMachineLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnStateMachineLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnStateMachineLogsMixin_EXPRESS_LOGS() CfnStateMachineExpressLogs {
	_init_.Initialize()
	var returns CfnStateMachineExpressLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineLogsMixin",
		"EXPRESS_LOGS",
		&returns,
	)
	return returns
}

func CfnStateMachineLogsMixin_STANDARD_LOGS() CfnStateMachineStandardLogs {
	_init_.Initialize()
	var returns CfnStateMachineStandardLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineLogsMixin",
		"STANDARD_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnStateMachineLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnStateMachineLogsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}


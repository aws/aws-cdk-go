package previewawspipesmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspipes/previewawspipesmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a pipe.
//
// Amazon EventBridge Pipes connect event sources to targets and reduces the need for specialized knowledge and integration code.
//
// > As an aid to help you jumpstart developing CloudFormation templates, the EventBridge console enables you to create templates from the existing pipes in your account. For more information, see [Generate an CloudFormation template from EventBridge Pipes](https://docs.aws.amazon.com/eventbridge/latest/userguide/pipes-generate-template.html) in the *Amazon EventBridge User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnPipeLogsMixin := awscdkmixinspreview.Mixins.NewCfnPipeLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pipes-pipe.html
//
type CfnPipeLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPipeLogsMixin
type jsiiProxy_CfnPipeLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPipeLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipeLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::Pipes::Pipe`.
func NewCfnPipeLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnPipeLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnPipeLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPipeLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::Pipes::Pipe`.
func NewCfnPipeLogsMixin_Override(c CfnPipeLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPipeLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipeLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipeLogsMixin_EXECUTION_LOGS() CfnPipeExecutionLogs {
	_init_.Initialize()
	var returns CfnPipeExecutionLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeLogsMixin",
		"EXECUTION_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipeLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPipeLogsMixin) Supports(construct constructs.IConstruct) *bool {
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


package previewawselementalinferencemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawselementalinference/previewawselementalinferencemixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a feed that receives media for inference processing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var logsDelivery ILogsDelivery
//
//   cfnFeedLogsMixin := awscdkmixinspreview.Mixins.NewCfnFeedLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elementalinference-feed.html
//
type CfnFeedLogsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFeedLogsMixin
type jsiiProxy_CfnFeedLogsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFeedLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFeedLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::ElementalInference::Feed`.
func NewCfnFeedLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnFeedLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnFeedLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFeedLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::ElementalInference::Feed`.
func NewCfnFeedLogsMixin_Override(c CfnFeedLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFeedLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFeedLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFeedLogsMixin_APPLICATION_LOGS() CfnFeedApplicationLogs {
	_init_.Initialize()
	var returns CfnFeedApplicationLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_elementalinference.mixins.CfnFeedLogsMixin",
		"APPLICATION_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFeedLogsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFeedLogsMixin) Supports(construct constructs.IConstruct) *bool {
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


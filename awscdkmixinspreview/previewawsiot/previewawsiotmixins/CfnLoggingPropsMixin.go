package previewawsiotmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiot/previewawsiotmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Configure logging.
//
// > If you already set the log function of AWS IoT Core , you can't deploy the Cloud Development Kit  to change the logging settings. You can change the logging settings by either:
// >
// > - Importing the existing logging resource into your CloudFormation stack, such as with the [infrastructure as code generator (IaC generator)](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/generate-IaC.html) .
// > - Calling `aws iot set-v2-logging-options --disable-all-logs` before creating a new CloudFormation stack. This command disables all AWS IoT logging. As a result, no AWS IoT logs will be delivered to Amazon CloudWatch until you re-enable logging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLoggingPropsMixin := awscdkmixinspreview.Mixins.NewCfnLoggingPropsMixin(&CfnLoggingMixinProps{
//   	AccountId: jsii.String("accountId"),
//   	DefaultLogLevel: jsii.String("defaultLogLevel"),
//   	RoleArn: jsii.String("roleArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-logging.html
//
type CfnLoggingPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLoggingMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLoggingPropsMixin
type jsiiProxy_CfnLoggingPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLoggingPropsMixin) Props() *CfnLoggingMixinProps {
	var returns *CfnLoggingMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLoggingPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::Logging`.
func NewCfnLoggingPropsMixin(props *CfnLoggingMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLoggingPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLoggingPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLoggingPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnLoggingPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::Logging`.
func NewCfnLoggingPropsMixin_Override(c CfnLoggingPropsMixin, props *CfnLoggingMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnLoggingPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLoggingPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLoggingPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnLoggingPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLoggingPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnLoggingPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLoggingPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnLoggingPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


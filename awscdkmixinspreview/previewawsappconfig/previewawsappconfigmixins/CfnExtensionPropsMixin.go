package previewawsappconfigmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappconfig/previewawsappconfigmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an AWS AppConfig extension.
//
// An extension augments your ability to inject logic or behavior at different points during the AWS AppConfig workflow of creating or deploying a configuration.
//
// You can create your own extensions or use the AWS authored extensions provided by AWS AppConfig . For an AWS AppConfig extension that uses AWS Lambda , you must create a Lambda function to perform any computation and processing defined in the extension. If you plan to create custom versions of the AWS authored notification extensions, you only need to specify an Amazon Resource Name (ARN) in the `Uri` field for the new extension version.
//
// - For a custom EventBridge notification extension, enter the ARN of the EventBridge default events in the `Uri` field.
// - For a custom Amazon SNS notification extension, enter the ARN of an Amazon SNS topic in the `Uri` field.
// - For a custom Amazon SQS notification extension, enter the ARN of an Amazon SQS message queue in the `Uri` field.
//
// For more information about extensions, see [Extending workflows](https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html) in the *AWS AppConfig User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var actions interface{}
//
//   cfnExtensionPropsMixin := awscdkmixinspreview.Mixins.NewCfnExtensionPropsMixin(&CfnExtensionMixinProps{
//   	Actions: actions,
//   	Description: jsii.String("description"),
//   	LatestVersionNumber: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Parameters: map[string]interface{}{
//   		"parametersKey": &ParameterProperty{
//   			"description": jsii.String("description"),
//   			"dynamic": jsii.Boolean(false),
//   			"required": jsii.Boolean(false),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appconfig-extension.html
//
type CfnExtensionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnExtensionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnExtensionPropsMixin
type jsiiProxy_CfnExtensionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnExtensionPropsMixin) Props() *CfnExtensionMixinProps {
	var returns *CfnExtensionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnExtensionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppConfig::Extension`.
func NewCfnExtensionPropsMixin(props *CfnExtensionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnExtensionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnExtensionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnExtensionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnExtensionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppConfig::Extension`.
func NewCfnExtensionPropsMixin_Override(c CfnExtensionPropsMixin, props *CfnExtensionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnExtensionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnExtensionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnExtensionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnExtensionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnExtensionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appconfig.mixins.CfnExtensionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnExtensionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnExtensionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


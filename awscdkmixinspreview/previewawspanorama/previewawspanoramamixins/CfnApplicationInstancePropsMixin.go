package previewawspanoramamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawspanorama/previewawspanoramamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// > End of support notice: On May 31, 2026, AWS will end support for AWS Panorama .
//
// After May 31, 2026,
// > you will no longer be able to access the AWS Panorama console or AWS Panorama resources. For more information, see [AWS Panorama end of support](https://docs.aws.amazon.com/panorama/latest/dev/panorama-end-of-support.html) .
//
// Creates an application instance and deploys it to a device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationInstancePropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationInstancePropsMixin(&CfnApplicationInstanceMixinProps{
//   	ApplicationInstanceIdToReplace: jsii.String("applicationInstanceIdToReplace"),
//   	DefaultRuntimeContextDevice: jsii.String("defaultRuntimeContextDevice"),
//   	Description: jsii.String("description"),
//   	ManifestOverridesPayload: &ManifestOverridesPayloadProperty{
//   		PayloadData: jsii.String("payloadData"),
//   	},
//   	ManifestPayload: &ManifestPayloadProperty{
//   		PayloadData: jsii.String("payloadData"),
//   	},
//   	Name: jsii.String("name"),
//   	RuntimeRoleArn: jsii.String("runtimeRoleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-panorama-applicationinstance.html
//
type CfnApplicationInstancePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationInstanceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationInstancePropsMixin
type jsiiProxy_CfnApplicationInstancePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationInstancePropsMixin) Props() *CfnApplicationInstanceMixinProps {
	var returns *CfnApplicationInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationInstancePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Panorama::ApplicationInstance`.
func NewCfnApplicationInstancePropsMixin(props *CfnApplicationInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_panorama.mixins.CfnApplicationInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Panorama::ApplicationInstance`.
func NewCfnApplicationInstancePropsMixin_Override(c CfnApplicationInstancePropsMixin, props *CfnApplicationInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_panorama.mixins.CfnApplicationInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_panorama.mixins.CfnApplicationInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_panorama.mixins.CfnApplicationInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationInstancePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnApplicationInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


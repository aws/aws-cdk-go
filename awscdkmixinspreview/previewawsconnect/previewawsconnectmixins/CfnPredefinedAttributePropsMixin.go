package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconnect/previewawsconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Textual or numeric value that describes an attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPredefinedAttributePropsMixin := awscdkmixinspreview.Mixins.NewCfnPredefinedAttributePropsMixin(&CfnPredefinedAttributeMixinProps{
//   	AttributeConfiguration: &AttributeConfigurationProperty{
//   		EnableValueValidationOnAssociation: jsii.Boolean(false),
//   		IsReadOnly: jsii.Boolean(false),
//   	},
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	Purposes: []*string{
//   		jsii.String("purposes"),
//   	},
//   	Values: &ValuesProperty{
//   		StringList: []*string{
//   			jsii.String("stringList"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-predefinedattribute.html
//
type CfnPredefinedAttributePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPredefinedAttributeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPredefinedAttributePropsMixin
type jsiiProxy_CfnPredefinedAttributePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPredefinedAttributePropsMixin) Props() *CfnPredefinedAttributeMixinProps {
	var returns *CfnPredefinedAttributeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPredefinedAttributePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::PredefinedAttribute`.
func NewCfnPredefinedAttributePropsMixin(props *CfnPredefinedAttributeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPredefinedAttributePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPredefinedAttributePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPredefinedAttributePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnPredefinedAttributePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::PredefinedAttribute`.
func NewCfnPredefinedAttributePropsMixin_Override(c CfnPredefinedAttributePropsMixin, props *CfnPredefinedAttributeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnPredefinedAttributePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPredefinedAttributePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPredefinedAttributePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnPredefinedAttributePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPredefinedAttributePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnPredefinedAttributePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPredefinedAttributePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPredefinedAttributePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package awstranslate

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awstranslate/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A custom terminology resource for Amazon Translate that enables customized translations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTerminologyPropsMixin := awscdkcfnpropertymixins.Aws_translate.NewCfnTerminologyPropsMixin(&CfnTerminologyMixinProps{
//   	Description: jsii.String("description"),
//   	EncryptionKey: &EncryptionKeyProperty{
//   		Id: jsii.String("id"),
//   		Type: jsii.String("type"),
//   	},
//   	MergeStrategy: jsii.String("mergeStrategy"),
//   	Name: jsii.String("name"),
//   	Tags: []TagsItemsProperty{
//   		&TagsItemsProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TerminologyData: &TerminologyDataProperty{
//   		Directionality: jsii.String("directionality"),
//   		File: jsii.String("file"),
//   		Format: jsii.String("format"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-translate-terminology.html
//
type CfnTerminologyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTerminologyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTerminologyPropsMixin
type jsiiProxy_CfnTerminologyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTerminologyPropsMixin) Props() *CfnTerminologyMixinProps {
	var returns *CfnTerminologyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTerminologyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Translate::Terminology`.
func NewCfnTerminologyPropsMixin(props *CfnTerminologyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTerminologyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTerminologyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTerminologyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_translate.CfnTerminologyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Translate::Terminology`.
func NewCfnTerminologyPropsMixin_Override(c CfnTerminologyPropsMixin, props *CfnTerminologyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_translate.CfnTerminologyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTerminologyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTerminologyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_translate.CfnTerminologyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTerminologyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_translate.CfnTerminologyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTerminologyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTerminologyPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


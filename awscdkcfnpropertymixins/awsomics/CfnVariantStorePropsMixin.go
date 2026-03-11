package awsomics

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a store for variant data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnVariantStorePropsMixin := awscdkcfnpropertymixins.Aws_omics.NewCfnVariantStorePropsMixin(&CfnVariantStoreMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Reference: &ReferenceItemProperty{
//   		ReferenceArn: jsii.String("referenceArn"),
//   	},
//   	SseConfig: &SseConfigProperty{
//   		KeyArn: jsii.String("keyArn"),
//   		Type: jsii.String("type"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-variantstore.html
//
type CfnVariantStorePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnVariantStoreMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVariantStorePropsMixin
type jsiiProxy_CfnVariantStorePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnVariantStorePropsMixin) Props() *CfnVariantStoreMixinProps {
	var returns *CfnVariantStoreMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVariantStorePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Omics::VariantStore`.
func NewCfnVariantStorePropsMixin(props *CfnVariantStoreMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnVariantStorePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVariantStorePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVariantStorePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_omics.CfnVariantStorePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Omics::VariantStore`.
func NewCfnVariantStorePropsMixin_Override(c CfnVariantStorePropsMixin, props *CfnVariantStoreMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_omics.CfnVariantStorePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnVariantStorePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVariantStorePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_omics.CfnVariantStorePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVariantStorePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_omics.CfnVariantStorePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVariantStorePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnVariantStorePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


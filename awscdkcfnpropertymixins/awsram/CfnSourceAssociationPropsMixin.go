package awsram

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsram/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates a specified source account with a resource share.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSourceAssociationPropsMixin := awscdkcfnpropertymixins.Aws_ram.NewCfnSourceAssociationPropsMixin(&CfnSourceAssociationMixinProps{
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//   	SourceId: jsii.String("sourceId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-sourceassociation.html
//
type CfnSourceAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSourceAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSourceAssociationPropsMixin
type jsiiProxy_CfnSourceAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSourceAssociationPropsMixin) Props() *CfnSourceAssociationMixinProps {
	var returns *CfnSourceAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RAM::SourceAssociation`.
func NewCfnSourceAssociationPropsMixin(props *CfnSourceAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSourceAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSourceAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSourceAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ram.CfnSourceAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RAM::SourceAssociation`.
func NewCfnSourceAssociationPropsMixin_Override(c CfnSourceAssociationPropsMixin, props *CfnSourceAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ram.CfnSourceAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSourceAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSourceAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ram.CfnSourceAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSourceAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ram.CfnSourceAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSourceAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSourceAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


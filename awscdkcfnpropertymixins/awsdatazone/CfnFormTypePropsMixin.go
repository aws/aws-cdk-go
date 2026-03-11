package awsdatazone

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The details of the metadata form type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnFormTypePropsMixin := awscdkcfnpropertymixins.Aws_datazone.NewCfnFormTypePropsMixin(&CfnFormTypeMixinProps{
//   	Description: jsii.String("description"),
//   	DomainIdentifier: jsii.String("domainIdentifier"),
//   	Model: &ModelProperty{
//   		Smithy: jsii.String("smithy"),
//   	},
//   	Name: jsii.String("name"),
//   	OwningProjectIdentifier: jsii.String("owningProjectIdentifier"),
//   	Status: jsii.String("status"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datazone-formtype.html
//
type CfnFormTypePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnFormTypeMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFormTypePropsMixin
type jsiiProxy_CfnFormTypePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnFormTypePropsMixin) Props() *CfnFormTypeMixinProps {
	var returns *CfnFormTypeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFormTypePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DataZone::FormType`.
func NewCfnFormTypePropsMixin(props *CfnFormTypeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnFormTypePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFormTypePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFormTypePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnFormTypePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DataZone::FormType`.
func NewCfnFormTypePropsMixin_Override(c CfnFormTypePropsMixin, props *CfnFormTypeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnFormTypePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnFormTypePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFormTypePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnFormTypePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFormTypePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_datazone.CfnFormTypePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFormTypePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFormTypePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


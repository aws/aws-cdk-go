package previewawsservicecatalogmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsservicecatalog/previewawsservicecatalogmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a `RESOURCE_UPDATE` constraint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceUpdateConstraintPropsMixin := awscdkmixinspreview.Mixins.NewCfnResourceUpdateConstraintPropsMixin(&CfnResourceUpdateConstraintMixinProps{
//   	AcceptLanguage: jsii.String("acceptLanguage"),
//   	Description: jsii.String("description"),
//   	PortfolioId: jsii.String("portfolioId"),
//   	ProductId: jsii.String("productId"),
//   	TagUpdateOnProvisionedProduct: jsii.String("tagUpdateOnProvisionedProduct"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-resourceupdateconstraint.html
//
type CfnResourceUpdateConstraintPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResourceUpdateConstraintMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResourceUpdateConstraintPropsMixin
type jsiiProxy_CfnResourceUpdateConstraintPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResourceUpdateConstraintPropsMixin) Props() *CfnResourceUpdateConstraintMixinProps {
	var returns *CfnResourceUpdateConstraintMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceUpdateConstraintPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ServiceCatalog::ResourceUpdateConstraint`.
func NewCfnResourceUpdateConstraintPropsMixin(props *CfnResourceUpdateConstraintMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResourceUpdateConstraintPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResourceUpdateConstraintPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResourceUpdateConstraintPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_servicecatalog.mixins.CfnResourceUpdateConstraintPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ServiceCatalog::ResourceUpdateConstraint`.
func NewCfnResourceUpdateConstraintPropsMixin_Override(c CfnResourceUpdateConstraintPropsMixin, props *CfnResourceUpdateConstraintMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_servicecatalog.mixins.CfnResourceUpdateConstraintPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResourceUpdateConstraintPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResourceUpdateConstraintPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_servicecatalog.mixins.CfnResourceUpdateConstraintPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceUpdateConstraintPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_servicecatalog.mixins.CfnResourceUpdateConstraintPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceUpdateConstraintPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnResourceUpdateConstraintPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


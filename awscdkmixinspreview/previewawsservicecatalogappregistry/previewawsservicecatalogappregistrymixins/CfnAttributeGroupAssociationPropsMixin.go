package previewawsservicecatalogappregistrymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsservicecatalogappregistry/previewawsservicecatalogappregistrymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates an attribute group with an application to augment the application's metadata with the group's attributes.
//
// This feature enables applications to be described with user-defined details that are machine-readable, such as third-party integrations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAttributeGroupAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnAttributeGroupAssociationPropsMixin(&CfnAttributeGroupAssociationMixinProps{
//   	Application: jsii.String("application"),
//   	AttributeGroup: jsii.String("attributeGroup"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroupassociation.html
//
type CfnAttributeGroupAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAttributeGroupAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAttributeGroupAssociationPropsMixin
type jsiiProxy_CfnAttributeGroupAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAttributeGroupAssociationPropsMixin) Props() *CfnAttributeGroupAssociationMixinProps {
	var returns *CfnAttributeGroupAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAttributeGroupAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation`.
func NewCfnAttributeGroupAssociationPropsMixin(props *CfnAttributeGroupAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAttributeGroupAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAttributeGroupAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAttributeGroupAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_servicecatalogappregistry.mixins.CfnAttributeGroupAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation`.
func NewCfnAttributeGroupAssociationPropsMixin_Override(c CfnAttributeGroupAssociationPropsMixin, props *CfnAttributeGroupAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_servicecatalogappregistry.mixins.CfnAttributeGroupAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAttributeGroupAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAttributeGroupAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_servicecatalogappregistry.mixins.CfnAttributeGroupAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAttributeGroupAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_servicecatalogappregistry.mixins.CfnAttributeGroupAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAttributeGroupAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAttributeGroupAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


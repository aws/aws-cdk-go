package previewawsappstreammixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappstream/previewawsappstreammixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates an application to an entitlement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationEntitlementAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnApplicationEntitlementAssociationPropsMixin(&CfnApplicationEntitlementAssociationMixinProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EntitlementName: jsii.String("entitlementName"),
//   	StackName: jsii.String("stackName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationentitlementassociation.html
//
type CfnApplicationEntitlementAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnApplicationEntitlementAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationEntitlementAssociationPropsMixin
type jsiiProxy_CfnApplicationEntitlementAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnApplicationEntitlementAssociationPropsMixin) Props() *CfnApplicationEntitlementAssociationMixinProps {
	var returns *CfnApplicationEntitlementAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnApplicationEntitlementAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppStream::ApplicationEntitlementAssociation`.
func NewCfnApplicationEntitlementAssociationPropsMixin(props *CfnApplicationEntitlementAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnApplicationEntitlementAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationEntitlementAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationEntitlementAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnApplicationEntitlementAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppStream::ApplicationEntitlementAssociation`.
func NewCfnApplicationEntitlementAssociationPropsMixin_Override(c CfnApplicationEntitlementAssociationPropsMixin, props *CfnApplicationEntitlementAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnApplicationEntitlementAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnApplicationEntitlementAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationEntitlementAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnApplicationEntitlementAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnApplicationEntitlementAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appstream.mixins.CfnApplicationEntitlementAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationEntitlementAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnApplicationEntitlementAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


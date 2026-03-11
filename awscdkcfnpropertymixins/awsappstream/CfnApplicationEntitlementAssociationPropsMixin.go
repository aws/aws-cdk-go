package awsappstream

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates an application to an entitlement.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnApplicationEntitlementAssociationPropsMixin := awscdkcfnpropertymixins.Aws_appstream.NewCfnApplicationEntitlementAssociationPropsMixin(&CfnApplicationEntitlementAssociationMixinProps{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EntitlementName: jsii.String("entitlementName"),
//   	StackName: jsii.String("stackName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-applicationentitlementassociation.html
//
type CfnApplicationEntitlementAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnApplicationEntitlementAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnApplicationEntitlementAssociationPropsMixin
type jsiiProxy_CfnApplicationEntitlementAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnApplicationEntitlementAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppStream::ApplicationEntitlementAssociation`.
func NewCfnApplicationEntitlementAssociationPropsMixin(props *CfnApplicationEntitlementAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnApplicationEntitlementAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnApplicationEntitlementAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnApplicationEntitlementAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appstream.CfnApplicationEntitlementAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppStream::ApplicationEntitlementAssociation`.
func NewCfnApplicationEntitlementAssociationPropsMixin_Override(c CfnApplicationEntitlementAssociationPropsMixin, props *CfnApplicationEntitlementAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_appstream.CfnApplicationEntitlementAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnApplicationEntitlementAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnApplicationEntitlementAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_appstream.CfnApplicationEntitlementAssociationPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_appstream.CfnApplicationEntitlementAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnApplicationEntitlementAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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


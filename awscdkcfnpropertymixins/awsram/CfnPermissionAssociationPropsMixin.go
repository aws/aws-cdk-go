package awsram

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsram/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates a specified AWS RAM permission with a resource share.
//
// You can only associate one permission with each resource type in a resource share.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPermissionAssociationPropsMixin := awscdkcfnpropertymixins.Aws_ram.NewCfnPermissionAssociationPropsMixin(&CfnPermissionAssociationMixinProps{
//   	PermissionArn: jsii.String("permissionArn"),
//   	Replace: jsii.Boolean(false),
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-permissionassociation.html
//
type CfnPermissionAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPermissionAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPermissionAssociationPropsMixin
type jsiiProxy_CfnPermissionAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPermissionAssociationPropsMixin) Props() *CfnPermissionAssociationMixinProps {
	var returns *CfnPermissionAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissionAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RAM::PermissionAssociation`.
func NewCfnPermissionAssociationPropsMixin(props *CfnPermissionAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPermissionAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPermissionAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPermissionAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ram.CfnPermissionAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RAM::PermissionAssociation`.
func NewCfnPermissionAssociationPropsMixin_Override(c CfnPermissionAssociationPropsMixin, props *CfnPermissionAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ram.CfnPermissionAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPermissionAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPermissionAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ram.CfnPermissionAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPermissionAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ram.CfnPermissionAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPermissionAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPermissionAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


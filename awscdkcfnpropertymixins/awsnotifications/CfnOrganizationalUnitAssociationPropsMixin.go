package awsnotifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Notifications::OrganizationalUnitAssociation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnOrganizationalUnitAssociationPropsMixin := awscdkcfnpropertymixins.Aws_notifications.NewCfnOrganizationalUnitAssociationPropsMixin(&CfnOrganizationalUnitAssociationMixinProps{
//   	NotificationConfigurationArn: jsii.String("notificationConfigurationArn"),
//   	OrganizationalUnitId: jsii.String("organizationalUnitId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-organizationalunitassociation.html
//
type CfnOrganizationalUnitAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnOrganizationalUnitAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOrganizationalUnitAssociationPropsMixin
type jsiiProxy_CfnOrganizationalUnitAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnOrganizationalUnitAssociationPropsMixin) Props() *CfnOrganizationalUnitAssociationMixinProps {
	var returns *CfnOrganizationalUnitAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOrganizationalUnitAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Notifications::OrganizationalUnitAssociation`.
func NewCfnOrganizationalUnitAssociationPropsMixin(props *CfnOrganizationalUnitAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnOrganizationalUnitAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOrganizationalUnitAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOrganizationalUnitAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_notifications.CfnOrganizationalUnitAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Notifications::OrganizationalUnitAssociation`.
func NewCfnOrganizationalUnitAssociationPropsMixin_Override(c CfnOrganizationalUnitAssociationPropsMixin, props *CfnOrganizationalUnitAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_notifications.CfnOrganizationalUnitAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnOrganizationalUnitAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOrganizationalUnitAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_notifications.CfnOrganizationalUnitAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOrganizationalUnitAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_notifications.CfnOrganizationalUnitAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOrganizationalUnitAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnOrganizationalUnitAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


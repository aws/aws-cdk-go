package mixinsawsnotifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsnotifications/mixinsawsnotifications/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates an Account Management Contact with a `ManagedNotificationConfiguration` for AWS User Notifications .
//
// For more information about AWS User Notifications , see the [AWS User Notifications User Guide](https://docs.aws.amazon.com/notifications/latest/userguide/what-is-service.html) . For more information about Account Management Contacts, see the [Account Management Reference Guide](https://docs.aws.amazon.com/accounts/latest/reference/API_AlternateContact.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnManagedNotificationAccountContactAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnManagedNotificationAccountContactAssociationPropsMixin(&CfnManagedNotificationAccountContactAssociationMixinProps{
//   	ContactIdentifier: jsii.String("contactIdentifier"),
//   	ManagedNotificationConfigurationArn: jsii.String("managedNotificationConfigurationArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-managednotificationaccountcontactassociation.html
//
type CfnManagedNotificationAccountContactAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnManagedNotificationAccountContactAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnManagedNotificationAccountContactAssociationPropsMixin
type jsiiProxy_CfnManagedNotificationAccountContactAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnManagedNotificationAccountContactAssociationPropsMixin) Props() *CfnManagedNotificationAccountContactAssociationMixinProps {
	var returns *CfnManagedNotificationAccountContactAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnManagedNotificationAccountContactAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Notifications::ManagedNotificationAccountContactAssociation`.
func NewCfnManagedNotificationAccountContactAssociationPropsMixin(props *CfnManagedNotificationAccountContactAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnManagedNotificationAccountContactAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnManagedNotificationAccountContactAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnManagedNotificationAccountContactAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_notifications.mixins.CfnManagedNotificationAccountContactAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Notifications::ManagedNotificationAccountContactAssociation`.
func NewCfnManagedNotificationAccountContactAssociationPropsMixin_Override(c CfnManagedNotificationAccountContactAssociationPropsMixin, props *CfnManagedNotificationAccountContactAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_notifications.mixins.CfnManagedNotificationAccountContactAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnManagedNotificationAccountContactAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnManagedNotificationAccountContactAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_notifications.mixins.CfnManagedNotificationAccountContactAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnManagedNotificationAccountContactAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_notifications.mixins.CfnManagedNotificationAccountContactAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnManagedNotificationAccountContactAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnManagedNotificationAccountContactAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


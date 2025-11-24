package mixinsawsnotifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsnotifications/mixinsawsnotifications/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates a `Channel` with a `ManagedNotificationAdditionalChannelAssociation` for AWS User Notifications .
//
// For more information about AWS User Notifications , see the [AWS User Notifications User Guide](https://docs.aws.amazon.com/notifications/latest/userguide/what-is-service.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnManagedNotificationAdditionalChannelAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnManagedNotificationAdditionalChannelAssociationPropsMixin(&CfnManagedNotificationAdditionalChannelAssociationMixinProps{
//   	ChannelArn: jsii.String("channelArn"),
//   	ManagedNotificationConfigurationArn: jsii.String("managedNotificationConfigurationArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-managednotificationadditionalchannelassociation.html
//
type CfnManagedNotificationAdditionalChannelAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnManagedNotificationAdditionalChannelAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnManagedNotificationAdditionalChannelAssociationPropsMixin
type jsiiProxy_CfnManagedNotificationAdditionalChannelAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnManagedNotificationAdditionalChannelAssociationPropsMixin) Props() *CfnManagedNotificationAdditionalChannelAssociationMixinProps {
	var returns *CfnManagedNotificationAdditionalChannelAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnManagedNotificationAdditionalChannelAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Notifications::ManagedNotificationAdditionalChannelAssociation`.
func NewCfnManagedNotificationAdditionalChannelAssociationPropsMixin(props *CfnManagedNotificationAdditionalChannelAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnManagedNotificationAdditionalChannelAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnManagedNotificationAdditionalChannelAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnManagedNotificationAdditionalChannelAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_notifications.mixins.CfnManagedNotificationAdditionalChannelAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Notifications::ManagedNotificationAdditionalChannelAssociation`.
func NewCfnManagedNotificationAdditionalChannelAssociationPropsMixin_Override(c CfnManagedNotificationAdditionalChannelAssociationPropsMixin, props *CfnManagedNotificationAdditionalChannelAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_notifications.mixins.CfnManagedNotificationAdditionalChannelAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnManagedNotificationAdditionalChannelAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnManagedNotificationAdditionalChannelAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_notifications.mixins.CfnManagedNotificationAdditionalChannelAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnManagedNotificationAdditionalChannelAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_notifications.mixins.CfnManagedNotificationAdditionalChannelAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnManagedNotificationAdditionalChannelAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnManagedNotificationAdditionalChannelAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


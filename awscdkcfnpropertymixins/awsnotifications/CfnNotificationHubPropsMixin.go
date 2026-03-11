package awsnotifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Configures a `NotificationHub` for AWS User Notifications .
//
// For more information about notification hub, see the [AWS User Notifications User Guide](https://docs.aws.amazon.com/notifications/latest/userguide/notification-hubs.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnNotificationHubPropsMixin := awscdkcfnpropertymixins.Aws_notifications.NewCfnNotificationHubPropsMixin(&CfnNotificationHubMixinProps{
//   	Region: jsii.String("region"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-notifications-notificationhub.html
//
type CfnNotificationHubPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnNotificationHubMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNotificationHubPropsMixin
type jsiiProxy_CfnNotificationHubPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnNotificationHubPropsMixin) Props() *CfnNotificationHubMixinProps {
	var returns *CfnNotificationHubMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationHubPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Notifications::NotificationHub`.
func NewCfnNotificationHubPropsMixin(props *CfnNotificationHubMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnNotificationHubPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNotificationHubPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNotificationHubPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_notifications.CfnNotificationHubPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Notifications::NotificationHub`.
func NewCfnNotificationHubPropsMixin_Override(c CfnNotificationHubPropsMixin, props *CfnNotificationHubMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_notifications.CfnNotificationHubPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnNotificationHubPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNotificationHubPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_notifications.CfnNotificationHubPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNotificationHubPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_notifications.CfnNotificationHubPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNotificationHubPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnNotificationHubPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


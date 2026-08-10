package awsusernotifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsusernotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for AWS User Notifications ManagedNotificationConfiguration.
//
// This is a read-only resource representing AWS-managed notification configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnManagedNotificationConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_usernotifications.NewCfnManagedNotificationConfigurationPropsMixin(&CfnManagedNotificationConfigurationMixinProps{
//   	Category: jsii.String("category"),
//   	SubCategory: jsii.String("subCategory"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-usernotifications-managednotificationconfiguration.html
//
type CfnManagedNotificationConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnManagedNotificationConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnManagedNotificationConfigurationPropsMixin
type jsiiProxy_CfnManagedNotificationConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnManagedNotificationConfigurationPropsMixin) Props() *CfnManagedNotificationConfigurationMixinProps {
	var returns *CfnManagedNotificationConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnManagedNotificationConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::UserNotifications::ManagedNotificationConfiguration`.
func NewCfnManagedNotificationConfigurationPropsMixin(props *CfnManagedNotificationConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnManagedNotificationConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnManagedNotificationConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnManagedNotificationConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_usernotifications.CfnManagedNotificationConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::UserNotifications::ManagedNotificationConfiguration`.
func NewCfnManagedNotificationConfigurationPropsMixin_Override(c CfnManagedNotificationConfigurationPropsMixin, props *CfnManagedNotificationConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_usernotifications.CfnManagedNotificationConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnManagedNotificationConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnManagedNotificationConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_usernotifications.CfnManagedNotificationConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnManagedNotificationConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_usernotifications.CfnManagedNotificationConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnManagedNotificationConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnManagedNotificationConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package mixinsawscodestarnotifications

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awscodestarnotifications/mixinsawscodestarnotifications/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a notification rule for a resource.
//
// The rule specifies the events you want notifications about and the targets (such as Amazon Simple Notification Service topics or  clients configured for Slack) where you want to receive them.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNotificationRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnNotificationRulePropsMixin(&CfnNotificationRuleMixinProps{
//   	CreatedBy: jsii.String("createdBy"),
//   	DetailType: jsii.String("detailType"),
//   	EventTypeId: jsii.String("eventTypeId"),
//   	EventTypeIds: []*string{
//   		jsii.String("eventTypeIds"),
//   	},
//   	Name: jsii.String("name"),
//   	Resource: jsii.String("resource"),
//   	Status: jsii.String("status"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TargetAddress: jsii.String("targetAddress"),
//   	Targets: []interface{}{
//   		&TargetProperty{
//   			TargetAddress: jsii.String("targetAddress"),
//   			TargetType: jsii.String("targetType"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html
//
type CfnNotificationRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNotificationRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNotificationRulePropsMixin
type jsiiProxy_CfnNotificationRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNotificationRulePropsMixin) Props() *CfnNotificationRuleMixinProps {
	var returns *CfnNotificationRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNotificationRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeStarNotifications::NotificationRule`.
func NewCfnNotificationRulePropsMixin(props *CfnNotificationRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNotificationRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNotificationRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNotificationRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codestarnotifications.mixins.CfnNotificationRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeStarNotifications::NotificationRule`.
func NewCfnNotificationRulePropsMixin_Override(c CfnNotificationRulePropsMixin, props *CfnNotificationRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codestarnotifications.mixins.CfnNotificationRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNotificationRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNotificationRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codestarnotifications.mixins.CfnNotificationRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNotificationRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codestarnotifications.mixins.CfnNotificationRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNotificationRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnNotificationRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


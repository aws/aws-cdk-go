package previewawsautoscalingmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsautoscaling/previewawsautoscalingmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::AutoScaling::LifecycleHook` resource specifies lifecycle hooks for an Auto Scaling group.
//
// These hooks let you create solutions that are aware of events in the Auto Scaling instance lifecycle, and then perform a custom action on instances when the corresponding lifecycle event occurs. A lifecycle hook provides a specified amount of time (one hour by default) to wait for the action to complete before the instance transitions to the next state.
//
// Use lifecycle hooks to prepare new instances for use or to delay them from being registered behind a load balancer before their configuration has been applied completely. You can also use lifecycle hooks to prepare running instances to be terminated by, for example, downloading logs or other data.
//
// For more information, see [Amazon EC2 Auto Scaling lifecycle hooks](https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html) in the *Amazon EC2 Auto Scaling User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLifecycleHookPropsMixin := awscdkmixinspreview.Mixins.NewCfnLifecycleHookPropsMixin(&CfnLifecycleHookMixinProps{
//   	AutoScalingGroupName: jsii.String("autoScalingGroupName"),
//   	DefaultResult: jsii.String("defaultResult"),
//   	HeartbeatTimeout: jsii.Number(123),
//   	LifecycleHookName: jsii.String("lifecycleHookName"),
//   	LifecycleTransition: jsii.String("lifecycleTransition"),
//   	NotificationMetadata: jsii.String("notificationMetadata"),
//   	NotificationTargetArn: jsii.String("notificationTargetArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-lifecyclehook.html
//
type CfnLifecycleHookPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLifecycleHookMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLifecycleHookPropsMixin
type jsiiProxy_CfnLifecycleHookPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLifecycleHookPropsMixin) Props() *CfnLifecycleHookMixinProps {
	var returns *CfnLifecycleHookMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLifecycleHookPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AutoScaling::LifecycleHook`.
func NewCfnLifecycleHookPropsMixin(props *CfnLifecycleHookMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLifecycleHookPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLifecycleHookPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLifecycleHookPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnLifecycleHookPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AutoScaling::LifecycleHook`.
func NewCfnLifecycleHookPropsMixin_Override(c CfnLifecycleHookPropsMixin, props *CfnLifecycleHookMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnLifecycleHookPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLifecycleHookPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLifecycleHookPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnLifecycleHookPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLifecycleHookPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_autoscaling.mixins.CfnLifecycleHookPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLifecycleHookPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLifecycleHookPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


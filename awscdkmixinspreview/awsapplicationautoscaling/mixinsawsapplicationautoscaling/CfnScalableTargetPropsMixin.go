package mixinsawsapplicationautoscaling

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsapplicationautoscaling/mixinsawsapplicationautoscaling/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ApplicationAutoScaling::ScalableTarget` resource specifies a resource that Application Auto Scaling can scale, such as an AWS::DynamoDB::Table or AWS::ECS::Service resource.
//
// For more information, see [Getting started](https://docs.aws.amazon.com/autoscaling/application/userguide/getting-started.html) in the *Application Auto Scaling User Guide* .
//
// > If the resource that you want Application Auto Scaling to scale is not yet created in your account, add a dependency on the resource when registering it as a scalable target using the [DependsOn](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScalableTargetPropsMixin := awscdkmixinspreview.Mixins.NewCfnScalableTargetPropsMixin(&CfnScalableTargetMixinProps{
//   	MaxCapacity: jsii.Number(123),
//   	MinCapacity: jsii.Number(123),
//   	ResourceId: jsii.String("resourceId"),
//   	RoleArn: jsii.String("roleArn"),
//   	ScalableDimension: jsii.String("scalableDimension"),
//   	ScheduledActions: []interface{}{
//   		&ScheduledActionProperty{
//   			EndTime: NewDate(),
//   			ScalableTargetAction: &ScalableTargetActionProperty{
//   				MaxCapacity: jsii.Number(123),
//   				MinCapacity: jsii.Number(123),
//   			},
//   			Schedule: jsii.String("schedule"),
//   			ScheduledActionName: jsii.String("scheduledActionName"),
//   			StartTime: NewDate(),
//   			Timezone: jsii.String("timezone"),
//   		},
//   	},
//   	ServiceNamespace: jsii.String("serviceNamespace"),
//   	SuspendedState: &SuspendedStateProperty{
//   		DynamicScalingInSuspended: jsii.Boolean(false),
//   		DynamicScalingOutSuspended: jsii.Boolean(false),
//   		ScheduledScalingSuspended: jsii.Boolean(false),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-applicationautoscaling-scalabletarget.html
//
type CfnScalableTargetPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnScalableTargetMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnScalableTargetPropsMixin
type jsiiProxy_CfnScalableTargetPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnScalableTargetPropsMixin) Props() *CfnScalableTargetMixinProps {
	var returns *CfnScalableTargetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScalableTargetPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ApplicationAutoScaling::ScalableTarget`.
func NewCfnScalableTargetPropsMixin(props *CfnScalableTargetMixinProps, options *mixins.CfnPropertyMixinOptions) CfnScalableTargetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnScalableTargetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScalableTargetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalableTargetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ApplicationAutoScaling::ScalableTarget`.
func NewCfnScalableTargetPropsMixin_Override(c CfnScalableTargetPropsMixin, props *CfnScalableTargetMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalableTargetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnScalableTargetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScalableTargetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalableTargetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScalableTargetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_applicationautoscaling.mixins.CfnScalableTargetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScalableTargetPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnScalableTargetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


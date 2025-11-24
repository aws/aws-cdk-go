package mixinsawsredshift

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsredshift/mixinsawsredshift/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a scheduled action.
//
// A scheduled action contains a schedule and an Amazon Redshift API action. For example, you can create a schedule of when to run the `ResizeCluster` API operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScheduledActionPropsMixin := awscdkmixinspreview.Mixins.NewCfnScheduledActionPropsMixin(&CfnScheduledActionMixinProps{
//   	Enable: jsii.Boolean(false),
//   	EndTime: jsii.String("endTime"),
//   	IamRole: jsii.String("iamRole"),
//   	Schedule: jsii.String("schedule"),
//   	ScheduledActionDescription: jsii.String("scheduledActionDescription"),
//   	ScheduledActionName: jsii.String("scheduledActionName"),
//   	StartTime: jsii.String("startTime"),
//   	TargetAction: &ScheduledActionTypeProperty{
//   		PauseCluster: &PauseClusterMessageProperty{
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   		ResizeCluster: &ResizeClusterMessageProperty{
//   			Classic: jsii.Boolean(false),
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
//   			ClusterType: jsii.String("clusterType"),
//   			NodeType: jsii.String("nodeType"),
//   			NumberOfNodes: jsii.Number(123),
//   		},
//   		ResumeCluster: &ResumeClusterMessageProperty{
//   			ClusterIdentifier: jsii.String("clusterIdentifier"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-scheduledaction.html
//
type CfnScheduledActionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnScheduledActionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnScheduledActionPropsMixin
type jsiiProxy_CfnScheduledActionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnScheduledActionPropsMixin) Props() *CfnScheduledActionMixinProps {
	var returns *CfnScheduledActionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnScheduledActionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Redshift::ScheduledAction`.
func NewCfnScheduledActionPropsMixin(props *CfnScheduledActionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnScheduledActionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnScheduledActionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnScheduledActionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnScheduledActionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Redshift::ScheduledAction`.
func NewCfnScheduledActionPropsMixin_Override(c CfnScheduledActionPropsMixin, props *CfnScheduledActionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnScheduledActionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnScheduledActionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnScheduledActionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnScheduledActionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnScheduledActionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_redshift.mixins.CfnScheduledActionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnScheduledActionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnScheduledActionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


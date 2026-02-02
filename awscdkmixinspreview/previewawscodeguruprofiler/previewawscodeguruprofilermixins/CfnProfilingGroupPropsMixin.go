package previewawscodeguruprofilermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodeguruprofiler/previewawscodeguruprofilermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a profiling group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var agentPermissions interface{}
//
//   cfnProfilingGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnProfilingGroupPropsMixin(&CfnProfilingGroupMixinProps{
//   	AgentPermissions: agentPermissions,
//   	AnomalyDetectionNotificationConfiguration: []interface{}{
//   		&ChannelProperty{
//   			ChannelId: jsii.String("channelId"),
//   			ChannelUri: jsii.String("channelUri"),
//   		},
//   	},
//   	ComputePlatform: jsii.String("computePlatform"),
//   	ProfilingGroupName: jsii.String("profilingGroupName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html
//
type CfnProfilingGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnProfilingGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProfilingGroupPropsMixin
type jsiiProxy_CfnProfilingGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnProfilingGroupPropsMixin) Props() *CfnProfilingGroupMixinProps {
	var returns *CfnProfilingGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProfilingGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeGuruProfiler::ProfilingGroup`.
func NewCfnProfilingGroupPropsMixin(props *CfnProfilingGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnProfilingGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProfilingGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProfilingGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.mixins.CfnProfilingGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeGuruProfiler::ProfilingGroup`.
func NewCfnProfilingGroupPropsMixin_Override(c CfnProfilingGroupPropsMixin, props *CfnProfilingGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.mixins.CfnProfilingGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnProfilingGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProfilingGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.mixins.CfnProfilingGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProfilingGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codeguruprofiler.mixins.CfnProfilingGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProfilingGroupPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnProfilingGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


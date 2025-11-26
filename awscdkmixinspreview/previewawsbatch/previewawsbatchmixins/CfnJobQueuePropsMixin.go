package previewawsbatchmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsbatch/previewawsbatchmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Batch::JobQueue` resource specifies the parameters for an AWS Batch job queue definition.
//
// For more information, see [Job Queues](https://docs.aws.amazon.com/batch/latest/userguide/job_queues.html) in the ** .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnJobQueuePropsMixin := awscdkmixinspreview.Mixins.NewCfnJobQueuePropsMixin(&CfnJobQueueMixinProps{
//   	ComputeEnvironmentOrder: []interface{}{
//   		&ComputeEnvironmentOrderProperty{
//   			ComputeEnvironment: jsii.String("computeEnvironment"),
//   			Order: jsii.Number(123),
//   		},
//   	},
//   	JobQueueName: jsii.String("jobQueueName"),
//   	JobQueueType: jsii.String("jobQueueType"),
//   	JobStateTimeLimitActions: []interface{}{
//   		&JobStateTimeLimitActionProperty{
//   			Action: jsii.String("action"),
//   			MaxTimeSeconds: jsii.Number(123),
//   			Reason: jsii.String("reason"),
//   			State: jsii.String("state"),
//   		},
//   	},
//   	Priority: jsii.Number(123),
//   	SchedulingPolicyArn: jsii.String("schedulingPolicyArn"),
//   	ServiceEnvironmentOrder: []interface{}{
//   		&ServiceEnvironmentOrderProperty{
//   			Order: jsii.Number(123),
//   			ServiceEnvironment: jsii.String("serviceEnvironment"),
//   		},
//   	},
//   	State: jsii.String("state"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobqueue.html
//
type CfnJobQueuePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnJobQueueMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnJobQueuePropsMixin
type jsiiProxy_CfnJobQueuePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnJobQueuePropsMixin) Props() *CfnJobQueueMixinProps {
	var returns *CfnJobQueueMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobQueuePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Batch::JobQueue`.
func NewCfnJobQueuePropsMixin(props *CfnJobQueueMixinProps, options *mixins.CfnPropertyMixinOptions) CfnJobQueuePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnJobQueuePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJobQueuePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobQueuePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Batch::JobQueue`.
func NewCfnJobQueuePropsMixin_Override(c CfnJobQueuePropsMixin, props *CfnJobQueueMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobQueuePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnJobQueuePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobQueuePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobQueuePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobQueuePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_batch.mixins.CfnJobQueuePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJobQueuePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnJobQueuePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


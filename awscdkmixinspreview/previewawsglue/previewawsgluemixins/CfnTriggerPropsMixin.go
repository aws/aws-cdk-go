package previewawsgluemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsglue/previewawsgluemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Glue::Trigger` resource specifies triggers that run AWS Glue jobs.
//
// For more information, see [Triggering Jobs in AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/trigger-job.html) and [Trigger Structure](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-trigger.html#aws-glue-api-jobs-trigger-Trigger) in the *AWS Glue Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var arguments_ interface{}
//   var tags interface{}
//
//   cfnTriggerPropsMixin := awscdkmixinspreview.Mixins.NewCfnTriggerPropsMixin(&CfnTriggerMixinProps{
//   	Actions: []interface{}{
//   		&ActionProperty{
//   			Arguments: arguments_,
//   			CrawlerName: jsii.String("crawlerName"),
//   			JobName: jsii.String("jobName"),
//   			NotificationProperty: &NotificationPropertyProperty{
//   				NotifyDelayAfter: jsii.Number(123),
//   			},
//   			SecurityConfiguration: jsii.String("securityConfiguration"),
//   			Timeout: jsii.Number(123),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EventBatchingCondition: &EventBatchingConditionProperty{
//   		BatchSize: jsii.Number(123),
//   		BatchWindow: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	Predicate: &PredicateProperty{
//   		Conditions: []interface{}{
//   			&ConditionProperty{
//   				CrawlerName: jsii.String("crawlerName"),
//   				CrawlState: jsii.String("crawlState"),
//   				JobName: jsii.String("jobName"),
//   				LogicalOperator: jsii.String("logicalOperator"),
//   				State: jsii.String("state"),
//   			},
//   		},
//   		Logical: jsii.String("logical"),
//   	},
//   	Schedule: jsii.String("schedule"),
//   	StartOnCreation: jsii.Boolean(false),
//   	Tags: tags,
//   	Type: jsii.String("type"),
//   	WorkflowName: jsii.String("workflowName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html
//
type CfnTriggerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTriggerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTriggerPropsMixin
type jsiiProxy_CfnTriggerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnTriggerPropsMixin) Props() *CfnTriggerMixinProps {
	var returns *CfnTriggerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTriggerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Glue::Trigger`.
func NewCfnTriggerPropsMixin(props *CfnTriggerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTriggerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTriggerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTriggerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Glue::Trigger`.
func NewCfnTriggerPropsMixin_Override(c CfnTriggerPropsMixin, props *CfnTriggerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTriggerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTriggerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTriggerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_glue.mixins.CfnTriggerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTriggerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTriggerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


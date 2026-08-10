package awsimagebuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a runtime instance of an EC2 Image Builder workflow.
//
// Workflow executions are created automatically when an image build runs a workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWorkflowExecutionPropsMixin := awscdkcfnpropertymixins.Aws_imagebuilder.NewCfnWorkflowExecutionPropsMixin(&CfnWorkflowExecutionMixinProps{
//   	ImageBuildVersionArn: jsii.String("imageBuildVersionArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-workflowexecution.html
//
type CfnWorkflowExecutionPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWorkflowExecutionMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkflowExecutionPropsMixin
type jsiiProxy_CfnWorkflowExecutionPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnWorkflowExecutionPropsMixin) Props() *CfnWorkflowExecutionMixinProps {
	var returns *CfnWorkflowExecutionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkflowExecutionPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ImageBuilder::WorkflowExecution`.
func NewCfnWorkflowExecutionPropsMixin(props *CfnWorkflowExecutionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWorkflowExecutionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkflowExecutionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkflowExecutionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnWorkflowExecutionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ImageBuilder::WorkflowExecution`.
func NewCfnWorkflowExecutionPropsMixin_Override(c CfnWorkflowExecutionPropsMixin, props *CfnWorkflowExecutionMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnWorkflowExecutionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWorkflowExecutionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkflowExecutionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnWorkflowExecutionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkflowExecutionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_imagebuilder.CfnWorkflowExecutionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkflowExecutionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWorkflowExecutionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


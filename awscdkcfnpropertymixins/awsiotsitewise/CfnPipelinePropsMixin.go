package awsiotsitewise

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiotsitewise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::IoTSiteWise::Pipeline.
//
// A pipeline defines a directed acyclic graph (DAG) of compute nodes, where each node references a task definition and can declare dependencies on other nodes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnPipelinePropsMixin := awscdkcfnpropertymixins.Aws_iotsitewise.NewCfnPipelinePropsMixin(&CfnPipelineMixinProps{
//   	Computations: []interface{}{
//   		&ComputeNodeProperty{
//   			ComputeNodeName: jsii.String("computeNodeName"),
//   			DependsOn: []*string{
//   				jsii.String("dependsOn"),
//   			},
//   			EnvironmentVariables: map[string]*string{
//   				"environmentVariablesKey": jsii.String("environmentVariables"),
//   			},
//   			TaskName: jsii.String("taskName"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	PipelineName: jsii.String("pipelineName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WorkspaceName: jsii.String("workspaceName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-pipeline.html
//
type CfnPipelinePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPipelineMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPipelinePropsMixin
type jsiiProxy_CfnPipelinePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPipelinePropsMixin) Props() *CfnPipelineMixinProps {
	var returns *CfnPipelineMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPipelinePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTSiteWise::Pipeline`.
func NewCfnPipelinePropsMixin(props *CfnPipelineMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnPipelinePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPipelinePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPipelinePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnPipelinePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTSiteWise::Pipeline`.
func NewCfnPipelinePropsMixin_Override(c CfnPipelinePropsMixin, props *CfnPipelineMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnPipelinePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPipelinePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPipelinePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnPipelinePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPipelinePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_iotsitewise.CfnPipelinePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPipelinePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPipelinePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


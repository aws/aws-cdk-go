package previewawsneptunegraphmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsneptunegraph/previewawsneptunegraphmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS ::NeptuneGraph::Graph` resource creates an  graph.
//
// is a memory-optimized graph database engine for analytics. For more information, see [](https://docs.aws.amazon.com/neptune-analytics/latest/userguide/what-is-neptune-analytics.html) .
//
// You can use `AWS ::NeptuneGraph::Graph.DeletionProtection` to help guard against unintended deletion of your graph.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGraphPropsMixin := awscdkmixinspreview.Mixins.NewCfnGraphPropsMixin(&CfnGraphMixinProps{
//   	DeletionProtection: jsii.Boolean(false),
//   	GraphName: jsii.String("graphName"),
//   	ProvisionedMemory: jsii.Number(123),
//   	PublicConnectivity: jsii.Boolean(false),
//   	ReplicaCount: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VectorSearchConfiguration: &VectorSearchConfigurationProperty{
//   		VectorSearchDimension: jsii.Number(123),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptunegraph-graph.html
//
type CfnGraphPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGraphMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGraphPropsMixin
type jsiiProxy_CfnGraphPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGraphPropsMixin) Props() *CfnGraphMixinProps {
	var returns *CfnGraphMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGraphPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NeptuneGraph::Graph`.
func NewCfnGraphPropsMixin(props *CfnGraphMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGraphPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGraphPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGraphPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_neptunegraph.mixins.CfnGraphPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NeptuneGraph::Graph`.
func NewCfnGraphPropsMixin_Override(c CfnGraphPropsMixin, props *CfnGraphMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_neptunegraph.mixins.CfnGraphPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGraphPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGraphPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_neptunegraph.mixins.CfnGraphPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGraphPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_neptunegraph.mixins.CfnGraphPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGraphPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnGraphPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package previewawsmemorydbmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmemorydb/previewawsmemorydbmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a multi-Region cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMultiRegionClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnMultiRegionClusterPropsMixin(&CfnMultiRegionClusterMixinProps{
//   	Description: jsii.String("description"),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	MultiRegionClusterNameSuffix: jsii.String("multiRegionClusterNameSuffix"),
//   	MultiRegionParameterGroupName: jsii.String("multiRegionParameterGroupName"),
//   	NodeType: jsii.String("nodeType"),
//   	NumShards: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TlsEnabled: jsii.Boolean(false),
//   	UpdateStrategy: jsii.String("updateStrategy"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-memorydb-multiregioncluster.html
//
type CfnMultiRegionClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMultiRegionClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMultiRegionClusterPropsMixin
type jsiiProxy_CfnMultiRegionClusterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMultiRegionClusterPropsMixin) Props() *CfnMultiRegionClusterMixinProps {
	var returns *CfnMultiRegionClusterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMultiRegionClusterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MemoryDB::MultiRegionCluster`.
func NewCfnMultiRegionClusterPropsMixin(props *CfnMultiRegionClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMultiRegionClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMultiRegionClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMultiRegionClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_memorydb.mixins.CfnMultiRegionClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MemoryDB::MultiRegionCluster`.
func NewCfnMultiRegionClusterPropsMixin_Override(c CfnMultiRegionClusterPropsMixin, props *CfnMultiRegionClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_memorydb.mixins.CfnMultiRegionClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMultiRegionClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMultiRegionClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_memorydb.mixins.CfnMultiRegionClusterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMultiRegionClusterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_memorydb.mixins.CfnMultiRegionClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMultiRegionClusterPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnMultiRegionClusterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


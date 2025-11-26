package previewawselasticachemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawselasticache/previewawselasticachemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Consists of a primary cluster that accepts writes and an associated secondary cluster that resides in a different Amazon region.
//
// The secondary cluster accepts only reads. The primary cluster automatically replicates updates to the secondary cluster.
//
// - The *GlobalReplicationGroupIdSuffix* represents the name of the Global datastore, which is what you use to associate a secondary cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGlobalReplicationGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnGlobalReplicationGroupPropsMixin(&CfnGlobalReplicationGroupMixinProps{
//   	AutomaticFailoverEnabled: jsii.Boolean(false),
//   	CacheNodeType: jsii.String("cacheNodeType"),
//   	CacheParameterGroupName: jsii.String("cacheParameterGroupName"),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	GlobalNodeGroupCount: jsii.Number(123),
//   	GlobalReplicationGroupDescription: jsii.String("globalReplicationGroupDescription"),
//   	GlobalReplicationGroupIdSuffix: jsii.String("globalReplicationGroupIdSuffix"),
//   	Members: []interface{}{
//   		&GlobalReplicationGroupMemberProperty{
//   			ReplicationGroupId: jsii.String("replicationGroupId"),
//   			ReplicationGroupRegion: jsii.String("replicationGroupRegion"),
//   			Role: jsii.String("role"),
//   		},
//   	},
//   	RegionalConfigurations: []interface{}{
//   		&RegionalConfigurationProperty{
//   			ReplicationGroupId: jsii.String("replicationGroupId"),
//   			ReplicationGroupRegion: jsii.String("replicationGroupRegion"),
//   			ReshardingConfigurations: []interface{}{
//   				&ReshardingConfigurationProperty{
//   					NodeGroupId: jsii.String("nodeGroupId"),
//   					PreferredAvailabilityZones: []*string{
//   						jsii.String("preferredAvailabilityZones"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-globalreplicationgroup.html
//
type CfnGlobalReplicationGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnGlobalReplicationGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnGlobalReplicationGroupPropsMixin
type jsiiProxy_CfnGlobalReplicationGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnGlobalReplicationGroupPropsMixin) Props() *CfnGlobalReplicationGroupMixinProps {
	var returns *CfnGlobalReplicationGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnGlobalReplicationGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElastiCache::GlobalReplicationGroup`.
func NewCfnGlobalReplicationGroupPropsMixin(props *CfnGlobalReplicationGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnGlobalReplicationGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnGlobalReplicationGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnGlobalReplicationGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnGlobalReplicationGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElastiCache::GlobalReplicationGroup`.
func NewCfnGlobalReplicationGroupPropsMixin_Override(c CfnGlobalReplicationGroupPropsMixin, props *CfnGlobalReplicationGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnGlobalReplicationGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnGlobalReplicationGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnGlobalReplicationGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnGlobalReplicationGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnGlobalReplicationGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnGlobalReplicationGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnGlobalReplicationGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnGlobalReplicationGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


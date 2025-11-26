package previewawselasticachemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawselasticache/previewawselasticachemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ElastiCache::CacheCluster` type creates an Amazon ElastiCache cache cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCacheClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnCacheClusterPropsMixin(&CfnCacheClusterMixinProps{
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	AzMode: jsii.String("azMode"),
//   	CacheNodeType: jsii.String("cacheNodeType"),
//   	CacheParameterGroupName: jsii.String("cacheParameterGroupName"),
//   	CacheSecurityGroupNames: []*string{
//   		jsii.String("cacheSecurityGroupNames"),
//   	},
//   	CacheSubnetGroupName: jsii.String("cacheSubnetGroupName"),
//   	ClusterName: jsii.String("clusterName"),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	IpDiscovery: jsii.String("ipDiscovery"),
//   	LogDeliveryConfigurations: []interface{}{
//   		&LogDeliveryConfigurationRequestProperty{
//   			DestinationDetails: &DestinationDetailsProperty{
//   				CloudWatchLogsDetails: &CloudWatchLogsDestinationDetailsProperty{
//   					LogGroup: jsii.String("logGroup"),
//   				},
//   				KinesisFirehoseDetails: &KinesisFirehoseDestinationDetailsProperty{
//   					DeliveryStream: jsii.String("deliveryStream"),
//   				},
//   			},
//   			DestinationType: jsii.String("destinationType"),
//   			LogFormat: jsii.String("logFormat"),
//   			LogType: jsii.String("logType"),
//   		},
//   	},
//   	NetworkType: jsii.String("networkType"),
//   	NotificationTopicArn: jsii.String("notificationTopicArn"),
//   	NumCacheNodes: jsii.Number(123),
//   	Port: jsii.Number(123),
//   	PreferredAvailabilityZone: jsii.String("preferredAvailabilityZone"),
//   	PreferredAvailabilityZones: []*string{
//   		jsii.String("preferredAvailabilityZones"),
//   	},
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	SnapshotArns: []*string{
//   		jsii.String("snapshotArns"),
//   	},
//   	SnapshotName: jsii.String("snapshotName"),
//   	SnapshotRetentionLimit: jsii.Number(123),
//   	SnapshotWindow: jsii.String("snapshotWindow"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitEncryptionEnabled: jsii.Boolean(false),
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-cachecluster.html
//
type CfnCacheClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCacheClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCacheClusterPropsMixin
type jsiiProxy_CfnCacheClusterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCacheClusterPropsMixin) Props() *CfnCacheClusterMixinProps {
	var returns *CfnCacheClusterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCacheClusterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ElastiCache::CacheCluster`.
func NewCfnCacheClusterPropsMixin(props *CfnCacheClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCacheClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCacheClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCacheClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ElastiCache::CacheCluster`.
func NewCfnCacheClusterPropsMixin_Override(c CfnCacheClusterPropsMixin, props *CfnCacheClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCacheClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCacheClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCacheClusterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCacheClusterPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCacheClusterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


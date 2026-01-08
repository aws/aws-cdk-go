package previewawselasticachemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawselasticache/previewawselasticachemixins/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslogs"
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
//   var logsDelivery ILogsDelivery
//
//   cfnCacheClusterLogsMixin := awscdkmixinspreview.Mixins.NewCfnCacheClusterLogsMixin(jsii.String("logType"), logsDelivery)
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-cachecluster.html
//
type CfnCacheClusterLogsMixin interface {
	core.Mixin
	core.IMixin
	LogDelivery() previewawslogs.ILogsDelivery
	LogType() *string
	// Apply vended logs configuration to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct (has vendedLogs property).
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCacheClusterLogsMixin
type jsiiProxy_CfnCacheClusterLogsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCacheClusterLogsMixin) LogDelivery() previewawslogs.ILogsDelivery {
	var returns previewawslogs.ILogsDelivery
	_jsii_.Get(
		j,
		"logDelivery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCacheClusterLogsMixin) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}


// Create a mixin to enable vended logs for `AWS::ElastiCache::CacheCluster`.
func NewCfnCacheClusterLogsMixin(logType *string, logDelivery previewawslogs.ILogsDelivery) CfnCacheClusterLogsMixin {
	_init_.Initialize()

	if err := validateNewCfnCacheClusterLogsMixinParameters(logType, logDelivery); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCacheClusterLogsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterLogsMixin",
		[]interface{}{logType, logDelivery},
		&j,
	)

	return &j
}

// Create a mixin to enable vended logs for `AWS::ElastiCache::CacheCluster`.
func NewCfnCacheClusterLogsMixin_Override(c CfnCacheClusterLogsMixin, logType *string, logDelivery previewawslogs.ILogsDelivery) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterLogsMixin",
		[]interface{}{logType, logDelivery},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCacheClusterLogsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCacheClusterLogsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterLogsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCacheClusterLogsMixin_ELASTICACHE_LOGS() CfnCacheClusterElasticacheLogs {
	_init_.Initialize()
	var returns CfnCacheClusterElasticacheLogs
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_elasticache.mixins.CfnCacheClusterLogsMixin",
		"ELASTICACHE_LOGS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCacheClusterLogsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCacheClusterLogsMixin) Supports(construct constructs.IConstruct) *bool {
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


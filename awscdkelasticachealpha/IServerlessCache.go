package awscdkelasticachealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdkelasticachealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a Serverless ElastiCache cache.
// Experimental.
type IServerlessCache interface {
	awsec2.IConnectable
	awscdk.IResource
	// Grant the given identity custom permissions.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant connect permissions to the cache.
	// Experimental.
	GrantConnect(grantee awsiam.IGrantable) awsiam.Grant
	// Return the given named metric for this cache.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for active connections.
	// Experimental.
	MetricActiveConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for cache hit count.
	// Experimental.
	MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for cache hit rate.
	// Experimental.
	MetricCacheHitRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for cache miss count.
	// Experimental.
	MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for data stored in the cache.
	// Experimental.
	MetricDataStored(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for network bytes in.
	// Experimental.
	MetricNetworkBytesIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for network bytes out.
	// Experimental.
	MetricNetworkBytesOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for ECPUs consumed.
	// Experimental.
	MetricProcessingUnitsConsumed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for read request latency.
	// Experimental.
	MetricReadRequestLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for write request latency.
	// Experimental.
	MetricWriteRequestLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARNs of backups restored in the cache.
	// Experimental.
	BackupArnsToRestore() *[]*string
	// The cache engine used by this cache.
	// Experimental.
	Engine() CacheEngine
	// The KMS key used for encryption.
	// Experimental.
	KmsKey() awskms.IKey
	// The security groups associated with this cache.
	// Experimental.
	SecurityGroups() *[]awsec2.ISecurityGroup
	// The ARN of the serverless cache.
	// Experimental.
	ServerlessCacheArn() *string
	// The name of the serverless cache.
	// Experimental.
	ServerlessCacheName() *string
	// The subnets this cache is deployed in.
	// Experimental.
	Subnets() *[]awsec2.ISubnet
	// The user group associated with this cache.
	// Experimental.
	UserGroup() IUserGroup
	// The VPC this cache is deployed in.
	// Experimental.
	Vpc() awsec2.IVpc
}

// The jsii proxy for IServerlessCache
type jsiiProxy_IServerlessCache struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IServerlessCache) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) GrantConnect(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantConnectParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantConnect",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) MetricActiveConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricActiveConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricActiveConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) MetricCacheHitCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCacheHitCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCacheHitCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) MetricCacheHitRate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCacheHitRateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCacheHitRate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) MetricCacheMissCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCacheMissCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCacheMissCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) MetricDataStored(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDataStoredParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDataStored",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) MetricNetworkBytesIn(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNetworkBytesInParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkBytesIn",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) MetricNetworkBytesOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNetworkBytesOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkBytesOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) MetricProcessingUnitsConsumed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricProcessingUnitsConsumedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricProcessingUnitsConsumed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) MetricReadRequestLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricReadRequestLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricReadRequestLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) MetricWriteRequestLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricWriteRequestLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricWriteRequestLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IServerlessCache) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IServerlessCache) BackupArnsToRestore() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"backupArnsToRestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) Engine() CacheEngine {
	var returns CacheEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) ServerlessCacheArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessCacheArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) ServerlessCacheName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverlessCacheName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) Subnets() *[]awsec2.ISubnet {
	var returns *[]awsec2.ISubnet
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) UserGroup() IUserGroup {
	var returns IUserGroup
	_jsii_.Get(
		j,
		"userGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IServerlessCache) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsrds/internal"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Create a clustered database with a given number of instances.
// Experimental.
type IDatabaseCluster interface {
	awsec2.IConnectable
	awscdk.IResource
	awssecretsmanager.ISecretAttachmentTarget
	// Add a new db proxy to this cluster.
	// Experimental.
	AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy
	// Return the given named metric for this DBCluster.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The percentage of CPU utilization.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of database connections in use.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of deadlocks in the database per second.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of time that the instance has been running, in seconds.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available random access memory, in bytes.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of local storage available, in bytes.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput received from clients by each instance, in bytes per second.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput both received from and transmitted to clients by each instance, in bytes per second.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of network throughput sent to clients by each instance, in bytes per second.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total amount of backup storage in bytes consumed by all Aurora snapshots outside its backup retention window.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total amount of backup storage in bytes for which you are billed.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of storage used by your Aurora DB instance, in bytes.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of billed read I/O operations from a cluster volume, reported at 5-minute intervals.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of write disk I/O operations to the cluster volume, reported at 5-minute intervals.
	//
	// Average over 5 minutes.
	// Experimental.
	MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The endpoint to use for read/write operations.
	// Experimental.
	ClusterEndpoint() Endpoint
	// Identifier of the cluster.
	// Experimental.
	ClusterIdentifier() *string
	// Endpoint to use for load-balanced read-only operations.
	// Experimental.
	ClusterReadEndpoint() Endpoint
	// The engine of this Cluster.
	//
	// May be not known for imported Clusters if it wasn't provided explicitly.
	// Experimental.
	Engine() IClusterEngine
	// Endpoints which address each individual replica.
	// Experimental.
	InstanceEndpoints() *[]Endpoint
	// Identifiers of the replicas.
	// Experimental.
	InstanceIdentifiers() *[]*string
}

// The jsii proxy for IDatabaseCluster
type jsiiProxy_IDatabaseCluster struct {
	internal.Type__awsec2IConnectable
	internal.Type__awscdkIResource
	internal.Type__awssecretsmanagerISecretAttachmentTarget
}

func (i *jsiiProxy_IDatabaseCluster) AddProxy(id *string, options *DatabaseProxyOptions) DatabaseProxy {
	if err := i.validateAddProxyParameters(id, options); err != nil {
		panic(err)
	}
	var returns DatabaseProxy

	_jsii_.Invoke(
		i,
		"addProxy",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IDatabaseCluster) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCPUUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricDatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDatabaseConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDatabaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricDeadlocks(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDeadlocksParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDeadlocks",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricEngineUptime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEngineUptimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEngineUptime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricFreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFreeableMemoryParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFreeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricFreeLocalStorage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFreeLocalStorageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFreeLocalStorage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricNetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNetworkReceiveThroughputParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkReceiveThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricNetworkThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNetworkThroughputParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricNetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNetworkTransmitThroughputParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkTransmitThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricSnapshotStorageUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSnapshotStorageUsedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSnapshotStorageUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricTotalBackupStorageBilled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTotalBackupStorageBilledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTotalBackupStorageBilled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricVolumeBytesUsed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricVolumeBytesUsedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricVolumeBytesUsed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricVolumeReadIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricVolumeReadIOPsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricVolumeReadIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) MetricVolumeWriteIOPs(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricVolumeWriteIOPsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricVolumeWriteIOPs",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDatabaseCluster) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IDatabaseCluster) AsSecretAttachmentTarget() *awssecretsmanager.SecretAttachmentTargetProps {
	var returns *awssecretsmanager.SecretAttachmentTargetProps

	_jsii_.Invoke(
		i,
		"asSecretAttachmentTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) ClusterReadEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"clusterReadEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Engine() IClusterEngine {
	var returns IClusterEngine
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) InstanceEndpoints() *[]Endpoint {
	var returns *[]Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) InstanceIdentifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceIdentifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseCluster) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


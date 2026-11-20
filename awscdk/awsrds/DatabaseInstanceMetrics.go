package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrds"
)

// CloudWatch metrics for a database instance.
//
// Works with any reference to a database instance: a `DatabaseInstance`, an
// imported instance or a `CfnDBInstance`.
//
// Example:
//   var instance IDatabaseInstance
//
//
//   metrics := rds.DatabaseInstanceMetrics_FromDatabaseInstance(instance)
//
//   // How far a read replica lags behind its source (average over 5 minutes)
//   replicaLag := metrics.ReplicaLag()
//
//   // The average amount of time taken per disk I/O read operation (average over 1 minute)
//   readLatency := metrics.ReadLatency(&MetricOptions{
//   	Period: awscdk.Duration_Seconds(jsii.Number(60)),
//   })
//
//   // Any other metric, such as an engine-specific one
//   binLogDiskUsage := metrics.Metric(jsii.String("BinLogDiskUsage"), &MetricOptions{
//   	Statistic: jsii.String("Maximum"),
//   })
//
type DatabaseInstanceMetrics interface {
	// The percent of burst-bucket I/O credits available.
	//
	// This metric is available for DB instances that use General Purpose SSD (gp2) storage.
	// Default: - average over 5 minutes.
	//
	BurstBalance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The percentage of CPU utilization.
	// Default: - average over 5 minutes.
	//
	CpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of client network connections to the database instance.
	// Default: - average over 5 minutes.
	//
	DatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of outstanding I/Os (read/write requests) waiting to access the disk.
	// Default: - average over 5 minutes.
	//
	DiskQueueDepth(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available random access memory.
	// Default: - average over 5 minutes.
	//
	FreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of available storage space.
	// Default: - average over 5 minutes.
	//
	FreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Return the given named metric for this database instance.
	//
	// If `dimensionsMap` is set, it replaces the default `DBInstanceIdentifier`
	// dimension, so the metric no longer refers to this instance unless the new
	// dimensions include it.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The incoming (receive) network traffic on the DB instance, including both customer database traffic and Amazon RDS traffic used for monitoring and replication.
	// Default: - average over 5 minutes.
	//
	NetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The outgoing (transmit) network traffic on the DB instance, including both customer database traffic and Amazon RDS traffic used for monitoring and replication.
	// Default: - average over 5 minutes.
	//
	NetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of disk read I/O operations per second.
	// Default: - average over 5 minutes.
	//
	ReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average amount of time taken per disk I/O read operation.
	// Default: - average over 5 minutes.
	//
	ReadLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of bytes read from disk per second.
	// Default: - average over 5 minutes.
	//
	ReadThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// For read replica configurations, the amount of time a read replica DB instance lags behind the source DB instance.
	//
	// This metric is available for MariaDB, Microsoft SQL Server, MySQL, Oracle, and PostgreSQL read replicas.
	// Default: - average over 5 minutes.
	//
	ReplicaLag(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The amount of swap space used on the DB instance.
	//
	// This metric is available for MariaDB, MySQL, Oracle, and PostgreSQL DB instances.
	// Default: - average over 5 minutes.
	//
	SwapUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of disk write I/O operations per second.
	// Default: - average over 5 minutes.
	//
	WriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average amount of time taken per disk I/O write operation.
	// Default: - average over 5 minutes.
	//
	WriteLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The average number of bytes written to disk per second.
	// Default: - average over 5 minutes.
	//
	WriteThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
}

// The jsii proxy struct for DatabaseInstanceMetrics
type jsiiProxy_DatabaseInstanceMetrics struct {
	_ byte // padding
}

// Creates the metrics for the given database instance.
func DatabaseInstanceMetrics_FromDatabaseInstance(instance interfacesawsrds.IDBInstanceRef) DatabaseInstanceMetrics {
	_init_.Initialize()

	if err := validateDatabaseInstanceMetrics_FromDatabaseInstanceParameters(instance); err != nil {
		panic(err)
	}
	var returns DatabaseInstanceMetrics

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.DatabaseInstanceMetrics",
		"fromDatabaseInstance",
		[]interface{}{instance},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) BurstBalance(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateBurstBalanceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"burstBalance",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) CpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"cpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) DatabaseConnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateDatabaseConnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"databaseConnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) DiskQueueDepth(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateDiskQueueDepthParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"diskQueueDepth",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) FreeableMemory(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateFreeableMemoryParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"freeableMemory",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) FreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateFreeStorageSpaceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"freeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) NetworkReceiveThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateNetworkReceiveThroughputParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"networkReceiveThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) NetworkTransmitThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateNetworkTransmitThroughputParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"networkTransmitThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) ReadIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateReadIOPSParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"readIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) ReadLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateReadLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"readLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) ReadThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateReadThroughputParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"readThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) ReplicaLag(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateReplicaLagParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"replicaLag",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) SwapUsage(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateSwapUsageParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"swapUsage",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) WriteIOPS(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateWriteIOPSParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"writeIOPS",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) WriteLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateWriteLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"writeLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseInstanceMetrics) WriteThroughput(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateWriteThroughputParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"writeThroughput",
		[]interface{}{props},
		&returns,
	)

	return returns
}


package awskinesis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesis/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// A Kinesis Stream.
type IStream interface {
	awscdk.IResource
	// Grant the indicated permissions on this stream to the provided IAM principal.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Grant read permissions for this stream and its contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to ues the key to decrypt the
	// contents of the stream will also be granted.
	GrantRead(grantee awsiam.IGrantable) awsiam.Grant
	// Grants read/write permissions for this stream and its contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to use the key for
	// encrypt/decrypt will also be granted.
	GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Grant write permissions for this stream and its contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to ues the key to encrypt the
	// contents of the stream will also be granted.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Return stream metric based from its metric name.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records retrieved from the shard, measured over the specified time period.
	//
	// Minimum, Maximum, and
	// Average statistics represent the records in a single GetRecords operation for the stream in the specified time
	// period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricGetRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of bytes retrieved from the Kinesis stream, measured over the specified time period.
	//
	// Minimum, Maximum,
	// and Average statistics represent the bytes in a single GetRecords operation for the stream in the specified time
	// period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricGetRecordsBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The age of the last record in all GetRecords calls made against a Kinesis stream, measured over the specified time period.
	//
	// Age is the difference between the current time and when the last record of the GetRecords call was written
	// to the stream. The Minimum and Maximum statistics can be used to track the progress of Kinesis consumer
	// applications. A value of zero indicates that the records being read are completely caught up with the stream.
	//
	// The metric defaults to maximum over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricGetRecordsIteratorAgeMilliseconds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time taken per GetRecords operation, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricGetRecordsLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of successful GetRecords operations per stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricGetRecordsSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of bytes successfully put to the Kinesis stream over the specified time period.
	//
	// This metric includes
	// bytes from PutRecord and PutRecords operations. Minimum, Maximum, and Average statistics represent the bytes in a
	// single put operation for the stream in the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records successfully put to the Kinesis stream over the specified time period.
	//
	// This metric includes
	// record counts from PutRecord and PutRecords operations. Minimum, Maximum, and Average statistics represent the
	// records in a single put operation for the stream in the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of bytes put to the Kinesis stream using the PutRecord operation over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time taken per PutRecord operation, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of bytes put to the Kinesis stream using the PutRecords operation over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records rejected due to internal failures in a PutRecords operation per Kinesis data stream, measured over the specified time period.
	//
	// Occasional internal failures are to be expected and should be retried.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsFailedRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time taken per PutRecords operation, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of PutRecords operations where at least one record succeeded, per Kinesis stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of successful records in a PutRecords operation per Kinesis data stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsSuccessfulRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records rejected due to throttling in a PutRecords operation per Kinesis data stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsThrottledRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The total number of records sent in a PutRecords operation per Kinesis data stream, measured over the specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordsTotalRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of successful PutRecord operations per Kinesis stream, measured over the specified time period.
	//
	// Average
	// reflects the percentage of successful writes to a stream.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricPutRecordSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of GetRecords calls throttled for the stream over the specified time period.
	//
	// The most commonly used
	// statistic for this metric is Average.
	//
	// When the Minimum statistic has a value of 1, all records were throttled for the stream during the specified time
	// period.
	//
	// When the Maximum statistic has a value of 0 (zero), no records were throttled for the stream during the specified
	// time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricReadProvisionedThroughputExceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The number of records rejected due to throttling for the stream over the specified time period.
	//
	// This metric
	// includes throttling from PutRecord and PutRecords operations.
	//
	// When the Minimum statistic has a non-zero value, records were being throttled for the stream during the specified
	// time period.
	//
	// When the Maximum statistic has a value of 0 (zero), no records were being throttled for the stream during the
	// specified time period.
	//
	// The metric defaults to average over 5 minutes, it can be changed by passing `statistic` and `period` properties.
	MetricWriteProvisionedThroughputExceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Optional KMS encryption key associated with this stream.
	EncryptionKey() awskms.IKey
	// The ARN of the stream.
	StreamArn() *string
	// The name of the stream.
	StreamName() *string
}

// The jsii proxy for IStream
type jsiiProxy_IStream struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IStream) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_IStream) GrantRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) GrantReadWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantWriteParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IStream) MetricGetRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricGetRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricGetRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricGetRecordsBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricGetRecordsBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricGetRecordsBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricGetRecordsIteratorAgeMilliseconds(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricGetRecordsIteratorAgeMillisecondsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricGetRecordsIteratorAgeMilliseconds",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricGetRecordsLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricGetRecordsLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricGetRecordsLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricGetRecordsSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricGetRecordsSuccessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricGetRecordsSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIncomingBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIncomingBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIncomingRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIncomingRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricPutRecordBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPutRecordBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPutRecordBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricPutRecordLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPutRecordLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPutRecordLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricPutRecordsBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPutRecordsBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPutRecordsBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricPutRecordsFailedRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPutRecordsFailedRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPutRecordsFailedRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricPutRecordsLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPutRecordsLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPutRecordsLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricPutRecordsSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPutRecordsSuccessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPutRecordsSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricPutRecordsSuccessfulRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPutRecordsSuccessfulRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPutRecordsSuccessfulRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricPutRecordsThrottledRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPutRecordsThrottledRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPutRecordsThrottledRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricPutRecordsTotalRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPutRecordsTotalRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPutRecordsTotalRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricPutRecordSuccess(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricPutRecordSuccessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricPutRecordSuccess",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricReadProvisionedThroughputExceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricReadProvisionedThroughputExceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricReadProvisionedThroughputExceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IStream) MetricWriteProvisionedThroughputExceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricWriteProvisionedThroughputExceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricWriteProvisionedThroughputExceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IStream) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStream) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStream) StreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamName",
		&returns,
	)
	return returns
}


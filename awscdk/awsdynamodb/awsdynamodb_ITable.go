package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// An interface that represents a DynamoDB Table - either created with the CDK, or an existing one.
type ITable interface {
	awscdk.IResource
	// Adds an IAM policy statement associated with this table to an IAM principal's policy.
	//
	// If `encryptionKey` is present, appropriate grants to the key needs to be added
	// separately using the `table.encryptionKey.grant*` methods.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Permits all DynamoDB operations ("dynamodb:*") to an IAM principal.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal all data read operations from this table: BatchGetItem, GetRecords, GetShardIterator, Query, GetItem, Scan.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	GrantReadData(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal to all data read/write operations to this table.
	//
	// BatchGetItem, GetRecords, GetShardIterator, Query, GetItem, Scan,
	// BatchWriteItem, PutItem, UpdateItem, DeleteItem
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	GrantReadWriteData(grantee awsiam.IGrantable) awsiam.Grant
	// Adds an IAM policy statement associated with this table's stream to an IAM principal's policy.
	//
	// If `encryptionKey` is present, appropriate grants to the key needs to be added
	// separately using the `table.encryptionKey.grant*` methods.
	GrantStream(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Permits an IAM principal all stream data read operations for this table's stream: DescribeStream, GetRecords, GetShardIterator, ListStreams.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	GrantStreamRead(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM Principal to list streams attached to current dynamodb table.
	GrantTableListStreams(grantee awsiam.IGrantable) awsiam.Grant
	// Permits an IAM principal all data write operations to this table: BatchWriteItem, PutItem, UpdateItem, DeleteItem.
	//
	// Appropriate grants will also be added to the customer-managed KMS key
	// if one was configured.
	GrantWriteData(grantee awsiam.IGrantable) awsiam.Grant
	// Metric for the number of Errors executing all Lambdas.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the conditional check failed requests.
	MetricConditionalCheckFailedRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the consumed read capacity units.
	MetricConsumedReadCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the consumed write capacity units.
	MetricConsumedWriteCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the successful request latency.
	MetricSuccessfulRequestLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the system errors this table.
	MetricSystemErrorsForOperations(props *SystemErrorsForOperationsMetricOptions) awscloudwatch.IMetric
	// Metric for throttled requests.
	// Deprecated: use `metricThrottledRequestsForOperations`.
	MetricThrottledRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for throttled requests.
	MetricThrottledRequestsForOperations(props *OperationsMetricOptions) awscloudwatch.IMetric
	// Metric for the user errors.
	MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Optional KMS encryption key associated with this table.
	EncryptionKey() awskms.IKey
	// Arn of the dynamodb table.
	TableArn() *string
	// Table name of the dynamodb table.
	TableName() *string
	// ARN of the table's stream, if there is one.
	TableStreamArn() *string
}

// The jsii proxy for ITable
type jsiiProxy_ITable struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_ITable) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_ITable) GrantFullAccess(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantFullAccessParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantFullAccess",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantReadData(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadDataParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantReadWriteData(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantReadWriteDataParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantReadWriteData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantStream(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := i.validateGrantStreamParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStream",
		args,
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantStreamRead(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantStreamReadParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantStreamRead",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantTableListStreams(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantTableListStreamsParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantTableListStreams",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) GrantWriteData(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantWriteDataParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWriteData",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_ITable) MetricConditionalCheckFailedRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricConditionalCheckFailedRequestsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricConditionalCheckFailedRequests",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricConsumedReadCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricConsumedReadCapacityUnitsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricConsumedReadCapacityUnits",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricConsumedWriteCapacityUnits(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricConsumedWriteCapacityUnitsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricConsumedWriteCapacityUnits",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricSuccessfulRequestLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSuccessfulRequestLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSuccessfulRequestLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricSystemErrorsForOperations(props *SystemErrorsForOperationsMetricOptions) awscloudwatch.IMetric {
	if err := i.validateMetricSystemErrorsForOperationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.IMetric

	_jsii_.Invoke(
		i,
		"metricSystemErrorsForOperations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricThrottledRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricThrottledRequestsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricThrottledRequests",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricThrottledRequestsForOperations(props *OperationsMetricOptions) awscloudwatch.IMetric {
	if err := i.validateMetricThrottledRequestsForOperationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.IMetric

	_jsii_.Invoke(
		i,
		"metricThrottledRequestsForOperations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ITable) MetricUserErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricUserErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricUserErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ITable) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) TableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITable) TableStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableStreamArn",
		&returns,
	)
	return returns
}


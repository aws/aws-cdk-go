package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslogs/internal"
)

// Experimental.
type ILogGroup interface {
	awsiam.IResourceWithPolicy
	// Create a new Metric Filter on this Log Group.
	// Experimental.
	AddMetricFilter(id *string, props *MetricFilterOptions) MetricFilter
	// Create a new Log Stream for this Log Group.
	// Experimental.
	AddStream(id *string, props *StreamOptions) LogStream
	// Create a new Subscription Filter on this Log Group.
	// Experimental.
	AddSubscriptionFilter(id *string, props *SubscriptionFilterOptions) SubscriptionFilter
	// Extract a metric from structured log events in the LogGroup.
	//
	// Creates a MetricFilter on this LogGroup that will extract the value
	// of the indicated JSON field in all records where it occurs.
	//
	// The metric will be available in CloudWatch Metrics under the
	// indicated namespace and name.
	//
	// Returns: A Metric object representing the extracted metric.
	// Experimental.
	ExtractMetric(jsonField *string, metricNamespace *string, metricName *string) awscloudwatch.Metric
	// Give the indicated permissions on this log group and all streams.
	// Experimental.
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	// Give permissions to write to create and write to streams in this log group.
	// Experimental.
	GrantWrite(grantee awsiam.IGrantable) awsiam.Grant
	// Public method to get the physical name of this log group.
	// Experimental.
	LogGroupPhysicalName() *string
	// The ARN of this log group, with ':*' appended.
	// Experimental.
	LogGroupArn() *string
	// The name of this log group.
	// Experimental.
	LogGroupName() *string
}

// The jsii proxy for ILogGroup
type jsiiProxy_ILogGroup struct {
	internal.Type__awsiamIResourceWithPolicy
}

func (i *jsiiProxy_ILogGroup) AddMetricFilter(id *string, props *MetricFilterOptions) MetricFilter {
	var returns MetricFilter

	_jsii_.Invoke(
		i,
		"addMetricFilter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) AddStream(id *string, props *StreamOptions) LogStream {
	var returns LogStream

	_jsii_.Invoke(
		i,
		"addStream",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) AddSubscriptionFilter(id *string, props *SubscriptionFilterOptions) SubscriptionFilter {
	var returns SubscriptionFilter

	_jsii_.Invoke(
		i,
		"addSubscriptionFilter",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) ExtractMetric(jsonField *string, metricNamespace *string, metricName *string) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"extractMetric",
		[]interface{}{jsonField, metricNamespace, metricName},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
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

func (i *jsiiProxy_ILogGroup) GrantWrite(grantee awsiam.IGrantable) awsiam.Grant {
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantWrite",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILogGroup) LogGroupPhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"logGroupPhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_ILogGroup) LogGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILogGroup) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}


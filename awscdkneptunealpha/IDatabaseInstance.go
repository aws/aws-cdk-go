package awscdkneptunealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdkneptunealpha/v2/internal"
)

// A database instance.
// Experimental.
type IDatabaseInstance interface {
	awscdk.IResource
	// Return the given named metric associated with this database instance.
	// See: https://docs.aws.amazon.com/neptune/latest/userguide/cw-dimensions.html
	//
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The instance endpoint address.
	// Experimental.
	DbInstanceEndpointAddress() *string
	// The instance endpoint port.
	// Experimental.
	DbInstanceEndpointPort() *string
	// The instance endpoint.
	// Experimental.
	InstanceEndpoint() Endpoint
	// The instance identifier.
	// Experimental.
	InstanceIdentifier() *string
}

// The jsii proxy for IDatabaseInstance
type jsiiProxy_IDatabaseInstance struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IDatabaseInstance) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (j *jsiiProxy_IDatabaseInstance) DbInstanceEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) DbInstanceEndpointPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceEndpointPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceEndpoint() Endpoint {
	var returns Endpoint
	_jsii_.Get(
		j,
		"instanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDatabaseInstance) InstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdentifier",
		&returns,
	)
	return returns
}


package awscdkapigatewayv2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/internal"
)

// Represents a API Gateway HTTP/WebSocket API.
// Deprecated.
type IApi interface {
	awscdk.IResource
	// Return the given named metric for this Api Gateway.
	// Default: - average over 5 minutes.
	//
	// Deprecated.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The default endpoint for an API.
	// Deprecated.
	ApiEndpoint() *string
	// The identifier of this API Gateway API.
	// Deprecated.
	ApiId() *string
}

// The jsii proxy for IApi
type jsiiProxy_IApi struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApi) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (j *jsiiProxy_IApi) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApi) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}


package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
)

// Represents a API Gateway HTTP/WebSocket API.
// Experimental.
type IApi interface {
	awscdk.IResource
	// Return the given named metric for this Api Gateway.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The default endpoint for an API.
	// Experimental.
	ApiEndpoint() *string
	// The identifier of this API Gateway API.
	// Experimental.
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


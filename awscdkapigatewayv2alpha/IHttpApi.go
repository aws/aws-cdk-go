package awscdkapigatewayv2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Represents an HTTP API.
// Experimental.
type IHttpApi interface {
	IApi
	// Add a new VpcLink.
	// Experimental.
	AddVpcLink(options *VpcLinkProps) VpcLink
	// Metric for the number of client-side errors captured in a given period.
	// Experimental.
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	// Experimental.
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the amount of data processed in bytes.
	// Experimental.
	MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	// Experimental.
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	// Experimental.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	// Experimental.
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The identifier of this API Gateway HTTP API.
	// Deprecated: - use apiId instead.
	HttpApiId() *string
}

// The jsii proxy for IHttpApi
type jsiiProxy_IHttpApi struct {
	jsiiProxy_IApi
}

func (i *jsiiProxy_IHttpApi) AddVpcLink(options *VpcLinkProps) VpcLink {
	if err := i.validateAddVpcLinkParameters(options); err != nil {
		panic(err)
	}
	var returns VpcLink

	_jsii_.Invoke(
		i,
		"addVpcLink",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricClientErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricClientError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDataProcessedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDataProcessed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIntegrationLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIntegrationLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IHttpApi) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricServerErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricServerError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IHttpApi) HttpApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpApiId",
		&returns,
	)
	return returns
}


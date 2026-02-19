package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the HttpStage.
type IHttpStage interface {
	IHttpStageRef
	IStage
	// Metric for the number of client-side errors captured in a given period.
	// Default: - sum over 5 minutes.
	//
	MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number API requests in a given period.
	// Default: - SampleCount over 5 minutes.
	//
	MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the amount of data processed in bytes.
	// Default: - sum over 5 minutes.
	//
	MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time between when API Gateway relays a request to the backend and when it receives a response from the backend.
	// Default: - no statistic.
	//
	MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The time between when API Gateway receives a request from a client and when it returns a response to the client.
	//
	// The latency includes the integration latency and other API Gateway overhead.
	// Default: - no statistic.
	//
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of server-side errors captured in a given period.
	// Default: - sum over 5 minutes.
	//
	MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The API this stage is associated to.
	Api() IHttpApi
	// The custom domain URL to this stage.
	DomainUrl() *string
}

// The jsii proxy for IHttpStage
type jsiiProxy_IHttpStage struct {
	jsiiProxy_IHttpStageRef
	jsiiProxy_IStage
}

func (i *jsiiProxy_IHttpStage) MetricClientError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IHttpStage) MetricCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IHttpStage) MetricDataProcessed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IHttpStage) MetricIntegrationLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IHttpStage) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IHttpStage) MetricServerError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IHttpStage) AddStageVariable(name *string, value *string) {
	if err := i.validateAddStageVariableParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addStageVariable",
		[]interface{}{name, value},
	)
}

func (i *jsiiProxy_IHttpStage) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IHttpStage) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IHttpStage) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IHttpStage) Api() IHttpApi {
	var returns IHttpApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpStage) DomainUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpStage) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpStage) IsHttpStage() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isHttpStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpStage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpStage) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpStage) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpStage) StageRef() *interfacesawsapigatewayv2.StageReference {
	var returns *interfacesawsapigatewayv2.StageReference
	_jsii_.Get(
		j,
		"stageRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpStage) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


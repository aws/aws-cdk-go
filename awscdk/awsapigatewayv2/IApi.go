package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a API Gateway HTTP/WebSocket API.
type IApi interface {
	interfacesawsapigatewayv2.IApiRef
	awscdk.IResource
	// Return the given named metric for this Api Gateway.
	// Default: - average over 5 minutes.
	//
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The default endpoint for an API.
	ApiEndpoint() *string
	// The identifier of this API Gateway API.
	ApiId() *string
}

// The jsii proxy for IApi
type jsiiProxy_IApi struct {
	internal.Type__interfacesawsapigatewayv2IApiRef
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

func (i *jsiiProxy_IApi) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IApi) ApiRef() *interfacesawsapigatewayv2.ApiReference {
	var returns *interfacesawsapigatewayv2.ApiReference
	_jsii_.Get(
		j,
		"apiRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApi) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApi) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


package awsmediaconnectalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
	"github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for Router Output.
// Experimental.
type IRouterOutput interface {
	awscdk.IResource
	interfacesawsmediaconnect.IRouterOutputRef
	// Create a CloudWatch metric for this router output.
	//
	// Router output metrics are dimensioned by `RouterOutputARN`. See the MediaConnect
	// documentation for available metric names (e.g. `RouterOutputBitRate`,
	// `RouterOutputTotalPackets`).
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of retransmitted packets requested through automatic repeat request (ARQ).
	//
	// Applies to RIST and SRT outputs only.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricArqRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the bitrate of the router output's payload.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricBitrate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the router output connection state (1 connected, 0 disconnected).
	//
	// Applies to SRT outputs only.
	// Default: - minimum over 60 seconds.
	//
	// Experimental.
	MetricConnected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number of packets sent by the router output.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricTotalPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The timestamp when the router output was created.
	// Experimental.
	CreatedAt() *string
	// The IP address of the router output.
	// Experimental.
	IpAddress() *string
	// The Amazon Resource Name (ARN) of the router output.
	// Experimental.
	RouterOutputArn() *string
	// The unique identifier of the router output.
	// Experimental.
	RouterOutputId() *string
	// The name of the router output.
	// Experimental.
	RouterOutputName() *string
	// The timestamp when the router output was last updated.
	// Experimental.
	UpdatedAt() *string
}

// The jsii proxy for IRouterOutput
type jsiiProxy_IRouterOutput struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsmediaconnectIRouterOutputRef
}

func (i *jsiiProxy_IRouterOutput) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRouterOutput) MetricArqRequests(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricArqRequestsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricArqRequests",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRouterOutput) MetricBitrate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricBitrateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricBitrate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRouterOutput) MetricConnected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricConnectedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricConnected",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRouterOutput) MetricTotalPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricTotalPacketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricTotalPackets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRouterOutput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRouterOutput) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRouterOutput) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutput) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutput) RouterOutputArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerOutputArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutput) RouterOutputId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerOutputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutput) RouterOutputName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerOutputName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutput) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutput) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutput) RouterOutputRef() *interfacesawsmediaconnect.RouterOutputReference {
	var returns *interfacesawsmediaconnect.RouterOutputReference
	_jsii_.Get(
		j,
		"routerOutputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutput) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


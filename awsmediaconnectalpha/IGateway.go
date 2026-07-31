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

// Interface for Gateway.
// Experimental.
type IGateway interface {
	interfacesawsmediaconnect.IGatewayRef
	awscdk.IResource
	// Create a CloudWatch metric for this gateway.
	//
	// The gateway ARN is set as the `GatewayARN` dimension automatically. For
	// metrics that require additional dimensions (such as `NetworkName`,
	// `InstanceId` for a specific appliance, or `BridgeSourceName`), pass them
	// via `props.dimensionsMap`.
	//
	// See the MediaConnect documentation for available metric names (for example
	// `IngressBridgeBitRate`, `EgressBridgeBitRate`, `IngressBridgeSourcePacketLossPercent`).
	// See: https://docs.aws.amazon.com/mediaconnect/latest/ug/monitor-with-cloudwatch-metrics-gateway-health.html
	//
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of packets dropped by egress bridges hosted on this gateway.
	//
	// Pair with {@link metricEgressBridgeTotalPackets} to chart a dropped-packet percentage.
	// For per-bridge data, pass `dimensionsMap: { BridgeSourceName: '...' }`.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricEgressBridgeDroppedPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number of packets sent from egress bridges hosted on this gateway.
	//
	// Pair with {@link metricEgressBridgeDroppedPackets} to chart an egress dropped-packet
	// percentage. For per-bridge data, pass `dimensionsMap: { BridgeSourceName: '...' }`.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricEgressBridgeTotalPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of packets dropped by ingress bridges hosted on this gateway.
	//
	// Pair with {@link metricIngressBridgeTotalPackets} to chart a dropped-packet percentage.
	// For per-bridge data, pass `dimensionsMap: { BridgeSourceName: '...' }`.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricIngressBridgeDroppedPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number of packets received by ingress bridges hosted on this gateway.
	//
	// Pair with {@link metricIngressBridgeDroppedPackets} to chart an ingress dropped-packet
	// percentage. For per-bridge data, pass `dimensionsMap: { BridgeSourceName: '...' }`.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricIngressBridgeTotalPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The Amazon Resource Name (ARN) of the gateway.
	// Experimental.
	GatewayArn() *string
	// The current state of the gateway.
	// Experimental.
	GatewayState() *string
}

// The jsii proxy for IGateway
type jsiiProxy_IGateway struct {
	internal.Type__interfacesawsmediaconnectIGatewayRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IGateway) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IGateway) MetricEgressBridgeDroppedPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEgressBridgeDroppedPacketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEgressBridgeDroppedPackets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) MetricEgressBridgeTotalPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEgressBridgeTotalPacketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEgressBridgeTotalPackets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) MetricIngressBridgeDroppedPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIngressBridgeDroppedPacketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIngressBridgeDroppedPackets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) MetricIngressBridgeTotalPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIngressBridgeTotalPacketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIngressBridgeTotalPackets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IGateway) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IGateway) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGateway) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) GatewayState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) GatewayRef() *interfacesawsmediaconnect.GatewayReference {
	var returns *interfacesawsmediaconnect.GatewayReference
	_jsii_.Get(
		j,
		"gatewayRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGateway) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


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

// Interface for Bridge.
// Experimental.
type IBridge interface {
	interfacesawsmediaconnect.IBridgeRef
	awscdk.IResource
	// Add a network output to this bridge (for egress bridges only).
	// Experimental.
	AddOutput(id *string, networkOutput *BridgeNetworkOutput) IBridgeOutput
	// Create a CloudWatch metric for this bridge.
	//
	// Bridge metrics are dimensioned by `BridgeARN`. See the MediaConnect
	// documentation for available metric names (e.g. `IngressBridgeBitRate`,
	// `EgressBridgeBitRate`, `IngressBridgePacketLossPercent`).
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number of times the bridge switches between sources when using the `FAILOVER` failover mode.
	//
	// Uses `IngressBridgeFailoverSwitches` for ingress bridges and
	// `EgressBridgeFailoverSwitches` for egress bridges.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricFailoverSwitches(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the bitrate of a specific bridge source.
	//
	// Uses `IngressBridgeSourceBitRate` for ingress bridges and
	// `EgressBridgeSourceBitRate` for egress bridges.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricSourceBitrate(bridgeSourceName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the percentage of packets lost on a specific bridge source.
	//
	// Uses `IngressBridgeSourcePacketLossPercent` for ingress bridges and
	// `EgressBridgeSourcePacketLossPercent` for egress bridges.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricSourcePacketLossPercent(bridgeSourceName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The Amazon Resource Name (ARN) of the bridge.
	// Experimental.
	BridgeArn() *string
	// The name of the bridge.
	// Experimental.
	BridgeName() *string
	// The current state of the bridge.
	// Experimental.
	BridgeState() *string
	// The type of bridge (ingress or egress).
	// Experimental.
	BridgeType() BridgeType
	// Failover Configuration for Bridge.
	// Experimental.
	IsFailoverEnabled() *bool
}

// The jsii proxy for IBridge
type jsiiProxy_IBridge struct {
	internal.Type__interfacesawsmediaconnectIBridgeRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IBridge) AddOutput(id *string, networkOutput *BridgeNetworkOutput) IBridgeOutput {
	if err := i.validateAddOutputParameters(id, networkOutput); err != nil {
		panic(err)
	}
	var returns IBridgeOutput

	_jsii_.Invoke(
		i,
		"addOutput",
		[]interface{}{id, networkOutput},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBridge) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IBridge) MetricFailoverSwitches(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFailoverSwitchesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFailoverSwitches",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBridge) MetricSourceBitrate(bridgeSourceName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourceBitrateParameters(bridgeSourceName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourceBitrate",
		[]interface{}{bridgeSourceName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBridge) MetricSourcePacketLossPercent(bridgeSourceName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourcePacketLossPercentParameters(bridgeSourceName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourcePacketLossPercent",
		[]interface{}{bridgeSourceName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IBridge) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IBridge) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IBridge) BridgeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bridgeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridge) BridgeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bridgeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridge) BridgeState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bridgeState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridge) BridgeType() BridgeType {
	var returns BridgeType
	_jsii_.Get(
		j,
		"bridgeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridge) IsFailoverEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFailoverEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridge) BridgeRef() *interfacesawsmediaconnect.BridgeReference {
	var returns *interfacesawsmediaconnect.BridgeReference
	_jsii_.Get(
		j,
		"bridgeRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridge) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridge) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBridge) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


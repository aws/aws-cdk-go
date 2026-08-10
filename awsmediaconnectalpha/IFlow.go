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

// Interface for Flow.
// Experimental.
type IFlow interface {
	interfacesawsmediaconnect.IFlowRef
	awscdk.IResource
	// Add an output to this flow.
	// Experimental.
	AddOutput(id *string, options *AddFlowOutputOptions) IFlowOutput
	// Create a CloudWatch metric for this flow.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the bitrate of content ingested into the flow.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricSourceBitrate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric indicating the connection state of the source (1 connected, 0 disconnected).
	//
	// Applies only to Zixi, SRT, and RIST sources.
	// Default: - minimum over 60 seconds.
	//
	// Experimental.
	MetricSourceConnected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times the source transitioned from connected to disconnected.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricSourceDisconnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of packets lost during transit, measured before any error correction.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricSourceDroppedPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the current network jitter of the source, measured in milliseconds.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricSourceJitter(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for packets that were not recovered by the flow source.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricSourceNotRecoveredPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the percentage of packets lost during transit, even if they were recovered.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricSourcePacketLossPercent(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the round-trip time between the source and MediaConnect.
	//
	// Applies only to RIST, Zixi, and SRT sources.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricSourceRoundTripTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric indicating which source is selected for ingest when using `Failover` failover mode.
	// Default: - maximum over 60 seconds.
	//
	// Experimental.
	MetricSourceSelected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number of packets received by the flow source.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricSourceTotalPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The IP address that the flow uses to send outbound content.
	// Experimental.
	EgressIp() *string
	// The Amazon Resource Name (ARN) of the flow.
	// Experimental.
	FlowArn() *string
	// The Availability Zone that the flow was created in.
	// Experimental.
	FlowAvailabilityZone() *string
	// Collection of grant methods for this flow.
	// Experimental.
	Grants() FlowGrants
	// Failover Configuration for flow.
	// Experimental.
	IsFailoverEnabled() *bool
	// The Amazon Resource Name (ARN) of the source defined on the flow.
	// Experimental.
	SourceArn() *string
	// The IP address that the flow listens on for incoming content.
	//
	// Available for listener-style source protocols (RTP, RTP-FEC, RIST, SRT listener,
	// Zixi push). Accessing this on SRT caller, entitlement, gateway bridge, router,
	// CDI, JPEG XS, or imported flows throws — those sources don't expose a listening
	// IP address.
	// Experimental.
	SourceIngestIp() *string
	// The port that the flow listens on for incoming content.
	//
	// Available for the same listener-style source protocols as `sourceIngestIp`.
	// Accessing this on SRT caller, entitlement, gateway bridge, router, CDI, JPEG XS,
	// or imported flows throws.
	// Experimental.
	SourceIngestPort() *string
	// The full ingest URL for the flow source, combining protocol, IP, and port. For example: `srt://203.0.113.10:5000`.
	//
	// Available for listener-style source protocols (RTP, RTP-FEC, RIST, SRT listener,
	// Zixi push) where the flow listens for an upstream sender. Accessing this on
	// SRT caller, entitlement, gateway bridge, router, CDI, JPEG XS, or imported
	// flows throws — those sources don't expose a single host:port ingest URL.
	// Experimental.
	SourceIngestUrl() *string
}

// The jsii proxy for IFlow
type jsiiProxy_IFlow struct {
	internal.Type__interfacesawsmediaconnectIFlowRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IFlow) AddOutput(id *string, options *AddFlowOutputOptions) IFlowOutput {
	if err := i.validateAddOutputParameters(id, options); err != nil {
		panic(err)
	}
	var returns IFlowOutput

	_jsii_.Invoke(
		i,
		"addOutput",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IFlow) MetricSourceBitrate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourceBitrateParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourceBitrate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricSourceConnected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourceConnectedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourceConnected",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricSourceDisconnections(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourceDisconnectionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourceDisconnections",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricSourceDroppedPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourceDroppedPacketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourceDroppedPackets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricSourceJitter(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourceJitterParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourceJitter",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricSourceNotRecoveredPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourceNotRecoveredPacketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourceNotRecoveredPackets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricSourcePacketLossPercent(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourcePacketLossPercentParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourcePacketLossPercent",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricSourceRoundTripTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourceRoundTripTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourceRoundTripTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricSourceSelected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourceSelectedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourceSelected",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) MetricSourceTotalPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSourceTotalPacketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSourceTotalPackets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFlow) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IFlow) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFlow) EgressIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) FlowArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) FlowAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) Grants() FlowGrants {
	var returns FlowGrants
	_jsii_.Get(
		j,
		"grants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) IsFailoverEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isFailoverEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) SourceIngestIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIngestIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) SourceIngestPort() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIngestPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) SourceIngestUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIngestUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) FlowRef() *interfacesawsmediaconnect.FlowReference {
	var returns *interfacesawsmediaconnect.FlowReference
	_jsii_.Get(
		j,
		"flowRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlow) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


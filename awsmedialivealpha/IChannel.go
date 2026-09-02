package awsmedialivealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive"
	"github.com/aws/aws-cdk-go/awsmedialivealpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a MediaLive Channel.
// Experimental.
type IChannel interface {
	interfacesawsmedialive.IChannelRef
	awscdk.IResource
	// Create a CloudWatch metric for this channel scoped to a specific pipeline.
	//
	// Channel metrics are published per-pipeline. `STANDARD` channels run two
	// redundant pipelines (`PIPELINE_0` and `PIPELINE_1`); to cover both, build
	// a metric for each. `SINGLE_PIPELINE` channels only publish on `PIPELINE_0`.
	// See the
	// {@link https://docs.aws.amazon.com/medialive/latest/ug/monitoring-eml-metrics.html MediaLive metrics docs}
	// for the full set of metric names and recommended statistics.
	// Experimental.
	Metric(metricName *string, pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number of active alerts on this channel.
	// Default: - max over 5 minutes.
	//
	// Experimental.
	MetricActiveAlerts(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for dropped frames.
	//
	// A non-zero value indicates the encoder cannot
	// keep up with the incoming video in real time.
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricDroppedFrames(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for fill milliseconds — the time MediaLive has filled the video output with fill frames because the input did not deliver content within the expected window.
	//
	// A non-zero value indicates an unhealthy input.
	// Default: - max over 5 minutes.
	//
	// Experimental.
	MetricFillMsec(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for input loss seconds (RTP and MediaConnect inputs only).
	// Default: - sum over 5 minutes.
	//
	// Experimental.
	MetricInputLossSeconds(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the input video frame rate (frames per second).
	// Default: - max over 5 minutes.
	//
	// Experimental.
	MetricInputVideoFrameRate(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the rate of inbound network traffic to MediaLive in Mbps.
	// Default: - average over 5 minutes.
	//
	// Experimental.
	MetricNetworkIn(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the rate of outbound network traffic from MediaLive in Mbps.
	// Default: - average over 5 minutes.
	//
	// Experimental.
	MetricNetworkOut(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for SVQ time (speed-vs-quality), expressed as a percent.
	//
	// Indicates
	// the share of time MediaLive reduced quality optimisations to keep up with
	// real-time output.
	// Default: - max over 5 minutes.
	//
	// Experimental.
	MetricSvqTime(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The ARN of the channel.
	// Experimental.
	ChannelArn() *string
	// The ID of the channel.
	// Experimental.
	ChannelId() *string
	// The IDs of the inputs attached to this channel.
	// Experimental.
	ChannelInputs() *[]*string
	// Collection of grant methods for this channel — start, stop, and update its schedule.
	// Experimental.
	Grants() ChannelGrants
}

// The jsii proxy for IChannel
type jsiiProxy_IChannel struct {
	internal.Type__interfacesawsmedialiveIChannelRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IChannel) Metric(metricName *string, pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, pipeline, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, pipeline, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricActiveAlerts(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricActiveAlertsParameters(pipeline, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricActiveAlerts",
		[]interface{}{pipeline, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricDroppedFrames(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDroppedFramesParameters(pipeline, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDroppedFrames",
		[]interface{}{pipeline, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricFillMsec(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricFillMsecParameters(pipeline, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricFillMsec",
		[]interface{}{pipeline, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricInputLossSeconds(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInputLossSecondsParameters(pipeline, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInputLossSeconds",
		[]interface{}{pipeline, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricInputVideoFrameRate(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInputVideoFrameRateParameters(pipeline, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInputVideoFrameRate",
		[]interface{}{pipeline, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricNetworkIn(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNetworkInParameters(pipeline, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkIn",
		[]interface{}{pipeline, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricNetworkOut(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNetworkOutParameters(pipeline, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNetworkOut",
		[]interface{}{pipeline, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricSvqTime(pipeline Pipeline, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricSvqTimeParameters(pipeline, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricSvqTime",
		[]interface{}{pipeline, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IChannel) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IChannel) ChannelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) ChannelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) ChannelInputs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"channelInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) Grants() ChannelGrants {
	var returns ChannelGrants
	_jsii_.Get(
		j,
		"grants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) ChannelRef() *interfacesawsmedialive.ChannelReference {
	var returns *interfacesawsmedialive.ChannelReference
	_jsii_.Get(
		j,
		"channelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


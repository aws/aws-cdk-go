package awsmediaconnectalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect"
	"github.com/aws/aws-cdk-go/awsmediaconnectalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for Router Input.
// Experimental.
type IRouterInput interface {
	awscdk.IResource
	interfacesawsmediaconnect.IRouterInputRef
	// Create a CloudWatch metric for this router input.
	//
	// Router input metrics are dimensioned by `RouterInputARN`. See the MediaConnect
	// documentation for available metric names (e.g. `RouterInputBitRate`,
	// `RouterInputNotRecoveredPackets`).
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the bitrate of the router input's payload.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricBitrate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the router input connection state (1 connected, 0 disconnected).
	//
	// Applies to SRT sources only.
	// Default: - minimum over 60 seconds.
	//
	// Experimental.
	MetricConnected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for continuity counter errors in the transport stream.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricContinuityCounterErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times the router input has switched sources in Failover mode.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricFailoverSwitches(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the recovery latency of the input stream.
	//
	// Applies to RIST, SRT, and RTP-FEC.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for packets lost in transit that were not recovered by error correction.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricNotRecoveredPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the total number of packets received by the router input.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricTotalPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The timestamp when the router input was created.
	// Experimental.
	CreatedAt() *string
	// The ingest endpoints (URL + port) where the router input listens.
	//
	// Returns one entry for standard protocol-based variants (RTP, RIST, SRT listener),
	// and one entry per source for failover and merge configurations built from those
	// protocols. For example a failover RTP input will return:
	//
	// ```
	// [
	//   { url: 'rtp://203.0.113.10:5000', port: 5000 },
	//   { url: 'rtp://203.0.113.10:5001', port: 5001 },
	// ]
	// ```
	//
	// Accessing this on SRT caller (where the router input dials out to a remote
	// source), MediaConnect Flow, MediaLive Channel, or imported inputs throws — those
	// variants do not expose host:port pairs the input listens on.
	// Experimental.
	Endpoints() *[]*RouterInputEndpoint
	// Collection of grant methods for this router input.
	// Experimental.
	Grants() RouterInputGrants
	// The IP address of the router input.
	// Experimental.
	IpAddress() *string
	// The Amazon Resource Name (ARN) of the router input.
	// Experimental.
	RouterInputArn() *string
	// The unique identifier of the router input.
	// Experimental.
	RouterInputId() *string
	// The Secrets Manager secret containing the transit encryption passphrase.
	//
	// Only available when the Router Input was created with explicit
	// `transitEncryption` configuration. Not available for
	// automatic encryption or imported inputs.
	// Experimental.
	TransitEncryptionSecret() awssecretsmanager.ISecret
	// The timestamp when the router input was last updated.
	// Experimental.
	UpdatedAt() *string
}

// The jsii proxy for IRouterInput
type jsiiProxy_IRouterInput struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsmediaconnectIRouterInputRef
}

func (i *jsiiProxy_IRouterInput) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRouterInput) MetricBitrate(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRouterInput) MetricConnected(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRouterInput) MetricContinuityCounterErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricContinuityCounterErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricContinuityCounterErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRouterInput) MetricFailoverSwitches(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRouterInput) MetricLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRouterInput) MetricNotRecoveredPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricNotRecoveredPacketsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricNotRecoveredPackets",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IRouterInput) MetricTotalPackets(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IRouterInput) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRouterInput) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRouterInput) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) Endpoints() *[]*RouterInputEndpoint {
	var returns *[]*RouterInputEndpoint
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) Grants() RouterInputGrants {
	var returns RouterInputGrants
	_jsii_.Get(
		j,
		"grants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) RouterInputArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerInputArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) RouterInputId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerInputId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) TransitEncryptionSecret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"transitEncryptionSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) RouterInputRef() *interfacesawsmediaconnect.RouterInputReference {
	var returns *interfacesawsmediaconnect.RouterInputReference
	_jsii_.Get(
		j,
		"routerInputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInput) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


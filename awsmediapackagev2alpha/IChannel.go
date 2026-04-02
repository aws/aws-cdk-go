package awsmediapackagev2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackagev2"
	"github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a MediaPackage V2 Channel.
// Experimental.
type IChannel interface {
	interfacesawsmediapackagev2.IChannelRef
	awscdk.IResource
	// Add Origin Endpoint for this Channel.
	// Experimental.
	AddOriginEndpoint(id *string, options *OriginEndpointOptions) OriginEndpoint
	// Configure channel policy.
	//
	// You can only add 1 ChannelPolicy to a Channel.
	// If you have already defined one, function will append the policy already created.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
	// Create a CloudWatch metric.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Egress Bytes.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricEgressBytes(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Egress Request Count.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricEgressRequestCount(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Egress Response time.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricEgressResponseTime(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Ingress Bytes.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricIngressBytes(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Ingress Request Count.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricIngressRequestCount(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Ingress response time.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricIngressResponseTime(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The Amazon Resource Name (ARN) associated with the resource.
	// Experimental.
	ChannelArn() *string
	// The channel group this channel belongs to.
	//
	// Only available when the channel was created in the same stack.
	// Undefined for imported channels.
	// Experimental.
	ChannelGroup() IChannelGroup
	// The name that describes the channel group.
	//
	// The name is the primary identifier for the channel group.
	// Experimental.
	ChannelGroupName() *string
	// The name that describes the channel.
	//
	// The name is the primary identifier for the channel.
	// Experimental.
	ChannelName() *string
	// The date and time the channel was created.
	// Experimental.
	CreatedAt() *string
	// Grants IAM resource policy to the role used to write to MediaPackage V2 Channel.
	// Experimental.
	Grants() ChannelGrants
	// The ingest endpoint URLs for the channel.
	// Experimental.
	IngestEndpointUrls() *[]*string
	// The date and time the channel was modified.
	// Experimental.
	ModifiedAt() *string
}

// The jsii proxy for IChannel
type jsiiProxy_IChannel struct {
	internal.Type__interfacesawsmediapackagev2IChannelRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IChannel) AddOriginEndpoint(id *string, options *OriginEndpointOptions) OriginEndpoint {
	if err := i.validateAddOriginEndpointParameters(id, options); err != nil {
		panic(err)
	}
	var returns OriginEndpoint

	_jsii_.Invoke(
		i,
		"addOriginEndpoint",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := i.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IChannel) MetricEgressBytes(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEgressBytesParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEgressBytes",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricEgressRequestCount(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEgressRequestCountParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEgressRequestCount",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricEgressResponseTime(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEgressResponseTimeParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEgressResponseTime",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricIngressBytes(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIngressBytesParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIngressBytes",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricIngressRequestCount(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIngressRequestCountParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIngressRequestCount",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannel) MetricIngressResponseTime(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIngressResponseTimeParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIngressResponseTime",
		[]interface{}{options},
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

func (j *jsiiProxy_IChannel) ChannelGroup() IChannelGroup {
	var returns IChannelGroup
	_jsii_.Get(
		j,
		"channelGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) ChannelGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
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

func (j *jsiiProxy_IChannel) IngestEndpointUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingestEndpointUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannel) ChannelRef() *interfacesawsmediapackagev2.ChannelReference {
	var returns *interfacesawsmediapackagev2.ChannelReference
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


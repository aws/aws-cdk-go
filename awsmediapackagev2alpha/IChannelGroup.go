package awsmediapackagev2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediapackagev2"
	"github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Interface for AWS Elemental MediaPackage V2 Channel Group.
// Experimental.
type IChannelGroup interface {
	interfacesawsmediapackagev2.IChannelGroupRef
	awscdk.IResource
	// Add Channel for this Channel Group.
	// Experimental.
	AddChannel(id *string, options *ChannelOptions) Channel
	// Create a CloudWatch metric.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Egress Bytes.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricEgressBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Egress Request Count.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricEgressRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Egress Response time.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricEgressResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Ingress Bytes.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricIngressBytes(options *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Ingress Request Count.
	// Default: - sum over 60 seconds.
	//
	// Experimental.
	MetricIngressRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns Metric for Ingress response time.
	// Default: - average over 60 seconds.
	//
	// Experimental.
	MetricIngressResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The Amazon Resource Name (ARN) associated with the resource.
	// Experimental.
	ChannelGroupArn() *string
	// The name that describes the channel group.
	//
	// The name is the primary identifier for the channel group.
	// Experimental.
	ChannelGroupName() *string
	// The date and time the channel group was created.
	// Experimental.
	CreatedAt() *string
	// The egress domain where packaged content is available.
	//
	// Use this as the origin domain when configuring a CDN such as Amazon CloudFront.
	// Experimental.
	EgressDomain() *string
	// The date and time the channel group was modified.
	// Experimental.
	ModifiedAt() *string
}

// The jsii proxy for IChannelGroup
type jsiiProxy_IChannelGroup struct {
	internal.Type__interfacesawsmediapackagev2IChannelGroupRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IChannelGroup) AddChannel(id *string, options *ChannelOptions) Channel {
	if err := i.validateAddChannelParameters(id, options); err != nil {
		panic(err)
	}
	var returns Channel

	_jsii_.Invoke(
		i,
		"addChannel",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannelGroup) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IChannelGroup) MetricEgressBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEgressBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEgressBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannelGroup) MetricEgressRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEgressRequestCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEgressRequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannelGroup) MetricEgressResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricEgressResponseTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricEgressResponseTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannelGroup) MetricIngressBytes(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IChannelGroup) MetricIngressRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIngressRequestCountParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIngressRequestCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannelGroup) MetricIngressResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricIngressResponseTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricIngressResponseTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IChannelGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IChannelGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IChannelGroup) ChannelGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelGroup) ChannelGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelGroup) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelGroup) EgressDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"egressDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelGroup) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelGroup) ChannelGroupRef() *interfacesawsmediapackagev2.ChannelGroupReference {
	var returns *interfacesawsmediapackagev2.ChannelGroupReference
	_jsii_.Get(
		j,
		"channelGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelGroup) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IChannelGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


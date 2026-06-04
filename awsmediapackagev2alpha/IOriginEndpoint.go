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

// Origin Endpoint interface.
// Experimental.
type IOriginEndpoint interface {
	interfacesawsmediapackagev2.IOriginEndpointRef
	awscdk.IResource
	// Configure origin endpoint policy.
	//
	// You can only add 1 OriginEndpointPolicy to an OriginEndpoint.
	// If you have already defined one, it will append to the policy already created.
	// Experimental.
	AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
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
	// The name of the channel group associated with the origin endpoint configuration.
	// Experimental.
	ChannelGroupName() *string
	// The channel name associated with the origin endpoint.
	// Experimental.
	ChannelName() *string
	// The date and time the origin endpoint was created.
	// Experimental.
	CreatedAt() *string
	// The DASH manifest URLs for the origin endpoint.
	// Experimental.
	DashManifestUrls() *[]*string
	// The HLS manifest URLs for the origin endpoint.
	// Experimental.
	HlsManifestUrls() *[]*string
	// The Low Latency HLS manifest URLs for the origin endpoint.
	// Experimental.
	LowLatencyHlsManifestUrls() *[]*string
	// The date and time the origin endpoint was modified.
	// Experimental.
	ModifiedAt() *string
	// The MSS manifest URLs for the origin endpoint.
	// Experimental.
	MssManifestUrls() *[]*string
	// The Amazon Resource Name (ARN) of the origin endpoint.
	// Experimental.
	OriginEndpointArn() *string
	// The name of the origin endpoint associated with the origin endpoint configuration.
	// Experimental.
	OriginEndpointName() *string
}

// The jsii proxy for IOriginEndpoint
type jsiiProxy_IOriginEndpoint struct {
	internal.Type__interfacesawsmediapackagev2IOriginEndpointRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IOriginEndpoint) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
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

func (i *jsiiProxy_IOriginEndpoint) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IOriginEndpoint) MetricEgressBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IOriginEndpoint) MetricEgressRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IOriginEndpoint) MetricEgressResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IOriginEndpoint) MetricIngressBytes(options *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IOriginEndpoint) MetricIngressRequestCount(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IOriginEndpoint) MetricIngressResponseTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (i *jsiiProxy_IOriginEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IOriginEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IOriginEndpoint) ChannelGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) DashManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dashManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) HlsManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hlsManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) LowLatencyHlsManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"lowLatencyHlsManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) ModifiedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) MssManifestUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mssManifestUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) OriginEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) OriginEndpointName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originEndpointName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) OriginEndpointRef() *interfacesawsmediapackagev2.OriginEndpointReference {
	var returns *interfacesawsmediapackagev2.OriginEndpointReference
	_jsii_.Get(
		j,
		"originEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOriginEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


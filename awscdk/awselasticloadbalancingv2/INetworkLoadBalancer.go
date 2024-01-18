package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A network load balancer.
type INetworkLoadBalancer interface {
	awsec2.IConnectable
	ILoadBalancerV2
	awsec2.IVpcEndpointServiceLoadBalancer
	// Add a listener to this load balancer.
	//
	// Returns: The newly created listener.
	AddListener(id *string, props *BaseNetworkListenerProps) NetworkListener
	// The type of IP addresses to use.
	// Default: IpAddressType.IPV4
	//
	IpAddressType() IpAddressType
	// All metrics available for this load balancer.
	Metrics() INetworkLoadBalancerMetrics
	// Security groups associated with this load balancer.
	SecurityGroups() *[]*string
	// The VPC this load balancer has been created in (if available).
	Vpc() awsec2.IVpc
}

// The jsii proxy for INetworkLoadBalancer
type jsiiProxy_INetworkLoadBalancer struct {
	internal.Type__awsec2IConnectable
	jsiiProxy_ILoadBalancerV2
	internal.Type__awsec2IVpcEndpointServiceLoadBalancer
}

func (i *jsiiProxy_INetworkLoadBalancer) AddListener(id *string, props *BaseNetworkListenerProps) NetworkListener {
	if err := i.validateAddListenerParameters(id, props); err != nil {
		panic(err)
	}
	var returns NetworkListener

	_jsii_.Invoke(
		i,
		"addListener",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_INetworkLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_INetworkLoadBalancer) IpAddressType() IpAddressType {
	var returns IpAddressType
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) Metrics() INetworkLoadBalancerMetrics {
	var returns INetworkLoadBalancerMetrics
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


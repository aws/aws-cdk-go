package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// An application load balancer.
type IApplicationLoadBalancer interface {
	awsec2.IConnectable
	ILoadBalancerV2
	// Add a new listener to this load balancer.
	AddListener(id *string, props *BaseApplicationListenerProps) ApplicationListener
	// The IP Address Type for this load balancer.
	// Default: IpAddressType.IPV4
	//
	IpAddressType() IpAddressType
	// A list of listeners that have been added to the load balancer.
	//
	// This list is only valid for owned constructs.
	Listeners() *[]ApplicationListener
	// The ARN of this load balancer.
	LoadBalancerArn() *string
	// All metrics available for this load balancer.
	Metrics() IApplicationLoadBalancerMetrics
	// The VPC this load balancer has been created in (if available).
	//
	// If this interface is the result of an import call to fromApplicationLoadBalancerAttributes,
	// the vpc attribute will be undefined unless specified in the optional properties of that method.
	Vpc() awsec2.IVpc
}

// The jsii proxy for IApplicationLoadBalancer
type jsiiProxy_IApplicationLoadBalancer struct {
	internal.Type__awsec2IConnectable
	jsiiProxy_ILoadBalancerV2
}

func (i *jsiiProxy_IApplicationLoadBalancer) AddListener(id *string, props *BaseApplicationListenerProps) ApplicationListener {
	if err := i.validateAddListenerParameters(id, props); err != nil {
		panic(err)
	}
	var returns ApplicationListener

	_jsii_.Invoke(
		i,
		"addListener",
		[]interface{}{id, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApplicationLoadBalancer) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IApplicationLoadBalancer) IpAddressType() IpAddressType {
	var returns IpAddressType
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Listeners() *[]ApplicationListener {
	var returns *[]ApplicationListener
	_jsii_.Get(
		j,
		"listeners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Metrics() IApplicationLoadBalancerMetrics {
	var returns IApplicationLoadBalancerMetrics
	_jsii_.Get(
		j,
		"metrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Vpc() awsec2.IVpc {
	var returns awsec2.IVpc
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) LoadBalancerCanonicalHostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerCanonicalHostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) LoadBalancerDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationLoadBalancer) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


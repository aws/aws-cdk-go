package awselasticloadbalancing

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// An EC2 instance that is the target for load balancing.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
//   	Vpc: Vpc,
//   })
//   // instance to add as the target for load balancer.
//   instance := NewInstance(stack, jsii.String("targetInstance"), map[string]interface{}{
//   	"vpc": vpc,
//   	"instanceType": InstanceType_of(InstanceClass_BURSTABLE2, InstanceSize_MICRO),
//   	"machineImage": NewAmazonLinuxImage(),
//   })
//   lb.AddTarget(elb.InstanceTarget(instance))
//
type InstanceTarget interface {
	ILoadBalancerTarget
	// The network connections associated with this resource.
	Connections() awsec2.Connections
	// Instance to register to.
	Instance() awsec2.Instance
	// Attach load-balanced target to a classic ELB.
	AttachToClassicLB(loadBalancer LoadBalancer)
}

// The jsii proxy struct for InstanceTarget
type jsiiProxy_InstanceTarget struct {
	jsiiProxy_ILoadBalancerTarget
}

func (j *jsiiProxy_InstanceTarget) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceTarget) Instance() awsec2.Instance {
	var returns awsec2.Instance
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}


// Create a new Instance target.
func NewInstanceTarget(instance awsec2.Instance) InstanceTarget {
	_init_.Initialize()

	if err := validateNewInstanceTargetParameters(instance); err != nil {
		panic(err)
	}
	j := jsiiProxy_InstanceTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.InstanceTarget",
		[]interface{}{instance},
		&j,
	)

	return &j
}

// Create a new Instance target.
func NewInstanceTarget_Override(i InstanceTarget, instance awsec2.Instance) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancing.InstanceTarget",
		[]interface{}{instance},
		i,
	)
}

func (i *jsiiProxy_InstanceTarget) AttachToClassicLB(loadBalancer LoadBalancer) {
	if err := i.validateAttachToClassicLBParameters(loadBalancer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}


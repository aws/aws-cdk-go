package awselasticloadbalancing

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// An EC2 instance that is the target for load balancing.
//
// Example:
//   var vpc iVpc
//
//   lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
//   	Vpc: Vpc,
//   	InternetFacing: jsii.Boolean(true),
//   })
//
//   // instance to add as the target for load balancer.
//   instance := ec2.NewInstance(this, jsii.String("targetInstance"), &InstanceProps{
//   	Vpc: vpc,
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
//   	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
//   		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
//   	}),
//   })
//   lb.AddTarget(elb.NewInstanceTarget(instance))
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


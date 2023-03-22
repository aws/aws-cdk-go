# Amazon Elastic Load Balancing Construct Library

The `@aws-cdk/aws-elasticloadbalancing` package provides constructs for configuring
classic load balancers.

## Configuring a Load Balancer

Load balancers send traffic to one or more AutoScalingGroups. Create a load
balancer, set up listeners and a health check, and supply the fleet(s) you want
to load balance to in the `targets` property.

```go
var vpc iVpc

var myAutoScalingGroup autoScalingGroup

lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
	HealthCheck: &HealthCheck{
		Port: jsii.Number(80),
	},
})
lb.AddTarget(myAutoScalingGroup)
lb.AddListener(&LoadBalancerListener{
	ExternalPort: jsii.Number(80),
})
```

The load balancer allows all connections by default. If you want to change that,
pass the `allowConnectionsFrom` property while setting up the listener:

```go
var mySecurityGroup securityGroup
var lb loadBalancer

lb.AddListener(&LoadBalancerListener{
	ExternalPort: jsii.Number(80),
	AllowConnectionsFrom: []iConnectable{
		mySecurityGroup,
	},
})
```

### Adding Ec2 Instance as a target for the load balancer

You can add an EC2 instance to the load balancer by calling using `new InstanceTarget` as the argument to `addTarget()`:

```go
// Example automatically generated from non-compiling source. May contain errors.
lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
	Vpc: Vpc,
})
// instance to add as the target for load balancer.
instance := NewInstance(stack, jsii.String("targetInstance"), map[string]interface{}{
	"vpc": vpc,
	"instanceType": InstanceType_of(InstanceClass_BURSTABLE2, InstanceSize_MICRO),
	"machineImage": NewAmazonLinuxImage(),
})
lb.AddTarget(elb.InstanceTarget(instance))
```

# Amazon Elastic Load Balancing Construct Library

The `aws-cdk-lib/aws-elasticloadbalancing` package provides constructs for configuring
classic load balancers.

## Configuring a Load Balancer

Load balancers send traffic to one or more AutoScalingGroups. Create a load
balancer, set up listeners and a health check, and supply the fleet(s) you want
to load balance to in the `targets` property. If you want the load balancer to be
accessible from the internet, set `internetFacing: true`.

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
var vpc iVpc

lb := elb.NewLoadBalancer(this, jsii.String("LB"), &LoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
})

// instance to add as the target for load balancer.
instance := ec2.NewInstance(this, jsii.String("targetInstance"), &InstanceProps{
	Vpc: vpc,
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_BURSTABLE2, ec2.InstanceSize_MICRO),
	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
	}),
})
lb.AddTarget(elb.NewInstanceTarget(instance))
```

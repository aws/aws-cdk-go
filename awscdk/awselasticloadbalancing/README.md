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

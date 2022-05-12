# AWS::GlobalAccelerator Construct Library

## Introduction

AWS Global Accelerator (AGA) is a service that improves the availability and
performance of your applications with local or global users.

It intercepts your user's network connection at an edge location close to
them, and routes it to one of potentially multiple, redundant backends across
the more reliable and less congested AWS global network.

AGA can be used to route traffic to Application Load Balancers, Network Load
Balancers, EC2 Instances and Elastic IP Addresses.

For more information, see the [AWS Global
Accelerator Developer Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/AWS_GlobalAccelerator.html).

## Example

Here's an example that sets up a Global Accelerator for two Application Load
Balancers in two different AWS Regions:

```go
// Create an Accelerator
accelerator := globalaccelerator.NewAccelerator(this, jsii.String("Accelerator"))

// Create a Listener
listener := accelerator.addListener(jsii.String("Listener"), &listenerOptions{
	portRanges: []portRange{
		&portRange{
			fromPort: jsii.Number(80),
		},
		&portRange{
			fromPort: jsii.Number(443),
		},
	},
})

// Import the Load Balancers
nlb1 := elbv2.networkLoadBalancer.fromNetworkLoadBalancerAttributes(this, jsii.String("NLB1"), &networkLoadBalancerAttributes{
	loadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-west-2:111111111111:loadbalancer/app/my-load-balancer1/e16bef66805b"),
})
nlb2 := elbv2.networkLoadBalancer.fromNetworkLoadBalancerAttributes(this, jsii.String("NLB2"), &networkLoadBalancerAttributes{
	loadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:ap-south-1:111111111111:loadbalancer/app/my-load-balancer2/5513dc2ea8a1"),
})

// Add one EndpointGroup for each Region we are targeting
listener.addEndpointGroup(jsii.String("Group1"), &endpointGroupOptions{
	endpoints: []iEndpoint{
		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb1),
	},
})
listener.addEndpointGroup(jsii.String("Group2"), &endpointGroupOptions{
	// Imported load balancers automatically calculate their Region from the ARN.
	// If you are load balancing to other resources, you must also pass a `region`
	// parameter here.
	endpoints: []*iEndpoint{
		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb2),
	},
})
```

## Concepts

The **Accelerator** construct defines a Global Accelerator resource.

An Accelerator includes one or more **Listeners** that accepts inbound
connections on one or more ports.

Each Listener has one or more **Endpoint Groups**, representing multiple
geographically distributed copies of your application. There is one Endpoint
Group per Region, and user traffic is routed to the closest Region by default.

An Endpoint Group consists of one or more **Endpoints**, which is where the
user traffic coming in on the Listener is ultimately sent. The Endpoint port
used is the same as the traffic came in on at the Listener, unless overridden.

## Types of Endpoints

There are 4 types of Endpoints, and they can be found in the
`@aws-cdk/aws-globalaccelerator-endpoints` package:

* Application Load Balancers
* Network Load Balancers
* EC2 Instances
* Elastic IP Addresses

### Application Load Balancers

```go
var alb applicationLoadBalancer
var listener listener


listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
	endpoints: []iEndpoint{
		ga_endpoints.NewApplicationLoadBalancerEndpoint(alb, &applicationLoadBalancerEndpointOptions{
			weight: jsii.Number(128),
			preserveClientIp: jsii.Boolean(true),
		}),
	},
})
```

### Network Load Balancers

```go
var nlb networkLoadBalancer
var listener listener


listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
	endpoints: []iEndpoint{
		ga_endpoints.NewNetworkLoadBalancerEndpoint(nlb, &networkLoadBalancerEndpointProps{
			weight: jsii.Number(128),
		}),
	},
})
```

### EC2 Instances

```go
var listener listener
var instance instance


listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
	endpoints: []iEndpoint{
		ga_endpoints.NewInstanceEndpoint(instance, &instanceEndpointProps{
			weight: jsii.Number(128),
			preserveClientIp: jsii.Boolean(true),
		}),
	},
})
```

### Elastic IP Addresses

```go
var listener listener
var eip cfnEIP


listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
	endpoints: []iEndpoint{
		ga_endpoints.NewCfnEipEndpoint(eip, &cfnEipEndpointProps{
			weight: jsii.Number(128),
		}),
	},
})
```

## Client IP Address Preservation and Security Groups

When using the `preserveClientIp` feature, AGA creates
**Elastic Network Interfaces** (ENIs) in your AWS account, that are
associated with a Security Group AGA creates for you. You can use the
security group created by AGA as a source group in other security groups
(such as those for EC2 instances or Elastic Load Balancers), if you want to
restrict incoming traffic to the AGA security group rules.

AGA creates a specific security group called `GlobalAccelerator` for each VPC
it has an ENI in (this behavior can not be changed). CloudFormation doesn't
support referencing the security group created by AGA, but this construct
library comes with a custom resource that enables you to reference the AGA
security group.

Call `endpointGroup.connectionsPeer()` to obtain a reference to the Security Group
which you can use in connection rules. You must pass a reference to the VPC in whose
context the security group will be looked up. Example:

```go
var listener listener

// Non-open ALB
var alb applicationLoadBalancer

// Remember that there is only one AGA security group per VPC.
var vpc vpc


endpointGroup := listener.addEndpointGroup(jsii.String("Group"), &endpointGroupOptions{
	endpoints: []iEndpoint{
		ga_endpoints.NewApplicationLoadBalancerEndpoint(alb, &applicationLoadBalancerEndpointOptions{
			preserveClientIp: jsii.Boolean(true),
		}),
	},
})
agaSg := endpointGroup.connectionsPeer(jsii.String("GlobalAcceleratorSG"), vpc)

// Allow connections from the AGA to the ALB
alb.connections.allowFrom(agaSg, ec2.port.tcp(jsii.Number(443)))
```

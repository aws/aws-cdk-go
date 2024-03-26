# Amazon Elastic Load Balancing V2 Construct Library

The `aws-cdk-lib/aws-elasticloadbalancingv2` package provides constructs for
configuring application and network load balancers.

For more information, see the AWS documentation for
[Application Load Balancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/introduction.html)
and [Network Load Balancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/introduction.html).

## Defining an Application Load Balancer

You define an application load balancer by creating an instance of
`ApplicationLoadBalancer`, adding a Listener to the load balancer
and adding Targets to the Listener:

```go
import "github.com/aws/aws-cdk-go/awscdk"
var asg autoScalingGroup
var vpc vpc


// Create the load balancer in a VPC. 'internetFacing' is 'false'
// by default, which creates an internal load balancer.
lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
})

// Add a listener and open up the load balancer's security group
// to the world.
listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
	Port: jsii.Number(80),

	// 'open: true' is the default, you can leave it out if you want. Set it
	// to 'false' and use `listener.connections` if you want to be selective
	// about who can access the load balancer.
	Open: jsii.Boolean(true),
})

// Create an AutoScaling group and add it as a load balancing
// target to the listener.
listener.AddTargets(jsii.String("ApplicationFleet"), &AddApplicationTargetsProps{
	Port: jsii.Number(8080),
	Targets: []iApplicationLoadBalancerTarget{
		asg,
	},
})
```

The security groups of the load balancer and the target are automatically
updated to allow the network traffic.

One (or more) security groups can be associated with the load balancer;
if a security group isn't provided, one will be automatically created.

```go
var vpc vpc


securityGroup1 := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup1"), &SecurityGroupProps{
	Vpc: Vpc,
})
lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
	SecurityGroup: securityGroup1,
})

securityGroup2 := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup2"), &SecurityGroupProps{
	Vpc: Vpc,
})
lb.AddSecurityGroup(securityGroup2)
```

### Conditions

It's possible to route traffic to targets based on conditions in the incoming
HTTP request. For example, the following will route requests to the indicated
AutoScalingGroup only if the requested host in the request is either for
`example.com/ok` or `example.com/path`:

```go
var listener applicationListener
var asg autoScalingGroup


listener.AddTargets(jsii.String("Example.Com Fleet"), &AddApplicationTargetsProps{
	Priority: jsii.Number(10),
	Conditions: []listenerCondition{
		elbv2.*listenerCondition_HostHeaders([]*string{
			jsii.String("example.com"),
		}),
		elbv2.*listenerCondition_PathPatterns([]*string{
			jsii.String("/ok"),
			jsii.String("/path"),
		}),
	},
	Port: jsii.Number(8080),
	Targets: []iApplicationLoadBalancerTarget{
		asg,
	},
})
```

A target with a condition contains either `pathPatterns` or `hostHeader`, or
both. If both are specified, both conditions must be met for the requests to
be routed to the given target. `priority` is a required field when you add
targets with conditions. The lowest number wins.

Every listener must have at least one target without conditions, which is
where all requests that didn't match any of the conditions will be sent.

### Convenience methods and more complex Actions

Routing traffic from a Load Balancer to a Target involves the following steps:

* Create a Target Group, register the Target into the Target Group
* Add an Action to the Listener which forwards traffic to the Target Group.

A new listener can be added to the Load Balancer by calling `addListener()`.
Listeners that have been added to the load balancer can be listed using the
`listeners` property.  Note that the `listeners` property will throw an Error
for imported or looked up Load Balancers.

Various methods on the `Listener` take care of this work for you to a greater
or lesser extent:

* `addTargets()` performs both steps: automatically creates a Target Group and the
  required Action.
* `addTargetGroups()` gives you more control: you create the Target Group (or
  Target Groups) yourself and the method creates Action that routes traffic to
  the Target Groups.
* `addAction()` gives you full control: you supply the Action and wire it up
  to the Target Groups yourself (or access one of the other ELB routing features).

Using `addAction()` gives you access to some of the features of an Elastic Load
Balancer that the other two convenience methods don't:

* **Routing stickiness**: use `ListenerAction.forward()` and supply a
  `stickinessDuration` to make sure requests are routed to the same target group
  for a given duration.
* **Weighted Target Groups**: use `ListenerAction.weightedForward()`
  to give different weights to different target groups.
* **Fixed Responses**: use `ListenerAction.fixedResponse()` to serve
  a static response (ALB only).
* **Redirects**: use `ListenerAction.redirect()` to serve an HTTP
  redirect response (ALB only).
* **Authentication**: use `ListenerAction.authenticateOidc()` to
  perform OpenID authentication before serving a request (see the
  `aws-cdk-lib/aws-elasticloadbalancingv2-actions` package for direct authentication
  integration with Cognito) (ALB only).

Here's an example of serving a fixed response at the `/ok` URL:

```go
var listener applicationListener


listener.AddAction(jsii.String("Fixed"), &AddApplicationActionProps{
	Priority: jsii.Number(10),
	Conditions: []listenerCondition{
		elbv2.*listenerCondition_PathPatterns([]*string{
			jsii.String("/ok"),
		}),
	},
	Action: elbv2.ListenerAction_FixedResponse(jsii.Number(200), &FixedResponseOptions{
		ContentType: jsii.String("text/plain"),
		MessageBody: jsii.String("OK"),
	}),
})
```

Here's an example of using OIDC authentication before forwarding to a TargetGroup:

```go
var listener applicationListener
var myTargetGroup applicationTargetGroup


listener.AddAction(jsii.String("DefaultAction"), &AddApplicationActionProps{
	Action: elbv2.ListenerAction_AuthenticateOidc(&AuthenticateOidcOptions{
		AuthorizationEndpoint: jsii.String("https://example.com/openid"),
		// Other OIDC properties here
		ClientId: jsii.String("..."),
		ClientSecret: awscdk.SecretValue_SecretsManager(jsii.String("...")),
		Issuer: jsii.String("..."),
		TokenEndpoint: jsii.String("..."),
		UserInfoEndpoint: jsii.String("..."),

		// Next
		Next: elbv2.ListenerAction_Forward([]iApplicationTargetGroup{
			myTargetGroup,
		}),
	}),
})
```

If you just want to redirect all incoming traffic on one port to another port, you can use the following code:

```go
var lb applicationLoadBalancer


lb.AddRedirect(&ApplicationLoadBalancerRedirectConfig{
	SourceProtocol: elbv2.ApplicationProtocol_HTTPS,
	SourcePort: jsii.Number(8443),
	TargetProtocol: elbv2.ApplicationProtocol_HTTP,
	TargetPort: jsii.Number(8080),
})
```

If you do not provide any options for this method, it redirects HTTP port 80 to HTTPS port 443.

By default all ingress traffic will be allowed on the source port. If you want to be more selective with your
ingress rules then set `open: false` and use the listener's `connections` object to selectively grant access to the listener.

### Application Load Balancer attributes

You can modify attributes of Application Load Balancers:

```go
var vpc vpc


lb := elbv2.NewApplicationLoadBalancer(this, jsii.String("LB"), &ApplicationLoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),

	// Whether HTTP/2 is enabled
	Http2Enabled: jsii.Boolean(false),

	// The idle timeout value, in seconds
	IdleTimeout: awscdk.Duration_Seconds(jsii.Number(1000)),

	// Whether HTTP headers with header fields thatare not valid
	// are removed by the load balancer (true), or routed to targets
	DropInvalidHeaderFields: jsii.Boolean(true),

	// How the load balancer handles requests that might
	// pose a security risk to your application
	DesyncMitigationMode: elbv2.DesyncMitigationMode_DEFENSIVE,

	// The type of IP addresses to use.
	IpAddressType: elbv2.IpAddressType_IPV4,

	// The duration of client keep-alive connections
	ClientKeepAlive: awscdk.Duration_*Seconds(jsii.Number(500)),

	// Whether cross-zone load balancing is enabled.
	CrossZoneEnabled: jsii.Boolean(true),

	// Whether the load balancer blocks traffic through the Internet Gateway (IGW).
	DenyAllIgwTraffic: jsii.Boolean(false),
})
```

For more information, see [Load balancer attributes](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/application-load-balancers.html#load-balancer-attributes)

## Defining a Network Load Balancer

Network Load Balancers are defined in a similar way to Application Load
Balancers:

```go
var vpc vpc
var asg autoScalingGroup
var sg1 iSecurityGroup
var sg2 iSecurityGroup


// Create the load balancer in a VPC. 'internetFacing' is 'false'
// by default, which creates an internal load balancer.
lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("LB"), &NetworkLoadBalancerProps{
	Vpc: Vpc,
	InternetFacing: jsii.Boolean(true),
	SecurityGroups: []*iSecurityGroup{
		sg1,
	},
})
lb.AddSecurityGroup(sg2)

// Add a listener on a particular port.
listener := lb.AddListener(jsii.String("Listener"), &BaseNetworkListenerProps{
	Port: jsii.Number(443),
})

// Add targets on a particular port.
listener.AddTargets(jsii.String("AppFleet"), &AddNetworkTargetsProps{
	Port: jsii.Number(443),
	Targets: []iNetworkLoadBalancerTarget{
		asg,
	},
})
```

### Enforce security group inbound rules on PrivateLink traffic for a Network Load Balancer

You can indicate whether to evaluate inbound security group rules for traffic
sent to a Network Load Balancer through AWS PrivateLink.
The evaluation is enabled by default.

```go
var vpc vpc


nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("LB"), &NetworkLoadBalancerProps{
	Vpc: Vpc,
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic: jsii.Boolean(true),
})
```

One thing to keep in mind is that network load balancers do not have security
groups, and no automatic security group configuration is done for you. You will
have to configure the security groups of the target yourself to allow traffic by
clients and/or load balancer instances, depending on your target types.  See
[Target Groups for your Network Load
Balancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html)
and [Register targets with your Target
Group](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/target-group-register-targets.html)
for more information.

### Dualstack Network Load Balancer

You can create a dualstack Network Load Balancer using the `ipAddressType` property:

```go
var vpc vpc


lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("LB"), &NetworkLoadBalancerProps{
	Vpc: Vpc,
	IpAddressType: elbv2.IpAddressType_DUAL_STACK,
})
```

You cannot add UDP or TCP_UDP listeners to a dualstack Network Load Balancer.

### Network Load Balancer attributes

You can modify attributes of Network Load Balancers:

```go
var vpc vpc


lb := elbv2.NewNetworkLoadBalancer(this, jsii.String("LB"), &NetworkLoadBalancerProps{
	Vpc: Vpc,
	// Whether deletion protection is enabled.
	DeletionProtection: jsii.Boolean(true),

	// Whether cross-zone load balancing is enabled.
	CrossZoneEnabled: jsii.Boolean(true),

	// Whether the load balancer blocks traffic through the Internet Gateway (IGW).
	DenyAllIgwTraffic: jsii.Boolean(false),

	// Indicates how traffic is distributed among the load balancer Availability Zones.
	ClientRoutingPolicy: elbv2.ClientRoutingPolicy_AVAILABILITY_ZONE_AFFINITY,
})
```

## Targets and Target Groups

Application and Network Load Balancers organize load balancing targets in Target
Groups. If you add your balancing targets (such as AutoScalingGroups, ECS
services or individual instances) to your listener directly, the appropriate
`TargetGroup` will be automatically created for you.

If you need more control over the Target Groups created, create an instance of
`ApplicationTargetGroup` or `NetworkTargetGroup`, add the members you desire,
and add it to the listener by calling `addTargetGroups` instead of `addTargets`.

`addTargets()` will always return the Target Group it just created for you:

```go
var listener networkListener
var asg1 autoScalingGroup
var asg2 autoScalingGroup


group := listener.AddTargets(jsii.String("AppFleet"), &AddNetworkTargetsProps{
	Port: jsii.Number(443),
	Targets: []iNetworkLoadBalancerTarget{
		asg1,
	},
})

group.AddTarget(asg2)
```

### Sticky sessions for your Application Load Balancer

By default, an Application Load Balancer routes each request independently to a registered target based on the chosen load-balancing algorithm. However, you can use the sticky session feature (also known as session affinity) to enable the load balancer to bind a user's session to a specific target. This ensures that all requests from the user during the session are sent to the same target. This feature is useful for servers that maintain state information in order to provide a continuous experience to clients. To use sticky sessions, the client must support cookies.

Application Load Balancers support both duration-based cookies (`lb_cookie`) and application-based cookies (`app_cookie`). The key to managing sticky sessions is determining how long your load balancer should consistently route the user's request to the same target. Sticky sessions are enabled at the target group level. You can use a combination of duration-based stickiness, application-based stickiness, and no stickiness across all of your target groups.

```go
var vpc vpc


// Target group with duration-based stickiness with load-balancer generated cookie
tg1 := elbv2.NewApplicationTargetGroup(this, jsii.String("TG1"), &ApplicationTargetGroupProps{
	TargetType: elbv2.TargetType_INSTANCE,
	Port: jsii.Number(80),
	StickinessCookieDuration: awscdk.Duration_Minutes(jsii.Number(5)),
	Vpc: Vpc,
})

// Target group with application-based stickiness
tg2 := elbv2.NewApplicationTargetGroup(this, jsii.String("TG2"), &ApplicationTargetGroupProps{
	TargetType: elbv2.TargetType_INSTANCE,
	Port: jsii.Number(80),
	StickinessCookieDuration: awscdk.Duration_*Minutes(jsii.Number(5)),
	StickinessCookieName: jsii.String("MyDeliciousCookie"),
	Vpc: Vpc,
})
```

### Slow start mode for your Application Load Balancer

By default, a target starts to receive its full share of requests as soon as it is registered with a target group and passes an initial health check. Using slow start mode gives targets time to warm up before the load balancer sends them a full share of requests.

After you enable slow start for a target group, its targets enter slow start mode when they are considered healthy by the target group. A target in slow start mode exits slow start mode when the configured slow start duration period elapses or the target becomes unhealthy. The load balancer linearly increases the number of requests that it can send to a target in slow start mode. After a healthy target exits slow start mode, the load balancer can send it a full share of requests.

The allowed range is 30-900 seconds (15 minutes). The default is 0 seconds (disabled).

```go
var vpc vpc


// Target group with slow start mode enabled
tg := elbv2.NewApplicationTargetGroup(this, jsii.String("TG"), &ApplicationTargetGroupProps{
	TargetType: elbv2.TargetType_INSTANCE,
	SlowStart: awscdk.Duration_Seconds(jsii.Number(60)),
	Port: jsii.Number(80),
	Vpc: Vpc,
})
```

For more information see: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/sticky-sessions.html#application-based-stickiness

### Setting the target group protocol version

By default, Application Load Balancers send requests to targets using HTTP/1.1. You can use the [protocol version](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-target-groups.html#target-group-protocol-version) to send requests to targets using HTTP/2 or gRPC.

```go
var vpc vpc


tg := elbv2.NewApplicationTargetGroup(this, jsii.String("TG"), &ApplicationTargetGroupProps{
	TargetType: elbv2.TargetType_IP,
	Port: jsii.Number(50051),
	Protocol: elbv2.ApplicationProtocol_HTTP,
	ProtocolVersion: elbv2.ApplicationProtocolVersion_GRPC,
	HealthCheck: &HealthCheck{
		Enabled: jsii.Boolean(true),
		HealthyGrpcCodes: jsii.String("0-99"),
	},
	Vpc: Vpc,
})
```

## Using Lambda Targets

To use a Lambda Function as a target, use the integration class in the
`aws-cdk-lib/aws-elasticloadbalancingv2-targets` package:

```go
import lambda "github.com/aws/aws-cdk-go/awscdk"
import targets "github.com/aws/aws-cdk-go/awscdk"

var lambdaFunction function
var lb applicationLoadBalancer


listener := lb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
	Port: jsii.Number(80),
})
listener.AddTargets(jsii.String("Targets"), &AddApplicationTargetsProps{
	Targets: []iApplicationLoadBalancerTarget{
		targets.NewLambdaTarget(lambdaFunction),
	},

	// For Lambda Targets, you need to explicitly enable health checks if you
	// want them.
	HealthCheck: &HealthCheck{
		Enabled: jsii.Boolean(true),
	},
})
```

Only a single Lambda function can be added to a single listener rule.

## Using Application Load Balancer Targets

To use a single application load balancer as a target for the network load balancer, use the integration class in the
`aws-cdk-lib/aws-elasticloadbalancingv2-targets` package:

```go
import targets "github.com/aws/aws-cdk-go/awscdk"
import "github.com/aws/aws-cdk-go/awscdk"
import patterns "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc


task := ecs.NewFargateTaskDefinition(this, jsii.String("Task"), &FargateTaskDefinitionProps{
	Cpu: jsii.Number(256),
	MemoryLimitMiB: jsii.Number(512),
})
task.AddContainer(jsii.String("nginx"), &ContainerDefinitionOptions{
	Image: ecs.ContainerImage_FromRegistry(jsii.String("public.ecr.aws/nginx/nginx:latest")),
	PortMappings: []portMapping{
		&portMapping{
			ContainerPort: jsii.Number(80),
		},
	},
})

svc := patterns.NewApplicationLoadBalancedFargateService(this, jsii.String("Service"), &ApplicationLoadBalancedFargateServiceProps{
	Vpc: Vpc,
	TaskDefinition: task,
	PublicLoadBalancer: jsii.Boolean(false),
})

nlb := elbv2.NewNetworkLoadBalancer(this, jsii.String("Nlb"), &NetworkLoadBalancerProps{
	Vpc: Vpc,
	CrossZoneEnabled: jsii.Boolean(true),
	InternetFacing: jsii.Boolean(true),
})

listener := nlb.AddListener(jsii.String("listener"), &BaseNetworkListenerProps{
	Port: jsii.Number(80),
})

listener.AddTargets(jsii.String("Targets"), &AddNetworkTargetsProps{
	Targets: []iNetworkLoadBalancerTarget{
		targets.NewAlbTarget(svc.loadBalancer, jsii.Number(80)),
	},
	Port: jsii.Number(80),
})

awscdk.NewCfnOutput(this, jsii.String("NlbEndpoint"), &CfnOutputProps{
	Value: fmt.Sprintf("http://%v", nlb.LoadBalancerDnsName),
})
```

Only the network load balancer is allowed to add the application load balancer as the target.

## Configuring Health Checks

Health checks are configured upon creation of a target group:

```go
var listener applicationListener
var asg autoScalingGroup


listener.AddTargets(jsii.String("AppFleet"), &AddApplicationTargetsProps{
	Port: jsii.Number(8080),
	Targets: []iApplicationLoadBalancerTarget{
		asg,
	},
	HealthCheck: &HealthCheck{
		Path: jsii.String("/ping"),
		Interval: awscdk.Duration_Minutes(jsii.Number(1)),
	},
})
```

The health check can also be configured after creation by calling
`configureHealthCheck()` on the created object.

No attempts are made to configure security groups for the port you're
configuring a health check for, but if the health check is on the same port
you're routing traffic to, the security group already allows the traffic.
If not, you will have to configure the security groups appropriately:

```go
var lb applicationLoadBalancer
var listener applicationListener
var asg autoScalingGroup


listener.AddTargets(jsii.String("AppFleet"), &AddApplicationTargetsProps{
	Port: jsii.Number(8080),
	Targets: []iApplicationLoadBalancerTarget{
		asg,
	},
	HealthCheck: &HealthCheck{
		Port: jsii.String("8088"),
	},
})

asg.connections.AllowFrom(lb, ec2.Port_Tcp(jsii.Number(8088)))
```

## Using a Load Balancer from a different Stack

If you want to put your Load Balancer and the Targets it is load balancing to in
different stacks, you may not be able to use the convenience methods
`loadBalancer.addListener()` and `listener.addTargets()`.

The reason is that these methods will create resources in the same Stack as the
object they're called on, which may lead to cyclic references between stacks.
Instead, you will have to create an `ApplicationListener` in the target stack,
or an empty `TargetGroup` in the load balancer stack that you attach your
service to.

For an example of the alternatives while load balancing to an ECS service, see the
[ecs/cross-stack-load-balancer
example](https://github.com/aws-samples/aws-cdk-examples/tree/master/typescript/ecs/cross-stack-load-balancer/).

## Protocol for Load Balancer Targets

Constructs that want to be a load balancer target should implement
`IApplicationLoadBalancerTarget` and/or `INetworkLoadBalancerTarget`, and
provide an implementation for the function `attachToXxxTargetGroup()`, which can
call functions on the load balancer and should return metadata about the
load balancing target:

```go
type myTarget struct {
}

func (this *myTarget) attachToApplicationTargetGroup(targetGroup applicationTargetGroup) loadBalancerTargetProps {
	// If we need to add security group rules
	// targetGroup.registerConnectable(...);
	return &loadBalancerTargetProps{
		TargetType: elbv2.TargetType_IP,
		TargetJson: map[string]interface{}{
			"id": jsii.String("1.2.3.4"),
			"port": jsii.Number(8080),
		},
	}
}
```

`targetType` should be one of `Instance` or `Ip`. If the target can be
directly added to the target group, `targetJson` should contain the `id` of
the target (either instance ID or IP address depending on the type) and
optionally a `port` or `availabilityZone` override.

Application load balancer targets can call `registerConnectable()` on the
target group to register themselves for addition to the load balancer's security
group rules.

If your load balancer target requires that the TargetGroup has been
associated with a LoadBalancer before registration can happen (such as is the
case for ECS Services for example), take a resource dependency on
`targetGroup.loadBalancerAttached` as follows:

```go
var resource resource
var targetGroup applicationTargetGroup


// Make sure that the listener has been created, and so the TargetGroup
// has been associated with the LoadBalancer, before 'resource' is created.

constructs.Node_Of(resource).AddDependency(targetGroup.loadBalancerAttached)
```

## Looking up Load Balancers and Listeners

You may look up load balancers and load balancer listeners by using one of the
following lookup methods:

* `ApplicationLoadBalancer.fromlookup(options)` - Look up an application load
  balancer.
* `ApplicationListener.fromLookup(options)` - Look up an application load
  balancer listener.
* `NetworkLoadBalancer.fromLookup(options)` - Look up a network load balancer.
* `NetworkListener.fromLookup(options)` - Look up a network load balancer
  listener.

### Load Balancer lookup options

You may look up a load balancer by ARN or by associated tags. When you look a
load balancer up by ARN, that load balancer will be returned unless CDK detects
that the load balancer is of the wrong type. When you look up a load balancer by
tags, CDK will return the load balancer matching all specified tags. If more
than one load balancer matches, CDK will throw an error requesting that you
provide more specific criteria.

**Look up a Application Load Balancer by ARN**

```go
loadBalancer := elbv2.ApplicationLoadBalancer_FromLookup(this, jsii.String("ALB"), &ApplicationLoadBalancerLookupOptions{
	LoadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-east-2:123456789012:loadbalancer/app/my-load-balancer/1234567890123456"),
})
```

**Look up an Application Load Balancer by tags**

```go
loadBalancer := elbv2.ApplicationLoadBalancer_FromLookup(this, jsii.String("ALB"), &ApplicationLoadBalancerLookupOptions{
	LoadBalancerTags: map[string]*string{
		// Finds a load balancer matching all tags.
		"some": jsii.String("tag"),
		"someother": jsii.String("tag"),
	},
})
```

## Load Balancer Listener lookup options

You may look up a load balancer listener by the following criteria:

* Associated load balancer ARN
* Associated load balancer tags
* Listener ARN
* Listener port
* Listener protocol

The lookup method will return the matching listener. If more than one listener
matches, CDK will throw an error requesting that you specify additional
criteria.

**Look up a Listener by associated Load Balancer, Port, and Protocol**

```go
listener := elbv2.ApplicationListener_FromLookup(this, jsii.String("ALBListener"), &ApplicationListenerLookupOptions{
	LoadBalancerArn: jsii.String("arn:aws:elasticloadbalancing:us-east-2:123456789012:loadbalancer/app/my-load-balancer/1234567890123456"),
	ListenerProtocol: elbv2.ApplicationProtocol_HTTPS,
	ListenerPort: jsii.Number(443),
})
```

**Look up a Listener by associated Load Balancer Tag, Port, and Protocol**

```go
listener := elbv2.ApplicationListener_FromLookup(this, jsii.String("ALBListener"), &ApplicationListenerLookupOptions{
	LoadBalancerTags: map[string]*string{
		"Cluster": jsii.String("MyClusterName"),
	},
	ListenerProtocol: elbv2.ApplicationProtocol_HTTPS,
	ListenerPort: jsii.Number(443),
})
```

**Look up a Network Listener by associated Load Balancer Tag, Port, and Protocol**

```go
listener := elbv2.NetworkListener_FromLookup(this, jsii.String("ALBListener"), &NetworkListenerLookupOptions{
	LoadBalancerTags: map[string]*string{
		"Cluster": jsii.String("MyClusterName"),
	},
	ListenerProtocol: elbv2.Protocol_TCP,
	ListenerPort: jsii.Number(12345),
})
```

## Metrics

You may create metrics for Load Balancers and Target Groups through the `metrics` attribute:

**Load Balancer:**

```go
var alb iApplicationLoadBalancer


albMetrics := alb.Metrics
metricConnectionCount := albMetrics.ActiveConnectionCount()
```

**Target Group:**

```go
var targetGroup iApplicationTargetGroup


targetGroupMetrics := targetGroup.Metrics
metricHealthyHostCount := targetGroupMetrics.HealthyHostCount()
```

Metrics are also available to imported resources:

```go
var stack stack


targetGroup := elbv2.ApplicationTargetGroup_FromTargetGroupAttributes(this, jsii.String("MyTargetGroup"), &TargetGroupAttributes{
	TargetGroupArn: awscdk.Fn_ImportValue(jsii.String("TargetGroupArn")),
	LoadBalancerArns: awscdk.Fn_*ImportValue(jsii.String("LoadBalancerArn")),
})

targetGroupMetrics := targetGroup.Metrics
```

Notice that TargetGroups must be imported by supplying the Load Balancer too, otherwise accessing the `metrics` will
throw an error:

```go
var stack stack

targetGroup := elbv2.ApplicationTargetGroup_FromTargetGroupAttributes(this, jsii.String("MyTargetGroup"), &TargetGroupAttributes{
	TargetGroupArn: awscdk.Fn_ImportValue(jsii.String("TargetGroupArn")),
})

targetGroupMetrics := targetGroup.Metrics
```

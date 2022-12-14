# Amazon EC2 Construct Library

The `@aws-cdk/aws-ec2` package contains primitives for setting up networking and
instances.

```go
import ec2 "github.com/aws/aws-cdk-go/awscdk"
```

## VPC

Most projects need a Virtual Private Cloud to provide security by means of
network partitioning. This is achieved by creating an instance of
`Vpc`:

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.NewVpc(this, jsii.String("VPC"))
```

All default constructs require EC2 instances to be launched inside a VPC, so
you should generally start by defining a VPC whenever you need to launch
instances for your project.

### Subnet Types

A VPC consists of one or more subnets that instances can be placed into. CDK
distinguishes three different subnet types:

* **Public (`SubnetType.PUBLIC`)** - public subnets connect directly to the Internet using an
  Internet Gateway. If you want your instances to have a public IP address
  and be directly reachable from the Internet, you must place them in a
  public subnet.
* **Private with Internet Access (`SubnetType.PRIVATE_WITH_EGRESS`)** - instances in private subnets are not directly routable from the
  Internet, and you must provide a way to connect out to the Internet.
  By default, a NAT gateway is created in every public subnet for maximum availability. Be
  aware that you will be charged for NAT gateways.
  Alternatively you can set `natGateways:0` and provide your own egress configuration (i.e through Transit Gateway)
* **Isolated (`SubnetType.PRIVATE_ISOLATED`)** - isolated subnets do not route from or to the Internet, and
  as such do not require NAT gateways. They can only connect to or be
  connected to from other instances in the same VPC. A default VPC configuration
  will not include isolated subnets,

A default VPC configuration will create public and **private** subnets. However, if
`natGateways:0` **and** `subnetConfiguration` is undefined, default VPC configuration
will create public and **isolated** subnets. See [*Advanced Subnet Configuration*](#advanced-subnet-configuration)
below for information on how to change the default subnet configuration.

Constructs using the VPC will "launch instances" (or more accurately, create
Elastic Network Interfaces) into one or more of the subnets. They all accept
a property called `subnetSelection` (sometimes called `vpcSubnets`) to allow
you to select in what subnet to place the ENIs, usually defaulting to
*private* subnets if the property is omitted.

If you would like to save on the cost of NAT gateways, you can use
*isolated* subnets instead of *private* subnets (as described in Advanced
*Subnet Configuration*). If you need private instances to have
internet connectivity, another option is to reduce the number of NAT gateways
created by setting the `natGateways` property to a lower value (the default
is one NAT gateway per availability zone). Be aware that this may have
availability implications for your application.

[Read more about
subnets](https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Subnets.html).

### Control over availability zones

By default, a VPC will spread over at most 3 Availability Zones available to
it. To change the number of Availability Zones that the VPC will spread over,
specify the `maxAzs` property when defining it.

The number of Availability Zones that are available depends on the *region*
and *account* of the Stack containing the VPC. If the [region and account are
specified](https://docs.aws.amazon.com/cdk/latest/guide/environments.html) on
the Stack, the CLI will [look up the existing Availability
Zones](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#using-regions-availability-zones-describe)
and get an accurate count. If region and account are not specified, the stack
could be deployed anywhere and it will have to make a safe choice, limiting
itself to 2 Availability Zones.

Therefore, to get the VPC to spread over 3 or more availability zones, you
must specify the environment where the stack will be deployed.

You can gain full control over the availability zones selection strategy by overriding the Stack's [`get availabilityZones()`](https://github.com/aws/aws-cdk/blob/main/packages/@aws-cdk/core/lib/stack.ts) method:

```text
// This example is only available in TypeScript

class MyStack extends Stack {

  constructor(scope: Construct, id: string, props?: StackProps) {
    super(scope, id, props);

    // ...
  }

  get availabilityZones(): string[] {
    return ['us-west-2a', 'us-west-2b'];
  }

}
```

Note that overriding the `get availabilityZones()` method will override the default behavior for all constructs defined within the Stack.

### Choosing subnets for resources

When creating resources that create Elastic Network Interfaces (such as
databases or instances), there is an option to choose which subnets to place
them in. For example, a VPC endpoint by default is placed into a subnet in
every availability zone, but you can override which subnets to use. The property
is typically called one of `subnets`, `vpcSubnets` or `subnetSelection`.

The example below will place the endpoint into two AZs (`us-east-1a` and `us-east-1c`),
in Isolated subnets:

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc


ec2.NewInterfaceVpcEndpoint(this, jsii.String("VPC Endpoint"), &interfaceVpcEndpointProps{
	vpc: vpc,
	service: ec2.NewInterfaceVpcEndpointService(jsii.String("com.amazonaws.vpce.us-east-1.vpce-svc-uuddlrlrbastrtsvc"), jsii.Number(443)),
	subnets: &subnetSelection{
		subnetType: ec2.subnetType_PRIVATE_ISOLATED,
		availabilityZones: []*string{
			jsii.String("us-east-1a"),
			jsii.String("us-east-1c"),
		},
	},
})
```

You can also specify specific subnet objects for granular control:

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc
var subnet1 subnet
var subnet2 subnet


ec2.NewInterfaceVpcEndpoint(this, jsii.String("VPC Endpoint"), &interfaceVpcEndpointProps{
	vpc: vpc,
	service: ec2.NewInterfaceVpcEndpointService(jsii.String("com.amazonaws.vpce.us-east-1.vpce-svc-uuddlrlrbastrtsvc"), jsii.Number(443)),
	subnets: &subnetSelection{
		subnets: []iSubnet{
			subnet1,
			subnet2,
		},
	},
})
```

Which subnets are selected is evaluated as follows:

* `subnets`: if specific subnet objects are supplied, these are selected, and no other
  logic is used.
* `subnetType`/`subnetGroupName`: otherwise, a set of subnets is selected by
  supplying either type or name:

  * `subnetType` will select all subnets of the given type.
  * `subnetGroupName` should be used to distinguish between multiple groups of subnets of
    the same type (for example, you may want to separate your application instances and your
    RDS instances into two distinct groups of Isolated subnets).
  * If neither are given, the first available subnet group of a given type that
    exists in the VPC will be used, in this order: Private, then Isolated, then Public.
    In short: by default ENIs will preferentially be placed in subnets not connected to
    the Internet.
* `availabilityZones`/`onePerAz`: finally, some availability-zone based filtering may be done.
  This filtering by availability zones will only be possible if the VPC has been created or
  looked up in a non-environment agnostic stack (so account and region have been set and
  availability zones have been looked up).

  * `availabilityZones`: only the specific subnets from the selected subnet groups that are
    in the given availability zones will be returned.
  * `onePerAz`: per availability zone, a maximum of one subnet will be returned (Useful for resource
    types that do not allow creating two ENIs in the same availability zone).
* `subnetFilters`: additional filtering on subnets using any number of user-provided filters which
  extend `SubnetFilter`.  The following methods on the `SubnetFilter` class can be used to create
  a filter:

  * `byIds`: chooses subnets from a list of ids
  * `availabilityZones`: chooses subnets in the provided list of availability zones
  * `onePerAz`: chooses at most one subnet per availability zone
  * `containsIpAddresses`: chooses a subnet which contains *any* of the listed ip addresses
  * `byCidrMask`: chooses subnets that have the provided CIDR netmask

### Using NAT instances

By default, the `Vpc` construct will create NAT *gateways* for you, which
are managed by AWS. If you would prefer to use your own managed NAT
*instances* instead, specify a different value for the `natGatewayProvider`
property, as follows:

```go
// Configure the `natGatewayProvider` when defining a Vpc
natGatewayProvider := ec2.natProvider.instance(&natInstanceProps{
	instanceType: ec2.NewInstanceType(jsii.String("t3.small")),
})

vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &vpcProps{
	natGatewayProvider: natGatewayProvider,

	// The 'natGateways' parameter now controls the number of NAT instances
	natGateways: jsii.Number(2),
})
```

The construct will automatically search for the most recent NAT gateway AMI.
If you prefer to use a custom AMI, use `machineImage: MachineImage.genericLinux({ ... })` and configure the right AMI ID for the
regions you want to deploy to.

By default, the NAT instances will route all traffic. To control what traffic
gets routed, pass a custom value for `defaultAllowedTraffic` and access the
`NatInstanceProvider.connections` member after having passed the NAT provider to
the VPC:

```go
// Example automatically generated from non-compiling source. May contain errors.
var instanceType instanceType


provider := ec2.natProvider.instance(&natInstanceProps{
	instanceType: instanceType,
	defaultAllowedTraffic: ec2.natTrafficDirection_OUTBOUND_ONLY,
})
ec2.NewVpc(this, jsii.String("TheVPC"), &vpcProps{
	natGatewayProvider: provider,
})
provider.connections.allowFrom(ec2.peer.ipv4(jsii.String("1.2.3.4/8")), ec2.port.tcp(jsii.Number(80)))
```

### Ip Address Management

The VPC spans a supernet IP range, which contains the non-overlapping IPs of its contained subnets. Possible sources for this IP range are:

* You specify an IP range directly by specifying a CIDR
* You allocate an IP range of a given size automatically from AWS IPAM

By default the Vpc will allocate the `10.0.0.0/16` address range which will be exhaustively spread across all subnets in the subnet configuration. This behavior can be changed by passing an object that implements `IIpAddresses` to the `ipAddress` property of a Vpc. See the subsequent sections for the options.

Be aware that if you don't explicitly reserve subnet groups in `subnetConfiguration`, the address space will be fully allocated! If you predict you may need to add more subnet groups later, add them early on and set `reserved: true` (see the "Advanced Subnet Configuration" section for more information).

#### Specifying a CIDR directly

Use `IpAddresses.cidr` to define a Cidr range for your Vpc directly in code:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"


ec2.NewVpc(stack, jsii.String("TheVPC"), &vpcProps{
	ipAddresses: ec2.ipAddresses.cidr(jsii.String("10.0.1.0/20")),
})
```

Space will be allocated to subnets in the following order:

* First, spaces is allocated for all subnets groups that explicitly have a `cidrMask` set as part of their configuration (including reserved subnets).
* Afterwards, any remaining space is divided evenly between the rest of the subnets (if any).

The argument to `IpAddresses.cidr` may not be a token, and concrete Cidr values are generated in the synthesized CloudFormation template.

#### Allocating an IP range from AWS IPAM

Amazon VPC IP Address Manager (IPAM) manages a large IP space, from which chunks can be allocated for use in the Vpc. For information on Amazon VPC IP Address Manager please see the [official documentation](https://docs.aws.amazon.com/vpc/latest/ipam/what-it-is-ipam.html). An example of allocating from AWS IPAM looks like this:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"

var pool cfnIPAMPool


ec2.NewVpc(stack, jsii.String("TheVPC"), &vpcProps{
	ipAddresses: ec2.ipAddresses.awsIpamAllocation(&awsIpamProps{
		ipv4IpamPoolId: pool.ref,
		ipv4NetmaskLength: jsii.Number(18),
		defaultSubnetIpv4NetmaskLength: jsii.Number(24),
	}),
})
```

`IpAddresses.awsIpamAllocation` requires the following:

* `ipv4IpamPoolId`, the id of an IPAM Pool from which the VPC range should be allocated.
* `ipv4NetmaskLength`, the size of the IP range that will be requested from the Pool at deploy time.
* `defaultSubnetIpv4NetmaskLength`, the size of subnets in groups that don't have `cidrMask` set.

With this method of IP address management, no attempt is made to guess at subnet group sizes or to exhaustively allocate the IP range. All subnet groups must have an explicit `cidrMask` set as part of their subnet configuration, or `defaultSubnetIpv4NetmaskLength` must be set for a default size. If not, synthesis will fail and you must provide one or the other.

### Reserving availability zones

There are situations where the IP space for availability zones will
need to be reserved. This is useful in situations where availability
zones would need to be added after the vpc is originally deployed,
without causing IP renumbering for availability zones subnets. The IP
space for reserving `n` availability zones can be done by setting the
`reservedAzs` to `n` in vpc props, as shown below:

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.NewVpc(this, jsii.String("TheVPC"), &vpcProps{
	cidr: jsii.String("10.0.0.0/21"),
	maxAzs: jsii.Number(3),
	reservedAzs: jsii.Number(1),
})
```

In the example above, the subnets for reserved availability zones is not
actually provisioned but its IP space is still reserved. If, in the future,
new availability zones needs to be provisioned, then we would decrement
the value of `reservedAzs` and increment the `maxAzs` or `availabilityZones`
accordingly. This action would not cause the IP address of subnets to get
renumbered, but rather the IP space that was previously reserved will be
used for the new availability zones subnets.

### Advanced Subnet Configuration

If the default VPC configuration (public and private subnets spanning the
size of the VPC) don't suffice for you, you can configure what subnets to
create by specifying the `subnetConfiguration` property. It allows you
to configure the number and size of all subnets. Specifying an advanced
subnet configuration could look like this:

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.NewVpc(this, jsii.String("TheVPC"), &vpcProps{
	// 'IpAddresses' configures the IP range and size of the entire VPC.
	// The IP space will be divided based on configuration for the subnets.
	ipAddresses: ipAddresses_Cidr(jsii.String("10.0.0.0/21")),

	// 'maxAzs' configures the maximum number of availability zones to use.
	// If you want to specify the exact availability zones you want the VPC
	// to use, use `availabilityZones` instead.
	maxAzs: jsii.Number(3),

	// 'subnetConfiguration' specifies the "subnet groups" to create.
	// Every subnet group will have a subnet for each AZ, so this
	// configuration will create `3 groups × 3 AZs = 9` subnets.
	subnetConfiguration: []subnetConfiguration{
		&subnetConfiguration{
			// 'subnetType' controls Internet access, as described above.
			subnetType: ec2.subnetType_PUBLIC,

			// 'name' is used to name this particular subnet group. You will have to
			// use the name for subnet selection if you have more than one subnet
			// group of the same type.
			name: jsii.String("Ingress"),

			// 'cidrMask' specifies the IP addresses in the range of of individual
			// subnets in the group. Each of the subnets in this group will contain
			// `2^(32 address bits - 24 subnet bits) - 2 reserved addresses = 254`
			// usable IP addresses.
			//
			// If 'cidrMask' is left out the available address space is evenly
			// divided across the remaining subnet groups.
			cidrMask: jsii.Number(24),
		},
		&subnetConfiguration{
			cidrMask: jsii.Number(24),
			name: jsii.String("Application"),
			subnetType: ec2.*subnetType_PRIVATE_WITH_EGRESS,
		},
		&subnetConfiguration{
			cidrMask: jsii.Number(28),
			name: jsii.String("Database"),
			subnetType: ec2.*subnetType_PRIVATE_ISOLATED,

			// 'reserved' can be used to reserve IP address space. No resources will
			// be created for this subnet, but the IP range will be kept available for
			// future creation of this subnet, or even for future subdivision.
			reserved: jsii.Boolean(true),
		},
	},
})
```

The example above is one possible configuration, but the user can use the
constructs above to implement many other network configurations.

The `Vpc` from the above configuration in a Region with three
availability zones will be the following:

Subnet Name       |Type      |IP Block      |AZ|Features
------------------|----------|--------------|--|--------
IngressSubnet1    |`PUBLIC`  |`10.0.0.0/24` |#1|NAT Gateway
IngressSubnet2    |`PUBLIC`  |`10.0.1.0/24` |#2|NAT Gateway
IngressSubnet3    |`PUBLIC`  |`10.0.2.0/24` |#3|NAT Gateway
ApplicationSubnet1|`PRIVATE` |`10.0.3.0/24` |#1|Route to NAT in IngressSubnet1
ApplicationSubnet2|`PRIVATE` |`10.0.4.0/24` |#2|Route to NAT in IngressSubnet2
ApplicationSubnet3|`PRIVATE` |`10.0.5.0/24` |#3|Route to NAT in IngressSubnet3
DatabaseSubnet1   |`ISOLATED`|`10.0.6.0/28` |#1|Only routes within the VPC
DatabaseSubnet2   |`ISOLATED`|`10.0.6.16/28`|#2|Only routes within the VPC
DatabaseSubnet3   |`ISOLATED`|`10.0.6.32/28`|#3|Only routes within the VPC

### Accessing the Internet Gateway

If you need access to the internet gateway, you can get its ID like so:

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc


igwId := vpc.internetGatewayId
```

For a VPC with only `ISOLATED` subnets, this value will be undefined.

This is only supported for VPCs created in the stack - currently you're
unable to get the ID for imported VPCs. To do that you'd have to specifically
look up the Internet Gateway by name, which would require knowing the name
beforehand.

This can be useful for configuring routing using a combination of gateways:
for more information see [Routing](#routing) below.

#### Routing

It's possible to add routes to any subnets using the `addRoute()` method. If for
example you want an isolated subnet to have a static route via the default
Internet Gateway created for the public subnet - perhaps for routing a VPN
connection - you can do so like this:

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.NewVpc(this, jsii.String("VPC"), &vpcProps{
	subnetConfiguration: []subnetConfiguration{
		&subnetConfiguration{
			subnetType: ec2.subnetType_PUBLIC,
			name: jsii.String("Public"),
		},
		&subnetConfiguration{
			subnetType: ec2.*subnetType_PRIVATE_ISOLATED,
			name: jsii.String("Isolated"),
		},
	},
})

(vpc.isolatedSubnets[0].(subnet)).addRoute(jsii.String("StaticRoute"), &addRouteOptions{
	routerId: vpc.internetGatewayId,
	routerType: ec2.routerType_GATEWAY,
	destinationCidrBlock: jsii.String("8.8.8.8/32"),
})
```

*Note that we cast to `Subnet` here because the list of subnets only returns an
`ISubnet`.*

### Reserving subnet IP space

There are situations where the IP space for a subnet or number of subnets
will need to be reserved. This is useful in situations where subnets would
need to be added after the vpc is originally deployed, without causing IP
renumbering for existing subnets. The IP space for a subnet may be reserved
by setting the `reserved` subnetConfiguration property to true, as shown
below:

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.NewVpc(this, jsii.String("TheVPC"), &vpcProps{
	natGateways: jsii.Number(1),
	subnetConfiguration: []subnetConfiguration{
		&subnetConfiguration{
			cidrMask: jsii.Number(26),
			name: jsii.String("Public"),
			subnetType: ec2.subnetType_PUBLIC,
		},
		&subnetConfiguration{
			cidrMask: jsii.Number(26),
			name: jsii.String("Application1"),
			subnetType: ec2.*subnetType_PRIVATE_WITH_EGRESS,
		},
		&subnetConfiguration{
			cidrMask: jsii.Number(26),
			name: jsii.String("Application2"),
			subnetType: ec2.*subnetType_PRIVATE_WITH_EGRESS,
			reserved: jsii.Boolean(true),
		},
		&subnetConfiguration{
			cidrMask: jsii.Number(27),
			name: jsii.String("Database"),
			subnetType: ec2.*subnetType_PRIVATE_ISOLATED,
		},
	},
})
```

In the example above, the subnet for Application2 is not actually provisioned
but its IP space is still reserved. If in the future this subnet needs to be
provisioned, then the `reserved: true` property should be removed. Reserving
parts of the IP space prevents the other subnets from getting renumbered.

### Sharing VPCs between stacks

If you are creating multiple `Stack`s inside the same CDK application, you
can reuse a VPC defined in one Stack in another by simply passing the VPC
instance around:

```go
/**
 * Stack1 creates the VPC
 */
type stack1 struct {
	stack
	vpc vpc
}

func newStack1(scope app, id *string, props stackProps) *stack1 {
	this := &stack1{}
	cdk.NewStack_Override(this, scope, id, props)

	this.vpc = ec2.NewVpc(this, jsii.String("VPC"))
	return this
}

type stack2Props struct {
	stackProps
	vpc *iVpc
}

/**
 * Stack2 consumes the VPC
 */
type stack2 struct {
	stack
}

func newStack2(scope app, id *string, props stack2Props) *stack2 {
	this := &stack2{}
	cdk.NewStack_Override(this, scope, id, props)

	// Pass the VPC to a construct that needs it
	// Pass the VPC to a construct that needs it
	NewConstructThatTakesAVpc(this, jsii.String("Construct"), &constructThatTakesAVpcProps{
		vpc: props.vpc,
	})
	return this
}

stack1 := NewStack1(app, jsii.String("Stack1"))
stack2 := NewStack2(app, jsii.String("Stack2"), &stack2Props{
	vpc: stack1.vpc,
})
```

### Importing an existing VPC

If your VPC is created outside your CDK app, you can use `Vpc.fromLookup()`.
The CDK CLI will search for the specified VPC in the the stack's region and
account, and import the subnet configuration. Looking up can be done by VPC
ID, but more flexibly by searching for a specific tag on the VPC.

Subnet types will be determined from the `aws-cdk:subnet-type` tag on the
subnet if it exists, or the presence of a route to an Internet Gateway
otherwise. Subnet names will be determined from the `aws-cdk:subnet-name` tag
on the subnet if it exists, or will mirror the subnet type otherwise (i.e.
a public subnet will have the name `"Public"`).

The result of the `Vpc.fromLookup()` operation will be written to a file
called `cdk.context.json`. You must commit this file to source control so
that the lookup values are available in non-privileged environments such
as CI build steps, and to ensure your template builds are repeatable.

Here's how `Vpc.fromLookup()` can be used:

```go
vpc := ec2.vpc.fromLookup(stack, jsii.String("VPC"), &vpcLookupOptions{
	// This imports the default VPC but you can also
	// specify a 'vpcName' or 'tags'.
	isDefault: jsii.Boolean(true),
})
```

`Vpc.fromLookup` is the recommended way to import VPCs. If for whatever
reason you do not want to use the context mechanism to look up a VPC at
synthesis time, you can also use `Vpc.fromVpcAttributes`. This has the
following limitations:

* Every subnet group in the VPC must have a subnet in each availability zone
  (for example, each AZ must have both a public and private subnet). Asymmetric
  VPCs are not supported.
* All VpcId, SubnetId, RouteTableId, ... parameters must either be known at
  synthesis time, or they must come from deploy-time list parameters whose
  deploy-time lengths are known at synthesis time.

Using `Vpc.fromVpcAttributes()` looks like this:

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.vpc.fromVpcAttributes(this, jsii.String("VPC"), &vpcAttributes{
	vpcId: jsii.String("vpc-1234"),
	availabilityZones: []*string{
		jsii.String("us-east-1a"),
		jsii.String("us-east-1b"),
	},

	// Either pass literals for all IDs
	publicSubnetIds: []*string{
		jsii.String("s-12345"),
		jsii.String("s-67890"),
	},

	// OR: import a list of known length
	privateSubnetIds: awscdk.Fn.importListValue(jsii.String("PrivateSubnetIds"), jsii.Number(2)),

	// OR: split an imported string to a list of known length
	isolatedSubnetIds: awscdk.Fn.split(jsii.String(","), ssm.stringParameter.valueForStringParameter(this, jsii.String("MyParameter")), jsii.Number(2)),
})
```

## Allowing Connections

In AWS, all network traffic in and out of **Elastic Network Interfaces** (ENIs)
is controlled by **Security Groups**. You can think of Security Groups as a
firewall with a set of rules. By default, Security Groups allow no incoming
(ingress) traffic and all outgoing (egress) traffic. You can add ingress rules
to them to allow incoming traffic streams. To exert fine-grained control over
egress traffic, set `allowAllOutbound: false` on the `SecurityGroup`, after
which you can add egress traffic rules.

You can manipulate Security Groups directly:

```go
mySecurityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &securityGroupProps{
	vpc: vpc,
	description: jsii.String("Allow ssh access to ec2 instances"),
	allowAllOutbound: jsii.Boolean(true),
})
mySecurityGroup.addIngressRule(ec2.peer.anyIpv4(), ec2.port.tcp(jsii.Number(22)), jsii.String("allow ssh access from the world"))
```

All constructs that create ENIs on your behalf (typically constructs that create
EC2 instances or other VPC-connected resources) will all have security groups
automatically assigned. Those constructs have an attribute called
**connections**, which is an object that makes it convenient to update the
security groups. If you want to allow connections between two constructs that
have security groups, you have to add an **Egress** rule to one Security Group,
and an **Ingress** rule to the other. The connections object will automatically
take care of this for you:

```go
// Example automatically generated from non-compiling source. May contain errors.
var loadBalancer applicationLoadBalancer
var appFleet autoScalingGroup
var dbFleet autoScalingGroup


// Allow connections from anywhere
loadBalancer.connections.allowFromAnyIpv4(ec2.port.tcp(jsii.Number(443)), jsii.String("Allow inbound HTTPS"))

// The same, but an explicit IP address
loadBalancer.connections.allowFrom(ec2.peer.ipv4(jsii.String("1.2.3.4/32")), ec2.port.tcp(jsii.Number(443)), jsii.String("Allow inbound HTTPS"))

// Allow connection between AutoScalingGroups
appFleet.connections.allowTo(dbFleet, ec2.port.tcp(jsii.Number(443)), jsii.String("App can call database"))
```

### Connection Peers

There are various classes that implement the connection peer part:

```go
// Example automatically generated from non-compiling source. May contain errors.
var appFleet autoScalingGroup
var dbFleet autoScalingGroup


// Simple connection peers
peer := ec2.peer.ipv4(jsii.String("10.0.0.0/16"))
peer = ec2.peer.anyIpv4()
peer = ec2.peer.ipv6(jsii.String("::0/0"))
peer = ec2.peer.anyIpv6()
peer = ec2.peer.prefixList(jsii.String("pl-12345"))
appFleet.connections.allowTo(peer, ec2.port.tcp(jsii.Number(443)), jsii.String("Allow outbound HTTPS"))
```

Any object that has a security group can itself be used as a connection peer:

```go
// Example automatically generated from non-compiling source. May contain errors.
var fleet1 autoScalingGroup
var fleet2 autoScalingGroup
var appFleet autoScalingGroup


// These automatically create appropriate ingress and egress rules in both security groups
fleet1.connections.allowTo(fleet2, ec2.port.tcp(jsii.Number(80)), jsii.String("Allow between fleets"))

appFleet.connections.allowFromAnyIpv4(ec2.port.tcp(jsii.Number(80)), jsii.String("Allow from load balancer"))
```

### Port Ranges

The connections that are allowed are specified by port ranges. A number of classes provide
the connection specifier:

```go
// Example automatically generated from non-compiling source. May contain errors.
ec2.port.tcp(jsii.Number(80))
ec2.port.tcpRange(jsii.Number(60000), jsii.Number(65535))
ec2.port.allTcp()
ec2.port.allIcmp()
ec2.port.allIcmpV6()
ec2.port.allTraffic()
```

> NOTE: Not all protocols have corresponding helper methods. In the absence of a helper method,
> you can instantiate `Port` yourself with your own settings. You are also welcome to contribute
> new helper methods.

### Default Ports

Some Constructs have default ports associated with them. For example, the
listener of a load balancer does (it's the public port), or instances of an
RDS database (it's the port the database is accepting connections on).

If the object you're calling the peering method on has a default port associated with it, you can call
`allowDefaultPortFrom()` and omit the port specifier. If the argument has an associated default port, call
`allowDefaultPortTo()`.

For example:

```go
// Example automatically generated from non-compiling source. May contain errors.
var listener applicationListener
var appFleet autoScalingGroup
var rdsDatabase databaseCluster


// Port implicit in listener
listener.connections.allowDefaultPortFromAnyIpv4(jsii.String("Allow public"))

// Port implicit in peer
appFleet.connections.allowDefaultPortTo(rdsDatabase, jsii.String("Fleet can access database"))
```

### Security group rules

By default, security group wills be added inline to the security group in the output cloud formation
template, if applicable.  This includes any static rules by ip address and port range.  This
optimization helps to minimize the size of the template.

In some environments this is not desirable, for example if your security group access is controlled
via tags. You can disable inline rules per security group or globally via the context key
`@aws-cdk/aws-ec2.securityGroupDisableInlineRules`.

```go
mySecurityGroupWithoutInlineRules := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &securityGroupProps{
	vpc: vpc,
	description: jsii.String("Allow ssh access to ec2 instances"),
	allowAllOutbound: jsii.Boolean(true),
	disableInlineRules: jsii.Boolean(true),
})
//This will add the rule as an external cloud formation construct
mySecurityGroupWithoutInlineRules.addIngressRule(ec2.peer.anyIpv4(), ec2.port.tcp(jsii.Number(22)), jsii.String("allow ssh access from the world"))
```

### Importing an existing security group

If you know the ID and the configuration of the security group to import, you can use `SecurityGroup.fromSecurityGroupId`:

```go
// Example automatically generated from non-compiling source. May contain errors.
sg := ec2.securityGroup.fromSecurityGroupId(this, jsii.String("SecurityGroupImport"), jsii.String("sg-1234"), &securityGroupImportOptions{
	allowAllOutbound: jsii.Boolean(true),
})
```

Alternatively, use lookup methods to import security groups if you do not know the ID or the configuration details. Method `SecurityGroup.fromLookupByName` looks up a security group if the security group ID is unknown.

```go
sg := ec2.securityGroup.fromLookupByName(this, jsii.String("SecurityGroupLookup"), jsii.String("security-group-name"), vpc)
```

If the security group ID is known and configuration details are unknown, use method `SecurityGroup.fromLookupById` instead. This method will lookup property `allowAllOutbound` from the current configuration of the security group.

```go
// Example automatically generated from non-compiling source. May contain errors.
sg := ec2.securityGroup.fromLookupById(this, jsii.String("SecurityGroupLookup"), jsii.String("sg-1234"))
```

The result of `SecurityGroup.fromLookupByName` and `SecurityGroup.fromLookupById` operations will be written to a file called `cdk.context.json`. You must commit this file to source control so that the lookup values are available in non-privileged environments such as CI build steps, and to ensure your template builds are repeatable.

### Cross Stack Connections

If you are attempting to add a connection from a peer in one stack to a peer in a different stack, sometimes it is necessary to ensure that you are making the connection in
a specific stack in order to avoid a cyclic reference. If there are no other dependencies between stacks then it will not matter in which stack you make
the connection, but if there are existing dependencies (i.e. stack1 already depends on stack2), then it is important to make the connection in the dependent stack (i.e. stack1).

Whenever you make a `connections` function call, the ingress and egress security group rules will be added to the stack that the calling object exists in.
So if you are doing something like `peer1.connections.allowFrom(peer2)`, then the security group rules (both ingress and egress) will be created in `peer1`'s Stack.

As an example, if we wanted to allow a connection from a security group in one stack (egress) to a security group in a different stack (ingress),
we would make the connection like:

**If Stack1 depends on Stack2**

```go
// Stack 1
var stack1 stack
var stack2 stack


sg1 := ec2.NewSecurityGroup(stack1, jsii.String("SG1"), &securityGroupProps{
	allowAllOutbound: jsii.Boolean(false),
	 // if this is `true` then no egress rule will be created
	vpc: vpc,
})

// Stack 2
sg2 := ec2.NewSecurityGroup(stack2, jsii.String("SG2"), &securityGroupProps{
	allowAllOutbound: jsii.Boolean(false),
	 // if this is `true` then no egress rule will be created
	vpc: vpc,
})

// `connections.allowTo` on `sg1` since we want the
// rules to be created in Stack1
sg1.connections.allowTo(sg2, ec2.port.tcp(jsii.Number(3333)))
```

In this case both the Ingress Rule for `sg2` and the Egress Rule for `sg1` will both be created
in `Stack 1` which avoids the cyclic reference.

**If Stack2 depends on Stack1**

```go
// Stack 1
var stack1 stack
var stack2 stack


sg1 := ec2.NewSecurityGroup(stack1, jsii.String("SG1"), &securityGroupProps{
	allowAllOutbound: jsii.Boolean(false),
	 // if this is `true` then no egress rule will be created
	vpc: vpc,
})

// Stack 2
sg2 := ec2.NewSecurityGroup(stack2, jsii.String("SG2"), &securityGroupProps{
	allowAllOutbound: jsii.Boolean(false),
	 // if this is `true` then no egress rule will be created
	vpc: vpc,
})

// `connections.allowFrom` on `sg2` since we want the
// rules to be created in Stack2
sg2.connections.allowFrom(sg1, ec2.port.tcp(jsii.Number(3333)))
```

In this case both the Ingress Rule for `sg2` and the Egress Rule for `sg1` will both be created
in `Stack 2` which avoids the cyclic reference.

## Machine Images (AMIs)

AMIs control the OS that gets launched when you start your EC2 instance. The EC2
library contains constructs to select the AMI you want to use.

Depending on the type of AMI, you select it a different way. Here are some
examples of things you might want to use:

```go
// Pick the right Amazon Linux edition. All arguments shown are optional
// and will default to these values when omitted.
amznLinux := ec2.machineImage.latestAmazonLinux(&amazonLinuxImageProps{
	generation: ec2.amazonLinuxGeneration_AMAZON_LINUX,
	edition: ec2.amazonLinuxEdition_STANDARD,
	virtualization: ec2.amazonLinuxVirt_HVM,
	storage: ec2.amazonLinuxStorage_GENERAL_PURPOSE,
	cpuType: ec2.amazonLinuxCpuType_X86_64,
})

// Pick a Windows edition to use
windows := ec2.machineImage.latestWindows(ec2.windowsVersion_WINDOWS_SERVER_2019_ENGLISH_FULL_BASE)

// Read AMI id from SSM parameter store
ssm := ec2.machineImage.fromSsmParameter(jsii.String("/my/ami"), &ssmParameterImageOptions{
	os: ec2.operatingSystemType_LINUX,
})

// Look up the most recent image matching a set of AMI filters.
// In this case, look up the NAT instance AMI, by using a wildcard
// in the 'name' field:
natAmi := ec2.machineImage.lookup(&lookupMachineImageProps{
	name: jsii.String("amzn-ami-vpc-nat-*"),
	owners: []*string{
		jsii.String("amazon"),
	},
})

// For other custom (Linux) images, instantiate a `GenericLinuxImage` with
// a map giving the AMI to in for each region:
linux := ec2.machineImage.genericLinux(map[string]*string{
	"us-east-1": jsii.String("ami-97785bed"),
	"eu-west-1": jsii.String("ami-12345678"),
})

// For other custom (Windows) images, instantiate a `GenericWindowsImage` with
// a map giving the AMI to in for each region:
genericWindows := ec2.machineImage.genericWindows(map[string]*string{
	"us-east-1": jsii.String("ami-97785bed"),
	"eu-west-1": jsii.String("ami-12345678"),
})
```

> NOTE: The AMIs selected by `MachineImage.lookup()` will be cached in
> `cdk.context.json`, so that your AutoScalingGroup instances aren't replaced while
> you are making unrelated changes to your CDK app.
>
> To query for the latest AMI again, remove the relevant cache entry from
> `cdk.context.json`, or use the `cdk context` command. For more information, see
> [Runtime Context](https://docs.aws.amazon.com/cdk/latest/guide/context.html) in the CDK
> developer guide.
>
> `MachineImage.genericLinux()`, `MachineImage.genericWindows()` will use `CfnMapping` in
> an agnostic stack.

## Special VPC configurations

### VPN connections to a VPC

Create your VPC with VPN connections by specifying the `vpnConnections` props (keys are construct `id`s):

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &vpcProps{
	vpnConnections: map[string]vpnConnectionOptions{
		"dynamic": &vpnConnectionOptions{
			 // Dynamic routing (BGP)
			"ip": jsii.String("1.2.3.4"),
		},
		"static": &vpnConnectionOptions{
			 // Static routing
			"ip": jsii.String("4.5.6.7"),
			"staticRoutes": []*string{
				jsii.String("192.168.10.0/24"),
				jsii.String("192.168.20.0/24"),
			},
		},
	},
})
```

To create a VPC that can accept VPN connections, set `vpnGateway` to `true`:

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &vpcProps{
	vpnGateway: jsii.Boolean(true),
})
```

VPN connections can then be added:

```go
vpc.addVpnConnection(jsii.String("Dynamic"), &vpnConnectionOptions{
	ip: jsii.String("1.2.3.4"),
})
```

By default, routes will be propagated on the route tables associated with the private subnets. If no
private subnets exist, isolated subnets are used. If no isolated subnets exist, public subnets are
used. Use the `Vpc` property `vpnRoutePropagation` to customize this behavior.

VPN connections expose [metrics (cloudwatch.Metric)](https://github.com/aws/aws-cdk/blob/main/packages/%40aws-cdk/aws-cloudwatch/README.md) across all tunnels in the account/region and per connection:

```go
// Across all tunnels in the account/region
allDataOut := ec2.vpnConnection.metricAllTunnelDataOut()

// For a specific vpn connection
vpnConnection := vpc.addVpnConnection(jsii.String("Dynamic"), &vpnConnectionOptions{
	ip: jsii.String("1.2.3.4"),
})
state := vpnConnection.metricTunnelState()
```

### VPC endpoints

A VPC endpoint enables you to privately connect your VPC to supported AWS services and VPC endpoint services powered by PrivateLink without requiring an internet gateway, NAT device, VPN connection, or AWS Direct Connect connection. Instances in your VPC do not require public IP addresses to communicate with resources in the service. Traffic between your VPC and the other service does not leave the Amazon network.

Endpoints are virtual devices. They are horizontally scaled, redundant, and highly available VPC components that allow communication between instances in your VPC and services without imposing availability risks or bandwidth constraints on your network traffic.

```go
// Add gateway endpoints when creating the VPC
vpc := ec2.NewVpc(this, jsii.String("MyVpc"), &vpcProps{
	gatewayEndpoints: map[string]gatewayVpcEndpointOptions{
		"S3": &gatewayVpcEndpointOptions{
			"service": ec2.GatewayVpcEndpointAwsService_S3(),
		},
	},
})

// Alternatively gateway endpoints can be added on the VPC
dynamoDbEndpoint := vpc.addGatewayEndpoint(jsii.String("DynamoDbEndpoint"), &gatewayVpcEndpointOptions{
	service: ec2.gatewayVpcEndpointAwsService_DYNAMODB(),
})

// This allows to customize the endpoint policy
dynamoDbEndpoint.addToPolicy(
iam.NewPolicyStatement(&policyStatementProps{
	 // Restrict to listing and describing tables
	principals: []iPrincipal{
		iam.NewAnyPrincipal(),
	},
	actions: []*string{
		jsii.String("dynamodb:DescribeTable"),
		jsii.String("dynamodb:ListTables"),
	},
	resources: []*string{
		jsii.String("*"),
	},
}))

// Add an interface endpoint
vpc.addInterfaceEndpoint(jsii.String("EcrDockerEndpoint"), &interfaceVpcEndpointOptions{
	service: ec2.interfaceVpcEndpointAwsService_ECR_DOCKER(),
})
```

By default, CDK will place a VPC endpoint in one subnet per AZ. If you wish to override the AZs CDK places the VPC endpoint in,
use the `subnets` parameter as follows:

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc


ec2.NewInterfaceVpcEndpoint(this, jsii.String("VPC Endpoint"), &interfaceVpcEndpointProps{
	vpc: vpc,
	service: ec2.NewInterfaceVpcEndpointService(jsii.String("com.amazonaws.vpce.us-east-1.vpce-svc-uuddlrlrbastrtsvc"), jsii.Number(443)),
	// Choose which availability zones to place the VPC endpoint in, based on
	// available AZs
	subnets: &subnetSelection{
		availabilityZones: []*string{
			jsii.String("us-east-1a"),
			jsii.String("us-east-1c"),
		},
	},
})
```

Per the [AWS documentation](https://aws.amazon.com/premiumsupport/knowledge-center/interface-endpoint-availability-zone/), not all
VPC endpoint services are available in all AZs. If you specify the parameter `lookupSupportedAzs`, CDK attempts to discover which
AZs an endpoint service is available in, and will ensure the VPC endpoint is not placed in a subnet that doesn't match those AZs.
These AZs will be stored in cdk.context.json.

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc


ec2.NewInterfaceVpcEndpoint(this, jsii.String("VPC Endpoint"), &interfaceVpcEndpointProps{
	vpc: vpc,
	service: ec2.NewInterfaceVpcEndpointService(jsii.String("com.amazonaws.vpce.us-east-1.vpce-svc-uuddlrlrbastrtsvc"), jsii.Number(443)),
	// Choose which availability zones to place the VPC endpoint in, based on
	// available AZs
	lookupSupportedAzs: jsii.Boolean(true),
})
```

Pre-defined AWS services are defined in the [InterfaceVpcEndpointAwsService](lib/vpc-endpoint.ts) class, and can be used to
create VPC endpoints without having to configure name, ports, etc. For example, a Keyspaces endpoint can be created for
use in your VPC:

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc


ec2.NewInterfaceVpcEndpoint(this, jsii.String("VPC Endpoint"), &interfaceVpcEndpointProps{
	vpc: vpc,
	service: ec2.interfaceVpcEndpointAwsService_KEYSPACES(),
})
```

#### Security groups for interface VPC endpoints

By default, interface VPC endpoints create a new security group and traffic is **not**
automatically allowed from the VPC CIDR.

Use the `connections` object to allow traffic to flow to the endpoint:

```go
// Example automatically generated from non-compiling source. May contain errors.
var myEndpoint interfaceVpcEndpoint


myEndpoint.connections.allowDefaultPortFromAnyIpv4()
```

Alternatively, existing security groups can be used by specifying the `securityGroups` prop.

### VPC endpoint services

A VPC endpoint service enables you to expose a Network Load Balancer(s) as a provider service to consumers, who connect to your service over a VPC endpoint. You can restrict access to your service via allowed principals (anything that extends ArnPrincipal), and require that new connections be manually accepted.

```go
// Example automatically generated from non-compiling source. May contain errors.
var networkLoadBalancer1 networkLoadBalancer
var networkLoadBalancer2 networkLoadBalancer


ec2.NewVpcEndpointService(this, jsii.String("EndpointService"), &vpcEndpointServiceProps{
	vpcEndpointServiceLoadBalancers: []iVpcEndpointServiceLoadBalancer{
		networkLoadBalancer1,
		networkLoadBalancer2,
	},
	acceptanceRequired: jsii.Boolean(true),
	allowedPrincipals: []arnPrincipal{
		iam.NewArnPrincipal(jsii.String("arn:aws:iam::123456789012:root")),
	},
})
```

Endpoint services support private DNS, which makes it easier for clients to connect to your service by automatically setting up DNS in their VPC.
You can enable private DNS on an endpoint service like so:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"
var zone hostedZone
var vpces vpcEndpointService


awscdk.NewVpcEndpointServiceDomainName(this, jsii.String("EndpointDomain"), &vpcEndpointServiceDomainNameProps{
	endpointService: vpces,
	domainName: jsii.String("my-stuff.aws-cdk.dev"),
	publicHostedZone: zone,
})
```

Note: The domain name must be owned (registered through Route53) by the account the endpoint service is in, or delegated to the account.
The VpcEndpointServiceDomainName will handle the AWS side of domain verification, the process for which can be found
[here](https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-services-dns-validation.html)

### Client VPN endpoint

AWS Client VPN is a managed client-based VPN service that enables you to securely access your AWS
resources and resources in your on-premises network. With Client VPN, you can access your resources
from any location using an OpenVPN-based VPN client.

Use the `addClientVpnEndpoint()` method to add a client VPN endpoint to a VPC:

```go
vpc.addClientVpnEndpoint(jsii.String("Endpoint"), &clientVpnEndpointOptions{
	cidr: jsii.String("10.100.0.0/16"),
	serverCertificateArn: jsii.String("arn:aws:acm:us-east-1:123456789012:certificate/server-certificate-id"),
	// Mutual authentication
	clientCertificateArn: jsii.String("arn:aws:acm:us-east-1:123456789012:certificate/client-certificate-id"),
	// User-based authentication
	userBasedAuthentication: ec2.clientVpnUserBasedAuthentication.federated(samlProvider),
})
```

The endpoint must use at least one [authentication method](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/client-authentication.html):

* Mutual authentication with a client certificate
* User-based authentication (directory or federated)

If user-based authentication is used, the [self-service portal URL](https://docs.aws.amazon.com/vpn/latest/clientvpn-user/self-service-portal.html)
is made available via a CloudFormation output.

By default, a new security group is created, and logging is enabled. Moreover, a rule to
authorize all users to the VPC CIDR is created.

To customize authorization rules, set the `authorizeAllUsersToVpcCidr` prop to `false`
and use `addAuthorizationRule()`:

```go
endpoint := vpc.addClientVpnEndpoint(jsii.String("Endpoint"), &clientVpnEndpointOptions{
	cidr: jsii.String("10.100.0.0/16"),
	serverCertificateArn: jsii.String("arn:aws:acm:us-east-1:123456789012:certificate/server-certificate-id"),
	userBasedAuthentication: ec2.clientVpnUserBasedAuthentication.federated(samlProvider),
	authorizeAllUsersToVpcCidr: jsii.Boolean(false),
})

endpoint.addAuthorizationRule(jsii.String("Rule"), &clientVpnAuthorizationRuleOptions{
	cidr: jsii.String("10.0.10.0/32"),
	groupId: jsii.String("group-id"),
})
```

Use `addRoute()` to configure network routes:

```go
endpoint := vpc.addClientVpnEndpoint(jsii.String("Endpoint"), &clientVpnEndpointOptions{
	cidr: jsii.String("10.100.0.0/16"),
	serverCertificateArn: jsii.String("arn:aws:acm:us-east-1:123456789012:certificate/server-certificate-id"),
	userBasedAuthentication: ec2.clientVpnUserBasedAuthentication.federated(samlProvider),
})

// Client-to-client access
endpoint.addRoute(jsii.String("Route"), &clientVpnRouteOptions{
	cidr: jsii.String("10.100.0.0/16"),
	target: ec2.clientVpnRouteTarget.local(),
})
```

Use the `connections` object of the endpoint to allow traffic to other security groups.

## Instances

You can use the `Instance` class to start up a single EC2 instance. For production setups, we recommend
you use an `AutoScalingGroup` from the `aws-autoscaling` module instead, as AutoScalingGroups will take
care of restarting your instance if it ever fails.

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc
var instanceType instanceType


// AWS Linux
// AWS Linux
ec2.NewInstance(this, jsii.String("Instance1"), &instanceProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: ec2.NewAmazonLinuxImage(),
})

// AWS Linux 2
// AWS Linux 2
ec2.NewInstance(this, jsii.String("Instance2"), &instanceProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
		generation: ec2.amazonLinuxGeneration_AMAZON_LINUX_2,
	}),
})

// AWS Linux 2 with kernel 5.x
// AWS Linux 2 with kernel 5.x
ec2.NewInstance(this, jsii.String("Instance3"), &instanceProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
		generation: ec2.*amazonLinuxGeneration_AMAZON_LINUX_2,
		kernel: ec2.amazonLinuxKernel_KERNEL5_X,
	}),
})

// AWS Linux 2022
// AWS Linux 2022
ec2.NewInstance(this, jsii.String("Instance4"), &instanceProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
		generation: ec2.*amazonLinuxGeneration_AMAZON_LINUX_2022,
	}),
})

// Graviton 3 Processor
// Graviton 3 Processor
ec2.NewInstance(this, jsii.String("Instance5"), &instanceProps{
	vpc: vpc,
	instanceType: ec2.*instanceType.of(ec2.instanceClass_C7G, ec2.instanceSize_LARGE),
	machineImage: ec2.NewAmazonLinuxImage(&amazonLinuxImageProps{
		generation: ec2.*amazonLinuxGeneration_AMAZON_LINUX_2,
		cpuType: ec2.amazonLinuxCpuType_ARM_64,
	}),
})
```

### Configuring Instances using CloudFormation Init (cfn-init)

CloudFormation Init allows you to configure your instances by writing files to them, installing software
packages, starting services and running arbitrary commands. By default, if any of the instance setup
commands throw an error; the deployment will fail and roll back to the previously known good state.
The following documentation also applies to `AutoScalingGroup`s.

For the full set of capabilities of this system, see the documentation for
[`AWS::CloudFormation::Init`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-init.html).
Here is an example of applying some configuration to an instance:

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


ec2.NewInstance(this, jsii.String("Instance"), &instanceProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// Showing the most complex setup, if you have simpler requirements
	// you can use `CloudFormationInit.fromElements()`.
	init: ec2.cloudFormationInit.fromConfigSets(&configSetProps{
		configSets: map[string][]*string{
			// Applies the configs below in this order
			"default": []*string{
				jsii.String("yumPreinstall"),
				jsii.String("config"),
			},
		},
		configs: map[string]initConfig{
			"yumPreinstall": ec2.NewInitConfig([]InitElement{
				ec2.InitPackage.yum(jsii.String("git")),
			}),
			"config": ec2.NewInitConfig([]InitElement{
				ec2.InitFile.fromObject(jsii.String("/etc/stack.json"), map[string]interface{}{
					"stackId": awscdk.*stack.of(this).stackId,
					"stackName": awscdk.*stack.of(this).stackName,
					"region": awscdk.*stack.of(this).region,
				}),
				ec2.InitGroup.fromName(jsii.String("my-group")),
				ec2.InitUser.fromName(jsii.String("my-user")),
				ec2.InitPackage.rpm(jsii.String("http://mirrors.ukfast.co.uk/sites/dl.fedoraproject.org/pub/epel/8/Everything/x86_64/Packages/r/rubygem-git-1.5.0-2.el8.noarch.rpm")),
			}),
		},
	}),
	initOptions: &applyCloudFormationInitOptions{
		// Optional, which configsets to activate (['default'] by default)
		configSets: []*string{
			jsii.String("default"),
		},

		// Optional, how long the installation is expected to take (5 minutes by default)
		timeout: awscdk.Duration.minutes(jsii.Number(30)),

		// Optional, whether to include the --url argument when running cfn-init and cfn-signal commands (false by default)
		includeUrl: jsii.Boolean(true),

		// Optional, whether to include the --role argument when running cfn-init and cfn-signal commands (false by default)
		includeRole: jsii.Boolean(true),
	},
})
```

You can have services restarted after the init process has made changes to the system.
To do that, instantiate an `InitServiceRestartHandle` and pass it to the config elements
that need to trigger the restart and the service itself. For example, the following
config writes a config file for nginx, extracts an archive to the root directory, and then
restarts nginx so that it picks up the new config and files:

```go
// Example automatically generated from non-compiling source. May contain errors.
var myBucket bucket


handle := ec2.NewInitServiceRestartHandle()

ec2.cloudFormationInit.fromElements(ec2.initFile.fromString(jsii.String("/etc/nginx/nginx.conf"), jsii.String("..."), &initFileOptions{
	serviceRestartHandles: []initServiceRestartHandle{
		handle,
	},
}), ec2.initSource.fromS3Object(jsii.String("/var/www/html"), myBucket, jsii.String("html.zip"), &initSourceOptions{
	serviceRestartHandles: []*initServiceRestartHandle{
		handle,
	},
}), ec2.initService.enable(jsii.String("nginx"), &initServiceOptions{
	serviceRestartHandle: handle,
}))
```

### Bastion Hosts

A bastion host functions as an instance used to access servers and resources in a VPC without open up the complete VPC on a network level.
You can use bastion hosts using a standard SSH connection targeting port 22 on the host. As an alternative, you can connect the SSH connection
feature of AWS Systems Manager Session Manager, which does not need an opened security group. (https://aws.amazon.com/about-aws/whats-new/2019/07/session-manager-launches-tunneling-support-for-ssh-and-scp/)

A default bastion host for use via SSM can be configured like:

```go
host := ec2.NewBastionHostLinux(this, jsii.String("BastionHost"), &bastionHostLinuxProps{
	vpc: vpc,
})
```

If you want to connect from the internet using SSH, you need to place the host into a public subnet. You can then configure allowed source hosts.

```go
host := ec2.NewBastionHostLinux(this, jsii.String("BastionHost"), &bastionHostLinuxProps{
	vpc: vpc,
	subnetSelection: &subnetSelection{
		subnetType: ec2.subnetType_PUBLIC,
	},
})
host.allowSshAccessFrom(ec2.peer.ipv4(jsii.String("1.2.3.4/32")))
```

As there are no SSH public keys deployed on this machine, you need to use [EC2 Instance Connect](https://aws.amazon.com/de/blogs/compute/new-using-amazon-ec2-instance-connect-for-ssh-access-to-your-ec2-instances/)
with the command `aws ec2-instance-connect send-ssh-public-key` to provide your SSH public key.

EBS volume for the bastion host can be encrypted like:

```go
host := ec2.NewBastionHostLinux(this, jsii.String("BastionHost"), &bastionHostLinuxProps{
	vpc: vpc,
	blockDevices: []blockDevice{
		&blockDevice{
			deviceName: jsii.String("EBSBastionHost"),
			volume: ec2.blockDeviceVolume.ebs(jsii.Number(10), &ebsDeviceOptions{
				encrypted: jsii.Boolean(true),
			}),
		},
	},
})
```

### Block Devices

To add EBS block device mappings, specify the `blockDevices` property. The following example sets the EBS-backed
root device (`/dev/sda1`) size to 50 GiB, and adds another EBS-backed device mapped to `/dev/sdm` that is 100 GiB in
size:

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


ec2.NewInstance(this, jsii.String("Instance"), &instanceProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// ...

	blockDevices: []blockDevice{
		&blockDevice{
			deviceName: jsii.String("/dev/sda1"),
			volume: ec2.blockDeviceVolume.ebs(jsii.Number(50)),
		},
		&blockDevice{
			deviceName: jsii.String("/dev/sdm"),
			volume: ec2.*blockDeviceVolume.ebs(jsii.Number(100)),
		},
	},
})
```

It is also possible to encrypt the block devices. In this example we will create an customer managed key encrypted EBS-backed root device:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


kmsKey := awscdk.NewKey(this, jsii.String("KmsKey"))

ec2.NewInstance(this, jsii.String("Instance"), &instanceProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// ...

	blockDevices: []blockDevice{
		&blockDevice{
			deviceName: jsii.String("/dev/sda1"),
			volume: ec2.blockDeviceVolume.ebs(jsii.Number(50), &ebsDeviceOptions{
				encrypted: jsii.Boolean(true),
				kmsKey: kmsKey,
			}),
		},
	},
})
```

### Volumes

Whereas a `BlockDeviceVolume` is an EBS volume that is created and destroyed as part of the creation and destruction of a specific instance. A `Volume` is for when you want an EBS volume separate from any particular instance. A `Volume` is an EBS block device that can be attached to, or detached from, any instance at any time. Some types of `Volume`s can also be attached to multiple instances at the same time to allow you to have shared storage between those instances.

A notable restriction is that a Volume can only be attached to instances in the same availability zone as the Volume itself.

The following demonstrates how to create a 500 GiB encrypted Volume in the `us-west-2a` availability zone, and give a role the ability to attach that Volume to a specific instance:

```go
// Example automatically generated from non-compiling source. May contain errors.
var instance instance
var role role


volume := ec2.NewVolume(this, jsii.String("Volume"), &volumeProps{
	availabilityZone: jsii.String("us-west-2a"),
	size: awscdk.Size.gibibytes(jsii.Number(500)),
	encrypted: jsii.Boolean(true),
})

volume.grantAttachVolume(role, []iInstance{
	instance,
})
```

#### Instances Attaching Volumes to Themselves

If you need to grant an instance the ability to attach/detach an EBS volume to/from itself, then using `grantAttachVolume` and `grantDetachVolume` as outlined above
will lead to an unresolvable circular reference between the instance role and the instance. In this case, use `grantAttachVolumeByResourceTag` and `grantDetachVolumeByResourceTag` as follows:

```go
// Example automatically generated from non-compiling source. May contain errors.
var instance instance
var volume volume


attachGrant := volume.grantAttachVolumeByResourceTag(instance.grantPrincipal, []construct{
	instance,
})
detachGrant := volume.grantDetachVolumeByResourceTag(instance.grantPrincipal, []construct{
	instance,
})
```

#### Attaching Volumes

The Amazon EC2 documentation for
[Linux Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AmazonEBS.html) and
[Windows Instances](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-volumes.html) contains information on how
to attach and detach your Volumes to/from instances, and how to format them for use.

The following is a sample skeleton of EC2 UserData that can be used to attach a Volume to the Linux instance that it is running on:

```go
// Example automatically generated from non-compiling source. May contain errors.
var instance instance
var volume volume


volume.grantAttachVolumeByResourceTag(instance.grantPrincipal, []construct{
	instance,
})
targetDevice := "/dev/xvdz"
instance.userData.addCommands(jsii.String("TOKEN=$(curl -SsfX PUT \"http://169.254.169.254/latest/api/token\" -H \"X-aws-ec2-metadata-token-ttl-seconds: 21600\")"), jsii.String("INSTANCE_ID=$(curl -SsfH \"X-aws-ec2-metadata-token: $TOKEN\" http://169.254.169.254/latest/meta-data/instance-id)"),
// Attach the volume to /dev/xvdz
fmt.Sprintf("aws --region %v ec2 attach-volume --volume-id %v --instance-id $INSTANCE_ID --device %v", awscdk.stack.of(this).region, volume.volumeId, targetDevice),
// Wait until the volume has attached
fmt.Sprintf("while ! test -e %v; do sleep 1; done", targetDevice))
```

#### Tagging Volumes

You can configure [tag propagation on volume creation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#cfn-ec2-instance-propagatetagstovolumeoncreation).

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


ec2.NewInstance(this, jsii.String("Instance"), &instanceProps{
	vpc: vpc,
	machineImage: machineImage,
	instanceType: instanceType,
	propagateTagsToVolumeOnCreation: jsii.Boolean(true),
})
```

#### Throughput on GP3 Volumes

You can specify the `throughput` of a GP3 volume from 125 (default) to 1000.

```go
// Example automatically generated from non-compiling source. May contain errors.
ec2.NewVolume(this, jsii.String("Volume"), &volumeProps{
	availabilityZone: jsii.String("us-east-1a"),
	size: cdk.size_Gibibytes(jsii.Number(125)),
	volumeType: ebsDeviceVolumeType_GP3,
	throughput: jsii.Number(125),
})
```

### Configuring Instance Metadata Service (IMDS)

#### Toggling IMDSv1

You can configure [EC2 Instance Metadata Service](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) options to either
allow both IMDSv1 and IMDSv2 or enforce IMDSv2 when interacting with the IMDS.

To do this for a single `Instance`, you can use the `requireImdsv2` property.
The example below demonstrates IMDSv2 being required on a single `Instance`:

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc
var instanceType instanceType
var machineImage iMachineImage


ec2.NewInstance(this, jsii.String("Instance"), &instanceProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: machineImage,

	// ...

	requireImdsv2: jsii.Boolean(true),
})
```

You can also use the either the `InstanceRequireImdsv2Aspect` for EC2 instances or the `LaunchTemplateRequireImdsv2Aspect` for EC2 launch templates
to apply the operation to multiple instances or launch templates, respectively.

The following example demonstrates how to use the `InstanceRequireImdsv2Aspect` to require IMDSv2 for all EC2 instances in a stack:

```go
// Example automatically generated from non-compiling source. May contain errors.
aspect := ec2.NewInstanceRequireImdsv2Aspect()
awscdk.Aspects.of(this).add(aspect)
```

## VPC Flow Logs

VPC Flow Logs is a feature that enables you to capture information about the IP traffic going to and from network interfaces in your VPC. Flow log data can be published to Amazon CloudWatch Logs and Amazon S3. After you've created a flow log, you can retrieve and view its data in the chosen destination. ([https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html)).

By default, a flow log will be created with CloudWatch Logs as the destination.

You can create a flow log like this:

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc


ec2.NewFlowLog(this, jsii.String("FlowLog"), &flowLogProps{
	resourceType: ec2.flowLogResourceType.fromVpc(vpc),
})
```

Or you can add a Flow Log to a VPC by using the addFlowLog method like this:

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.NewVpc(this, jsii.String("Vpc"))

vpc.addFlowLog(jsii.String("FlowLog"))
```

You can also add multiple flow logs with different destinations.

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.NewVpc(this, jsii.String("Vpc"))

vpc.addFlowLog(jsii.String("FlowLogS3"), &flowLogOptions{
	destination: ec2.flowLogDestination.toS3(),
})

// Only reject traffic and interval every minute.
vpc.addFlowLog(jsii.String("FlowLogCloudWatch"), &flowLogOptions{
	trafficType: ec2.flowLogTrafficType_REJECT,
	maxAggregationInterval: flowLogMaxAggregationInterval_ONE_MINUTE,
})
```

### Custom Formatting

You can also custom format flow logs.

```go
// Example automatically generated from non-compiling source. May contain errors.
vpc := ec2.NewVpc(this, jsii.String("Vpc"))

vpc.addFlowLog(jsii.String("FlowLog"), &flowLogOptions{
	logFormat: []logFormat{
		ec2.*logFormat_DST_PORT(),
		ec2.*logFormat_SRC_PORT(),
	},
})

// If you just want to add a field to the default field
vpc.addFlowLog(jsii.String("FlowLog"), &flowLogOptions{
	logFormat: []*logFormat{
		ec2.*logFormat_VERSION(),
		ec2.*logFormat_ALL_DEFAULT_FIELDS(),
	},
})

// If AWS CDK does not support the new fields
vpc.addFlowLog(jsii.String("FlowLog"), &flowLogOptions{
	logFormat: []*logFormat{
		ec2.*logFormat_SRC_PORT(),
		ec2.*logFormat.custom(jsii.String("${new-field}")),
	},
})
```

By default, the CDK will create the necessary resources for the destination. For the CloudWatch Logs destination
it will create a CloudWatch Logs Log Group as well as the IAM role with the necessary permissions to publish to
the log group. In the case of an S3 destination, it will create the S3 bucket.

If you want to customize any of the destination resources you can provide your own as part of the `destination`.

*CloudWatch Logs*

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc


logGroup := logs.NewLogGroup(this, jsii.String("MyCustomLogGroup"))

role := iam.NewRole(this, jsii.String("MyCustomRole"), &roleProps{
	assumedBy: iam.NewServicePrincipal(jsii.String("vpc-flow-logs.amazonaws.com")),
})

ec2.NewFlowLog(this, jsii.String("FlowLog"), &flowLogProps{
	resourceType: ec2.flowLogResourceType.fromVpc(vpc),
	destination: ec2.flowLogDestination.toCloudWatchLogs(logGroup, role),
})
```

*S3*

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc


bucket := s3.NewBucket(this, jsii.String("MyCustomBucket"))

ec2.NewFlowLog(this, jsii.String("FlowLog"), &flowLogProps{
	resourceType: ec2.flowLogResourceType.fromVpc(vpc),
	destination: ec2.flowLogDestination.toS3(bucket),
})

ec2.NewFlowLog(this, jsii.String("FlowLogWithKeyPrefix"), &flowLogProps{
	resourceType: ec2.*flowLogResourceType.fromVpc(vpc),
	destination: ec2.*flowLogDestination.toS3(bucket, jsii.String("prefix/")),
})
```

When the S3 destination is configured, AWS will automatically create an S3 bucket policy
that allows the service to write logs to the bucket. This makes it impossible to later update
that bucket policy. To have CDK create the bucket policy so that future updates can be made,
the `@aws-cdk/aws-s3:createDefaultLoggingPolicy` [feature flag](https://docs.aws.amazon.com/cdk/v2/guide/featureflags.html) can be used. This can be set
in the `cdk.json` file.

```json
{
  "context": {
    "@aws-cdk/aws-s3:createDefaultLoggingPolicy": true
  }
}
```

## User Data

User data enables you to run a script when your instances start up.  In order to configure these scripts you can add commands directly to the script
or you can use the UserData's convenience functions to aid in the creation of your script.

A user data could be configured to run a script found in an asset through the following:

```go
// Example automatically generated from non-compiling source. May contain errors.
import "github.com/aws/aws-cdk-go/awscdk"

var instance instance


asset := awscdk.NewAsset(this, jsii.String("Asset"), &assetProps{
	path: jsii.String("./configure.sh"),
})

localPath := instance.userData.addS3DownloadCommand(&s3DownloadOptions{
	bucket: asset.bucket,
	bucketKey: asset.s3ObjectKey,
	region: jsii.String("us-east-1"),
})
instance.userData.addExecuteFileCommand(&executeFileOptions{
	filePath: localPath,
	arguments: jsii.String("--verbose -y"),
})
asset.grantRead(instance.role)
```

### Persisting user data

By default, EC2 UserData is run once on only the first time that an instance is started. It is possible to make the
user data script run on every start of the instance.

When creating a Windows UserData you can use the `persist` option to set whether or not to add
`<persist>true</persist>` [to the user data script](https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-windows-user-data.html#user-data-scripts). it can be used as follows:

```go
// Example automatically generated from non-compiling source. May contain errors.
windowsUserData := userData_ForWindows(map[string]*bool{
	"persist": jsii.Boolean(true),
})
```

For a Linux instance, this can be accomplished by using a Multipart user data to configure cloud-config as detailed
in: https://aws.amazon.com/premiumsupport/knowledge-center/execute-user-data-ec2/

### Multipart user data

In addition, to above the `MultipartUserData` can be used to change instance startup behavior. Multipart user data are composed
from separate parts forming archive. The most common parts are scripts executed during instance set-up. However, there are other
kinds, too.

The advantage of multipart archive is in flexibility when it's needed to add additional parts or to use specialized parts to
fine tune instance startup. Some services (like AWS Batch) support only `MultipartUserData`.

The parts can be executed at different moment of instance start-up and can serve a different purpose. This is controlled by `contentType` property.
For common scripts, `text/x-shellscript; charset="utf-8"` can be used as content type.

In order to create archive the `MultipartUserData` has to be instantiated. Than, user can add parts to multipart archive using `addPart`. The `MultipartBody` contains methods supporting creation of body parts.

If the very custom part is required, it can be created using `MultipartUserData.fromRawBody`, in this case full control over content type,
transfer encoding, and body properties is given to the user.

Below is an example for creating multipart user data with single body part responsible for installing `awscli` and configuring maximum size
of storage used by Docker containers:

```go
// Example automatically generated from non-compiling source. May contain errors.
bootHookConf := ec2.userData.forLinux()
bootHookConf.addCommands(jsii.String("cloud-init-per once docker_options echo 'OPTIONS=\"${OPTIONS} --storage-opt dm.basesize=40G\"' >> /etc/sysconfig/docker"))

setupCommands := ec2.userData.forLinux()
setupCommands.addCommands(jsii.String("sudo yum install awscli && echo Packages installed らと > /var/tmp/setup"))

multipartUserData := ec2.NewMultipartUserData()
// The docker has to be configured at early stage, so content type is overridden to boothook
multipartUserData.addPart(ec2.multipartBody.fromUserData(bootHookConf, jsii.String("text/cloud-boothook; charset=\"us-ascii\"")))
// Execute the rest of setup
multipartUserData.addPart(ec2.multipartBody.fromUserData(setupCommands))

ec2.NewLaunchTemplate(this, jsii.String(""), &launchTemplateProps{
	userData: multipartUserData,
	blockDevices: []blockDevice{
	},
})
```

For more information see
[Specifying Multiple User Data Blocks Using a MIME Multi Part Archive](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/bootstrap_container_instance.html#multi-part_user_data)

#### Using add*Command on MultipartUserData

To use the `add*Command` methods, that are inherited from the `UserData` interface, on `MultipartUserData` you must add a part
to the `MultipartUserData` and designate it as the receiver for these methods. This is accomplished by using the `addUserDataPart()`
method on `MultipartUserData` with the `makeDefault` argument set to `true`:

```go
// Example automatically generated from non-compiling source. May contain errors.
multipartUserData := ec2.NewMultipartUserData()
commandsUserData := ec2.userData.forLinux()
multipartUserData.addUserDataPart(commandsUserData, ec2.multipartBody_SHELL_SCRIPT(), jsii.Boolean(true))

// Adding commands to the multipartUserData adds them to commandsUserData, and vice-versa.
multipartUserData.addCommands(jsii.String("touch /root/multi.txt"))
commandsUserData.addCommands(jsii.String("touch /root/userdata.txt"))
```

When used on an EC2 instance, the above `multipartUserData` will create both `multi.txt` and `userdata.txt` in `/root`.

## Importing existing subnet

To import an existing Subnet, call `Subnet.fromSubnetAttributes()` or
`Subnet.fromSubnetId()`. Only if you supply the subnet's Availability Zone
and Route Table Ids when calling `Subnet.fromSubnetAttributes()` will you be
able to use the CDK features that use these values (such as selecting one
subnet per AZ).

Importing an existing subnet looks like this:

```go
// Example automatically generated from non-compiling source. May contain errors.
// Supply all properties
subnet1 := ec2.subnet.fromSubnetAttributes(this, jsii.String("SubnetFromAttributes"), &subnetAttributes{
	subnetId: jsii.String("s-1234"),
	availabilityZone: jsii.String("pub-az-4465"),
	routeTableId: jsii.String("rt-145"),
})

// Supply only subnet id
subnet2 := ec2.subnet.fromSubnetId(this, jsii.String("SubnetFromId"), jsii.String("s-1234"))
```

## Launch Templates

A Launch Template is a standardized template that contains the configuration information to launch an instance.
They can be used when launching instances on their own, through Amazon EC2 Auto Scaling, EC2 Fleet, and Spot Fleet.
Launch templates enable you to store launch parameters so that you do not have to specify them every time you launch
an instance. For information on Launch Templates please see the
[official documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html).

The following demonstrates how to create a launch template with an Amazon Machine Image, and security group.

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc


template := ec2.NewLaunchTemplate(this, jsii.String("LaunchTemplate"), &launchTemplateProps{
	machineImage: ec2.machineImage.latestAmazonLinux(),
	securityGroup: ec2.NewSecurityGroup(this, jsii.String("LaunchTemplateSG"), &securityGroupProps{
		vpc: vpc,
	}),
})
```

And the following demonstrates how to enable metadata options support.

```go
// Example automatically generated from non-compiling source. May contain errors.
ec2.NewLaunchTemplate(this, jsii.String("LaunchTemplate"), &launchTemplateProps{
	httpEndpoint: jsii.Boolean(true),
	httpProtocolIpv6: jsii.Boolean(true),
	httpPutResponseHopLimit: jsii.Number(1),
	httpTokens: ec2.launchTemplateHttpTokens_REQUIRED,
	instanceMetadataTags: jsii.Boolean(true),
})
```

## Detailed Monitoring

The following demonstrates how to enable [Detailed Monitoring](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch-new.html) for an EC2 instance. Keep in mind that Detailed Monitoring results in [additional charges](http://aws.amazon.com/cloudwatch/pricing/).

```go
// Example automatically generated from non-compiling source. May contain errors.
var vpc vpc
var instanceType instanceType


ec2.NewInstance(this, jsii.String("Instance1"), &instanceProps{
	vpc: vpc,
	instanceType: instanceType,
	machineImage: ec2.NewAmazonLinuxImage(),
	detailedMonitoring: jsii.Boolean(true),
})
```

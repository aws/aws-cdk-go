# Amazon VpcV2 Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## VpcV2

`VpcV2` is a re-write of the [`ec2.Vpc`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.Vpc.html) construct. This new construct enables higher level of customization
on the VPC being created. `VpcV2` implements the existing [`IVpc`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.IVpc.html), therefore,
`VpcV2` is compatible with other constructs that accepts `IVpc` (e.g. [`ApplicationLoadBalancer`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer.html#construct-props)).

To create a VPC with both IPv4 and IPv6 support:

```go
stack := awscdk.Newstack()
vpc_v2.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
	PrimaryAddressBlock: vpc_v2.IpAddresses_Ipv4(jsii.String("10.0.0.0/24")),
	SecondaryAddressBlocks: []iIpAddresses{
		vpc_v2.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
			CidrBlockName: jsii.String("AmazonProvidedIpv6"),
		}),
	},
})
```

`VpcV2` does not automatically create subnets or allocate IP addresses, which is different from the `Vpc` construct.

Importing existing VPC in an account into CDK as a `VpcV2` is not yet supported.

## SubnetV2

`SubnetV2` is a re-write of the [`ec2.Subnet`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.Subnet.html) construct.
This new construct can be used to add subnets to a `VpcV2` instance:

```go
stack := awscdk.Newstack()
myVpc := vpc_v2.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
	SecondaryAddressBlocks: []iIpAddresses{
		vpc_v2.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
			CidrBlockName: jsii.String("AmazonProvidedIp"),
		}),
	},
})

vpc_v2.NewSubnetV2(this, jsii.String("subnetA"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("us-east-1a"),
	Ipv4CidrBlock: vpc_v2.NewIpCidr(jsii.String("10.0.0.0/24")),
	Ipv6CidrBlock: vpc_v2.NewIpCidr(jsii.String("2a05:d02c:25:4000::/60")),
	SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
})
```

Same as `VpcV2`, importing existing subnets is not yet supported.

## IP Addresses Management

By default `VpcV2` uses `10.0.0.0/16` as the primary CIDR if none is defined.
Additional CIDRs can be adding to the VPC via the `secondaryAddressBlocks` prop.
The following example illustrates the different options of defining the address blocks:

```go
stack := awscdk.Newstack()
ipam := awsec2alpha.NewIpam(this, jsii.String("Ipam"), &IpamProps{
	OperatingRegion: []*string{
		jsii.String("us-west-1"),
	},
})
ipamPublicPool := ipam.PublicScope.AddPool(jsii.String("PublicPoolA"), &PoolOptions{
	AddressFamily: vpc_v2.AddressFamily_IP_V6,
	AwsService: awsec2alpha.AwsServiceName_EC2,
	Locale: jsii.String("us-west-1"),
	PublicIpSource: vpc_v2.IpamPoolPublicIpSource_AMAZON,
})
ipamPublicPool.ProvisionCidr(jsii.String("PublicPoolACidrA"), &IpamPoolCidrProvisioningOptions{
	NetmaskLength: jsii.Number(52),
})

ipamPrivatePool := ipam.PrivateScope.AddPool(jsii.String("PrivatePoolA"), &PoolOptions{
	AddressFamily: vpc_v2.AddressFamily_IP_V4,
})
ipamPrivatePool.ProvisionCidr(jsii.String("PrivatePoolACidrA"), &IpamPoolCidrProvisioningOptions{
	NetmaskLength: jsii.Number(8),
})

vpc_v2.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
	PrimaryAddressBlock: vpc_v2.IpAddresses_Ipv4(jsii.String("10.0.0.0/24")),
	SecondaryAddressBlocks: []iIpAddresses{
		vpc_v2.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
			CidrBlockName: jsii.String("AmazonIpv6"),
		}),
		vpc_v2.IpAddresses_Ipv6Ipam(&IpamOptions{
			IpamPool: ipamPublicPool,
			NetmaskLength: jsii.Number(52),
			CidrBlockName: jsii.String("ipv6Ipam"),
		}),
		vpc_v2.IpAddresses_Ipv4Ipam(&IpamOptions{
			IpamPool: ipamPrivatePool,
			NetmaskLength: jsii.Number(8),
			CidrBlockName: jsii.String("ipv4Ipam"),
		}),
	},
})
```

Since `VpcV2` does not create subnets automatically, users have full control over IP addresses allocation across subnets.

## Routing

`RouteTable` is a new construct that allows for route tables to be customized in a variety of ways. For instance, the following example shows how a custom route table can be created and appended to a subnet:

```go
myVpc := vpc_v2.NewVpcV2(this, jsii.String("Vpc"))
routeTable := vpc_v2.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})
subnet := vpc_v2.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	RouteTable: RouteTable,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
})
```

`Route`s can be created to link subnets to various different AWS services via gateways and endpoints. Each unique route target has its own dedicated construct that can be routed to a given subnet via the `Route` construct. An example using the `InternetGateway` construct can be seen below:

```go
stack := awscdk.Newstack()
myVpc := vpc_v2.NewVpcV2(this, jsii.String("Vpc"))
routeTable := vpc_v2.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})
subnet := vpc_v2.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
})

igw := vpc_v2.NewInternetGateway(this, jsii.String("IGW"), &InternetGatewayProps{
	Vpc: myVpc,
})
vpc_v2.NewRoute(this, jsii.String("IgwRoute"), &RouteProps{
	RouteTable: RouteTable,
	Destination: jsii.String("0.0.0.0/0"),
	Target: map[string]iRouteTarget{
		"gateway": igw,
	},
})
```

Other route targets may require a deeper set of parameters to set up properly. For instance, the example below illustrates how to set up a `NatGateway`:

```go
myVpc := vpc_v2.NewVpcV2(this, jsii.String("Vpc"))
routeTable := vpc_v2.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})
subnet := vpc_v2.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
})

natgw := vpc_v2.NewNatGateway(this, jsii.String("NatGW"), &NatGatewayProps{
	Subnet: subnet,
	Vpc: myVpc,
	ConnectivityType: awsec2alpha.NatConnectivityType_PRIVATE,
	PrivateIpAddress: jsii.String("10.0.0.42"),
})
vpc_v2.NewRoute(this, jsii.String("NatGwRoute"), &RouteProps{
	RouteTable: RouteTable,
	Destination: jsii.String("0.0.0.0/0"),
	Target: map[string]iRouteTarget{
		"gateway": natgw,
	},
})
```

It is also possible to set up endpoints connecting other AWS services. For instance, the example below illustrates the linking of a Dynamo DB endpoint via the existing `ec2.GatewayVpcEndpoint` construct as a route target:

```go
myVpc := vpc_v2.NewVpcV2(this, jsii.String("Vpc"))
routeTable := vpc_v2.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})
subnet := vpc_v2.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: ec2.SubnetType_PRIVATE,
})

dynamoEndpoint := ec2.NewGatewayVpcEndpoint(this, jsii.String("DynamoEndpoint"), &GatewayVpcEndpointProps{
	Service: ec2.GatewayVpcEndpointAwsService_DYNAMODB(),
	Vpc: myVpc,
	Subnets: []subnetSelection{
		subnet,
	},
})
vpc_v2.NewRoute(this, jsii.String("DynamoDBRoute"), &RouteProps{
	RouteTable: RouteTable,
	Destination: jsii.String("0.0.0.0/0"),
	Target: map[string]iVpcEndpoint{
		"endpoint": dynamoEndpoint,
	},
})
```

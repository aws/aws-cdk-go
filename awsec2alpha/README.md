# Amazon VpcV2 Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Developer Preview](https://img.shields.io/badge/cdk--constructs-developer--preview-informational.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are in **developer preview** before they
> become stable. We will only make breaking changes to address unforeseen API issues. Therefore,
> these APIs are not subject to [Semantic Versioning](https://semver.org/), and breaking changes
> will be announced in release notes. This means that while you may use them, you may need to
> update your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

## VpcV2

`VpcV2` is a re-write of the [`ec2.Vpc`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.Vpc.html) construct. This new construct enables higher level of customization
on the VPC being created. `VpcV2` implements the existing [`IVpc`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.IVpc.html), therefore,
`VpcV2` is compatible with other constructs that accepts `IVpc` (e.g. [`ApplicationLoadBalancer`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_elasticloadbalancingv2.ApplicationLoadBalancer.html#construct-props)).

`VpcV2` supports the addition of both primary and secondary addresses. The primary address must be an IPv4 address, which can be specified as a CIDR string or assigned from an IPAM pool. Secondary addresses can be either IPv4 or IPv6.
By default, `VpcV2` assigns `10.0.0.0/16` as the primary CIDR if no other CIDR is specified.

Below is an example of creating a VPC with both IPv4 and IPv6 support:

```go
stack := awscdk.Newstack()
awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.0.0.0/24")),
	SecondaryAddressBlocks: []IIpAddresses{
		awsec2alpha.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
			CidrBlockName: jsii.String("AmazonProvidedIpv6"),
		}),
	},
})
```

`VpcV2` does not automatically create subnets or allocate IP addresses, which is different from the `Vpc` construct.

## SubnetV2

`SubnetV2` is a re-write of the [`ec2.Subnet`](https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_ec2.Subnet.html) construct.
This new construct can be used to add subnets to a `VpcV2` instance:
Note: When defining a subnet with `SubnetV2`, CDK automatically creates a new route table, unless a route table is explicitly provided as an input to the construct.
To enable the `mapPublicIpOnLaunch` feature (which is `false` by default), set the property to `true` when creating the subnet.

```go
stack := awscdk.Newstack()
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
	SecondaryAddressBlocks: []IIpAddresses{
		awsec2alpha.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
			CidrBlockName: jsii.String("AmazonProvidedIp"),
		}),
	},
})

awsec2alpha.NewSubnetV2(this, jsii.String("subnetA"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("us-east-1a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	Ipv6CidrBlock: awsec2alpha.NewIpCidr(jsii.String("2a05:d02c:25:4000::/60")),
	SubnetType: awscdk.SubnetType_PUBLIC,
	MapPublicIpOnLaunch: jsii.Boolean(true),
})
```

Since `VpcV2` does not create subnets automatically, users have full control over IP addresses allocation across subnets.

## IP Addresses Management

Additional CIDRs can be added to the VPC via the `secondaryAddressBlocks` property.
The following example illustrates the options of defining these secondary address blocks using `IPAM`:

Note: There’s currently an issue with IPAM pool deletion that may affect the `cdk --destroy` command. This is because IPAM takes time to detect when the IP address pool has been deallocated after the VPC is deleted. The current workaround is to wait until the IP address is fully deallocated from the pool before retrying the deletion. Below command can be used to check allocations for a pool using CLI

```shell
aws ec2 get-ipam-pool-allocations --ipam-pool-id <ipam-pool-id>
```

Ref: https://docs.aws.amazon.com/cli/latest/reference/ec2/get-ipam-pool-allocations.html

```go
stack := awscdk.Newstack()
ipam := awsec2alpha.NewIpam(this, jsii.String("Ipam"), &IpamProps{
	OperatingRegions: []*string{
		jsii.String("us-west-1"),
	},
})
ipamPublicPool := ipam.PublicScope.AddPool(jsii.String("PublicPoolA"), &PoolOptions{
	AddressFamily: awsec2alpha.AddressFamily_IP_V6,
	AwsService: awsec2alpha.AwsServiceName_EC2,
	Locale: jsii.String("us-west-1"),
	PublicIpSource: awsec2alpha.IpamPoolPublicIpSource_AMAZON,
})
ipamPublicPool.ProvisionCidr(jsii.String("PublicPoolACidrA"), &IpamPoolCidrProvisioningOptions{
	NetmaskLength: jsii.Number(52),
})

ipamPrivatePool := ipam.PrivateScope.AddPool(jsii.String("PrivatePoolA"), &PoolOptions{
	AddressFamily: awsec2alpha.AddressFamily_IP_V4,
})
ipamPrivatePool.ProvisionCidr(jsii.String("PrivatePoolACidrA"), &IpamPoolCidrProvisioningOptions{
	NetmaskLength: jsii.Number(8),
})

awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.0.0.0/24")),
	SecondaryAddressBlocks: []IIpAddresses{
		awsec2alpha.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
			CidrBlockName: jsii.String("AmazonIpv6"),
		}),
		awsec2alpha.IpAddresses_Ipv6Ipam(&IpamOptions{
			IpamPool: ipamPublicPool,
			NetmaskLength: jsii.Number(52),
			CidrBlockName: jsii.String("ipv6Ipam"),
		}),
		awsec2alpha.IpAddresses_Ipv4Ipam(&IpamOptions{
			IpamPool: ipamPrivatePool,
			NetmaskLength: jsii.Number(8),
			CidrBlockName: jsii.String("ipv4Ipam"),
		}),
	},
})
```

### Bring your own IPv6 addresses (BYOIP)

If you have your own IP address that you would like to use with EC2, you can set up an IPv6 pool via the AWS CLI, and use that pool ID in your application.

Once you have certified your IP address block with an ROA and have obtained an X-509 certificate, you can run the following command to provision your CIDR block in your AWS account:

```shell
aws ec2 provision-byoip-cidr --region <region> --cidr <your CIDR block> --cidr-authorization-context Message="1|aws|<account>|<your CIDR block>|<expiration date>|SHA256".Signature="<signature>"
```

When your BYOIP CIDR is provisioned, you can run the following command to retrieve your IPv6 pool ID, which will be used in your VPC declaration:

```shell
aws ec2 describe-byoip-cidr --region <region>
```

For more help on setting up your IPv6 address, please review the [EC2 Documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html).

Once you have provisioned your address block, you can use the IPv6 in your VPC as follows:

```go
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.1.0.0/16")),
	SecondaryAddressBlocks: []IIpAddresses{
		awsec2alpha.IpAddresses_Ipv6ByoipPool(&Ipv6PoolSecondaryAddressProps{
			CidrBlockName: jsii.String("MyByoipCidrBlock"),
			Ipv6PoolId: jsii.String("ipv6pool-ec2-someHashValue"),
			Ipv6CidrBlock: jsii.String("2001:db8::/32"),
		}),
	},
	EnableDnsHostnames: jsii.Boolean(true),
	EnableDnsSupport: jsii.Boolean(true),
})
```

## Routing

`RouteTable` is a new construct that allows for route tables to be customized in a variety of ways. Using this construct, a customized route table can be added to the subnets defined using `SubnetV2`.
For instance, the following example shows how a custom route table can be created and appended to a `SubnetV2`:

```go
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})
subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	RouteTable: RouteTable,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PRIVATE_ISOLATED,
})
```

`Routes` can be created to link subnets to various different AWS services via gateways and endpoints. Each unique route target has its own dedicated construct that can be routed to a given subnet via the `Route` construct. An example using the `InternetGateway` construct can be seen below:

```go
stack := awscdk.Newstack()
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})
subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PRIVATE_ISOLATED,
})

igw := awsec2alpha.NewInternetGateway(this, jsii.String("IGW"), &InternetGatewayProps{
	Vpc: myVpc,
})
awsec2alpha.NewRoute(this, jsii.String("IgwRoute"), &RouteProps{
	RouteTable: RouteTable,
	Destination: jsii.String("0.0.0.0/0"),
	Target: map[string]iRouteTarget{
		"gateway": igw,
	},
})
```

Alternatively, `Routes` can also be created via method `addRoute` in the `RouteTable` class. An example using the `EgressOnlyInternetGateway` construct can be seen below:
Note: `EgressOnlyInternetGateway` can only be used to set up outbound IPv6 routing.

```go
stack := awscdk.Newstack()
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.1.0.0/16")),
	SecondaryAddressBlocks: []IIpAddresses{
		awsec2alpha.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
			CidrBlockName: jsii.String("AmazonProvided"),
		}),
	},
})

eigw := awsec2alpha.NewEgressOnlyInternetGateway(this, jsii.String("EIGW"), &EgressOnlyInternetGatewayProps{
	Vpc: myVpc,
})

routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})

routeTable.AddRoute(jsii.String("EIGW"), jsii.String("::/0"), map[string]iRouteTarget{
	"gateway": eigw,
})
```

Other route targets may require a deeper set of parameters to set up properly. For instance, the example below illustrates how to set up a `NatGateway`:

```go
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})
subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PRIVATE_ISOLATED,
})

natgw := awsec2alpha.NewNatGateway(this, jsii.String("NatGW"), &NatGatewayProps{
	Subnet: subnet,
	Vpc: myVpc,
	ConnectivityType: awsec2alpha.NatConnectivityType_PRIVATE,
	PrivateIpAddress: jsii.String("10.0.0.42"),
})
awsec2alpha.NewRoute(this, jsii.String("NatGwRoute"), &RouteProps{
	RouteTable: RouteTable,
	Destination: jsii.String("0.0.0.0/0"),
	Target: map[string]iRouteTarget{
		"gateway": natgw,
	},
})
```

It is also possible to set up endpoints connecting other AWS services. For instance, the example below illustrates the linking of a Dynamo DB endpoint via the existing `ec2.GatewayVpcEndpoint` construct as a route target:

```go
stack := awscdk.Newstack()
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})
subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PRIVATE,
})

dynamoEndpoint := ec2.NewGatewayVpcEndpoint(this, jsii.String("DynamoEndpoint"), &GatewayVpcEndpointProps{
	Service: ec2.GatewayVpcEndpointAwsService_DYNAMODB(),
	Vpc: myVpc,
	Subnets: []SubnetSelection{
		subnet,
	},
})
awsec2alpha.NewRoute(this, jsii.String("DynamoDBRoute"), &RouteProps{
	RouteTable: RouteTable,
	Destination: jsii.String("0.0.0.0/0"),
	Target: map[string]iVpcEndpoint{
		"endpoint": dynamoEndpoint,
	},
})
```

## VPC Peering Connection

VPC peering connection allows you to connect two VPCs and route traffic between them using private IP addresses. The VpcV2 construct supports creating VPC peering connections through the `VPCPeeringConnection` construct from the `route` module.

Peering Connection cannot be established between two VPCs with overlapping CIDR ranges. Please make sure the two VPC CIDRs do not overlap with each other else it will throw an error.

For more information, see [What is VPC peering?](https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html).

The following show examples of how to create a peering connection between two VPCs for all possible combinations of same-account or cross-account, and same-region or cross-region configurations.

Note: You cannot create a VPC peering connection between VPCs that have matching or overlapping CIDR blocks

**Case 1: Same Account and Same Region Peering Connection**

```go
stack := awscdk.Newstack()

vpcA := awsec2alpha.NewVpcV2(this, jsii.String("VpcA"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.0.0.0/16")),
})

vpcB := awsec2alpha.NewVpcV2(this, jsii.String("VpcB"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_*Ipv4(jsii.String("10.1.0.0/16")),
})

peeringConnection := vpcA.CreatePeeringConnection(jsii.String("sameAccountSameRegionPeering"), &VPCPeeringConnectionOptions{
	AcceptorVpc: vpcB,
})
```

**Case 2: Same Account and Cross Region Peering Connection**

There is no difference from Case 1 when calling `createPeeringConnection`. The only change is that one of the VPCs are created in another stack with a different region. To establish cross region VPC peering connection, acceptorVpc needs to be imported to the requestor VPC stack using `fromVpcV2Attributes` method.

```go
app := awscdk.NewApp()

stackA := awscdk.Newstack(app, jsii.String("VpcStackA"), &StackProps{
	Env: &Environment{
		Account: jsii.String("000000000000"),
		Region: jsii.String("us-east-1"),
	},
})
stackB := awscdk.Newstack(app, jsii.String("VpcStackB"), &StackProps{
	Env: &Environment{
		Account: jsii.String("000000000000"),
		Region: jsii.String("us-west-2"),
	},
})

vpcA := awsec2alpha.NewVpcV2(stackA, jsii.String("VpcA"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.0.0.0/16")),
})

awsec2alpha.NewVpcV2(stackB, jsii.String("VpcB"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_*Ipv4(jsii.String("10.1.0.0/16")),
})

vpcB := awsec2alpha.VpcV2_FromVpcV2Attributes(stackA, jsii.String("ImportedVpcB"), &VpcV2Attributes{
	VpcId: jsii.String("MockVpcBid"),
	VpcCidrBlock: jsii.String("10.1.0.0/16"),
	Region: jsii.String("us-west-2"),
	OwnerAccountId: jsii.String("000000000000"),
})

peeringConnection := vpcA.CreatePeeringConnection(jsii.String("sameAccountCrossRegionPeering"), &VPCPeeringConnectionOptions{
	AcceptorVpc: vpcB,
})
```

**Case 3: Cross Account Peering Connection**

For cross-account connections, the acceptor account needs an IAM role that grants the requestor account permission to initiate the connection. Create a new IAM role in the acceptor account using method `createAcceptorVpcRole` to provide the necessary permissions.

Once role is created in account, provide role arn for field `peerRoleArn` under method `createPeeringConnection`

```go
stack := awscdk.Newstack()

acceptorVpc := awsec2alpha.NewVpcV2(this, jsii.String("VpcA"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.0.0.0/16")),
})

acceptorRoleArn := acceptorVpc.CreateAcceptorVpcRole(jsii.String("000000000000"))
```

After creating an IAM role in the acceptor account, we can initiate the peering connection request from the requestor VPC. Import acceptorVpc to the stack using `fromVpcV2Attributes` method, it is recommended to specify owner account id of the acceptor VPC in case of cross account peering connection, if acceptor VPC is hosted in different region provide region value for import as well.
The following code snippet demonstrates how to set up VPC peering between two VPCs in different AWS accounts using CDK:

```go
stack := awscdk.Newstack()

acceptorVpc := awsec2alpha.VpcV2_FromVpcV2Attributes(this, jsii.String("acceptorVpc"), &VpcV2Attributes{
	VpcId: jsii.String("vpc-XXXX"),
	VpcCidrBlock: jsii.String("10.0.0.0/16"),
	Region: jsii.String("us-east-2"),
	OwnerAccountId: jsii.String("111111111111"),
})

acceptorRoleArn := "arn:aws:iam::111111111111:role/VpcPeeringRole"

requestorVpc := awsec2alpha.NewVpcV2(this, jsii.String("VpcB"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.1.0.0/16")),
})

peeringConnection := requestorVpc.CreatePeeringConnection(jsii.String("crossAccountCrossRegionPeering"), &VPCPeeringConnectionOptions{
	AcceptorVpc: acceptorVpc,
	PeerRoleArn: acceptorRoleArn,
})
```

### Route Table Configuration

After establishing the VPC peering connection, routes must be added to the respective route tables in the VPCs to enable traffic flow. If a route is added to the requestor stack, information will be able to flow from the requestor VPC to the acceptor VPC, but not in the reverse direction. For bi-directional communication, routes need to be added in both VPCs from their respective stacks.

For more information, see [Update your route tables for a VPC peering connection](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-routing.html).

```go
stack := awscdk.Newstack()

acceptorVpc := awsec2alpha.NewVpcV2(this, jsii.String("VpcA"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.0.0.0/16")),
})

requestorVpc := awsec2alpha.NewVpcV2(this, jsii.String("VpcB"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_*Ipv4(jsii.String("10.1.0.0/16")),
})

peeringConnection := requestorVpc.CreatePeeringConnection(jsii.String("peeringConnection"), &VPCPeeringConnectionOptions{
	AcceptorVpc: acceptorVpc,
})

routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: requestorVpc,
})

routeTable.AddRoute(jsii.String("vpcPeeringRoute"), jsii.String("10.0.0.0/16"), map[string]iRouteTarget{
	"gateway": peeringConnection,
})
```

This can also be done using AWS CLI. For more information, see [create-route](https://docs.aws.amazon.com/cli/latest/reference/ec2/create-route.html).

```bash
# Add a route to the requestor VPC route table
aws ec2 create-route --route-table-id rtb-requestor --destination-cidr-block 10.0.0.0/16 --vpc-peering-connection-id pcx-xxxxxxxx

# For bi-directional add a route in the acceptor vpc account as well
aws ec2 create-route --route-table-id rtb-acceptor --destination-cidr-block 10.1.0.0/16 --vpc-peering-connection-id pcx-xxxxxxxx
```

### Deleting the Peering Connection

To delete a VPC peering connection, use the following command:

```bash
aws ec2 delete-vpc-peering-connection --vpc-peering-connection-id pcx-xxxxxxxx
```

For more information, see [Delete a VPC peering connection](https://docs.aws.amazon.com/vpc/latest/peering/create-vpc-peering-connection.html#delete-vpc-peering-connection).

## Adding Egress-Only Internet Gateway to VPC

An egress-only internet gateway is a horizontally scaled, redundant, and highly available VPC component that allows outbound communication over IPv6 from instances in your VPC to the internet, and prevents the internet from initiating an IPv6 connection with your instances.

For more information see [Enable outbound IPv6 traffic using an egress-only internet gateway](https://docs.aws.amazon.com/vpc/latest/userguide/egress-only-internet-gateway.html).

VpcV2 supports adding an egress only internet gateway to VPC using the `addEgressOnlyInternetGateway` method.

By default, this method sets up a route to all outbound IPv6 address ranges, unless a specific destination is provided by the user. It can only be configured for IPv6-enabled VPCs.
The `Subnets` parameter accepts a `SubnetFilter`, which can be based on a `SubnetType` in VpcV2. A new route will be added to the route tables of all subnets that match this filter.

```go
stack := awscdk.Newstack()
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.1.0.0/16")),
	SecondaryAddressBlocks: []IIpAddresses{
		awsec2alpha.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
			CidrBlockName: jsii.String("AmazonProvided"),
		}),
	},
})
routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})
subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	Ipv6CidrBlock: awsec2alpha.NewIpCidr(jsii.String("2001:db8:1::/64")),
	SubnetType: awscdk.SubnetType_PRIVATE,
})

myVpc.AddEgressOnlyInternetGateway(&EgressOnlyInternetGatewayOptions{
	Subnets: []SubnetSelection{
		&SubnetSelection{
			SubnetType: awscdk.SubnetType_PRIVATE,
		},
	},
	Destination: jsii.String("::/60"),
})
```

## Adding NATGateway to the VPC

A NAT gateway is a Network Address Translation (NAT) service.You can use a NAT gateway so that instances in a private subnet can connect to services outside your VPC but external services cannot initiate a connection with those instances.

For more information, see [NAT gateway basics](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html).

When you create a NAT gateway, you specify one of the following connectivity types:

**Public – (Default)**: Instances in private subnets can connect to the internet through a public NAT gateway, but cannot receive unsolicited inbound connections from the internet

**Private**: Instances in private subnets can connect to other VPCs or your on-premises network through a private NAT gateway.

To define the NAT gateway connectivity type as `ConnectivityType.Public`, you need to ensure that there is an IGW(Internet Gateway) attached to the subnet's VPC.
Since a NATGW is associated with a particular subnet, providing `subnet` field in the input props is mandatory.

Additionally, you can set up a route in any route table with the target set to the NAT Gateway. The function `addNatGateway` returns a `NATGateway` object that you can reference later.

The code example below provides the definition for adding a NAT gateway to your subnet:

```go
stack := awscdk.Newstack()
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
	Vpc: myVpc,
})
subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PUBLIC,
})

myVpc.AddInternetGateway()
myVpc.AddNatGateway(&NatGatewayOptions{
	Subnet: subnet,
	ConnectivityType: awsec2alpha.NatConnectivityType_PUBLIC,
})
```

## Enable VPNGateway for the VPC

A virtual private gateway is the endpoint on the VPC side of your VPN connection.

For more information, see [What is AWS Site-to-Site VPN?](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html).

VPN route propagation is a feature in Amazon Web Services (AWS) that automatically updates route tables in your Virtual Private Cloud (VPC) with routes learned from a VPN connection.

To enable VPN route propagation, use the `vpnRoutePropagation` property to specify the subnets as an input to the function. VPN route propagation will then be enabled for each subnet with the corresponding route table IDs.

Additionally, you can set up a route in any route table with the target set to the VPN Gateway. The function `enableVpnGatewayV2` returns a `VPNGatewayV2` object that you can reference later.

The code example below provides the definition for setting up a VPN gateway with `vpnRoutePropagation` enabled:

```go
stack := awscdk.Newstack()
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
vpnGateway := myVpc.EnableVpnGatewayV2(&VPNGatewayV2Options{
	VpnRoutePropagation: []SubnetSelection{
		&SubnetSelection{
			SubnetType: awscdk.SubnetType_PUBLIC,
		},
	},
	Type: awscdk.VpnConnectionType_IPSEC_1,
})

routeTable := awsec2alpha.NewRouteTable(stack, jsii.String("routeTable"), &RouteTableProps{
	Vpc: myVpc,
})

awsec2alpha.NewRoute(stack, jsii.String("route"), &RouteProps{
	Destination: jsii.String("172.31.0.0/24"),
	Target: map[string]iRouteTarget{
		"gateway": vpnGateway,
	},
	RouteTable: routeTable,
})
```

## Adding InternetGateway to the VPC

An internet gateway is a horizontally scaled, redundant, and highly available VPC component that allows communication between your VPC and the internet. It supports both IPv4 and IPv6 traffic.

For more information, see [Enable VPC internet access using internet gateways](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-igw-internet-access.html).

You can add an internet gateway to a VPC using `addInternetGateway` method. By default, this method creates a route in all Public Subnets with outbound destination set to `0.0.0.0` for IPv4 and `::0` for IPv6 enabled VPC.
Instead of using the default settings, you can configure a custom destination range by providing an optional input `destination` to the method.
In addition to the custom IP range, you can also choose to filter subnets where default routes should be created.

The code example below shows how to add an internet gateway with a custom outbound destination IP range:

```go
stack := awscdk.Newstack()
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))

subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PUBLIC,
})

myVpc.AddInternetGateway(&InternetGatewayOptions{
	Ipv4Destination: jsii.String("192.168.0.0/16"),
})
```

The following code examples demonstrates how to add an internet gateway with a custom outbound destination IP range for specific subnets:

```go
stack := awscdk.Newstack()
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))

mySubnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PUBLIC,
})

myVpc.AddInternetGateway(&InternetGatewayOptions{
	Ipv4Destination: jsii.String("192.168.0.0/16"),
	Subnets: []SubnetSelection{
		mySubnet,
	},
})
```

```go
stack := awscdk.Newstack()
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))

myVpc.AddInternetGateway(&InternetGatewayOptions{
	Ipv4Destination: jsii.String("192.168.0.0/16"),
	Subnets: []SubnetSelection{
		&SubnetSelection{
			SubnetType: awscdk.SubnetType_PRIVATE_WITH_EGRESS,
		},
	},
})
```

## Importing an existing VPC

You can import an existing VPC and its subnets using the `VpcV2.fromVpcV2Attributes()` method or an individual subnet using `SubnetV2.fromSubnetV2Attributes()` method.

### Importing a VPC

To import an existing VPC, use the `VpcV2.fromVpcV2Attributes()` method. You'll need to provide the VPC ID, primary CIDR block, and information about the subnets. You can import secondary address as well created through IPAM, BYOIP(IPv4) or enabled through Amazon Provided IPv6. You must provide VPC Id and its primary CIDR block for importing it.

If you wish to add a new subnet to imported VPC, new subnet's IP range(IPv4) will be validated against provided secondary and primary address block to confirm that it is within the the range of VPC.

Here's an example of importing a VPC with only the required parameters

```go
stack := awscdk.Newstack()

importedVpc := awsec2alpha.VpcV2_FromVpcV2Attributes(stack, jsii.String("ImportedVpc"), &VpcV2Attributes{
	VpcId: jsii.String("mockVpcID"),
	VpcCidrBlock: jsii.String("10.0.0.0/16"),
})
```

In case of cross account or cross region VPC, its recommended to provide region and ownerAccountId so that these values for the VPC can be used to populate correct arn value for the VPC. If a VPC region and account ID is not provided, then region and account configured in the stack will be used. Furthermore, these fields will be referenced later while setting up VPC peering connection, so its necessary to set these fields to a correct value.

Below is an example of importing a cross region and cross account VPC, VPC arn for this case would be 'arn:aws:ec2:us-west-2:123456789012:vpc/mockVpcID'

```go
stack := awscdk.Newstack()

//Importing a cross account or cross region VPC
importedVpc := awsec2alpha.VpcV2_FromVpcV2Attributes(stack, jsii.String("ImportedVpc"), &VpcV2Attributes{
	VpcId: jsii.String("mockVpcID"),
	VpcCidrBlock: jsii.String("10.0.0.0/16"),
	OwnerAccountId: jsii.String("123456789012"),
	Region: jsii.String("us-west-2"),
})
```

Here's an example of how to import a VPC with multiple CIDR blocks, IPv6 support, and different subnet types:

In this example, we're importing a VPC with:

* A primary CIDR block (10.1.0.0/16)
* One secondary IPv4 CIDR block (10.2.0.0/16)
* Two secondary address using IPAM pool (IPv4 and IPv6)
* VPC has Amazon-provided IPv6 CIDR enabled
* An isolated subnet in us-west-2a
* A public subnet in us-west-2b

```go
stack := awscdk.Newstack()

importedVpc := awsec2alpha.VpcV2_FromVpcV2Attributes(this, jsii.String("ImportedVPC"), &VpcV2Attributes{
	VpcId: jsii.String("vpc-XXX"),
	VpcCidrBlock: jsii.String("10.1.0.0/16"),
	SecondaryCidrBlocks: []VPCCidrBlockattributes{
		&VPCCidrBlockattributes{
			CidrBlock: jsii.String("10.2.0.0/16"),
			CidrBlockName: jsii.String("ImportedBlock1"),
		},
		&VPCCidrBlockattributes{
			Ipv6IpamPoolId: jsii.String("ipam-pool-XXX"),
			Ipv6NetmaskLength: jsii.Number(52),
			CidrBlockName: jsii.String("ImportedIpamIpv6"),
		},
		&VPCCidrBlockattributes{
			Ipv4IpamPoolId: jsii.String("ipam-pool-XXX"),
			Ipv4IpamProvisionedCidrs: []*string{
				jsii.String("10.2.0.0/16"),
			},
			CidrBlockName: jsii.String("ImportedIpamIpv4"),
		},
		&VPCCidrBlockattributes{
			AmazonProvidedIpv6CidrBlock: jsii.Boolean(true),
		},
	},
	Subnets: []SubnetV2Attributes{
		&SubnetV2Attributes{
			SubnetName: jsii.String("IsolatedSubnet2"),
			SubnetId: jsii.String("subnet-03cd773c0fe08ed26"),
			SubnetType: awscdk.SubnetType_PRIVATE_ISOLATED,
			AvailabilityZone: jsii.String("us-west-2a"),
			Ipv4CidrBlock: jsii.String("10.2.0.0/24"),
			RouteTableId: jsii.String("rtb-0871c310f98da2cbb"),
		},
		&SubnetV2Attributes{
			SubnetId: jsii.String("subnet-0fa477e01db27d820"),
			SubnetType: awscdk.SubnetType_PUBLIC,
			AvailabilityZone: jsii.String("us-west-2b"),
			Ipv4CidrBlock: jsii.String("10.3.0.0/24"),
			RouteTableId: jsii.String("rtb-014f3043098fe4b96"),
		},
	},
})

// You can now use the imported VPC in your stack

// Adding a new subnet to the imported VPC
importedSubnet := awsec2alpha.NewSubnetV2(this, jsii.String("NewSubnet"), &SubnetV2Props{
	AvailabilityZone: jsii.String("us-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.2.2.0/24")),
	Vpc: importedVpc,
	SubnetType: awscdk.SubnetType_PUBLIC,
})

// Adding gateways to the imported VPC
importedVpc.AddInternetGateway()
importedVpc.AddNatGateway(&NatGatewayOptions{
	Subnet: importedSubnet,
})
importedVpc.AddEgressOnlyInternetGateway()
```

You can add more subnets as needed by including additional entries in the `isolatedSubnets`, `publicSubnets`, or other subnet type arrays (e.g., `privateSubnets`).

### Importing Subnets

You can also import individual subnets using the `SubnetV2.fromSubnetV2Attributes()` method. This is useful when you need to work with specific subnets independently of a VPC.

Here's an example of how to import a subnet:

```go
awsec2alpha.SubnetV2_FromSubnetV2Attributes(this, jsii.String("ImportedSubnet"), &SubnetV2Attributes{
	SubnetId: jsii.String("subnet-0123456789abcdef0"),
	AvailabilityZone: jsii.String("us-west-2a"),
	Ipv4CidrBlock: jsii.String("10.2.0.0/24"),
	RouteTableId: jsii.String("rtb-0871c310f98da2cbb"),
	SubnetType: awscdk.SubnetType_PRIVATE_ISOLATED,
})
```

By importing existing VPCs and subnets, you can easily integrate your existing AWS infrastructure with new resources created through CDK. This is particularly useful when you need to work with pre-existing network configurations or when you're migrating existing infrastructure to CDK.

### Tagging VPC and its components

By default, when a resource name is given to the construct, it automatically adds a tag with the key `Name` and the value set to the provided resource name. To add additional custom tags, use the Tag Manager, like this: `Tags.of(myConstruct).add('key', 'value');`.

For example, if the `vpcName` is set to `TestVpc`, the following code will add a tag to the VPC with `key: Name` and `value: TestVpc`.

```go
vpc := awsec2alpha.NewVpcV2(this, jsii.String("VPC-integ-test-tag"), &VpcV2Props{
	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.1.0.0/16")),
	EnableDnsHostnames: jsii.Boolean(true),
	EnableDnsSupport: jsii.Boolean(true),
	VpcName: jsii.String("CDKintegTestVPC"),
})

// Add custom tags if needed
awscdk.Tags_Of(vpc).Add(jsii.String("Environment"), jsii.String("Production"))
```

## Transit Gateway

The AWS Transit Gateway construct library allows you to create and configure Transit Gateway resources using AWS CDK.

See [AWS Transit Gateway Docs](docs.aws.amazon.com/vpc/latest/tgw/what-is-transit-gateway.html) for more info.

### Overview

The Transit Gateway construct (`TransitGateway`) is the main entry point for creating and managing your Transit Gateway infrastructure. It provides methods to create route tables, attach VPCs, and configure cross-account access.

The Transit Gateway construct library provides four main constructs:

* `TransitGateway`: The central hub for your network connections
* `TransitGatewayRouteTable`: Manages routing between attached networks
* `TransitGatewayVpcAttachment`: Connects VPCs to the Transit Gateway
* `TransitGatewayRoute`: Defines routing rules within your Transit Gateway

### Basic Usage

To create a minimal deployable `TransitGateway`:

```go
transitGateway := awsec2alpha.NewTransitGateway(this, jsii.String("MyTransitGateway"))
```

### Default Transit Gateway Route Table

By default, `TransitGateway` is created with a default `TransitGatewayRouteTable`, for which automatic Associations and automatic Propagations are enabled.

> Note: When you create a default Transit Gateway in AWS Console, a default Transit Gateway Route Table is automatically created by AWS. However, when using the CDK Transit Gateway L2 construct, the underlying L1 construct is configured with `defaultRouteTableAssociation` and `defaultRouteTablePropagation` explicitly disabled. This ensures that AWS does not create the default route table, allowing the CDK to define a custom default route table instead.
>
> As a result, in the AWS Console, the **Default association route table** and **Default propagation route table** settings will appear as disabled. Despite this, the CDK still provides automatic association and propagation functionality through its internal implementation, which can be controlled using the `defaultRouteTableAssociation` and `defaultRouteTablePropagation` properties within the CDK.

You can disable the automatic Association/Propagation on the default `TransitGatewayRouteTable` via the `TransitGateway` properties. This will still create a default route table for you:

```go
transitGateway := awsec2alpha.NewTransitGateway(this, jsii.String("MyTransitGateway"), &TransitGatewayProps{
	DefaultRouteTableAssociation: jsii.Boolean(false),
	DefaultRouteTablePropagation: jsii.Boolean(false),
})
```

### Transit Gateway Route Table Management

Add additional Transit Gateway Route Tables using the `addRouteTable()` method:

```go
transitGateway := awsec2alpha.NewTransitGateway(this, jsii.String("MyTransitGateway"))

routeTable := transitGateway.addRouteTable(jsii.String("CustomRouteTable"))
```

### Attaching VPCs to the Transit Gateway

Currently only VPC to Transit Gateway attachments are supported.

Create an attachment from a VPC to the Transit Gateway using the `attachVpc()` method:

```go
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
subnet1 := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PUBLIC,
})

subnet2 := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.1.0/24")),
	SubnetType: awscdk.SubnetType_PUBLIC,
})

transitGateway := awsec2alpha.NewTransitGateway(this, jsii.String("MyTransitGateway"))

// Create a basic attachment
attachment := transitGateway.attachVpc(jsii.String("VpcAttachment"), &AttachVpcOptions{
	Vpc: myVpc,
	Subnets: []ISubnet{
		subnet1,
		subnet2,
	},
})

// Create an attachment with optional parameters
attachmentWithOptions := transitGateway.attachVpc(jsii.String("VpcAttachmentWithOptions"), &AttachVpcOptions{
	Vpc: myVpc,
	Subnets: []ISubnet{
		subnet1,
	},
	VpcAttachmentOptions: map[string]*bool{
		"dnsSupport": jsii.Boolean(true),
		"applianceModeSupport": jsii.Boolean(true),
		"ipv6Support": jsii.Boolean(true),
		"securityGroupReferencingSupport": jsii.Boolean(true),
	},
})
```

If you want to automatically associate and propagate routes with transit gateway route tables, you can pass the `associationRouteTable` and `propagationRouteTables` parameters. This will automatically create the necessary associations and propagations based on the provided route tables.

```go
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
subnet1 := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PUBLIC,
})

subnet2 := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.1.0/24")),
	SubnetType: awscdk.SubnetType_PUBLIC,
})

transitGateway := awsec2alpha.NewTransitGateway(this, jsii.String("MyTransitGateway"))
associationRouteTable := transitGateway.addRouteTable(jsii.String("AssociationRouteTable"))
propagationRouteTable1 := transitGateway.addRouteTable(jsii.String("PropagationRouteTable1"))
propagationRouteTable2 := transitGateway.addRouteTable(jsii.String("PropagationRouteTable2"))

// Create an attachment with automatically created association + propagations
attachmentWithRoutes := transitGateway.attachVpc(jsii.String("VpcAttachment"), &AttachVpcOptions{
	Vpc: myVpc,
	Subnets: []ISubnet{
		subnet1,
		subnet2,
	},
	AssociationRouteTable: associationRouteTable,
	PropagationRouteTables: []ITransitGatewayRouteTable{
		propagationRouteTable1,
		propagationRouteTable2,
	},
})
```

In this example, the `associationRouteTable` is set to `associationRouteTable`, and `propagationRouteTables` is set to an array containing `propagationRouteTable1` and `propagationRouteTable2`. This triggers the automatic creation of route table associations and route propagations between the Transit Gateway and the specified route tables.

### Adding static routes to the route table

Add static routes using either the `addRoute()` method to add an active route or `addBlackholeRoute()` to add a blackhole route:

```go
transitGateway := awsec2alpha.NewTransitGateway(this, jsii.String("MyTransitGateway"))
routeTable := transitGateway.addRouteTable(jsii.String("CustomRouteTable"))

myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PUBLIC,
})

attachment := transitGateway.attachVpc(jsii.String("VpcAttachment"), &AttachVpcOptions{
	Vpc: myVpc,
	Subnets: []ISubnet{
		subnet,
	},
})

// Add a static route to direct traffic
routeTable.AddRoute(jsii.String("StaticRoute"), attachment, jsii.String("10.0.0.0/16"))

// Block unwanted traffic with a blackhole route
routeTable.AddBlackholeRoute(jsii.String("BlackholeRoute"), jsii.String("172.16.0.0/16"))
```

### Route Table Associations and Propagations

Configure route table associations and enable route propagation:

```go
transitGateway := awsec2alpha.NewTransitGateway(this, jsii.String("MyTransitGateway"))
routeTable := transitGateway.addRouteTable(jsii.String("CustomRouteTable"))
myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
	Vpc: myVpc,
	AvailabilityZone: jsii.String("eu-west-2a"),
	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
	SubnetType: awscdk.SubnetType_PUBLIC,
})
attachment := transitGateway.attachVpc(jsii.String("VpcAttachment"), &AttachVpcOptions{
	Vpc: myVpc,
	Subnets: []ISubnet{
		subnet,
	},
})

// Associate an attachment with a route table
routeTable.AddAssociation(jsii.String("Association"), attachment)

// Enable route propagation for an attachment
routeTable.EnablePropagation(jsii.String("Propagation"), attachment)
```

**Associations** — The linking of a Transit Gateway attachment to a specific route table, which determines which routes that attachment will use for routing decisions.

**Propagation** — The automatic advertisement of routes from an attachment to a route table, allowing the route table to learn about available network destinations.

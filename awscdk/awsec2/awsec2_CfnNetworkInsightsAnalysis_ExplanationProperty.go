package awsec2


// Describes an explanation code for an unreachable path.
//
// For more information, see [Reachability Analyzer explanation codes](https://docs.aws.amazon.com/vpc/latest/reachability/explanation-codes.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   explanationProperty := &explanationProperty{
//   	acl: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	aclRule: &analysisAclRuleProperty{
//   		cidr: jsii.String("cidr"),
//   		egress: jsii.Boolean(false),
//   		portRange: &portRangeProperty{
//   			from: jsii.Number(123),
//   			to: jsii.Number(123),
//   		},
//   		protocol: jsii.String("protocol"),
//   		ruleAction: jsii.String("ruleAction"),
//   		ruleNumber: jsii.Number(123),
//   	},
//   	address: jsii.String("address"),
//   	addresses: []*string{
//   		jsii.String("addresses"),
//   	},
//   	attachedTo: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   	classicLoadBalancerListener: &analysisLoadBalancerListenerProperty{
//   		instancePort: jsii.Number(123),
//   		loadBalancerPort: jsii.Number(123),
//   	},
//   	component: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	componentAccount: jsii.String("componentAccount"),
//   	componentRegion: jsii.String("componentRegion"),
//   	customerGateway: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	destination: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	destinationVpc: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	direction: jsii.String("direction"),
//   	elasticLoadBalancerListener: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	explanationCode: jsii.String("explanationCode"),
//   	ingressRouteTable: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	internetGateway: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	loadBalancerArn: jsii.String("loadBalancerArn"),
//   	loadBalancerListenerPort: jsii.Number(123),
//   	loadBalancerTarget: &analysisLoadBalancerTargetProperty{
//   		address: jsii.String("address"),
//   		availabilityZone: jsii.String("availabilityZone"),
//   		instance: &analysisComponentProperty{
//   			arn: jsii.String("arn"),
//   			id: jsii.String("id"),
//   		},
//   		port: jsii.Number(123),
//   	},
//   	loadBalancerTargetGroup: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	loadBalancerTargetGroups: []interface{}{
//   		&analysisComponentProperty{
//   			arn: jsii.String("arn"),
//   			id: jsii.String("id"),
//   		},
//   	},
//   	loadBalancerTargetPort: jsii.Number(123),
//   	missingComponent: jsii.String("missingComponent"),
//   	natGateway: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	networkInterface: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	packetField: jsii.String("packetField"),
//   	port: jsii.Number(123),
//   	portRanges: []interface{}{
//   		&portRangeProperty{
//   			from: jsii.Number(123),
//   			to: jsii.Number(123),
//   		},
//   	},
//   	prefixList: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	routeTable: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	routeTableRoute: &analysisRouteTableRouteProperty{
//   		destinationCidr: jsii.String("destinationCidr"),
//   		destinationPrefixListId: jsii.String("destinationPrefixListId"),
//   		egressOnlyInternetGatewayId: jsii.String("egressOnlyInternetGatewayId"),
//   		gatewayId: jsii.String("gatewayId"),
//   		instanceId: jsii.String("instanceId"),
//   		natGatewayId: jsii.String("natGatewayId"),
//   		networkInterfaceId: jsii.String("networkInterfaceId"),
//   		origin: jsii.String("origin"),
//   		state: jsii.String("state"),
//   		transitGatewayId: jsii.String("transitGatewayId"),
//   		vpcPeeringConnectionId: jsii.String("vpcPeeringConnectionId"),
//   	},
//   	securityGroup: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	securityGroupRule: &analysisSecurityGroupRuleProperty{
//   		cidr: jsii.String("cidr"),
//   		direction: jsii.String("direction"),
//   		portRange: &portRangeProperty{
//   			from: jsii.Number(123),
//   			to: jsii.Number(123),
//   		},
//   		prefixListId: jsii.String("prefixListId"),
//   		protocol: jsii.String("protocol"),
//   		securityGroupId: jsii.String("securityGroupId"),
//   	},
//   	securityGroups: []interface{}{
//   		&analysisComponentProperty{
//   			arn: jsii.String("arn"),
//   			id: jsii.String("id"),
//   		},
//   	},
//   	sourceVpc: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	state: jsii.String("state"),
//   	subnet: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	subnetRouteTable: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	transitGateway: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	transitGatewayAttachment: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	transitGatewayRouteTable: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	transitGatewayRouteTableRoute: &transitGatewayRouteTableRouteProperty{
//   		attachmentId: jsii.String("attachmentId"),
//   		destinationCidr: jsii.String("destinationCidr"),
//   		prefixListId: jsii.String("prefixListId"),
//   		resourceId: jsii.String("resourceId"),
//   		resourceType: jsii.String("resourceType"),
//   		routeOrigin: jsii.String("routeOrigin"),
//   		state: jsii.String("state"),
//   	},
//   	vpc: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	vpcEndpoint: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	vpcPeeringConnection: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	vpnConnection: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	vpnGateway: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   }
//
type CfnNetworkInsightsAnalysis_ExplanationProperty struct {
	// The network ACL.
	Acl interface{} `field:"optional" json:"acl" yaml:"acl"`
	// The network ACL rule.
	AclRule interface{} `field:"optional" json:"aclRule" yaml:"aclRule"`
	// The IPv4 address, in CIDR notation.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The IPv4 addresses, in CIDR notation.
	Addresses *[]*string `field:"optional" json:"addresses" yaml:"addresses"`
	// The resource to which the component is attached.
	AttachedTo interface{} `field:"optional" json:"attachedTo" yaml:"attachedTo"`
	// The Availability Zones.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The CIDR ranges.
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	// The listener for a Classic Load Balancer.
	ClassicLoadBalancerListener interface{} `field:"optional" json:"classicLoadBalancerListener" yaml:"classicLoadBalancerListener"`
	// The component.
	Component interface{} `field:"optional" json:"component" yaml:"component"`
	// `CfnNetworkInsightsAnalysis.ExplanationProperty.ComponentAccount`.
	ComponentAccount *string `field:"optional" json:"componentAccount" yaml:"componentAccount"`
	// `CfnNetworkInsightsAnalysis.ExplanationProperty.ComponentRegion`.
	ComponentRegion *string `field:"optional" json:"componentRegion" yaml:"componentRegion"`
	// The customer gateway.
	CustomerGateway interface{} `field:"optional" json:"customerGateway" yaml:"customerGateway"`
	// The destination.
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The destination VPC.
	DestinationVpc interface{} `field:"optional" json:"destinationVpc" yaml:"destinationVpc"`
	// The direction. The following are the possible values:.
	//
	// - egress
	// - ingress.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// The load balancer listener.
	ElasticLoadBalancerListener interface{} `field:"optional" json:"elasticLoadBalancerListener" yaml:"elasticLoadBalancerListener"`
	// The explanation code.
	ExplanationCode *string `field:"optional" json:"explanationCode" yaml:"explanationCode"`
	// The route table.
	IngressRouteTable interface{} `field:"optional" json:"ingressRouteTable" yaml:"ingressRouteTable"`
	// The internet gateway.
	InternetGateway interface{} `field:"optional" json:"internetGateway" yaml:"internetGateway"`
	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The listener port of the load balancer.
	LoadBalancerListenerPort *float64 `field:"optional" json:"loadBalancerListenerPort" yaml:"loadBalancerListenerPort"`
	// The target.
	LoadBalancerTarget interface{} `field:"optional" json:"loadBalancerTarget" yaml:"loadBalancerTarget"`
	// The target group.
	LoadBalancerTargetGroup interface{} `field:"optional" json:"loadBalancerTargetGroup" yaml:"loadBalancerTargetGroup"`
	// The target groups.
	LoadBalancerTargetGroups interface{} `field:"optional" json:"loadBalancerTargetGroups" yaml:"loadBalancerTargetGroups"`
	// The target port.
	LoadBalancerTargetPort *float64 `field:"optional" json:"loadBalancerTargetPort" yaml:"loadBalancerTargetPort"`
	// The missing component.
	MissingComponent *string `field:"optional" json:"missingComponent" yaml:"missingComponent"`
	// The NAT gateway.
	NatGateway interface{} `field:"optional" json:"natGateway" yaml:"natGateway"`
	// The network interface.
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// The packet field.
	PacketField *string `field:"optional" json:"packetField" yaml:"packetField"`
	// The port.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The port ranges.
	PortRanges interface{} `field:"optional" json:"portRanges" yaml:"portRanges"`
	// The prefix list.
	PrefixList interface{} `field:"optional" json:"prefixList" yaml:"prefixList"`
	// The protocols.
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// The route table.
	RouteTable interface{} `field:"optional" json:"routeTable" yaml:"routeTable"`
	// The route table route.
	RouteTableRoute interface{} `field:"optional" json:"routeTableRoute" yaml:"routeTableRoute"`
	// The security group.
	SecurityGroup interface{} `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The security group rule.
	SecurityGroupRule interface{} `field:"optional" json:"securityGroupRule" yaml:"securityGroupRule"`
	// The security groups.
	SecurityGroups interface{} `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The source VPC.
	SourceVpc interface{} `field:"optional" json:"sourceVpc" yaml:"sourceVpc"`
	// The state.
	State *string `field:"optional" json:"state" yaml:"state"`
	// The subnet.
	Subnet interface{} `field:"optional" json:"subnet" yaml:"subnet"`
	// The route table for the subnet.
	SubnetRouteTable interface{} `field:"optional" json:"subnetRouteTable" yaml:"subnetRouteTable"`
	// The transit gateway.
	TransitGateway interface{} `field:"optional" json:"transitGateway" yaml:"transitGateway"`
	// The transit gateway attachment.
	TransitGatewayAttachment interface{} `field:"optional" json:"transitGatewayAttachment" yaml:"transitGatewayAttachment"`
	// The transit gateway route table.
	TransitGatewayRouteTable interface{} `field:"optional" json:"transitGatewayRouteTable" yaml:"transitGatewayRouteTable"`
	// The transit gateway route table route.
	TransitGatewayRouteTableRoute interface{} `field:"optional" json:"transitGatewayRouteTableRoute" yaml:"transitGatewayRouteTableRoute"`
	// The component VPC.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
	// The VPC endpoint.
	VpcEndpoint interface{} `field:"optional" json:"vpcEndpoint" yaml:"vpcEndpoint"`
	// The VPC peering connection.
	VpcPeeringConnection interface{} `field:"optional" json:"vpcPeeringConnection" yaml:"vpcPeeringConnection"`
	// The VPN connection.
	VpnConnection interface{} `field:"optional" json:"vpnConnection" yaml:"vpnConnection"`
	// The VPN gateway.
	VpnGateway interface{} `field:"optional" json:"vpnGateway" yaml:"vpnGateway"`
}


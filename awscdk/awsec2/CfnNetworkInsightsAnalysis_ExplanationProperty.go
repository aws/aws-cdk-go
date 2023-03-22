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
//   explanationProperty := &ExplanationProperty{
//   	Acl: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	AclRule: &AnalysisAclRuleProperty{
//   		Cidr: jsii.String("cidr"),
//   		Egress: jsii.Boolean(false),
//   		PortRange: &PortRangeProperty{
//   			From: jsii.Number(123),
//   			To: jsii.Number(123),
//   		},
//   		Protocol: jsii.String("protocol"),
//   		RuleAction: jsii.String("ruleAction"),
//   		RuleNumber: jsii.Number(123),
//   	},
//   	Address: jsii.String("address"),
//   	Addresses: []*string{
//   		jsii.String("addresses"),
//   	},
//   	AttachedTo: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	Cidrs: []*string{
//   		jsii.String("cidrs"),
//   	},
//   	ClassicLoadBalancerListener: &AnalysisLoadBalancerListenerProperty{
//   		InstancePort: jsii.Number(123),
//   		LoadBalancerPort: jsii.Number(123),
//   	},
//   	Component: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	ComponentAccount: jsii.String("componentAccount"),
//   	ComponentRegion: jsii.String("componentRegion"),
//   	CustomerGateway: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	Destination: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	DestinationVpc: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	Direction: jsii.String("direction"),
//   	ElasticLoadBalancerListener: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	ExplanationCode: jsii.String("explanationCode"),
//   	IngressRouteTable: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	InternetGateway: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	LoadBalancerArn: jsii.String("loadBalancerArn"),
//   	LoadBalancerListenerPort: jsii.Number(123),
//   	LoadBalancerTarget: &AnalysisLoadBalancerTargetProperty{
//   		Address: jsii.String("address"),
//   		AvailabilityZone: jsii.String("availabilityZone"),
//   		Instance: &AnalysisComponentProperty{
//   			Arn: jsii.String("arn"),
//   			Id: jsii.String("id"),
//   		},
//   		Port: jsii.Number(123),
//   	},
//   	LoadBalancerTargetGroup: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	LoadBalancerTargetGroups: []interface{}{
//   		&AnalysisComponentProperty{
//   			Arn: jsii.String("arn"),
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	LoadBalancerTargetPort: jsii.Number(123),
//   	MissingComponent: jsii.String("missingComponent"),
//   	NatGateway: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	NetworkInterface: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	PacketField: jsii.String("packetField"),
//   	Port: jsii.Number(123),
//   	PortRanges: []interface{}{
//   		&PortRangeProperty{
//   			From: jsii.Number(123),
//   			To: jsii.Number(123),
//   		},
//   	},
//   	PrefixList: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	Protocols: []*string{
//   		jsii.String("protocols"),
//   	},
//   	RouteTable: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	RouteTableRoute: &AnalysisRouteTableRouteProperty{
//   		DestinationCidr: jsii.String("destinationCidr"),
//   		DestinationPrefixListId: jsii.String("destinationPrefixListId"),
//   		EgressOnlyInternetGatewayId: jsii.String("egressOnlyInternetGatewayId"),
//   		GatewayId: jsii.String("gatewayId"),
//   		InstanceId: jsii.String("instanceId"),
//   		NatGatewayId: jsii.String("natGatewayId"),
//   		NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   		Origin: jsii.String("origin"),
//   		State: jsii.String("state"),
//   		TransitGatewayId: jsii.String("transitGatewayId"),
//   		VpcPeeringConnectionId: jsii.String("vpcPeeringConnectionId"),
//   	},
//   	SecurityGroup: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	SecurityGroupRule: &AnalysisSecurityGroupRuleProperty{
//   		Cidr: jsii.String("cidr"),
//   		Direction: jsii.String("direction"),
//   		PortRange: &PortRangeProperty{
//   			From: jsii.Number(123),
//   			To: jsii.Number(123),
//   		},
//   		PrefixListId: jsii.String("prefixListId"),
//   		Protocol: jsii.String("protocol"),
//   		SecurityGroupId: jsii.String("securityGroupId"),
//   	},
//   	SecurityGroups: []interface{}{
//   		&AnalysisComponentProperty{
//   			Arn: jsii.String("arn"),
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	SourceVpc: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	State: jsii.String("state"),
//   	Subnet: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	SubnetRouteTable: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	TransitGateway: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	TransitGatewayAttachment: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	TransitGatewayRouteTable: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	TransitGatewayRouteTableRoute: &TransitGatewayRouteTableRouteProperty{
//   		AttachmentId: jsii.String("attachmentId"),
//   		DestinationCidr: jsii.String("destinationCidr"),
//   		PrefixListId: jsii.String("prefixListId"),
//   		ResourceId: jsii.String("resourceId"),
//   		ResourceType: jsii.String("resourceType"),
//   		RouteOrigin: jsii.String("routeOrigin"),
//   		State: jsii.String("state"),
//   	},
//   	Vpc: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	VpcEndpoint: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	VpcPeeringConnection: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	VpnConnection: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	VpnGateway: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
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
	// The AWS account for the component.
	ComponentAccount *string `field:"optional" json:"componentAccount" yaml:"componentAccount"`
	// The Region for the component.
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


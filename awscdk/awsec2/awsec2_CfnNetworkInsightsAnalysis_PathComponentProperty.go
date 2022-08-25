package awsec2


// Describes a path component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathComponentProperty := &pathComponentProperty{
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
//   	additionalDetails: []interface{}{
//   		&additionalDetailProperty{
//   			additionalDetailType: jsii.String("additionalDetailType"),
//   			component: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   		},
//   	},
//   	component: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	destinationVpc: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	elasticLoadBalancerListener: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	explanations: []interface{}{
//   		&explanationProperty{
//   			acl: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			aclRule: &analysisAclRuleProperty{
//   				cidr: jsii.String("cidr"),
//   				egress: jsii.Boolean(false),
//   				portRange: &portRangeProperty{
//   					from: jsii.Number(123),
//   					to: jsii.Number(123),
//   				},
//   				protocol: jsii.String("protocol"),
//   				ruleAction: jsii.String("ruleAction"),
//   				ruleNumber: jsii.Number(123),
//   			},
//   			address: jsii.String("address"),
//   			addresses: []*string{
//   				jsii.String("addresses"),
//   			},
//   			attachedTo: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			availabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   			classicLoadBalancerListener: &analysisLoadBalancerListenerProperty{
//   				instancePort: jsii.Number(123),
//   				loadBalancerPort: jsii.Number(123),
//   			},
//   			component: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			componentAccount: jsii.String("componentAccount"),
//   			componentRegion: jsii.String("componentRegion"),
//   			customerGateway: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			destination: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			destinationVpc: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			direction: jsii.String("direction"),
//   			elasticLoadBalancerListener: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			explanationCode: jsii.String("explanationCode"),
//   			ingressRouteTable: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			internetGateway: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			loadBalancerArn: jsii.String("loadBalancerArn"),
//   			loadBalancerListenerPort: jsii.Number(123),
//   			loadBalancerTarget: &analysisLoadBalancerTargetProperty{
//   				address: jsii.String("address"),
//   				availabilityZone: jsii.String("availabilityZone"),
//   				instance: &analysisComponentProperty{
//   					arn: jsii.String("arn"),
//   					id: jsii.String("id"),
//   				},
//   				port: jsii.Number(123),
//   			},
//   			loadBalancerTargetGroup: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			loadBalancerTargetGroups: []interface{}{
//   				&analysisComponentProperty{
//   					arn: jsii.String("arn"),
//   					id: jsii.String("id"),
//   				},
//   			},
//   			loadBalancerTargetPort: jsii.Number(123),
//   			missingComponent: jsii.String("missingComponent"),
//   			natGateway: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			networkInterface: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			packetField: jsii.String("packetField"),
//   			port: jsii.Number(123),
//   			portRanges: []interface{}{
//   				&portRangeProperty{
//   					from: jsii.Number(123),
//   					to: jsii.Number(123),
//   				},
//   			},
//   			prefixList: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			protocols: []*string{
//   				jsii.String("protocols"),
//   			},
//   			routeTable: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			routeTableRoute: &analysisRouteTableRouteProperty{
//   				destinationCidr: jsii.String("destinationCidr"),
//   				destinationPrefixListId: jsii.String("destinationPrefixListId"),
//   				egressOnlyInternetGatewayId: jsii.String("egressOnlyInternetGatewayId"),
//   				gatewayId: jsii.String("gatewayId"),
//   				instanceId: jsii.String("instanceId"),
//   				natGatewayId: jsii.String("natGatewayId"),
//   				networkInterfaceId: jsii.String("networkInterfaceId"),
//   				origin: jsii.String("origin"),
//   				state: jsii.String("state"),
//   				transitGatewayId: jsii.String("transitGatewayId"),
//   				vpcPeeringConnectionId: jsii.String("vpcPeeringConnectionId"),
//   			},
//   			securityGroup: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			securityGroupRule: &analysisSecurityGroupRuleProperty{
//   				cidr: jsii.String("cidr"),
//   				direction: jsii.String("direction"),
//   				portRange: &portRangeProperty{
//   					from: jsii.Number(123),
//   					to: jsii.Number(123),
//   				},
//   				prefixListId: jsii.String("prefixListId"),
//   				protocol: jsii.String("protocol"),
//   				securityGroupId: jsii.String("securityGroupId"),
//   			},
//   			securityGroups: []interface{}{
//   				&analysisComponentProperty{
//   					arn: jsii.String("arn"),
//   					id: jsii.String("id"),
//   				},
//   			},
//   			sourceVpc: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			state: jsii.String("state"),
//   			subnet: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			subnetRouteTable: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			transitGateway: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			transitGatewayAttachment: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			transitGatewayRouteTable: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			transitGatewayRouteTableRoute: &transitGatewayRouteTableRouteProperty{
//   				attachmentId: jsii.String("attachmentId"),
//   				destinationCidr: jsii.String("destinationCidr"),
//   				prefixListId: jsii.String("prefixListId"),
//   				resourceId: jsii.String("resourceId"),
//   				resourceType: jsii.String("resourceType"),
//   				routeOrigin: jsii.String("routeOrigin"),
//   				state: jsii.String("state"),
//   			},
//   			vpc: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			vpcEndpoint: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			vpcPeeringConnection: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			vpnConnection: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   			vpnGateway: &analysisComponentProperty{
//   				arn: jsii.String("arn"),
//   				id: jsii.String("id"),
//   			},
//   		},
//   	},
//   	inboundHeader: &analysisPacketHeaderProperty{
//   		destinationAddresses: []*string{
//   			jsii.String("destinationAddresses"),
//   		},
//   		destinationPortRanges: []interface{}{
//   			&portRangeProperty{
//   				from: jsii.Number(123),
//   				to: jsii.Number(123),
//   			},
//   		},
//   		protocol: jsii.String("protocol"),
//   		sourceAddresses: []*string{
//   			jsii.String("sourceAddresses"),
//   		},
//   		sourcePortRanges: []interface{}{
//   			&portRangeProperty{
//   				from: jsii.Number(123),
//   				to: jsii.Number(123),
//   			},
//   		},
//   	},
//   	outboundHeader: &analysisPacketHeaderProperty{
//   		destinationAddresses: []*string{
//   			jsii.String("destinationAddresses"),
//   		},
//   		destinationPortRanges: []interface{}{
//   			&portRangeProperty{
//   				from: jsii.Number(123),
//   				to: jsii.Number(123),
//   			},
//   		},
//   		protocol: jsii.String("protocol"),
//   		sourceAddresses: []*string{
//   			jsii.String("sourceAddresses"),
//   		},
//   		sourcePortRanges: []interface{}{
//   			&portRangeProperty{
//   				from: jsii.Number(123),
//   				to: jsii.Number(123),
//   			},
//   		},
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
//   	sequenceNumber: jsii.Number(123),
//   	sourceVpc: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	subnet: &analysisComponentProperty{
//   		arn: jsii.String("arn"),
//   		id: jsii.String("id"),
//   	},
//   	transitGateway: &analysisComponentProperty{
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
//   }
//
type CfnNetworkInsightsAnalysis_PathComponentProperty struct {
	// The network ACL rule.
	AclRule interface{} `field:"optional" json:"aclRule" yaml:"aclRule"`
	// `CfnNetworkInsightsAnalysis.PathComponentProperty.AdditionalDetails`.
	AdditionalDetails interface{} `field:"optional" json:"additionalDetails" yaml:"additionalDetails"`
	// The component.
	Component interface{} `field:"optional" json:"component" yaml:"component"`
	// The destination VPC.
	DestinationVpc interface{} `field:"optional" json:"destinationVpc" yaml:"destinationVpc"`
	// `CfnNetworkInsightsAnalysis.PathComponentProperty.ElasticLoadBalancerListener`.
	ElasticLoadBalancerListener interface{} `field:"optional" json:"elasticLoadBalancerListener" yaml:"elasticLoadBalancerListener"`
	// `CfnNetworkInsightsAnalysis.PathComponentProperty.Explanations`.
	Explanations interface{} `field:"optional" json:"explanations" yaml:"explanations"`
	// The inbound header.
	InboundHeader interface{} `field:"optional" json:"inboundHeader" yaml:"inboundHeader"`
	// The outbound header.
	OutboundHeader interface{} `field:"optional" json:"outboundHeader" yaml:"outboundHeader"`
	// The route table route.
	RouteTableRoute interface{} `field:"optional" json:"routeTableRoute" yaml:"routeTableRoute"`
	// The security group rule.
	SecurityGroupRule interface{} `field:"optional" json:"securityGroupRule" yaml:"securityGroupRule"`
	// The sequence number.
	SequenceNumber *float64 `field:"optional" json:"sequenceNumber" yaml:"sequenceNumber"`
	// The source VPC.
	SourceVpc interface{} `field:"optional" json:"sourceVpc" yaml:"sourceVpc"`
	// The subnet.
	Subnet interface{} `field:"optional" json:"subnet" yaml:"subnet"`
	// `CfnNetworkInsightsAnalysis.PathComponentProperty.TransitGateway`.
	TransitGateway interface{} `field:"optional" json:"transitGateway" yaml:"transitGateway"`
	// The route in a transit gateway route table.
	TransitGatewayRouteTableRoute interface{} `field:"optional" json:"transitGatewayRouteTableRoute" yaml:"transitGatewayRouteTableRoute"`
	// The component VPC.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}


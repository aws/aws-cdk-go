package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLoadBalancer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLoadBalancerProps := &cfnLoadBalancerProps{
//   	ipAddressType: jsii.String("ipAddressType"),
//   	loadBalancerAttributes: []interface{}{
//   		&loadBalancerAttributeProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	scheme: jsii.String("scheme"),
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	subnetMappings: []interface{}{
//   		&subnetMappingProperty{
//   			subnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			allocationId: jsii.String("allocationId"),
//   			iPv6Address: jsii.String("iPv6Address"),
//   			privateIPv4Address: jsii.String("privateIPv4Address"),
//   		},
//   	},
//   	subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnLoadBalancerProps struct {
	// The IP address type.
	//
	// The possible values are `ipv4` (for IPv4 addresses) and `dualstack` (for IPv4 and IPv6 addresses). You can’t specify `dualstack` for a load balancer with a UDP or TCP_UDP listener.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The load balancer attributes.
	LoadBalancerAttributes interface{} `field:"optional" json:"loadBalancerAttributes" yaml:"loadBalancerAttributes"`
	// The name of the load balancer.
	//
	// This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with "internal-".
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The nodes of an Internet-facing load balancer have public IP addresses.
	//
	// The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.
	//
	// The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.
	//
	// The default is an Internet-facing load balancer.
	//
	// You cannot specify a scheme for a Gateway Load Balancer.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// [Application Load Balancers] The IDs of the security groups for the load balancer.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The IDs of the public subnets.
	//
	// You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.
	//
	// [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.
	//
	// [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	//
	// [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
	//
	// [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.
	//
	// [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP addresses for your subnets.
	SubnetMappings interface{} `field:"optional" json:"subnetMappings" yaml:"subnetMappings"`
	// The IDs of the public subnets.
	//
	// You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.
	//
	// [Application Load Balancers] You must specify subnets from at least two Availability Zones.
	//
	// [Application Load Balancers on Outposts] You must specify one Outpost subnet.
	//
	// [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
	//
	// [Network Load Balancers] You can specify subnets from one or more Availability Zones.
	//
	// [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// The tags to assign to the load balancer.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of load balancer.
	//
	// The default is `application` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}


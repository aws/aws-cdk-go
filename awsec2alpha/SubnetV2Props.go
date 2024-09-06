package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties to define subnet for VPC.
//
// Example:
//   stack := awscdk.Newstack()
//   myVpc := vpc_v2.NewVpcV2(this, jsii.String("Vpc"))
//   routeTable := vpc_v2.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
//   	Vpc: myVpc,
//   })
//   subnet := vpc_v2.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
//   	Vpc: myVpc,
//   	AvailabilityZone: jsii.String("eu-west-2a"),
//   	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
//   	SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//   })
//
//   igw := vpc_v2.NewInternetGateway(this, jsii.String("IGW"), &InternetGatewayProps{
//   	Vpc: myVpc,
//   })
//   vpc_v2.NewRoute(this, jsii.String("IgwRoute"), &RouteProps{
//   	RouteTable: RouteTable,
//   	Destination: jsii.String("0.0.0.0/0"),
//   	Target: map[string]iRouteTarget{
//   		"gateway": igw,
//   	},
//   })
//
// Experimental.
type SubnetV2Props struct {
	// Custom AZ for the subnet.
	// Experimental.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// ipv4 cidr to assign to this subnet.
	//
	// See https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-cidrblock
	// Experimental.
	Ipv4CidrBlock IpCidr `field:"required" json:"ipv4CidrBlock" yaml:"ipv4CidrBlock"`
	// The type of Subnet to configure.
	//
	// The Subnet type will control the ability to route and connect to the
	// Internet.
	//
	// TODO: Add validation check `subnetType` when adding resources (e.g. cannot add NatGateway to private)
	// Experimental.
	SubnetType awsec2.SubnetType `field:"required" json:"subnetType" yaml:"subnetType"`
	// VPC Prop.
	// Experimental.
	Vpc IVpcV2 `field:"required" json:"vpc" yaml:"vpc"`
	// Indicates whether a network interface created in this subnet receives an IPv6 address.
	//
	// If you specify AssignIpv6AddressOnCreation, you must also specify Ipv6CidrBlock.
	// Default: false.
	//
	// Experimental.
	AssignIpv6AddressOnCreation *bool `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// Ipv6 CIDR Range for subnet.
	// Default: No Ipv6 address.
	//
	// Experimental.
	Ipv6CidrBlock IpCidr `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Custom Route for subnet.
	// Default: Default route table.
	//
	// Experimental.
	RouteTable awsec2.IRouteTable `field:"optional" json:"routeTable" yaml:"routeTable"`
	// Subnet name.
	// Default: none.
	//
	// Experimental.
	SubnetName *string `field:"optional" json:"subnetName" yaml:"subnetName"`
}


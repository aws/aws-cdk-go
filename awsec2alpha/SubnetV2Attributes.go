package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties required to import a subnet.
//
// Example:
//   awsec2alpha.SubnetV2_FromSubnetV2Attributes(this, jsii.String("ImportedSubnet"), &SubnetV2Attributes{
//   	SubnetId: jsii.String("subnet-0123456789abcdef0"),
//   	AvailabilityZone: jsii.String("us-west-2a"),
//   	Ipv4CidrBlock: jsii.String("10.2.0.0/24"),
//   	RouteTableId: jsii.String("rtb-0871c310f98da2cbb"),
//   	SubnetType: awscdk.SubnetType_PRIVATE_ISOLATED,
//   })
//
// Experimental.
type SubnetV2Attributes struct {
	// The Availability Zone this subnet is located in.
	// Default: - No AZ information, cannot use AZ selection features.
	//
	// Experimental.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The IPv4 CIDR block associated with the subnet.
	// Default: - No CIDR information, cannot use CIDR filter features.
	//
	// Experimental.
	Ipv4CidrBlock *string `field:"required" json:"ipv4CidrBlock" yaml:"ipv4CidrBlock"`
	// The subnetId for this particular subnet.
	// Experimental.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The type of subnet (public or private) that this subnet represents.
	// Experimental.
	SubnetType awsec2.SubnetType `field:"required" json:"subnetType" yaml:"subnetType"`
	// The IPv4 CIDR block associated with the subnet.
	// Default: - No CIDR information, cannot use CIDR filter features.
	//
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// The ID of the route table for this particular subnet.
	// Default: - No route table information, cannot create VPC endpoints.
	//
	// Experimental.
	RouteTableId *string `field:"optional" json:"routeTableId" yaml:"routeTableId"`
	// Name of the given subnet.
	// Default: - no subnet name.
	//
	// Experimental.
	SubnetName *string `field:"optional" json:"subnetName" yaml:"subnetName"`
}


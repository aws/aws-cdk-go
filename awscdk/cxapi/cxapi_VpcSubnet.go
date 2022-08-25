package cxapi


// A subnet representation that the VPC provider uses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcSubnet := &vpcSubnet{
//   	availabilityZone: jsii.String("availabilityZone"),
//   	routeTableId: jsii.String("routeTableId"),
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	cidr: jsii.String("cidr"),
//   }
//
type VpcSubnet struct {
	// The code of the availability zone this subnet is in (for example, 'us-west-2a').
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The identifier of the route table for this subnet.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
	// The identifier of the subnet.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// CIDR range of the subnet.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}


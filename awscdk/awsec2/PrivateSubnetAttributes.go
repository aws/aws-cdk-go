package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateSubnetAttributes := &PrivateSubnetAttributes{
//   	SubnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	Ipv4CidrBlock: jsii.String("ipv4CidrBlock"),
//   	RouteTableId: jsii.String("routeTableId"),
//   }
//
// Experimental.
type PrivateSubnetAttributes struct {
	// The subnetId for this particular subnet.
	// Experimental.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The Availability Zone the subnet is located in.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The IPv4 CIDR block associated with the subnet.
	// Experimental.
	Ipv4CidrBlock *string `field:"optional" json:"ipv4CidrBlock" yaml:"ipv4CidrBlock"`
	// The ID of the route table for this particular subnet.
	// Experimental.
	RouteTableId *string `field:"optional" json:"routeTableId" yaml:"routeTableId"`
}


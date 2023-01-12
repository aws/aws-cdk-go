package awsec2


// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   // Supply all properties
//   subnet1 := ec2.subnet.fromSubnetAttributes(this, jsii.String("SubnetFromAttributes"), &subnetAttributes{
//   	subnetId: jsii.String("s-1234"),
//   	availabilityZone: jsii.String("pub-az-4465"),
//   	routeTableId: jsii.String("rt-145"),
//   })
//
//   // Supply only subnet id
//   subnet2 := ec2.subnet.fromSubnetId(this, jsii.String("SubnetFromId"), jsii.String("s-1234"))
//
type SubnetAttributes struct {
	// The subnetId for this particular subnet.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The Availability Zone the subnet is located in.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The IPv4 CIDR block associated with the subnet.
	Ipv4CidrBlock *string `field:"optional" json:"ipv4CidrBlock" yaml:"ipv4CidrBlock"`
	// The ID of the route table for this particular subnet.
	RouteTableId *string `field:"optional" json:"routeTableId" yaml:"routeTableId"`
}


package cxapi


// A subnet representation that the VPC provider uses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcSubnet := &VpcSubnet{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	RouteTableId: jsii.String("routeTableId"),
//   	SubnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	Cidr: jsii.String("cidr"),
//   }
//
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type VpcSubnet struct {
	// The code of the availability zone this subnet is in (for example, 'us-west-2a').
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The identifier of the route table for this subnet.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
	// The identifier of the subnet.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// CIDR range of the subnet.
	// Default: - CIDR information not available.
	//
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}


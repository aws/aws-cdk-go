package cxapi


// A group of subnets returned by the VPC provider.
//
// The included subnets do NOT have to be symmetric!
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcSubnetGroup := &VpcSubnetGroup{
//   	Name: jsii.String("name"),
//   	Subnets: []VpcSubnet{
//   		&VpcSubnet{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			RouteTableId: jsii.String("routeTableId"),
//   			SubnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	Type: awscdk.Cx_api.VpcSubnetGroupType_PUBLIC,
//   }
//
// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
type VpcSubnetGroup struct {
	// The name of the subnet group, determined by looking at the tags of of the subnets that belong to it.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The subnets that are part of this group.
	//
	// There is no condition that the subnets have to be symmetric
	// in the group.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Subnets *[]*VpcSubnet `field:"required" json:"subnets" yaml:"subnets"`
	// The type of the subnet group.
	// Deprecated: The definition of this type has moved to `@aws-cdk/cloud-assembly-api`.
	Type VpcSubnetGroupType `field:"required" json:"type" yaml:"type"`
}


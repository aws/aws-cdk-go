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
//   vpcSubnetGroup := &vpcSubnetGroup{
//   	name: jsii.String("name"),
//   	subnets: []vpcSubnet{
//   		&vpcSubnet{
//   			availabilityZone: jsii.String("availabilityZone"),
//   			routeTableId: jsii.String("routeTableId"),
//   			subnetId: jsii.String("subnetId"),
//
//   			// the properties below are optional
//   			cidr: jsii.String("cidr"),
//   		},
//   	},
//   	type: awscdk.Cx_api.vpcSubnetGroupType_PUBLIC,
//   }
//
type VpcSubnetGroup struct {
	// The name of the subnet group, determined by looking at the tags of of the subnets that belong to it.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The subnets that are part of this group.
	//
	// There is no condition that the subnets have to be symmetric
	// in the group.
	Subnets *[]*VpcSubnet `field:"required" json:"subnets" yaml:"subnets"`
	// The type of the subnet group.
	Type VpcSubnetGroupType `field:"required" json:"type" yaml:"type"`
}


package awsec2


// Properties for defining a `CfnSubnetRouteTableAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubnetRouteTableAssociationProps := &cfnSubnetRouteTableAssociationProps{
//   	routeTableId: jsii.String("routeTableId"),
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnSubnetRouteTableAssociationProps struct {
	// The ID of the route table.
	//
	// The physical ID changes when the route table ID is changed.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
	// The ID of the subnet.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}


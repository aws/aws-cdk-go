package awsec2


// Properties for defining a `CfnSubnetRouteTableAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubnetRouteTableAssociationProps := &CfnSubnetRouteTableAssociationProps{
//   	RouteTableId: jsii.String("routeTableId"),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetroutetableassociation.html
//
type CfnSubnetRouteTableAssociationProps struct {
	// The ID of the route table.
	//
	// The physical ID changes when the route table ID is changed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetroutetableassociation.html#cfn-ec2-subnetroutetableassociation-routetableid
	//
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
	// The ID of the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetroutetableassociation.html#cfn-ec2-subnetroutetableassociation-subnetid
	//
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}


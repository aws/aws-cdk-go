package awsec2


// Properties for defining a `CfnIpPoolRouteTableAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIpPoolRouteTableAssociationProps := &CfnIpPoolRouteTableAssociationProps{
//   	PublicIpv4Pool: jsii.String("publicIpv4Pool"),
//   	RouteTableId: jsii.String("routeTableId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ippoolroutetableassociation.html
//
type CfnIpPoolRouteTableAssociationProps struct {
	// The ID of a public IPv4 address pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ippoolroutetableassociation.html#cfn-ec2-ippoolroutetableassociation-publicipv4pool
	//
	PublicIpv4Pool *string `field:"required" json:"publicIpv4Pool" yaml:"publicIpv4Pool"`
	// The ID of a route table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ippoolroutetableassociation.html#cfn-ec2-ippoolroutetableassociation-routetableid
	//
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}


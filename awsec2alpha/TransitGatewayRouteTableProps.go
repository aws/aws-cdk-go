package awsec2alpha


// Common properties for creating a Transit Gateway Route Table resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//
//   var transitGateway TransitGateway
//
//   transitGatewayRouteTableProps := &TransitGatewayRouteTableProps{
//   	TransitGateway: transitGateway,
//
//   	// the properties below are optional
//   	TransitGatewayRouteTableName: jsii.String("transitGatewayRouteTableName"),
//   }
//
// Experimental.
type TransitGatewayRouteTableProps struct {
	// The Transit Gateway that this route table belongs to.
	// Experimental.
	TransitGateway ITransitGateway `field:"required" json:"transitGateway" yaml:"transitGateway"`
	// Physical name of this Transit Gateway Route Table.
	// Default: - Assigned by CloudFormation.
	//
	// Experimental.
	TransitGatewayRouteTableName *string `field:"optional" json:"transitGatewayRouteTableName" yaml:"transitGatewayRouteTableName"`
}


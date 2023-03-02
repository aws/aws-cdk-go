package awsec2


// Properties for defining a `CfnTransitGatewayRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayRouteProps := &cfnTransitGatewayRouteProps{
//   	transitGatewayRouteTableId: jsii.String("transitGatewayRouteTableId"),
//
//   	// the properties below are optional
//   	blackhole: jsii.Boolean(false),
//   	destinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	transitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   }
//
type CfnTransitGatewayRouteProps struct {
	// The ID of the transit gateway route table.
	TransitGatewayRouteTableId *string `field:"required" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
	// Indicates whether to drop traffic that matches this route.
	Blackhole interface{} `field:"optional" json:"blackhole" yaml:"blackhole"`
	// The CIDR block used for destination matches.
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the attachment.
	TransitGatewayAttachmentId *string `field:"optional" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
}


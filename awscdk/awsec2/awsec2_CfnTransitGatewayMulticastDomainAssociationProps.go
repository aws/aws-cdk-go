package awsec2


// Properties for defining a `CfnTransitGatewayMulticastDomainAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayMulticastDomainAssociationProps := &cfnTransitGatewayMulticastDomainAssociationProps{
//   	subnetId: jsii.String("subnetId"),
//   	transitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	transitGatewayMulticastDomainId: jsii.String("transitGatewayMulticastDomainId"),
//   }
//
type CfnTransitGatewayMulticastDomainAssociationProps struct {
	// The IDs of the subnets to associate with the transit gateway multicast domain.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string `field:"required" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}


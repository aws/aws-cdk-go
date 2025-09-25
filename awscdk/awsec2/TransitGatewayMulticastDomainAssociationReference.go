package awsec2


// A reference to a TransitGatewayMulticastDomainAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayMulticastDomainAssociationReference := &TransitGatewayMulticastDomainAssociationReference{
//   	SubnetId: jsii.String("subnetId"),
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	TransitGatewayMulticastDomainId: jsii.String("transitGatewayMulticastDomainId"),
//   }
//
type TransitGatewayMulticastDomainAssociationReference struct {
	// The SubnetId of the TransitGatewayMulticastDomainAssociation resource.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The TransitGatewayAttachmentId of the TransitGatewayMulticastDomainAssociation resource.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The TransitGatewayMulticastDomainId of the TransitGatewayMulticastDomainAssociation resource.
	TransitGatewayMulticastDomainId *string `field:"required" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}


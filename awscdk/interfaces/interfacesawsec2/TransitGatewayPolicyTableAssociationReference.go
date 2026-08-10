package interfacesawsec2


// A reference to a TransitGatewayPolicyTableAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayPolicyTableAssociationReference := &TransitGatewayPolicyTableAssociationReference{
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	TransitGatewayPolicyTableId: jsii.String("transitGatewayPolicyTableId"),
//   }
//
type TransitGatewayPolicyTableAssociationReference struct {
	// The TransitGatewayAttachmentId of the TransitGatewayPolicyTableAssociation resource.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// The TransitGatewayPolicyTableId of the TransitGatewayPolicyTableAssociation resource.
	TransitGatewayPolicyTableId *string `field:"required" json:"transitGatewayPolicyTableId" yaml:"transitGatewayPolicyTableId"`
}


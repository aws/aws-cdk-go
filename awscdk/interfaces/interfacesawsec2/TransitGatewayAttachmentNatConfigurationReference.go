package interfacesawsec2


// A reference to a TransitGatewayAttachmentNatConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayAttachmentNatConfigurationReference := &TransitGatewayAttachmentNatConfigurationReference{
//   	TransitGatewayAttachmentNatConfigurationArn: jsii.String("transitGatewayAttachmentNatConfigurationArn"),
//   }
//
type TransitGatewayAttachmentNatConfigurationReference struct {
	// The Arn of the TransitGatewayAttachmentNatConfiguration resource.
	TransitGatewayAttachmentNatConfigurationArn *string `field:"required" json:"transitGatewayAttachmentNatConfigurationArn" yaml:"transitGatewayAttachmentNatConfigurationArn"`
}


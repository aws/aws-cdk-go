package awsnetworkmanager


// Describes a core network Connect attachment options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectAttachmentOptionsProperty := &connectAttachmentOptionsProperty{
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnConnectAttachment_ConnectAttachmentOptionsProperty struct {
	// The protocol used for the attachment connection.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}


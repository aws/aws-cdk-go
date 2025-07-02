package awsnetworkmanager


// Describes a core network Connect attachment options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectAttachmentOptionsProperty := &ConnectAttachmentOptionsProperty{
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectattachment-connectattachmentoptions.html
//
type CfnConnectAttachment_ConnectAttachmentOptionsProperty struct {
	// The protocol used for the attachment connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkmanager-connectattachment-connectattachmentoptions.html#cfn-networkmanager-connectattachment-connectattachmentoptions-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}


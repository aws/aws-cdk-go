package awskinesisfirehose


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentIdOptionsProperty := &DocumentIdOptionsProperty{
//   	DefaultDocumentIdFormat: jsii.String("defaultDocumentIdFormat"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-documentidoptions.html
//
type CfnDeliveryStream_DocumentIdOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-documentidoptions.html#cfn-kinesisfirehose-deliverystream-documentidoptions-defaultdocumentidformat
	//
	DefaultDocumentIdFormat *string `field:"required" json:"defaultDocumentIdFormat" yaml:"defaultDocumentIdFormat"`
}


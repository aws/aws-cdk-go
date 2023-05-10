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
type CfnDeliveryStream_DocumentIdOptionsProperty struct {
	// `CfnDeliveryStream.DocumentIdOptionsProperty.DefaultDocumentIdFormat`.
	DefaultDocumentIdFormat *string `field:"required" json:"defaultDocumentIdFormat" yaml:"defaultDocumentIdFormat"`
}


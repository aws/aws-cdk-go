package previewawsb2bimixins


// Contains options for processing inbound EDI files.
//
// These options allow for customizing how incoming EDI documents are processed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inboundEdiOptionsProperty := &InboundEdiOptionsProperty{
//   	X12: &X12InboundEdiOptionsProperty{
//   		AcknowledgmentOptions: &X12AcknowledgmentOptionsProperty{
//   			FunctionalAcknowledgment: jsii.String("functionalAcknowledgment"),
//   			TechnicalAcknowledgment: jsii.String("technicalAcknowledgment"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-inboundedioptions.html
//
type CfnPartnershipPropsMixin_InboundEdiOptionsProperty struct {
	// A structure that contains X12-specific options for processing inbound X12 EDI files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-inboundedioptions.html#cfn-b2bi-partnership-inboundedioptions-x12
	//
	X12 interface{} `field:"optional" json:"x12" yaml:"x12"`
}


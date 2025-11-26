package previewawsb2bimixins


// A wrapper structure for an X12 definition object.
//
// the X12 envelope ensures the integrity of the data and the efficiency of the information exchange. The X12 message structure has hierarchical levels. From highest to the lowest, they are:
//
// - Interchange Envelope
// - Functional Group
// - Transaction Set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   x12EnvelopeProperty := &X12EnvelopeProperty{
//   	Common: &X12OutboundEdiHeadersProperty{
//   		ControlNumbers: &X12ControlNumbersProperty{
//   			StartingFunctionalGroupControlNumber: jsii.Number(123),
//   			StartingInterchangeControlNumber: jsii.Number(123),
//   			StartingTransactionSetControlNumber: jsii.Number(123),
//   		},
//   		Delimiters: &X12DelimitersProperty{
//   			ComponentSeparator: jsii.String("componentSeparator"),
//   			DataElementSeparator: jsii.String("dataElementSeparator"),
//   			SegmentTerminator: jsii.String("segmentTerminator"),
//   		},
//   		FunctionalGroupHeaders: &X12FunctionalGroupHeadersProperty{
//   			ApplicationReceiverCode: jsii.String("applicationReceiverCode"),
//   			ApplicationSenderCode: jsii.String("applicationSenderCode"),
//   			ResponsibleAgencyCode: jsii.String("responsibleAgencyCode"),
//   		},
//   		Gs05TimeFormat: jsii.String("gs05TimeFormat"),
//   		InterchangeControlHeaders: &X12InterchangeControlHeadersProperty{
//   			AcknowledgmentRequestedCode: jsii.String("acknowledgmentRequestedCode"),
//   			ReceiverId: jsii.String("receiverId"),
//   			ReceiverIdQualifier: jsii.String("receiverIdQualifier"),
//   			RepetitionSeparator: jsii.String("repetitionSeparator"),
//   			SenderId: jsii.String("senderId"),
//   			SenderIdQualifier: jsii.String("senderIdQualifier"),
//   			UsageIndicatorCode: jsii.String("usageIndicatorCode"),
//   		},
//   		ValidateEdi: jsii.Boolean(false),
//   	},
//   	WrapOptions: &WrapOptionsProperty{
//   		LineLength: jsii.Number(123),
//   		LineTerminator: jsii.String("lineTerminator"),
//   		WrapBy: jsii.String("wrapBy"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12envelope.html
//
type CfnPartnershipPropsMixin_X12EnvelopeProperty struct {
	// A container for the X12 outbound EDI headers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12envelope.html#cfn-b2bi-partnership-x12envelope-common
	//
	Common interface{} `field:"optional" json:"common" yaml:"common"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12envelope.html#cfn-b2bi-partnership-x12envelope-wrapoptions
	//
	WrapOptions interface{} `field:"optional" json:"wrapOptions" yaml:"wrapOptions"`
}


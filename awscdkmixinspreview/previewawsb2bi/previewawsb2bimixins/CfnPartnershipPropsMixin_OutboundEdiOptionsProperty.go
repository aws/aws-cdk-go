package previewawsb2bimixins


// A container for outbound EDI options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outboundEdiOptionsProperty := &OutboundEdiOptionsProperty{
//   	X12: &X12EnvelopeProperty{
//   		Common: &X12OutboundEdiHeadersProperty{
//   			ControlNumbers: &X12ControlNumbersProperty{
//   				StartingFunctionalGroupControlNumber: jsii.Number(123),
//   				StartingInterchangeControlNumber: jsii.Number(123),
//   				StartingTransactionSetControlNumber: jsii.Number(123),
//   			},
//   			Delimiters: &X12DelimitersProperty{
//   				ComponentSeparator: jsii.String("componentSeparator"),
//   				DataElementSeparator: jsii.String("dataElementSeparator"),
//   				SegmentTerminator: jsii.String("segmentTerminator"),
//   			},
//   			FunctionalGroupHeaders: &X12FunctionalGroupHeadersProperty{
//   				ApplicationReceiverCode: jsii.String("applicationReceiverCode"),
//   				ApplicationSenderCode: jsii.String("applicationSenderCode"),
//   				ResponsibleAgencyCode: jsii.String("responsibleAgencyCode"),
//   			},
//   			Gs05TimeFormat: jsii.String("gs05TimeFormat"),
//   			InterchangeControlHeaders: &X12InterchangeControlHeadersProperty{
//   				AcknowledgmentRequestedCode: jsii.String("acknowledgmentRequestedCode"),
//   				ReceiverId: jsii.String("receiverId"),
//   				ReceiverIdQualifier: jsii.String("receiverIdQualifier"),
//   				RepetitionSeparator: jsii.String("repetitionSeparator"),
//   				SenderId: jsii.String("senderId"),
//   				SenderIdQualifier: jsii.String("senderIdQualifier"),
//   				UsageIndicatorCode: jsii.String("usageIndicatorCode"),
//   			},
//   			ValidateEdi: jsii.Boolean(false),
//   		},
//   		WrapOptions: &WrapOptionsProperty{
//   			LineLength: jsii.Number(123),
//   			LineTerminator: jsii.String("lineTerminator"),
//   			WrapBy: jsii.String("wrapBy"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-outboundedioptions.html
//
type CfnPartnershipPropsMixin_OutboundEdiOptionsProperty struct {
	// A structure that contains an X12 envelope structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-outboundedioptions.html#cfn-b2bi-partnership-outboundedioptions-x12
	//
	X12 interface{} `field:"optional" json:"x12" yaml:"x12"`
}


package awsb2bi


// Contains the details for an Outbound EDI capability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capabilityOptionsProperty := &CapabilityOptionsProperty{
//   	InboundEdi: &InboundEdiOptionsProperty{
//   		X12: &X12InboundEdiOptionsProperty{
//   			AcknowledgmentOptions: &X12AcknowledgmentOptionsProperty{
//   				FunctionalAcknowledgment: jsii.String("functionalAcknowledgment"),
//   				TechnicalAcknowledgment: jsii.String("technicalAcknowledgment"),
//   			},
//   		},
//   	},
//   	OutboundEdi: &OutboundEdiOptionsProperty{
//   		X12: &X12EnvelopeProperty{
//   			Common: &X12OutboundEdiHeadersProperty{
//   				ControlNumbers: &X12ControlNumbersProperty{
//   					StartingFunctionalGroupControlNumber: jsii.Number(123),
//   					StartingInterchangeControlNumber: jsii.Number(123),
//   					StartingTransactionSetControlNumber: jsii.Number(123),
//   				},
//   				Delimiters: &X12DelimitersProperty{
//   					ComponentSeparator: jsii.String("componentSeparator"),
//   					DataElementSeparator: jsii.String("dataElementSeparator"),
//   					SegmentTerminator: jsii.String("segmentTerminator"),
//   				},
//   				FunctionalGroupHeaders: &X12FunctionalGroupHeadersProperty{
//   					ApplicationReceiverCode: jsii.String("applicationReceiverCode"),
//   					ApplicationSenderCode: jsii.String("applicationSenderCode"),
//   					ResponsibleAgencyCode: jsii.String("responsibleAgencyCode"),
//   				},
//   				Gs05TimeFormat: jsii.String("gs05TimeFormat"),
//   				InterchangeControlHeaders: &X12InterchangeControlHeadersProperty{
//   					AcknowledgmentRequestedCode: jsii.String("acknowledgmentRequestedCode"),
//   					ReceiverId: jsii.String("receiverId"),
//   					ReceiverIdQualifier: jsii.String("receiverIdQualifier"),
//   					RepetitionSeparator: jsii.String("repetitionSeparator"),
//   					SenderId: jsii.String("senderId"),
//   					SenderIdQualifier: jsii.String("senderIdQualifier"),
//   					UsageIndicatorCode: jsii.String("usageIndicatorCode"),
//   				},
//   				ValidateEdi: jsii.Boolean(false),
//   			},
//   			WrapOptions: &WrapOptionsProperty{
//   				LineLength: jsii.Number(123),
//   				LineTerminator: jsii.String("lineTerminator"),
//   				WrapBy: jsii.String("wrapBy"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-capabilityoptions.html
//
type CfnPartnership_CapabilityOptionsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-capabilityoptions.html#cfn-b2bi-partnership-capabilityoptions-inboundedi
	//
	InboundEdi interface{} `field:"optional" json:"inboundEdi" yaml:"inboundEdi"`
	// A structure that contains the outbound EDI options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-capabilityoptions.html#cfn-b2bi-partnership-capabilityoptions-outboundedi
	//
	OutboundEdi interface{} `field:"optional" json:"outboundEdi" yaml:"outboundEdi"`
}


package awsb2bi


// A container for outbound EDI options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outboundEdiOptionsProperty := &OutboundEdiOptionsProperty{
//   	X12: &X12EnvelopeProperty{
//   		Common: &X12OutboundEdiHeadersProperty{
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
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-outboundedioptions.html
//
type CfnPartnership_OutboundEdiOptionsProperty struct {
	// A structure that contains an X12 envelope structure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-outboundedioptions.html#cfn-b2bi-partnership-outboundedioptions-x12
	//
	X12 interface{} `field:"required" json:"x12" yaml:"x12"`
}


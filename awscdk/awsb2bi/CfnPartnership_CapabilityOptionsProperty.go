package awsb2bi


// Contains the details for an Outbound EDI capability.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   capabilityOptionsProperty := &CapabilityOptionsProperty{
//   	OutboundEdi: &OutboundEdiOptionsProperty{
//   		X12: &X12EnvelopeProperty{
//   			Common: &X12OutboundEdiHeadersProperty{
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
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-capabilityoptions.html
//
type CfnPartnership_CapabilityOptionsProperty struct {
	// A structure that contains the outbound EDI options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-capabilityoptions.html#cfn-b2bi-partnership-capabilityoptions-outboundedi
	//
	OutboundEdi interface{} `field:"optional" json:"outboundEdi" yaml:"outboundEdi"`
}


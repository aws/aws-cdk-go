package mixinsawsb2bi


// A structure containing the details for an outbound EDI object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   x12OutboundEdiHeadersProperty := &X12OutboundEdiHeadersProperty{
//   	ControlNumbers: &X12ControlNumbersProperty{
//   		StartingFunctionalGroupControlNumber: jsii.Number(123),
//   		StartingInterchangeControlNumber: jsii.Number(123),
//   		StartingTransactionSetControlNumber: jsii.Number(123),
//   	},
//   	Delimiters: &X12DelimitersProperty{
//   		ComponentSeparator: jsii.String("componentSeparator"),
//   		DataElementSeparator: jsii.String("dataElementSeparator"),
//   		SegmentTerminator: jsii.String("segmentTerminator"),
//   	},
//   	FunctionalGroupHeaders: &X12FunctionalGroupHeadersProperty{
//   		ApplicationReceiverCode: jsii.String("applicationReceiverCode"),
//   		ApplicationSenderCode: jsii.String("applicationSenderCode"),
//   		ResponsibleAgencyCode: jsii.String("responsibleAgencyCode"),
//   	},
//   	Gs05TimeFormat: jsii.String("gs05TimeFormat"),
//   	InterchangeControlHeaders: &X12InterchangeControlHeadersProperty{
//   		AcknowledgmentRequestedCode: jsii.String("acknowledgmentRequestedCode"),
//   		ReceiverId: jsii.String("receiverId"),
//   		ReceiverIdQualifier: jsii.String("receiverIdQualifier"),
//   		RepetitionSeparator: jsii.String("repetitionSeparator"),
//   		SenderId: jsii.String("senderId"),
//   		SenderIdQualifier: jsii.String("senderIdQualifier"),
//   		UsageIndicatorCode: jsii.String("usageIndicatorCode"),
//   	},
//   	ValidateEdi: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12outboundediheaders.html
//
type CfnPartnershipPropsMixin_X12OutboundEdiHeadersProperty struct {
	// Specifies control number configuration for outbound X12 EDI headers.
	//
	// These settings determine the starting values for interchange, functional group, and transaction set control numbers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12outboundediheaders.html#cfn-b2bi-partnership-x12outboundediheaders-controlnumbers
	//
	ControlNumbers interface{} `field:"optional" json:"controlNumbers" yaml:"controlNumbers"`
	// The delimiters, for example semicolon ( `;` ), that separates sections of the headers for the X12 object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12outboundediheaders.html#cfn-b2bi-partnership-x12outboundediheaders-delimiters
	//
	Delimiters interface{} `field:"optional" json:"delimiters" yaml:"delimiters"`
	// The functional group headers for the X12 object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12outboundediheaders.html#cfn-b2bi-partnership-x12outboundediheaders-functionalgroupheaders
	//
	FunctionalGroupHeaders interface{} `field:"optional" json:"functionalGroupHeaders" yaml:"functionalGroupHeaders"`
	// Specifies the time format in the GS05 element (time) of the functional group header.
	//
	// The following formats use 24-hour clock time:
	//
	// - `HHMM` - Hours and minutes
	// - `HHMMSS` - Hours, minutes, and seconds
	// - `HHMMSSDD` - Hours, minutes, seconds, and decimal seconds
	//
	// Where:
	//
	// - `HH` - Hours (00-23)
	// - `MM` - Minutes (00-59)
	// - `SS` - Seconds (00-59)
	// - `DD` - Hundredths of seconds (00-99).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12outboundediheaders.html#cfn-b2bi-partnership-x12outboundediheaders-gs05timeformat
	//
	Gs05TimeFormat *string `field:"optional" json:"gs05TimeFormat" yaml:"gs05TimeFormat"`
	// In X12 EDI messages, delimiters are used to mark the end of segments or elements, and are defined in the interchange control header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12outboundediheaders.html#cfn-b2bi-partnership-x12outboundediheaders-interchangecontrolheaders
	//
	InterchangeControlHeaders interface{} `field:"optional" json:"interchangeControlHeaders" yaml:"interchangeControlHeaders"`
	// Specifies whether or not to validate the EDI for this X12 object: `TRUE` or `FALSE` .
	//
	// When enabled, this performs both standard EDI validation and applies any configured custom validation rules including element length constraints, code list validations, and element requirement checks. Validation results are returned in the response validation messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12outboundediheaders.html#cfn-b2bi-partnership-x12outboundediheaders-validateedi
	//
	ValidateEdi interface{} `field:"optional" json:"validateEdi" yaml:"validateEdi"`
}


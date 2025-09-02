package awsb2bi


// In X12, the Interchange Control Header is the first segment of an EDI document and is part of the Interchange Envelope.
//
// It contains information about the sender and receiver, the date and time of transmission, and the X12 version being used. It also includes delivery information, such as the sender and receiver IDs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   x12InterchangeControlHeadersProperty := &X12InterchangeControlHeadersProperty{
//   	AcknowledgmentRequestedCode: jsii.String("acknowledgmentRequestedCode"),
//   	ReceiverId: jsii.String("receiverId"),
//   	ReceiverIdQualifier: jsii.String("receiverIdQualifier"),
//   	RepetitionSeparator: jsii.String("repetitionSeparator"),
//   	SenderId: jsii.String("senderId"),
//   	SenderIdQualifier: jsii.String("senderIdQualifier"),
//   	UsageIndicatorCode: jsii.String("usageIndicatorCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html
//
type CfnPartnership_X12InterchangeControlHeadersProperty struct {
	// Located at position ISA-14 in the header.
	//
	// The value "1" indicates that the sender is requesting an interchange acknowledgment at receipt of the interchange. The value "0" is used otherwise.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-acknowledgmentrequestedcode
	//
	AcknowledgmentRequestedCode *string `field:"optional" json:"acknowledgmentRequestedCode" yaml:"acknowledgmentRequestedCode"`
	// Located at position ISA-08 in the header.
	//
	// This value (along with the `receiverIdQualifier` ) identifies the intended recipient of the interchange.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-receiverid
	//
	ReceiverId *string `field:"optional" json:"receiverId" yaml:"receiverId"`
	// Located at position ISA-07 in the header.
	//
	// Qualifier for the receiver ID. Together, the ID and qualifier uniquely identify the receiving trading partner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-receiveridqualifier
	//
	ReceiverIdQualifier *string `field:"optional" json:"receiverIdQualifier" yaml:"receiverIdQualifier"`
	// Located at position ISA-11 in the header.
	//
	// This string makes it easier when you need to group similar adjacent element values together without using extra segments.
	//
	// > This parameter is only honored for version greater than 401 ( `VERSION_4010` and higher).
	// >
	// > For versions less than 401, this field is called [StandardsId](https://docs.aws.amazon.com/https://www.stedi.com/edi/x12-004010/segment/ISA#ISA-11) , in which case our service sets the value to `U` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-repetitionseparator
	//
	RepetitionSeparator *string `field:"optional" json:"repetitionSeparator" yaml:"repetitionSeparator"`
	// Located at position ISA-06 in the header.
	//
	// This value (along with the `senderIdQualifier` ) identifies the sender of the interchange.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-senderid
	//
	SenderId *string `field:"optional" json:"senderId" yaml:"senderId"`
	// Located at position ISA-05 in the header.
	//
	// Qualifier for the sender ID. Together, the ID and qualifier uniquely identify the sending trading partner.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-senderidqualifier
	//
	SenderIdQualifier *string `field:"optional" json:"senderIdQualifier" yaml:"senderIdQualifier"`
	// Located at position ISA-15 in the header. Specifies how this interchange is being used:.
	//
	// - `T` indicates this interchange is for testing.
	// - `P` indicates this interchange is for production.
	// - `I` indicates this interchange is informational.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-usageindicatorcode
	//
	UsageIndicatorCode *string `field:"optional" json:"usageIndicatorCode" yaml:"usageIndicatorCode"`
}


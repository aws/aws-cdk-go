package awsb2bi


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-acknowledgmentrequestedcode
	//
	AcknowledgmentRequestedCode *string `field:"optional" json:"acknowledgmentRequestedCode" yaml:"acknowledgmentRequestedCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-receiverid
	//
	ReceiverId *string `field:"optional" json:"receiverId" yaml:"receiverId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-receiveridqualifier
	//
	ReceiverIdQualifier *string `field:"optional" json:"receiverIdQualifier" yaml:"receiverIdQualifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-repetitionseparator
	//
	RepetitionSeparator *string `field:"optional" json:"repetitionSeparator" yaml:"repetitionSeparator"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-senderid
	//
	SenderId *string `field:"optional" json:"senderId" yaml:"senderId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-senderidqualifier
	//
	SenderIdQualifier *string `field:"optional" json:"senderIdQualifier" yaml:"senderIdQualifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-b2bi-partnership-x12interchangecontrolheaders.html#cfn-b2bi-partnership-x12interchangecontrolheaders-usageindicatorcode
	//
	UsageIndicatorCode *string `field:"optional" json:"usageIndicatorCode" yaml:"usageIndicatorCode"`
}


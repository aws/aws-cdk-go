package awseventsv2


// A message attribute value and its data type, for Amazon SQS and Amazon SNS targets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageAttributeValueProperty := &MessageAttributeValueProperty{
//   	DataType: jsii.String("dataType"),
//
//   	// the properties below are optional
//   	BinaryValue: jsii.String("binaryValue"),
//   	StringValue: jsii.String("stringValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-messageattributevalue.html
//
type CfnSubscriber_MessageAttributeValueProperty struct {
	// The attribute data type.
	//
	// For Amazon SQS targets, specify String, Number, or Binary, optionally with a custom label suffix such as Number.float. For Amazon SNS targets, specify String, String.Array, Number, or Binary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-messageattributevalue.html#cfn-eventsv2-subscriber-messageattributevalue-datatype
	//
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The attribute value for the Binary data type, Base64-encoded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-messageattributevalue.html#cfn-eventsv2-subscriber-messageattributevalue-binaryvalue
	//
	BinaryValue *string `field:"optional" json:"binaryValue" yaml:"binaryValue"`
	// The attribute value for the String and Number data types (and String.Array for Amazon SNS targets).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-messageattributevalue.html#cfn-eventsv2-subscriber-messageattributevalue-stringvalue
	//
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}


package awsiotevents


// Sends information about the detector model instance and the event that triggered the action to an Amazon Kinesis Data Firehose delivery stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firehoseProperty := &FirehoseProperty{
//   	DeliveryStreamName: jsii.String("deliveryStreamName"),
//
//   	// the properties below are optional
//   	Payload: &PayloadProperty{
//   		ContentExpression: jsii.String("contentExpression"),
//   		Type: jsii.String("type"),
//   	},
//   	Separator: jsii.String("separator"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-firehose.html
//
type CfnDetectorModel_FirehoseProperty struct {
	// The name of the Kinesis Data Firehose delivery stream where the data is written.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-firehose.html#cfn-iotevents-detectormodel-firehose-deliverystreamname
	//
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
	// You can configure the action payload when you send a message to an Amazon Data Firehose delivery stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-firehose.html#cfn-iotevents-detectormodel-firehose-payload
	//
	Payload interface{} `field:"optional" json:"payload" yaml:"payload"`
	// A character separator that is used to separate records written to the Kinesis Data Firehose delivery stream.
	//
	// Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-firehose.html#cfn-iotevents-detectormodel-firehose-separator
	//
	Separator *string `field:"optional" json:"separator" yaml:"separator"`
}


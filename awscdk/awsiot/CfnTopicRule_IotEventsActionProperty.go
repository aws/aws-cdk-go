package awsiot


// Sends an input to an AWS IoT Events detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotEventsActionProperty := &IotEventsActionProperty{
//   	InputName: jsii.String("inputName"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	BatchMode: jsii.Boolean(false),
//   	MessageId: jsii.String("messageId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-ioteventsaction.html
//
type CfnTopicRule_IotEventsActionProperty struct {
	// The name of the AWS IoT Events input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-ioteventsaction.html#cfn-iot-topicrule-ioteventsaction-inputname
	//
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// The ARN of the role that grants AWS IoT permission to send an input to an AWS IoT Events detector.
	//
	// ("Action":"iotevents:BatchPutMessage").
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-ioteventsaction.html#cfn-iot-topicrule-ioteventsaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Whether to process the event actions as a batch. The default value is `false` .
	//
	// When `batchMode` is `true` , you can't specify a `messageId` .
	//
	// When `batchMode` is `true` and the rule SQL statement evaluates to an Array, each Array element is treated as a separate message when Events by calling [`BatchPutMessage`](https://docs.aws.amazon.com/iotevents/latest/apireference/API_iotevents-data_BatchPutMessage.html) . The resulting array can't have more than 10 messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-ioteventsaction.html#cfn-iot-topicrule-ioteventsaction-batchmode
	//
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
	// The ID of the message. The default `messageId` is a new UUID value.
	//
	// When `batchMode` is `true` , you can't specify a `messageId` --a new UUID value will be assigned.
	//
	// Assign a value to this property to ensure that only one input (message) with a given `messageId` will be processed by an AWS IoT Events detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-ioteventsaction.html#cfn-iot-topicrule-ioteventsaction-messageid
	//
	MessageId *string `field:"optional" json:"messageId" yaml:"messageId"`
}


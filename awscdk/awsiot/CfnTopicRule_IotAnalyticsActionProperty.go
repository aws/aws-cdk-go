package awsiot


// Sends message data to an AWS IoT Analytics channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotAnalyticsActionProperty := &IotAnalyticsActionProperty{
//   	ChannelName: jsii.String("channelName"),
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	BatchMode: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotanalyticsaction.html
//
type CfnTopicRule_IotAnalyticsActionProperty struct {
	// The name of the IoT Analytics channel to which message data will be sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotanalyticsaction.html#cfn-iot-topicrule-iotanalyticsaction-channelname
	//
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The ARN of the role which has a policy that grants IoT Analytics permission to send message data via IoT Analytics (iotanalytics:BatchPutMessage).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotanalyticsaction.html#cfn-iot-topicrule-iotanalyticsaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Whether to process the action as a batch. The default value is `false` .
	//
	// When `batchMode` is `true` and the rule SQL statement evaluates to an Array, each Array element is delivered as a separate message when passed by [`BatchPutMessage`](https://docs.aws.amazon.com/iotanalytics/latest/APIReference/API_BatchPutMessage.html) The resulting array can't have more than 100 messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotanalyticsaction.html#cfn-iot-topicrule-iotanalyticsaction-batchmode
	//
	BatchMode interface{} `field:"optional" json:"batchMode" yaml:"batchMode"`
}


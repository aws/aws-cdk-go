package previewawsiotfleetwisemixins


// The MQTT topic to which the AWS IoT FleetWise campaign routes data.
//
// For more information, see [Device communication protocols](https://docs.aws.amazon.com/iot/latest/developerguide/protocols.html) in the *AWS IoT Core Developer Guide* .
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mqttTopicConfigProperty := &MqttTopicConfigProperty{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	MqttTopicArn: jsii.String("mqttTopicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-mqtttopicconfig.html
//
type CfnCampaignPropsMixin_MqttTopicConfigProperty struct {
	// The ARN of the role that grants AWS IoT FleetWise permission to access and act on messages sent to the MQTT topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-mqtttopicconfig.html#cfn-iotfleetwise-campaign-mqtttopicconfig-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The ARN of the MQTT topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-mqtttopicconfig.html#cfn-iotfleetwise-campaign-mqtttopicconfig-mqtttopicarn
	//
	MqttTopicArn *string `field:"optional" json:"mqttTopicArn" yaml:"mqttTopicArn"`
}


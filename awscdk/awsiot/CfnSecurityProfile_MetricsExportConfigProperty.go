package awsiot


// Specifies the MQTT topic and role ARN required for metric export.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsExportConfigProperty := &MetricsExportConfigProperty{
//   	MqttTopic: jsii.String("mqttTopic"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricsexportconfig.html
//
type CfnSecurityProfile_MetricsExportConfigProperty struct {
	// The MQTT topic that Device Defender Detect should publish messages to for metrics export.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricsexportconfig.html#cfn-iot-securityprofile-metricsexportconfig-mqtttopic
	//
	MqttTopic *string `field:"required" json:"mqttTopic" yaml:"mqttTopic"`
	// This role ARN has permission to publish MQTT messages, after which Device Defender Detect can assume the role and publish messages on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-metricsexportconfig.html#cfn-iot-securityprofile-metricsexportconfig-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}


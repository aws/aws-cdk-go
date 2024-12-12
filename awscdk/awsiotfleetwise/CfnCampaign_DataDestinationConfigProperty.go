package awsiotfleetwise


// The destination where the AWS IoT FleetWise campaign sends data.
//
// You can send data to be stored in Amazon S3 or Amazon Timestream .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataDestinationConfigProperty := &DataDestinationConfigProperty{
//   	MqttTopicConfig: &MqttTopicConfigProperty{
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//   		MqttTopicArn: jsii.String("mqttTopicArn"),
//   	},
//   	S3Config: &S3ConfigProperty{
//   		BucketArn: jsii.String("bucketArn"),
//
//   		// the properties below are optional
//   		DataFormat: jsii.String("dataFormat"),
//   		Prefix: jsii.String("prefix"),
//   		StorageCompressionFormat: jsii.String("storageCompressionFormat"),
//   	},
//   	TimestreamConfig: &TimestreamConfigProperty{
//   		ExecutionRoleArn: jsii.String("executionRoleArn"),
//   		TimestreamTableArn: jsii.String("timestreamTableArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datadestinationconfig.html
//
type CfnCampaign_DataDestinationConfigProperty struct {
	// The MQTT topic to which the AWS IoT FleetWise campaign routes data.
	//
	// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datadestinationconfig.html#cfn-iotfleetwise-campaign-datadestinationconfig-mqtttopicconfig
	//
	MqttTopicConfig interface{} `field:"optional" json:"mqttTopicConfig" yaml:"mqttTopicConfig"`
	// The Amazon S3 bucket where the AWS IoT FleetWise campaign sends data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datadestinationconfig.html#cfn-iotfleetwise-campaign-datadestinationconfig-s3config
	//
	S3Config interface{} `field:"optional" json:"s3Config" yaml:"s3Config"`
	// The Amazon Timestream table where the campaign sends data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datadestinationconfig.html#cfn-iotfleetwise-campaign-datadestinationconfig-timestreamconfig
	//
	TimestreamConfig interface{} `field:"optional" json:"timestreamConfig" yaml:"timestreamConfig"`
}


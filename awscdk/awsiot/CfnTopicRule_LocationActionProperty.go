package awsiot


// Describes an action to send device location updates from an MQTT message to an Amazon Location tracker resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationActionProperty := &LocationActionProperty{
//   	DeviceId: jsii.String("deviceId"),
//   	Latitude: jsii.String("latitude"),
//   	Longitude: jsii.String("longitude"),
//   	RoleArn: jsii.String("roleArn"),
//   	TrackerName: jsii.String("trackerName"),
//
//   	// the properties below are optional
//   	Timestamp: &TimestampProperty{
//   		Value: jsii.String("value"),
//
//   		// the properties below are optional
//   		Unit: jsii.String("unit"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-locationaction.html
//
type CfnTopicRule_LocationActionProperty struct {
	// The unique ID of the device providing the location data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-locationaction.html#cfn-iot-topicrule-locationaction-deviceid
	//
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// A string that evaluates to a double value that represents the latitude of the device's location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-locationaction.html#cfn-iot-topicrule-locationaction-latitude
	//
	Latitude *string `field:"required" json:"latitude" yaml:"latitude"`
	// A string that evaluates to a double value that represents the longitude of the device's location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-locationaction.html#cfn-iot-topicrule-locationaction-longitude
	//
	Longitude *string `field:"required" json:"longitude" yaml:"longitude"`
	// The IAM role that grants permission to write to the Amazon Location resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-locationaction.html#cfn-iot-topicrule-locationaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the tracker resource in Amazon Location in which the location is updated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-locationaction.html#cfn-iot-topicrule-locationaction-trackername
	//
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
	// The time that the location data was sampled.
	//
	// The default value is the time the MQTT message was processed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-locationaction.html#cfn-iot-topicrule-locationaction-timestamp
	//
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
}


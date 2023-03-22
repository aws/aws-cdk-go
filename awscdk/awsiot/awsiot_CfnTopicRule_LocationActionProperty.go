package awsiot


// Describes an action to send device location updates from an MQTT message to an Amazon Location tracker resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationActionProperty := &locationActionProperty{
//   	deviceId: jsii.String("deviceId"),
//   	latitude: jsii.String("latitude"),
//   	longitude: jsii.String("longitude"),
//   	roleArn: jsii.String("roleArn"),
//   	trackerName: jsii.String("trackerName"),
//
//   	// the properties below are optional
//   	timestamp: NewDate(),
//   }
//
type CfnTopicRule_LocationActionProperty struct {
	// The unique ID of the device providing the location data.
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// A string that evaluates to a double value that represents the latitude of the device's location.
	Latitude *string `field:"required" json:"latitude" yaml:"latitude"`
	// A string that evaluates to a double value that represents the longitude of the device's location.
	Longitude *string `field:"required" json:"longitude" yaml:"longitude"`
	// The IAM role that grants permission to write to the Amazon Location resource.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the tracker resource in Amazon Location in which the location is updated.
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
	// The time that the location data was sampled.
	//
	// The default value is the time the MQTT message was processed.
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
}


package awsiot


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
	// `CfnTopicRule.LocationActionProperty.DeviceId`.
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// `CfnTopicRule.LocationActionProperty.Latitude`.
	Latitude *string `field:"required" json:"latitude" yaml:"latitude"`
	// `CfnTopicRule.LocationActionProperty.Longitude`.
	Longitude *string `field:"required" json:"longitude" yaml:"longitude"`
	// `CfnTopicRule.LocationActionProperty.RoleArn`.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// `CfnTopicRule.LocationActionProperty.TrackerName`.
	TrackerName *string `field:"required" json:"trackerName" yaml:"trackerName"`
	// `CfnTopicRule.LocationActionProperty.Timestamp`.
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
}


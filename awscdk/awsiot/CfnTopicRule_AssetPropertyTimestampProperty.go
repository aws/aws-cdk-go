package awsiot


// An asset property timestamp entry containing the following information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyTimestampProperty := &AssetPropertyTimestampProperty{
//   	TimeInSeconds: jsii.String("timeInSeconds"),
//
//   	// the properties below are optional
//   	OffsetInNanos: jsii.String("offsetInNanos"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertytimestamp.html
//
type CfnTopicRule_AssetPropertyTimestampProperty struct {
	// A string that contains the time in seconds since epoch.
	//
	// Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertytimestamp.html#cfn-iot-topicrule-assetpropertytimestamp-timeinseconds
	//
	TimeInSeconds *string `field:"required" json:"timeInSeconds" yaml:"timeInSeconds"`
	// Optional.
	//
	// A string that contains the nanosecond time offset. Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertytimestamp.html#cfn-iot-topicrule-assetpropertytimestamp-offsetinnanos
	//
	OffsetInNanos *string `field:"optional" json:"offsetInNanos" yaml:"offsetInNanos"`
}


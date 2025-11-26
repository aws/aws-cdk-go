package previewawsiotmixins


// An asset property timestamp entry containing the following information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   assetPropertyTimestampProperty := &AssetPropertyTimestampProperty{
//   	OffsetInNanos: jsii.String("offsetInNanos"),
//   	TimeInSeconds: jsii.String("timeInSeconds"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertytimestamp.html
//
type CfnTopicRulePropsMixin_AssetPropertyTimestampProperty struct {
	// Optional.
	//
	// A string that contains the nanosecond time offset. Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertytimestamp.html#cfn-iot-topicrule-assetpropertytimestamp-offsetinnanos
	//
	OffsetInNanos *string `field:"optional" json:"offsetInNanos" yaml:"offsetInNanos"`
	// A string that contains the time in seconds since epoch.
	//
	// Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-assetpropertytimestamp.html#cfn-iot-topicrule-assetpropertytimestamp-timeinseconds
	//
	TimeInSeconds *string `field:"optional" json:"timeInSeconds" yaml:"timeInSeconds"`
}


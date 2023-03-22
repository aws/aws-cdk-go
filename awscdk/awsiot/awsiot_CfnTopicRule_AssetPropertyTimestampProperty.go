package awsiot


// An asset property timestamp entry containing the following information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyTimestampProperty := &assetPropertyTimestampProperty{
//   	timeInSeconds: jsii.String("timeInSeconds"),
//
//   	// the properties below are optional
//   	offsetInNanos: jsii.String("offsetInNanos"),
//   }
//
type CfnTopicRule_AssetPropertyTimestampProperty struct {
	// A string that contains the time in seconds since epoch.
	//
	// Accepts substitution templates.
	TimeInSeconds *string `field:"required" json:"timeInSeconds" yaml:"timeInSeconds"`
	// Optional.
	//
	// A string that contains the nanosecond time offset. Accepts substitution templates.
	OffsetInNanos *string `field:"optional" json:"offsetInNanos" yaml:"offsetInNanos"`
}


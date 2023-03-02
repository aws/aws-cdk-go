package awsiot


// An asset property value entry containing the following information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyValueProperty := &assetPropertyValueProperty{
//   	timestamp: &assetPropertyTimestampProperty{
//   		timeInSeconds: jsii.String("timeInSeconds"),
//
//   		// the properties below are optional
//   		offsetInNanos: jsii.String("offsetInNanos"),
//   	},
//   	value: &assetPropertyVariantProperty{
//   		booleanValue: jsii.String("booleanValue"),
//   		doubleValue: jsii.String("doubleValue"),
//   		integerValue: jsii.String("integerValue"),
//   		stringValue: jsii.String("stringValue"),
//   	},
//
//   	// the properties below are optional
//   	quality: jsii.String("quality"),
//   }
//
type CfnTopicRule_AssetPropertyValueProperty struct {
	// The asset property value timestamp.
	Timestamp interface{} `field:"required" json:"timestamp" yaml:"timestamp"`
	// The value of the asset property.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// Optional.
	//
	// A string that describes the quality of the value. Accepts substitution templates. Must be `GOOD` , `BAD` , or `UNCERTAIN` .
	Quality *string `field:"optional" json:"quality" yaml:"quality"`
}


package awsiotevents


// A structure that contains value information. For more information, see [AssetPropertyValue](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_AssetPropertyValue.html) in the *AWS IoT SiteWise API Reference* .
//
// You must use expressions for all parameters in `AssetPropertyValue` . The expressions accept literals, operators, functions, references, and substitution templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `quality` parameter can be `'GOOD'` .
// - For references, you must specify either variables or input values. For example, the value for the `quality` parameter can be `$input.TemperatureInput.sensorData.quality` .
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyValueProperty := &assetPropertyValueProperty{
//   	value: &assetPropertyVariantProperty{
//   		booleanValue: jsii.String("booleanValue"),
//   		doubleValue: jsii.String("doubleValue"),
//   		integerValue: jsii.String("integerValue"),
//   		stringValue: jsii.String("stringValue"),
//   	},
//
//   	// the properties below are optional
//   	quality: jsii.String("quality"),
//   	timestamp: &assetPropertyTimestampProperty{
//   		timeInSeconds: jsii.String("timeInSeconds"),
//
//   		// the properties below are optional
//   		offsetInNanos: jsii.String("offsetInNanos"),
//   	},
//   }
//
type CfnAlarmModel_AssetPropertyValueProperty struct {
	// The value to send to an asset property.
	Value interface{} `field:"required" json:"value" yaml:"value"`
	// The quality of the asset property value.
	//
	// The value must be `'GOOD'` , `'BAD'` , or `'UNCERTAIN'` .
	Quality *string `field:"optional" json:"quality" yaml:"quality"`
	// The timestamp associated with the asset property value.
	//
	// The default is the current event time.
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
}


package previewawsioteventsmixins


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
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   assetPropertyValueProperty := &AssetPropertyValueProperty{
//   	Quality: jsii.String("quality"),
//   	Timestamp: &AssetPropertyTimestampProperty{
//   		OffsetInNanos: jsii.String("offsetInNanos"),
//   		TimeInSeconds: jsii.String("timeInSeconds"),
//   	},
//   	Value: &AssetPropertyVariantProperty{
//   		BooleanValue: jsii.String("booleanValue"),
//   		DoubleValue: jsii.String("doubleValue"),
//   		IntegerValue: jsii.String("integerValue"),
//   		StringValue: jsii.String("stringValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-assetpropertyvalue.html
//
type CfnAlarmModelPropsMixin_AssetPropertyValueProperty struct {
	// The quality of the asset property value.
	//
	// The value must be `'GOOD'` , `'BAD'` , or `'UNCERTAIN'` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-assetpropertyvalue.html#cfn-iotevents-alarmmodel-assetpropertyvalue-quality
	//
	Quality *string `field:"optional" json:"quality" yaml:"quality"`
	// The timestamp associated with the asset property value.
	//
	// The default is the current event time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-assetpropertyvalue.html#cfn-iotevents-alarmmodel-assetpropertyvalue-timestamp
	//
	Timestamp interface{} `field:"optional" json:"timestamp" yaml:"timestamp"`
	// The value to send to an asset property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-assetpropertyvalue.html#cfn-iotevents-alarmmodel-assetpropertyvalue-value
	//
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}


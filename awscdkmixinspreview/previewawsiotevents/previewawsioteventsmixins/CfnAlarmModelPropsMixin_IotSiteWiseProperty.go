package previewawsioteventsmixins


// Sends information about the detector model instance and the event that triggered the action to a specified asset property in AWS IoT SiteWise .
//
// You must use expressions for all parameters in `IotSiteWiseAction` . The expressions accept literals, operators, functions, references, and substitutions templates.
//
// **Examples** - For literal values, the expressions must contain single quotes. For example, the value for the `propertyAlias` parameter can be `'/company/windfarm/3/turbine/7/temperature'` .
// - For references, you must specify either variables or input values. For example, the value for the `assetId` parameter can be `$input.TurbineInput.assetId1` .
// - For a substitution template, you must use `${}` , and the template must be in single quotes. A substitution template can also contain a combination of literals, operators, functions, references, and substitution templates.
//
// In the following example, the value for the `propertyAlias` parameter uses a substitution template.
//
// `'company/windfarm/${$input.TemperatureInput.sensorData.windfarmID}/turbine/ ${$input.TemperatureInput.sensorData.turbineID}/temperature'`
//
// You must specify either `propertyAlias` or both `assetId` and `propertyId` to identify the target asset property in AWS IoT SiteWise .
//
// For more information, see [Expressions](https://docs.aws.amazon.com/iotevents/latest/developerguide/iotevents-expressions.html) in the *AWS IoT Events Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iotSiteWiseProperty := &IotSiteWiseProperty{
//   	AssetId: jsii.String("assetId"),
//   	EntryId: jsii.String("entryId"),
//   	PropertyAlias: jsii.String("propertyAlias"),
//   	PropertyId: jsii.String("propertyId"),
//   	PropertyValue: &AssetPropertyValueProperty{
//   		Quality: jsii.String("quality"),
//   		Timestamp: &AssetPropertyTimestampProperty{
//   			OffsetInNanos: jsii.String("offsetInNanos"),
//   			TimeInSeconds: jsii.String("timeInSeconds"),
//   		},
//   		Value: &AssetPropertyVariantProperty{
//   			BooleanValue: jsii.String("booleanValue"),
//   			DoubleValue: jsii.String("doubleValue"),
//   			IntegerValue: jsii.String("integerValue"),
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html
//
type CfnAlarmModelPropsMixin_IotSiteWiseProperty struct {
	// The ID of the asset that has the specified property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html#cfn-iotevents-alarmmodel-iotsitewise-assetid
	//
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// A unique identifier for this entry.
	//
	// You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html#cfn-iotevents-alarmmodel-iotsitewise-entryid
	//
	EntryId *string `field:"optional" json:"entryId" yaml:"entryId"`
	// The alias of the asset property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html#cfn-iotevents-alarmmodel-iotsitewise-propertyalias
	//
	PropertyAlias *string `field:"optional" json:"propertyAlias" yaml:"propertyAlias"`
	// The ID of the asset property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html#cfn-iotevents-alarmmodel-iotsitewise-propertyid
	//
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
	// The value to send to the asset property.
	//
	// This value contains timestamp, quality, and value (TQV) information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-alarmmodel-iotsitewise.html#cfn-iotevents-alarmmodel-iotsitewise-propertyvalue
	//
	PropertyValue interface{} `field:"optional" json:"propertyValue" yaml:"propertyValue"`
}


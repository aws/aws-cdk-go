package awsiotevents


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
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotSiteWiseProperty := &iotSiteWiseProperty{
//   	assetId: jsii.String("assetId"),
//   	entryId: jsii.String("entryId"),
//   	propertyAlias: jsii.String("propertyAlias"),
//   	propertyId: jsii.String("propertyId"),
//   	propertyValue: &assetPropertyValueProperty{
//   		value: &assetPropertyVariantProperty{
//   			booleanValue: jsii.String("booleanValue"),
//   			doubleValue: jsii.String("doubleValue"),
//   			integerValue: jsii.String("integerValue"),
//   			stringValue: jsii.String("stringValue"),
//   		},
//
//   		// the properties below are optional
//   		quality: jsii.String("quality"),
//   		timestamp: &assetPropertyTimestampProperty{
//   			timeInSeconds: jsii.String("timeInSeconds"),
//
//   			// the properties below are optional
//   			offsetInNanos: jsii.String("offsetInNanos"),
//   		},
//   	},
//   }
//
type CfnAlarmModel_IotSiteWiseProperty struct {
	// The ID of the asset that has the specified property.
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// A unique identifier for this entry.
	//
	// You can use the entry ID to track which data entry causes an error in case of failure. The default is a new unique identifier.
	EntryId *string `field:"optional" json:"entryId" yaml:"entryId"`
	// The alias of the asset property.
	PropertyAlias *string `field:"optional" json:"propertyAlias" yaml:"propertyAlias"`
	// The ID of the asset property.
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
	// The value to send to the asset property.
	//
	// This value contains timestamp, quality, and value (TQV) information.
	PropertyValue interface{} `field:"optional" json:"propertyValue" yaml:"propertyValue"`
}


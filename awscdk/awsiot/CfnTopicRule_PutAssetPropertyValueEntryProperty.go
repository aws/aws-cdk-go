package awsiot


// An asset property value entry containing the following information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   putAssetPropertyValueEntryProperty := &PutAssetPropertyValueEntryProperty{
//   	PropertyValues: []interface{}{
//   		&AssetPropertyValueProperty{
//   			Timestamp: &AssetPropertyTimestampProperty{
//   				TimeInSeconds: jsii.String("timeInSeconds"),
//
//   				// the properties below are optional
//   				OffsetInNanos: jsii.String("offsetInNanos"),
//   			},
//   			Value: &AssetPropertyVariantProperty{
//   				BooleanValue: jsii.String("booleanValue"),
//   				DoubleValue: jsii.String("doubleValue"),
//   				IntegerValue: jsii.String("integerValue"),
//   				StringValue: jsii.String("stringValue"),
//   			},
//
//   			// the properties below are optional
//   			Quality: jsii.String("quality"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AssetId: jsii.String("assetId"),
//   	EntryId: jsii.String("entryId"),
//   	PropertyAlias: jsii.String("propertyAlias"),
//   	PropertyId: jsii.String("propertyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putassetpropertyvalueentry.html
//
type CfnTopicRule_PutAssetPropertyValueEntryProperty struct {
	// A list of property values to insert that each contain timestamp, quality, and value (TQV) information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putassetpropertyvalueentry.html#cfn-iot-topicrule-putassetpropertyvalueentry-propertyvalues
	//
	PropertyValues interface{} `field:"required" json:"propertyValues" yaml:"propertyValues"`
	// The ID of the AWS IoT SiteWise asset.
	//
	// You must specify either a `propertyAlias` or both an `aliasId` and a `propertyId` . Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putassetpropertyvalueentry.html#cfn-iot-topicrule-putassetpropertyvalueentry-assetid
	//
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// Optional.
	//
	// A unique identifier for this entry that you can define to better track which message caused an error in case of failure. Accepts substitution templates. Defaults to a new UUID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putassetpropertyvalueentry.html#cfn-iot-topicrule-putassetpropertyvalueentry-entryid
	//
	EntryId *string `field:"optional" json:"entryId" yaml:"entryId"`
	// The name of the property alias associated with your asset property.
	//
	// You must specify either a `propertyAlias` or both an `aliasId` and a `propertyId` . Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putassetpropertyvalueentry.html#cfn-iot-topicrule-putassetpropertyvalueentry-propertyalias
	//
	PropertyAlias *string `field:"optional" json:"propertyAlias" yaml:"propertyAlias"`
	// The ID of the asset's property.
	//
	// You must specify either a `propertyAlias` or both an `aliasId` and a `propertyId` . Accepts substitution templates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-putassetpropertyvalueentry.html#cfn-iot-topicrule-putassetpropertyvalueentry-propertyid
	//
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
}


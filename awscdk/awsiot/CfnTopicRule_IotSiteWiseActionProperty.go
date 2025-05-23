package awsiot


// Describes an action to send data from an MQTT message that triggered the rule to AWS IoT SiteWise asset properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotSiteWiseActionProperty := &IotSiteWiseActionProperty{
//   	PutAssetPropertyValueEntries: []interface{}{
//   		&PutAssetPropertyValueEntryProperty{
//   			PropertyValues: []interface{}{
//   				&AssetPropertyValueProperty{
//   					Timestamp: &AssetPropertyTimestampProperty{
//   						TimeInSeconds: jsii.String("timeInSeconds"),
//
//   						// the properties below are optional
//   						OffsetInNanos: jsii.String("offsetInNanos"),
//   					},
//   					Value: &AssetPropertyVariantProperty{
//   						BooleanValue: jsii.String("booleanValue"),
//   						DoubleValue: jsii.String("doubleValue"),
//   						IntegerValue: jsii.String("integerValue"),
//   						StringValue: jsii.String("stringValue"),
//   					},
//
//   					// the properties below are optional
//   					Quality: jsii.String("quality"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			AssetId: jsii.String("assetId"),
//   			EntryId: jsii.String("entryId"),
//   			PropertyAlias: jsii.String("propertyAlias"),
//   			PropertyId: jsii.String("propertyId"),
//   		},
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotsitewiseaction.html
//
type CfnTopicRule_IotSiteWiseActionProperty struct {
	// A list of asset property value entries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotsitewiseaction.html#cfn-iot-topicrule-iotsitewiseaction-putassetpropertyvalueentries
	//
	PutAssetPropertyValueEntries interface{} `field:"required" json:"putAssetPropertyValueEntries" yaml:"putAssetPropertyValueEntries"`
	// The ARN of the role that grants AWS IoT permission to send an asset property value to AWS IoT SiteWise.
	//
	// ( `"Action": "iotsitewise:BatchPutAssetPropertyValue"` ). The trust policy can restrict access to specific asset hierarchy paths.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-iotsitewiseaction.html#cfn-iot-topicrule-iotsitewiseaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}


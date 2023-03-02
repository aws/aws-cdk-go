package awsiot


// Describes an action to send data from an MQTT message that triggered the rule to AWS IoT SiteWise asset properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iotSiteWiseActionProperty := &iotSiteWiseActionProperty{
//   	putAssetPropertyValueEntries: []interface{}{
//   		&putAssetPropertyValueEntryProperty{
//   			propertyValues: []interface{}{
//   				&assetPropertyValueProperty{
//   					timestamp: &assetPropertyTimestampProperty{
//   						timeInSeconds: jsii.String("timeInSeconds"),
//
//   						// the properties below are optional
//   						offsetInNanos: jsii.String("offsetInNanos"),
//   					},
//   					value: &assetPropertyVariantProperty{
//   						booleanValue: jsii.String("booleanValue"),
//   						doubleValue: jsii.String("doubleValue"),
//   						integerValue: jsii.String("integerValue"),
//   						stringValue: jsii.String("stringValue"),
//   					},
//
//   					// the properties below are optional
//   					quality: jsii.String("quality"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			assetId: jsii.String("assetId"),
//   			entryId: jsii.String("entryId"),
//   			propertyAlias: jsii.String("propertyAlias"),
//   			propertyId: jsii.String("propertyId"),
//   		},
//   	},
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnTopicRule_IotSiteWiseActionProperty struct {
	// A list of asset property value entries.
	PutAssetPropertyValueEntries interface{} `field:"required" json:"putAssetPropertyValueEntries" yaml:"putAssetPropertyValueEntries"`
	// The ARN of the role that grants AWS IoT permission to send an asset property value to AWS IoT SiteWise.
	//
	// ( `"Action": "iotsitewise:BatchPutAssetPropertyValue"` ). The trust policy can restrict access to specific asset hierarchy paths.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}


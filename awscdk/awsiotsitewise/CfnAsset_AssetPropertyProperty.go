package awsiotsitewise


// Contains asset property information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyProperty := &AssetPropertyProperty{
//   	Alias: jsii.String("alias"),
//   	ExternalId: jsii.String("externalId"),
//   	Id: jsii.String("id"),
//   	LogicalId: jsii.String("logicalId"),
//   	NotificationState: jsii.String("notificationState"),
//   	Unit: jsii.String("unit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assetproperty.html
//
type CfnAsset_AssetPropertyProperty struct {
	// The property alias that identifies the property, such as an OPC-UA server data stream path (for example, `/company/windfarm/3/turbine/7/temperature` ).
	//
	// For more information, see [Mapping industrial data streams to asset properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html) in the *AWS IoT SiteWise User Guide* .
	//
	// The property alias must have 1-1000 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assetproperty.html#cfn-iotsitewise-asset-assetproperty-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// String-friendly customer provided external ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assetproperty.html#cfn-iotsitewise-asset-assetproperty-externalid
	//
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Customer provided actual UUID for property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assetproperty.html#cfn-iotsitewise-asset-assetproperty-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The `LogicalID` of the asset property.
	//
	// The maximum length is 256 characters, with the pattern `[^\u0000-\u001F\u007F]+` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assetproperty.html#cfn-iotsitewise-asset-assetproperty-logicalid
	//
	LogicalId *string `field:"optional" json:"logicalId" yaml:"logicalId"`
	// The MQTT notification state ( `ENABLED` or `DISABLED` ) for this asset property.
	//
	// When the notification state is `ENABLED` , AWS IoT SiteWise publishes property value updates to a unique MQTT topic. For more information, see [Interacting with other services](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/interact-with-other-services.html) in the *AWS IoT SiteWise User Guide* .
	//
	// If you omit this parameter, the notification state is set to `DISABLED` .
	//
	// > You must use all caps for the NotificationState parameter. If you use lower case letters, you will receive a schema validation error.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assetproperty.html#cfn-iotsitewise-asset-assetproperty-notificationstate
	//
	NotificationState *string `field:"optional" json:"notificationState" yaml:"notificationState"`
	// The unit (such as `Newtons` or `RPM` ) of the asset property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-asset-assetproperty.html#cfn-iotsitewise-asset-assetproperty-unit
	//
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}


package awsiotsitewise


// Contains asset property information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assetPropertyProperty := &assetPropertyProperty{
//   	logicalId: jsii.String("logicalId"),
//
//   	// the properties below are optional
//   	alias: jsii.String("alias"),
//   	notificationState: jsii.String("notificationState"),
//   	unit: jsii.String("unit"),
//   }
//
type CfnAsset_AssetPropertyProperty struct {
	// The `LogicalID` of the asset property.
	//
	// The maximum length is 256 characters, with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	LogicalId *string `field:"required" json:"logicalId" yaml:"logicalId"`
	// The property alias that identifies the property, such as an OPC-UA server data stream path (for example, `/company/windfarm/3/turbine/7/temperature` ).
	//
	// For more information, see [Mapping industrial data streams to asset properties](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/connect-data-streams.html) in the *AWS IoT SiteWise User Guide* .
	//
	// The property alias must have 1-1000 characters.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The MQTT notification state (enabled or disabled) for this asset property.
	//
	// When the notification state is enabled, AWS IoT SiteWise publishes property value updates to a unique MQTT topic. For more information, see [Interacting with other services](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/interact-with-other-services.html) in the *AWS IoT SiteWise User Guide* .
	//
	// If you omit this parameter, the notification state is set to `DISABLED` .
	NotificationState *string `field:"optional" json:"notificationState" yaml:"notificationState"`
	// `CfnAsset.AssetPropertyProperty.Unit`.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}


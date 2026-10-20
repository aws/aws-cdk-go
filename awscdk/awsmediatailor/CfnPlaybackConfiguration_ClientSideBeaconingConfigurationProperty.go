package awsmediatailor


// The beaconing settings that apply to client-side reporting sessions: whether AWS Elemental MediaTailor includes its beacons in the ad tracking response, and which player operation events it reports on.
//
// If you omit this object, MediaTailor uses INSIGHTS reporting mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientSideBeaconingConfigurationProperty := &ClientSideBeaconingConfigurationProperty{
//   	ReportingMode: jsii.String("reportingMode"),
//
//   	// the properties below are optional
//   	AdditionalEventTypes: []*string{
//   		jsii.String("additionalEventTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-clientsidebeaconingconfiguration.html
//
type CfnPlaybackConfiguration_ClientSideBeaconingConfigurationProperty struct {
	// Specifies whether MediaTailor includes its beacons in the ad tracking response.
	//
	// INSIGHTS includes MediaTailor beacons; DISABLED does not. Values are case-sensitive. This setting is required when ClientSide is present. If BeaconingConfiguration or ClientSide is omitted, MediaTailor uses INSIGHTS. To keep beaconing off, retain DISABLED in subsequent updates.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-clientsidebeaconingconfiguration.html#cfn-mediatailor-playbackconfiguration-clientsidebeaconingconfiguration-reportingmode
	//
	ReportingMode *string `field:"required" json:"reportingMode" yaml:"reportingMode"`
	// The player operation events to report on, in addition to the ad progress events that MediaTailor always reports on.
	//
	// Valid values are MUTE, UNMUTE, PAUSE, and SKIP. The default is an empty list. This parameter is valid only when ReportingMode is INSIGHTS. MediaTailor rejects the request if you specify a value while ReportingMode is DISABLED, or if you specify duplicate values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-clientsidebeaconingconfiguration.html#cfn-mediatailor-playbackconfiguration-clientsidebeaconingconfiguration-additionaleventtypes
	//
	AdditionalEventTypes *[]*string `field:"optional" json:"additionalEventTypes" yaml:"additionalEventTypes"`
}


package awsmediatailor


// The beaconing configuration for a playback configuration.
//
// Beaconing controls whether AWS Elemental MediaTailor includes its own beacons in the ad tracking response, in addition to the ad server beacons. If you omit this setting, MediaTailor uses INSIGHTS.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   beaconingConfigurationProperty := &BeaconingConfigurationProperty{
//   	ClientSide: &ClientSideBeaconingConfigurationProperty{
//   		AdditionalEventTypes: []*string{
//   			jsii.String("additionalEventTypes"),
//   		},
//   		ReportingMode: jsii.String("reportingMode"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-beaconingconfiguration.html
//
type CfnPlaybackConfigurationPropsMixin_BeaconingConfigurationProperty struct {
	// The beaconing settings that apply to client-side reporting sessions: whether AWS Elemental MediaTailor includes its beacons in the ad tracking response, and which player operation events it reports on.
	//
	// If you omit this object, MediaTailor uses INSIGHTS reporting mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-beaconingconfiguration.html#cfn-mediatailor-playbackconfiguration-beaconingconfiguration-clientside
	//
	ClientSide interface{} `field:"optional" json:"clientSide" yaml:"clientSide"`
}


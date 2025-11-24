package mixinsawsmediatailor


// Settings for customizing what events are included in logs for interactions with the origin server.
//
// For more information about manifest service logs, including descriptions of the event types, see [MediaTailor manifest logs description and event types](https://docs.aws.amazon.com/mediatailor/latest/ug/log-types.html) in AWS Elemental MediaTailor User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   manifestServiceInteractionLogProperty := &ManifestServiceInteractionLogProperty{
//   	ExcludeEventTypes: []*string{
//   		jsii.String("excludeEventTypes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-manifestserviceinteractionlog.html
//
type CfnPlaybackConfigurationPropsMixin_ManifestServiceInteractionLogProperty struct {
	// Indicates that MediaTailor won't emit the selected events in the logs for playback sessions that are initialized with this configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-manifestserviceinteractionlog.html#cfn-mediatailor-playbackconfiguration-manifestserviceinteractionlog-excludeeventtypes
	//
	ExcludeEventTypes *[]*string `field:"optional" json:"excludeEventTypes" yaml:"excludeEventTypes"`
}


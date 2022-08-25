package awsmediatailor


// The configuration for avail suppression, also known as ad suppression.
//
// For more information about ad suppression, see [Ad Suppression](https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availSuppressionProperty := &availSuppressionProperty{
//   	mode: jsii.String("mode"),
//   	value: jsii.String("value"),
//   }
//
type CfnPlaybackConfiguration_AvailSuppressionProperty struct {
	// Sets the ad suppression mode.
	//
	// By default, ad suppression is off and all ad breaks are filled with ads or slate. When Mode is set to BEHIND_LIVE_EDGE, ad suppression is active and MediaTailor won't fill ad breaks on or behind the ad suppression Value time in the manifest lookback window.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// A live edge offset time in HH:MM:SS.
	//
	// MediaTailor won't fill ad breaks on or behind this time in the manifest lookback window. If Value is set to 00:00:00, it is in sync with the live edge, and MediaTailor won't fill any ad breaks on or behind the live edge. If you set a Value time, MediaTailor won't fill any ad breaks on or behind this time in the manifest lookback window. For example, if you set 00:45:00, then MediaTailor will fill ad breaks that occur within 45 minutes behind the live edge, but won't fill ad breaks on or behind 45 minutes behind the live edge.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


package awsmediatailor


// The configuration for manifest processing rules.
//
// Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   manifestProcessingRulesProperty := &manifestProcessingRulesProperty{
//   	adMarkerPassthrough: &adMarkerPassthroughProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnPlaybackConfiguration_ManifestProcessingRulesProperty struct {
	// For HLS, when set to `true` , MediaTailor passes through `EXT-X-CUE-IN` , `EXT-X-CUE-OUT` , and `EXT-X-SPLICEPOINT-SCTE35` ad markers from the origin manifest to the MediaTailor personalized manifest.
	//
	// No logic is applied to these ad markers. For example, if `EXT-X-CUE-OUT` has a value of `60` , but no ads are filled for that ad break, MediaTailor will not set the value to `0` .
	AdMarkerPassthrough interface{} `field:"optional" json:"adMarkerPassthrough" yaml:"adMarkerPassthrough"`
}


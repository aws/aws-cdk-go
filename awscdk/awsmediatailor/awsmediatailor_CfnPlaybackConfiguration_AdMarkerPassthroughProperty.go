package awsmediatailor


// For HLS, when set to `true` , MediaTailor passes through `EXT-X-CUE-IN` , `EXT-X-CUE-OUT` , and `EXT-X-SPLICEPOINT-SCTE35` ad markers from the origin manifest to the MediaTailor personalized manifest.
//
// No logic is applied to these ad markers. For example, if `EXT-X-CUE-OUT` has a value of `60` , but no ads are filled for that ad break, MediaTailor will not set the value to `0` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adMarkerPassthroughProperty := &adMarkerPassthroughProperty{
//   	enabled: jsii.Boolean(false),
//   }
//
type CfnPlaybackConfiguration_AdMarkerPassthroughProperty struct {
	// Enables ad marker passthrough for your configuration.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}


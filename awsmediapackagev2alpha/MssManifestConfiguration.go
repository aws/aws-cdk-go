package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The MSS manifest configuration associated with the origin endpoint.
//
// Example:
//   var channel Channel
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Ism(),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Mss(&MssManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   			ManifestWindow: awscdk.Duration_Seconds(jsii.Number(60)),
//   			ManifestLayout: awsmediapackagev2alpha.MssManifestLayout_COMPACT,
//   		}),
//   	},
//   })
//
// Experimental.
type MssManifestConfiguration struct {
	// The name of the manifest associated with the MSS manifest configuration.
	// Experimental.
	ManifestName *string `field:"required" json:"manifestName" yaml:"manifestName"`
	// Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	//
	// https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filter-query-parameters.html
	// Default: - No filter configuration.
	//
	// Experimental.
	FilterConfiguration *FilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// The layout of the MSS manifest.
	// Default: MssManifestLayout.FULL
	//
	// Experimental.
	ManifestLayout MssManifestLayout `field:"optional" json:"manifestLayout" yaml:"manifestLayout"`
	// The total duration (in seconds) of the manifest's content.
	// Default: 60.
	//
	// Experimental.
	ManifestWindow awscdk.Duration `field:"optional" json:"manifestWindow" yaml:"manifestWindow"`
}


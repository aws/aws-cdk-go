package awsmediapackagev2alpha


// Choose how SCTE-35 ad markers are included in HLS and LL-HLS manifests.
//
// Example:
//   var channel Channel
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   			ManifestWindow: awscdk.Duration_Seconds(jsii.Number(60)),
//   			ProgramDateTimeInterval: awscdk.Duration_*Seconds(jsii.Number(60)),
//   			ScteAdMarkerHls: awsmediapackagev2alpha.AdMarkerHls_DATERANGE,
//   		}),
//   	},
//   })
//
// Experimental.
type AdMarkerHls string

const (
	// Insert EXT-X-DATERANGE tags to signal ad content.
	// Experimental.
	AdMarkerHls_DATERANGE AdMarkerHls = "DATERANGE"
	// Insert enhanced SCTE-35 tags for ad content.
	// Experimental.
	AdMarkerHls_SCTE35_ENHANCED AdMarkerHls = "SCTE35_ENHANCED"
)


package awsmediapackagev2alpha


// The layout of the MSS manifest.
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
type MssManifestLayout string

const (
	// Full manifest layout.
	// Experimental.
	MssManifestLayout_FULL MssManifestLayout = "FULL"
	// Compact manifest layout.
	// Experimental.
	MssManifestLayout_COMPACT MssManifestLayout = "COMPACT"
)


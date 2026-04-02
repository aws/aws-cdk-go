package awsmediapackagev2alpha


// Encryption methods for TS container type.
//
// Example:
//   var channel Channel
//   var spekeRole IRole
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("TsEndpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Ts(&TsSegmentProps{
//   		Encryption: awsmediapackagev2alpha.TsEncryption_Speke(&TsSpekeEncryptionProps{
//   			Method: awsmediapackagev2alpha.TsEncryptionMethod_SAMPLE_AES,
//   			ResourceId: jsii.String("my-content-id"),
//   			Url: jsii.String("https://example.com/speke"),
//   			Role: spekeRole,
//   		}),
//   	}),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   })
//
// Experimental.
type TsEncryptionMethod string

const (
	// AES-128 encryption - requires Clear Key AES 128 DRM system.
	// Experimental.
	TsEncryptionMethod_AES_128 TsEncryptionMethod = "AES_128"
	// Sample-level AES encryption - requires FairPlay DRM system.
	// Experimental.
	TsEncryptionMethod_SAMPLE_AES TsEncryptionMethod = "SAMPLE_AES"
)


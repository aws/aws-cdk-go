package awsmediapackagev2alpha


// Encryption methods for CMAF container type.
//
// Example:
//   var channel Channel
//   var spekeRole IRole
//   var certificate ICertificate
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(&CmafSegmentProps{
//   		Encryption: awsmediapackagev2alpha.CmafEncryption_Speke(&CmafSpekeEncryptionProps{
//   			Method: awsmediapackagev2alpha.CmafEncryptionMethod_CBCS,
//   			DrmSystems: []CmafDrmSystem{
//   				awsmediapackagev2alpha.CmafDrmSystem_FAIRPLAY,
//   			},
//   			ResourceId: jsii.String("my-content-id"),
//   			Url: jsii.String("https://example.com/speke"),
//   			Role: spekeRole,
//   			Certificate: *Certificate,
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
type CmafEncryptionMethod string

const (
	// Common Encryption Scheme (CENC) - compatible with PlayReady, Widevine, and Irdeto DRM systems.
	// Experimental.
	CmafEncryptionMethod_CENC CmafEncryptionMethod = "CENC"
	// Common Encryption Scheme with CBCS mode - compatible with PlayReady, Widevine, and FairPlay DRM systems.
	// Experimental.
	CmafEncryptionMethod_CBCS CmafEncryptionMethod = "CBCS"
)


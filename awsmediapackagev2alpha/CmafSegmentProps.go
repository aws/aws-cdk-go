package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CMAF segment configuration.
//
// Example:
//   var channel Channel
//   var spekeRole IRole
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(&CmafSegmentProps{
//   		Encryption: awsmediapackagev2alpha.CmafEncryption_Speke(&CmafSpekeEncryptionProps{
//   			Method: awsmediapackagev2alpha.CmafEncryptionMethod_CBCS,
//   			DrmSystems: []CmafDrmSystem{
//   				awsmediapackagev2alpha.CmafDrmSystem_FAIRPLAY,
//   				awsmediapackagev2alpha.CmafDrmSystem_WIDEVINE,
//   			},
//   			ResourceId: jsii.String("my-content-id"),
//   			Url: jsii.String("https://example.com/speke"),
//   			Role: spekeRole,
//   			KeyRotationInterval: awscdk.Duration_Seconds(jsii.Number(300)),
//   			AudioPreset: awsmediapackagev2alpha.PresetSpeke20Audio_PRESET_AUDIO_2,
//   			VideoPreset: awsmediapackagev2alpha.PresetSpeke20Video_PRESET_VIDEO_2,
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
type CmafSegmentProps struct {
	// Duration of each segment.
	// Default: Duration.seconds(6)
	//
	// Experimental.
	Duration awscdk.Duration `field:"optional" json:"duration" yaml:"duration"`
	// Whether to include I-frame-only streams.
	// Default: false.
	//
	// Experimental.
	IncludeIframeOnlyStreams *bool `field:"optional" json:"includeIframeOnlyStreams" yaml:"includeIframeOnlyStreams"`
	// Name of the segment.
	// Default: 'segment'.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Encryption configuration for the CMAF segment.
	//
	// Use `CmafEncryption.speke()` to create the configuration.
	// Default: - No encryption.
	//
	// Experimental.
	Encryption CmafEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// SCTE-35 message types to treat as ad markers.
	// Default: - no filtering.
	//
	// Experimental.
	ScteFilter *[]ScteMessageType `field:"optional" json:"scteFilter" yaml:"scteFilter"`
	// Controls whether SCTE-35 messages are included in segment files.
	// Default: - SCTE-35 messages are not included in segments.
	//
	// Experimental.
	ScteInSegments ScteInSegments `field:"optional" json:"scteInSegments" yaml:"scteInSegments"`
}


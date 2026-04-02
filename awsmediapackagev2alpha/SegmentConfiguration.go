package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The segment configuration, including the segment name, duration, and other configuration values.
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
type SegmentConfiguration struct {
	// The container type for this segment (TS or CMAF).
	// Experimental.
	ContainerType ContainerType `field:"required" json:"containerType" yaml:"containerType"`
	// Encryption configuration for the segment.
	// Default: - No encryption.
	//
	// Experimental.
	Encryption EncryptionConfiguration `field:"optional" json:"encryption" yaml:"encryption"`
	// Whether the segment includes I-frame-only streams.
	// Default: undefined - Not specified.
	//
	// Experimental.
	IncludeIframeOnlyStreams *bool `field:"optional" json:"includeIframeOnlyStreams" yaml:"includeIframeOnlyStreams"`
	// The SCTE-35 configuration associated with the segment.
	//
	// The SCTE-35 message types that you want to be treated as ad markers in the output.
	// Default: - No SCTE filtering.
	//
	// Experimental.
	ScteFilter *[]ScteMessageType `field:"optional" json:"scteFilter" yaml:"scteFilter"`
	// Controls whether SCTE-35 messages are included in segment files.
	// Default: - SCTE-35 messages are not included in segments.
	//
	// Experimental.
	ScteInSegments ScteInSegments `field:"optional" json:"scteInSegments" yaml:"scteInSegments"`
	// The duration of the segments.
	// Default: 6.
	//
	// Experimental.
	SegmentDuration awscdk.Duration `field:"optional" json:"segmentDuration" yaml:"segmentDuration"`
	// The name of the segment associated with the origin endpoint.
	// Default: segment.
	//
	// Experimental.
	SegmentName *string `field:"optional" json:"segmentName" yaml:"segmentName"`
	// Whether the segment includes DVB subtitles.
	// Default: false.
	//
	// Experimental.
	TsIncludeDvbSubtitles *bool `field:"optional" json:"tsIncludeDvbSubtitles" yaml:"tsIncludeDvbSubtitles"`
	// Whether the segment is an audio rendition group.
	// Default: false.
	//
	// Experimental.
	TsUseAudioRenditionGroup *bool `field:"optional" json:"tsUseAudioRenditionGroup" yaml:"tsUseAudioRenditionGroup"`
}


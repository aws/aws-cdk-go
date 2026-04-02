package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for TS (Transport Stream) segment configuration.
//
// Example:
//   var channel Channel
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("TsEndpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Ts(&TsSegmentProps{
//   		Duration: awscdk.Duration_Seconds(jsii.Number(6)),
//   		Name: jsii.String("segment"),
//   		IncludeDvbSubtitles: jsii.Boolean(true),
//   		UseAudioRenditionGroup: jsii.Boolean(true),
//   		IncludeIframeOnlyStreams: jsii.Boolean(false),
//   		ScteFilter: []ScteMessageType{
//   			awsmediapackagev2alpha.ScteMessageType_BREAK,
//   			awsmediapackagev2alpha.ScteMessageType_DISTRIBUTOR_ADVERTISEMENT,
//   		},
//   	}),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   })
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("CmafEndpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(&CmafSegmentProps{
//   		Duration: awscdk.Duration_*Seconds(jsii.Number(6)),
//   		Name: jsii.String("segment"),
//   		IncludeIframeOnlyStreams: jsii.Boolean(true),
//   		ScteFilter: []ScteMessageType{
//   			awsmediapackagev2alpha.ScteMessageType_DISTRIBUTOR_ADVERTISEMENT,
//   		},
//   	}),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_*Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   })
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_*Cmaf(),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_*Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   })
//
// Experimental.
type TsSegmentProps struct {
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
	// Encryption configuration for the TS segment.
	//
	// Use `TsEncryption.speke()` to create the configuration.
	// Default: - No encryption.
	//
	// Experimental.
	Encryption TsEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Whether to include DVB subtitles.
	// Default: false.
	//
	// Experimental.
	IncludeDvbSubtitles *bool `field:"optional" json:"includeDvbSubtitles" yaml:"includeDvbSubtitles"`
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
	// Whether to use audio rendition groups.
	// Default: false.
	//
	// Experimental.
	UseAudioRenditionGroup *bool `field:"optional" json:"useAudioRenditionGroup" yaml:"useAudioRenditionGroup"`
}


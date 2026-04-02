package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for ISM (Microsoft Smooth Streaming) segment configuration.
//
// Example:
//   var channel Channel
//   var spekeRole IRole
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("IsmEndpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Ism(&IsmSegmentProps{
//   		Encryption: awsmediapackagev2alpha.IsmEncryption_Speke(&IsmSpekeEncryptionProps{
//   			ResourceId: jsii.String("my-content-id"),
//   			Url: jsii.String("https://example.com/speke"),
//   			Role: spekeRole,
//   		}),
//   	}),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Mss(&MssManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   })
//
// Experimental.
type IsmSegmentProps struct {
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
	// Encryption configuration for the ISM segment.
	//
	// Use `IsmEncryption.speke()` to create the configuration.
	// Default: - No encryption.
	//
	// Experimental.
	Encryption IsmEncryption `field:"optional" json:"encryption" yaml:"encryption"`
}


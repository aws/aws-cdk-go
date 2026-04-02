package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Specify a low-latency HTTP live streaming (LL-HLS) manifest configuration.
//
// Example:
//   var channel Channel
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_LowLatencyHLS(&LowLatencyHlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   			ManifestWindow: awscdk.Duration_Seconds(jsii.Number(30)),
//   			ProgramDateTimeInterval: awscdk.Duration_*Seconds(jsii.Number(5)),
//   			ChildManifestName: jsii.String("child"),
//   		}),
//   	},
//   })
//
// Experimental.
type LowLatencyHlsManifestConfiguration struct {
	// A short string that's appended to the endpoint URL.
	//
	// The manifest name creates a unique path to this endpoint.
	// If you don't enter a value, MediaPackage uses the default manifest name, index. MediaPackage automatically inserts the format extension, such as .m3u8.
	// You can't use the same manifest name if you use HLS manifest and low-latency HLS manifest.
	// The manifestName on the HLSManifest object overrides the manifestName you provided on the originEndpoint object.
	// Experimental.
	ManifestName *string `field:"required" json:"manifestName" yaml:"manifestName"`
	// The name of the child manifest associated with the low-latency HLS (LL-HLS) manifest configuration of the origin endpoint.
	// Default: - No child manifest name specified.
	//
	// Experimental.
	ChildManifestName *string `field:"optional" json:"childManifestName" yaml:"childManifestName"`
	// Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	//
	// https://docs.aws.amazon.com/mediapackage/latest/userguide/manifest-filter-query-parameters.html
	// Default: - No filter configuration.
	//
	// Experimental.
	FilterConfiguration *FilterConfiguration `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// The total duration (in seconds) of the manifest's content.
	// Default: 60.
	//
	// Experimental.
	ManifestWindow awscdk.Duration `field:"optional" json:"manifestWindow" yaml:"manifestWindow"`
	// Inserts EXT-X-PROGRAM-DATE-TIME tags in the output manifest at the interval that you specify.
	//
	// If you don't enter an interval, EXT-X-PROGRAM-DATE-TIME tags aren't included in the manifest.
	// The tags sync the stream to the wall clock so that viewers can seek to a specific time in the playback timeline on the player.
	// Default: - No program date time interval.
	//
	// Experimental.
	ProgramDateTimeInterval awscdk.Duration `field:"optional" json:"programDateTimeInterval" yaml:"programDateTimeInterval"`
	// The SCTE-35 HLS configuration associated with the low-latency HLS (LL-HLS) manifest configuration of the origin endpoint.
	// Default: - No SCTE ad marker configuration.
	//
	// Experimental.
	ScteAdMarkerHls AdMarkerHls `field:"optional" json:"scteAdMarkerHls" yaml:"scteAdMarkerHls"`
	// Insert EXT-X-START tag in the manifest with the configured settings.
	// Default: - No start tag.
	//
	// Experimental.
	StartTag StartTag `field:"optional" json:"startTag" yaml:"startTag"`
	// When enabled, MediaPackage URL-encodes the query string for API requests for LL-HLS child manifests to comply with AWS Signature Version 4 (SigV4) signature signing protocol.
	//
	// For more information, see AWS Signature Version 4 for API requests in AWS Identity and Access Management User Guide.
	// Default: false.
	//
	// Experimental.
	UrlEncodeChildManifest *bool `field:"optional" json:"urlEncodeChildManifest" yaml:"urlEncodeChildManifest"`
}


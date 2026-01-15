package awsmediapackagev2


// Specify a low-latency HTTP live streaming (LL-HLS) manifest configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lowLatencyHlsManifestConfigurationProperty := &LowLatencyHlsManifestConfigurationProperty{
//   	ManifestName: jsii.String("manifestName"),
//
//   	// the properties below are optional
//   	ChildManifestName: jsii.String("childManifestName"),
//   	FilterConfiguration: &FilterConfigurationProperty{
//   		ClipStartTime: jsii.String("clipStartTime"),
//   		DrmSettings: jsii.String("drmSettings"),
//   		End: jsii.String("end"),
//   		ManifestFilter: jsii.String("manifestFilter"),
//   		Start: jsii.String("start"),
//   		TimeDelaySeconds: jsii.Number(123),
//   	},
//   	ManifestWindowSeconds: jsii.Number(123),
//   	ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   	ScteHls: &ScteHlsProperty{
//   		AdMarkerHls: jsii.String("adMarkerHls"),
//   	},
//   	StartTag: &StartTagProperty{
//   		TimeOffset: jsii.Number(123),
//
//   		// the properties below are optional
//   		Precise: jsii.Boolean(false),
//   	},
//   	Url: jsii.String("url"),
//   	UrlEncodeChildManifest: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.html
//
type CfnOriginEndpoint_LowLatencyHlsManifestConfigurationProperty struct {
	// A short string that's appended to the endpoint URL.
	//
	// The manifest name creates a unique path to this endpoint. If you don't enter a value, MediaPackage uses the default manifest name, `index` . MediaPackage automatically inserts the format extension, such as `.m3u8` . You can't use the same manifest name if you use HLS manifest and low-latency HLS manifest. The `manifestName` on the `HLSManifest` object overrides the `manifestName` you provided on the `originEndpoint` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration-manifestname
	//
	ManifestName *string `field:"required" json:"manifestName" yaml:"manifestName"`
	// The name of the child manifest associated with the low-latency HLS (LL-HLS) manifest configuration of the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration-childmanifestname
	//
	ChildManifestName *string `field:"optional" json:"childManifestName" yaml:"childManifestName"`
	// Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration-filterconfiguration
	//
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// The total duration (in seconds) of the manifest's content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration-manifestwindowseconds
	//
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// Inserts `EXT-X-PROGRAM-DATE-TIME` tags in the output manifest at the interval that you specify.
	//
	// If you don't enter an interval, `EXT-X-PROGRAM-DATE-TIME` tags aren't included in the manifest. The tags sync the stream to the wall clock so that viewers can seek to a specific time in the playback timeline on the player.
	//
	// Irrespective of this parameter, if any `ID3Timed` metadata is in the HLS input, MediaPackage passes through that metadata to the HLS output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration-programdatetimeintervalseconds
	//
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// The SCTE-35 HLS configuration associated with the low-latency HLS (LL-HLS) manifest configuration of the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration-sctehls
	//
	ScteHls interface{} `field:"optional" json:"scteHls" yaml:"scteHls"`
	// To insert an EXT-X-START tag in your HLS playlist, specify a StartTag configuration object with a valid TimeOffset.
	//
	// When you do, you can also optionally specify whether to include a PRECISE value in the EXT-X-START tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration-starttag
	//
	StartTag interface{} `field:"optional" json:"startTag" yaml:"startTag"`
	// The URL of the low-latency HLS (LL-HLS) manifest configuration of the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
	// When enabled, MediaPackage URL-encodes the query string for API requests for LL-HLS child manifests to comply with AWS Signature Version 4 (SigV4) signature signing protocol.
	//
	// For more information, see [AWS Signature Version 4 for API requests](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_sigv.html) in *AWS Identity and Access Management User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifestconfiguration-urlencodechildmanifest
	//
	UrlEncodeChildManifest interface{} `field:"optional" json:"urlEncodeChildManifest" yaml:"urlEncodeChildManifest"`
}


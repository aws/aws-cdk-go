package awsmediapackagev2


// The HLS manifest configuration associated with the origin endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsManifestConfigurationProperty := &HlsManifestConfigurationProperty{
//   	ManifestName: jsii.String("manifestName"),
//
//   	// the properties below are optional
//   	ChildManifestName: jsii.String("childManifestName"),
//   	FilterConfiguration: &FilterConfigurationProperty{
//   		ClipStartTime: jsii.String("clipStartTime"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.html
//
type CfnOriginEndpoint_HlsManifestConfigurationProperty struct {
	// The name of the manifest associated with the HLS manifest configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-hlsmanifestconfiguration-manifestname
	//
	ManifestName *string `field:"required" json:"manifestName" yaml:"manifestName"`
	// The name of the child manifest associated with the HLS manifest configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-hlsmanifestconfiguration-childmanifestname
	//
	ChildManifestName *string `field:"optional" json:"childManifestName" yaml:"childManifestName"`
	// Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-hlsmanifestconfiguration-filterconfiguration
	//
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// The duration of the manifest window, in seconds, for the HLS manifest configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-hlsmanifestconfiguration-manifestwindowseconds
	//
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
	// The `EXT-X-PROGRAM-DATE-TIME` interval, in seconds, associated with the HLS manifest configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-hlsmanifestconfiguration-programdatetimeintervalseconds
	//
	ProgramDateTimeIntervalSeconds *float64 `field:"optional" json:"programDateTimeIntervalSeconds" yaml:"programDateTimeIntervalSeconds"`
	// THE SCTE-35 HLS configuration associated with the HLS manifest configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-hlsmanifestconfiguration-sctehls
	//
	ScteHls interface{} `field:"optional" json:"scteHls" yaml:"scteHls"`
	// To insert an EXT-X-START tag in your HLS playlist, specify a StartTag configuration object with a valid TimeOffset.
	//
	// When you do, you can also optionally specify whether to include a PRECISE value in the EXT-X-START tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-hlsmanifestconfiguration-starttag
	//
	StartTag interface{} `field:"optional" json:"startTag" yaml:"startTag"`
	// The URL of the HLS manifest configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-hlsmanifestconfiguration-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
	// When enabled, MediaPackage URL-encodes the query string for API requests for HLS child manifests to comply with AWS Signature Version 4 (SigV4) signature signing protocol.
	//
	// For more information, see [AWS Signature Version 4 for API requests](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_sigv.html) in *AWS Identity and Access Management User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-hlsmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-hlsmanifestconfiguration-urlencodechildmanifest
	//
	UrlEncodeChildManifest interface{} `field:"optional" json:"urlEncodeChildManifest" yaml:"urlEncodeChildManifest"`
}


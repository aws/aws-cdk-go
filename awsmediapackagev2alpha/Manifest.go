package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Manifest to add to Origin Endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var manifestFilter ManifestFilter
//
//   manifest := mediapackagev2_alpha.Manifest_Dash(&DashManifestConfiguration{
//   	ManifestName: jsii.String("manifestName"),
//
//   	// the properties below are optional
//   	BaseUrls: []DashBaseUrlProperty{
//   		&DashBaseUrlProperty{
//   			Url: jsii.String("url"),
//
//   			// the properties below are optional
//   			DvbPriority: jsii.Number(123),
//   			DvbWeight: jsii.Number(123),
//   			ServiceLocation: jsii.String("serviceLocation"),
//   		},
//   	},
//   	Compactness: mediapackagev2_alpha.DashManifestCompactness_STANDARD,
//   	DrmSignalling: mediapackagev2_alpha.DrmSignalling_INDIVIDUAL,
//   	DvbSettings: &DashDvbSettings{
//   		ErrorMetrics: []DashDvbMetricsReporting{
//   			&DashDvbMetricsReporting{
//   				ReportingUrl: jsii.String("reportingUrl"),
//
//   				// the properties below are optional
//   				Probability: jsii.Number(123),
//   			},
//   		},
//   		FontDownload: &DashDvbFontDownload{
//   			FontFamily: jsii.String("fontFamily"),
//   			MimeType: jsii.String("mimeType"),
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	FilterConfiguration: &FilterConfiguration{
//   		ClipStartTime: NewDate(),
//   		DrmSettings: []eXCLUDE_SESSION_KEYS{
//   			mediapackagev2_alpha.DrmSettingsKey_*eXCLUDE_SESSION_KEYS,
//   		},
//   		End: NewDate(),
//   		ManifestFilter: []ManifestFilter{
//   			manifestFilter,
//   		},
//   		Start: NewDate(),
//   		TimeDelay: cdk.Duration_Minutes(jsii.Number(30)),
//   	},
//   	ManifestWindow: cdk.Duration_*Minutes(jsii.Number(30)),
//   	MinBufferTime: cdk.Duration_*Minutes(jsii.Number(30)),
//   	MinUpdatePeriod: cdk.Duration_*Minutes(jsii.Number(30)),
//   	PeriodTriggers: []DashPeriodTriggers{
//   		mediapackagev2_alpha.DashPeriodTriggers_AVAILS,
//   	},
//   	Profiles: []*string{
//   		jsii.String("profiles"),
//   	},
//   	ProgramInformation: &DashProgramInformation{
//   		Copyright: jsii.String("copyright"),
//   		LanguageCode: jsii.String("languageCode"),
//   		MoreInformationUrl: jsii.String("moreInformationUrl"),
//   		Source: jsii.String("source"),
//   		Title: jsii.String("title"),
//   	},
//   	ScteDashAdMarker: mediapackagev2_alpha.AdMarkerDash_BINARY,
//   	SegmentTemplateFormat: mediapackagev2_alpha.SegmentTemplateFormat_NUMBER_WITH_TIMELINE,
//   	SubtitleConfiguration: &DashSubtitleConfiguration{
//   		TtmlConfiguration: &DashTtmlConfiguration{
//   			TtmlProfile: mediapackagev2_alpha.TtmlProfile_IMSC_1,
//   		},
//   	},
//   	SuggestedPresentationDelay: cdk.Duration_*Minutes(jsii.Number(30)),
//   	UtcTimingMode: mediapackagev2_alpha.DashUtcTimingMode_HTTP_HEAD,
//   	UtcTimingSource: jsii.String("utcTimingSource"),
//   })
//
// Experimental.
type Manifest interface {
}

// The jsii proxy struct for Manifest
type jsiiProxy_Manifest struct {
	_ byte // padding
}

// Experimental.
func NewManifest_Override(m Manifest) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.Manifest",
		nil, // no parameters
		m,
	)
}

// Specify a manifest configuration for DASH.
//
// **Note:** DASH manifests require CMAF container type.
// Use with `Segment.cmaf()`.
// Experimental.
func Manifest_Dash(manifest *DashManifestConfiguration) Manifest {
	_init_.Initialize()

	if err := validateManifest_DashParameters(manifest); err != nil {
		panic(err)
	}
	var returns Manifest

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.Manifest",
		"dash",
		[]interface{}{manifest},
		&returns,
	)

	return returns
}

// Specify a manifest configuration for HLS.
//
// **Note:** HLS manifests require TS or CMAF container type.
// Use with `Segment.ts()` or `Segment.cmaf()`.
// Experimental.
func Manifest_Hls(manifest *HlsManifestConfiguration) Manifest {
	_init_.Initialize()

	if err := validateManifest_HlsParameters(manifest); err != nil {
		panic(err)
	}
	var returns Manifest

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.Manifest",
		"hls",
		[]interface{}{manifest},
		&returns,
	)

	return returns
}

// Specify a manifest configuration for Low Latency HLS.
//
// **Note:** Low Latency HLS manifests require TS or CMAF container type.
// Use with `Segment.ts()` or `Segment.cmaf()`.
// Experimental.
func Manifest_LowLatencyHLS(manifest *LowLatencyHlsManifestConfiguration) Manifest {
	_init_.Initialize()

	if err := validateManifest_LowLatencyHLSParameters(manifest); err != nil {
		panic(err)
	}
	var returns Manifest

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.Manifest",
		"lowLatencyHLS",
		[]interface{}{manifest},
		&returns,
	)

	return returns
}

// Specify a manifest configuration for MSS.
//
// **Note:** MSS manifests require ISM container type.
// Use with `Segment.ism()`.
// Experimental.
func Manifest_Mss(manifest *MssManifestConfiguration) Manifest {
	_init_.Initialize()

	if err := validateManifest_MssParameters(manifest); err != nil {
		panic(err)
	}
	var returns Manifest

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.Manifest",
		"mss",
		[]interface{}{manifest},
		&returns,
	)

	return returns
}


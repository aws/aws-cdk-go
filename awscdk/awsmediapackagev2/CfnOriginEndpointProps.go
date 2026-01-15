package awsmediapackagev2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOriginEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginEndpointProps := &CfnOriginEndpointProps{
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   	ContainerType: jsii.String("containerType"),
//   	OriginEndpointName: jsii.String("originEndpointName"),
//
//   	// the properties below are optional
//   	DashManifests: []interface{}{
//   		&DashManifestConfigurationProperty{
//   			ManifestName: jsii.String("manifestName"),
//
//   			// the properties below are optional
//   			BaseUrls: []interface{}{
//   				&DashBaseUrlProperty{
//   					Url: jsii.String("url"),
//
//   					// the properties below are optional
//   					DvbPriority: jsii.Number(123),
//   					DvbWeight: jsii.Number(123),
//   					ServiceLocation: jsii.String("serviceLocation"),
//   				},
//   			},
//   			Compactness: jsii.String("compactness"),
//   			DrmSignaling: jsii.String("drmSignaling"),
//   			DvbSettings: &DashDvbSettingsProperty{
//   				ErrorMetrics: []interface{}{
//   					&DashDvbMetricsReportingProperty{
//   						ReportingUrl: jsii.String("reportingUrl"),
//
//   						// the properties below are optional
//   						Probability: jsii.Number(123),
//   					},
//   				},
//   				FontDownload: &DashDvbFontDownloadProperty{
//   					FontFamily: jsii.String("fontFamily"),
//   					MimeType: jsii.String("mimeType"),
//   					Url: jsii.String("url"),
//   				},
//   			},
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				DrmSettings: jsii.String("drmSettings"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestWindowSeconds: jsii.Number(123),
//   			MinBufferTimeSeconds: jsii.Number(123),
//   			MinUpdatePeriodSeconds: jsii.Number(123),
//   			PeriodTriggers: []*string{
//   				jsii.String("periodTriggers"),
//   			},
//   			Profiles: []*string{
//   				jsii.String("profiles"),
//   			},
//   			ProgramInformation: &DashProgramInformationProperty{
//   				Copyright: jsii.String("copyright"),
//   				LanguageCode: jsii.String("languageCode"),
//   				MoreInformationUrl: jsii.String("moreInformationUrl"),
//   				Source: jsii.String("source"),
//   				Title: jsii.String("title"),
//   			},
//   			ScteDash: &ScteDashProperty{
//   				AdMarkerDash: jsii.String("adMarkerDash"),
//   			},
//   			SegmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   			SubtitleConfiguration: &DashSubtitleConfigurationProperty{
//   				TtmlConfiguration: &DashTtmlConfigurationProperty{
//   					TtmlProfile: jsii.String("ttmlProfile"),
//   				},
//   			},
//   			SuggestedPresentationDelaySeconds: jsii.Number(123),
//   			UtcTiming: &DashUtcTimingProperty{
//   				TimingMode: jsii.String("timingMode"),
//   				TimingSource: jsii.String("timingSource"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	ForceEndpointErrorConfiguration: &ForceEndpointErrorConfigurationProperty{
//   		EndpointErrorConditions: []*string{
//   			jsii.String("endpointErrorConditions"),
//   		},
//   	},
//   	HlsManifests: []interface{}{
//   		&HlsManifestConfigurationProperty{
//   			ManifestName: jsii.String("manifestName"),
//
//   			// the properties below are optional
//   			ChildManifestName: jsii.String("childManifestName"),
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				DrmSettings: jsii.String("drmSettings"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestWindowSeconds: jsii.Number(123),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			ScteHls: &ScteHlsProperty{
//   				AdMarkerHls: jsii.String("adMarkerHls"),
//   			},
//   			StartTag: &StartTagProperty{
//   				TimeOffset: jsii.Number(123),
//
//   				// the properties below are optional
//   				Precise: jsii.Boolean(false),
//   			},
//   			Url: jsii.String("url"),
//   			UrlEncodeChildManifest: jsii.Boolean(false),
//   		},
//   	},
//   	LowLatencyHlsManifests: []interface{}{
//   		&LowLatencyHlsManifestConfigurationProperty{
//   			ManifestName: jsii.String("manifestName"),
//
//   			// the properties below are optional
//   			ChildManifestName: jsii.String("childManifestName"),
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				DrmSettings: jsii.String("drmSettings"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestWindowSeconds: jsii.Number(123),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			ScteHls: &ScteHlsProperty{
//   				AdMarkerHls: jsii.String("adMarkerHls"),
//   			},
//   			StartTag: &StartTagProperty{
//   				TimeOffset: jsii.Number(123),
//
//   				// the properties below are optional
//   				Precise: jsii.Boolean(false),
//   			},
//   			Url: jsii.String("url"),
//   			UrlEncodeChildManifest: jsii.Boolean(false),
//   		},
//   	},
//   	MssManifests: []interface{}{
//   		&MssManifestConfigurationProperty{
//   			ManifestName: jsii.String("manifestName"),
//
//   			// the properties below are optional
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				DrmSettings: jsii.String("drmSettings"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestLayout: jsii.String("manifestLayout"),
//   			ManifestWindowSeconds: jsii.Number(123),
//   		},
//   	},
//   	Segment: &SegmentProperty{
//   		Encryption: &EncryptionProperty{
//   			EncryptionMethod: &EncryptionMethodProperty{
//   				CmafEncryptionMethod: jsii.String("cmafEncryptionMethod"),
//   				IsmEncryptionMethod: jsii.String("ismEncryptionMethod"),
//   				TsEncryptionMethod: jsii.String("tsEncryptionMethod"),
//   			},
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				DrmSystems: []*string{
//   					jsii.String("drmSystems"),
//   				},
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				CertificateArn: jsii.String("certificateArn"),
//   			},
//
//   			// the properties below are optional
//   			CmafExcludeSegmentDrmMetadata: jsii.Boolean(false),
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			KeyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		IncludeIframeOnlyStreams: jsii.Boolean(false),
//   		Scte: &ScteProperty{
//   			ScteFilter: []*string{
//   				jsii.String("scteFilter"),
//   			},
//   			ScteInSegments: jsii.String("scteInSegments"),
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   		SegmentName: jsii.String("segmentName"),
//   		TsIncludeDvbSubtitles: jsii.Boolean(false),
//   		TsUseAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	StartoverWindowSeconds: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html
//
type CfnOriginEndpointProps struct {
	// The name of the channel group associated with the origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-channelgroupname
	//
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// The channel name associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-channelname
	//
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// The container type associated with the origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-containertype
	//
	ContainerType *string `field:"required" json:"containerType" yaml:"containerType"`
	// The name of the origin endpoint associated with the origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-originendpointname
	//
	OriginEndpointName *string `field:"required" json:"originEndpointName" yaml:"originEndpointName"`
	// A DASH manifest configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-dashmanifests
	//
	DashManifests interface{} `field:"optional" json:"dashManifests" yaml:"dashManifests"`
	// The description associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The failover settings for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-forceendpointerrorconfiguration
	//
	ForceEndpointErrorConfiguration interface{} `field:"optional" json:"forceEndpointErrorConfiguration" yaml:"forceEndpointErrorConfiguration"`
	// The HLS manifests associated with the origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-hlsmanifests
	//
	HlsManifests interface{} `field:"optional" json:"hlsManifests" yaml:"hlsManifests"`
	// The low-latency HLS (LL-HLS) manifests associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifests
	//
	LowLatencyHlsManifests interface{} `field:"optional" json:"lowLatencyHlsManifests" yaml:"lowLatencyHlsManifests"`
	// A list of Microsoft Smooth Streaming (MSS) manifest configurations associated with the origin endpoint.
	//
	// Each configuration represents a different MSS streaming option available from this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-mssmanifests
	//
	MssManifests interface{} `field:"optional" json:"mssManifests" yaml:"mssManifests"`
	// The segment associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-segment
	//
	Segment interface{} `field:"optional" json:"segment" yaml:"segment"`
	// The size of the window (in seconds) to specify a window of the live stream that's available for on-demand viewing.
	//
	// Viewers can start-over or catch-up on content that falls within the window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-startoverwindowseconds
	//
	StartoverWindowSeconds *float64 `field:"optional" json:"startoverWindowSeconds" yaml:"startoverWindowSeconds"`
	// The tags associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


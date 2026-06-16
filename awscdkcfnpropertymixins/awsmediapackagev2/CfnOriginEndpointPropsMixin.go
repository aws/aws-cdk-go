package awsmediapackagev2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediapackagev2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the configuration parameters for a MediaPackage V2 origin endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnOriginEndpointPropsMixin := awscdkcfnpropertymixins.Aws_mediapackagev2.NewCfnOriginEndpointPropsMixin(&CfnOriginEndpointMixinProps{
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   	ContainerType: jsii.String("containerType"),
//   	DashManifests: []interface{}{
//   		&DashManifestConfigurationProperty{
//   			AudioTimelinePattern: jsii.String("audioTimelinePattern"),
//   			AvailabilityStartTimeConfiguration: &DashAvailabilityStartTimeConfigurationProperty{
//   				FixedAvailabilityStartTime: jsii.String("fixedAvailabilityStartTime"),
//   			},
//   			BaseUrls: []interface{}{
//   				&DashBaseUrlProperty{
//   					DvbPriority: jsii.Number(123),
//   					DvbWeight: jsii.Number(123),
//   					ServiceLocation: jsii.String("serviceLocation"),
//   					Url: jsii.String("url"),
//   				},
//   			},
//   			Compactness: jsii.String("compactness"),
//   			DrmSignaling: jsii.String("drmSignaling"),
//   			DvbSettings: &DashDvbSettingsProperty{
//   				ErrorMetrics: []interface{}{
//   					&DashDvbMetricsReportingProperty{
//   						Probability: jsii.Number(123),
//   						ReportingUrl: jsii.String("reportingUrl"),
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
//   			ManifestName: jsii.String("manifestName"),
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
//   				ScteInManifests: jsii.String("scteInManifests"),
//   			},
//   			SegmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   			SubtitleConfiguration: &DashSubtitleConfigurationProperty{
//   				TtmlConfiguration: &DashTtmlConfigurationProperty{
//   					TtmlProfile: jsii.String("ttmlProfile"),
//   				},
//   			},
//   			SuggestedPresentationDelaySeconds: jsii.Number(123),
//   			UriPathType: jsii.String("uriPathType"),
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
//   			ChildManifestName: jsii.String("childManifestName"),
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				DrmSettings: jsii.String("drmSettings"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestName: jsii.String("manifestName"),
//   			ManifestWindowSeconds: jsii.Number(123),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			ScteHls: &ScteHlsProperty{
//   				AdMarkerHls: jsii.String("adMarkerHls"),
//   				ScteInManifests: jsii.String("scteInManifests"),
//   			},
//   			StartTag: &StartTagProperty{
//   				Precise: jsii.Boolean(false),
//   				TimeOffset: jsii.Number(123),
//   			},
//   			UriPathType: jsii.String("uriPathType"),
//   			Url: jsii.String("url"),
//   			UrlEncodeChildManifest: jsii.Boolean(false),
//   		},
//   	},
//   	LowLatencyHlsManifests: []interface{}{
//   		&LowLatencyHlsManifestConfigurationProperty{
//   			ChildManifestName: jsii.String("childManifestName"),
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				DrmSettings: jsii.String("drmSettings"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestName: jsii.String("manifestName"),
//   			ManifestWindowSeconds: jsii.Number(123),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			ScteHls: &ScteHlsProperty{
//   				AdMarkerHls: jsii.String("adMarkerHls"),
//   				ScteInManifests: jsii.String("scteInManifests"),
//   			},
//   			StartTag: &StartTagProperty{
//   				Precise: jsii.Boolean(false),
//   				TimeOffset: jsii.Number(123),
//   			},
//   			UriPathType: jsii.String("uriPathType"),
//   			Url: jsii.String("url"),
//   			UrlEncodeChildManifest: jsii.Boolean(false),
//   		},
//   	},
//   	MssManifests: []interface{}{
//   		&MssManifestConfigurationProperty{
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				ClipStartTime: jsii.String("clipStartTime"),
//   				DrmSettings: jsii.String("drmSettings"),
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestLayout: jsii.String("manifestLayout"),
//   			ManifestName: jsii.String("manifestName"),
//   			ManifestWindowSeconds: jsii.Number(123),
//   		},
//   	},
//   	OriginEndpointName: jsii.String("originEndpointName"),
//   	Segment: &SegmentProperty{
//   		Encryption: &EncryptionProperty{
//   			CmafExcludeSegmentDrmMetadata: jsii.Boolean(false),
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			EncryptionMethod: &EncryptionMethodProperty{
//   				CmafEncryptionMethod: jsii.String("cmafEncryptionMethod"),
//   				IsmEncryptionMethod: jsii.String("ismEncryptionMethod"),
//   				TsEncryptionMethod: jsii.String("tsEncryptionMethod"),
//   			},
//   			KeyRotationIntervalSeconds: jsii.Number(123),
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				CertificateArn: jsii.String("certificateArn"),
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
//   			},
//   		},
//   		IncludeIframeOnlyStreams: jsii.Boolean(false),
//   		Scte: &ScteProperty{
//   			CustomAdTypes: []*string{
//   				jsii.String("customAdTypes"),
//   			},
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
//   	UriSeparator: jsii.String("uriSeparator"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html
//
type CfnOriginEndpointPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnOriginEndpointMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOriginEndpointPropsMixin
type jsiiProxy_CfnOriginEndpointPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnOriginEndpointPropsMixin) Props() *CfnOriginEndpointMixinProps {
	var returns *CfnOriginEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginEndpointPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaPackageV2::OriginEndpoint`.
func NewCfnOriginEndpointPropsMixin(props *CfnOriginEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnOriginEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOriginEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOriginEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnOriginEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaPackageV2::OriginEndpoint`.
func NewCfnOriginEndpointPropsMixin_Override(c CfnOriginEndpointPropsMixin, props *CfnOriginEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnOriginEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnOriginEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOriginEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnOriginEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOriginEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediapackagev2.CfnOriginEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOriginEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnOriginEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}


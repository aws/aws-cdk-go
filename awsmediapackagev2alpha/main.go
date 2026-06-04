// The CDK construct library for MediaPackageV2
package awsmediapackagev2alpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.AdMarkerDash",
		reflect.TypeOf((*AdMarkerDash)(nil)).Elem(),
		map[string]interface{}{
			"BINARY": AdMarkerDash_BINARY,
			"XML": AdMarkerDash_XML,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.AdMarkerHls",
		reflect.TypeOf((*AdMarkerHls)(nil)).Elem(),
		map[string]interface{}{
			"DATERANGE": AdMarkerHls_DATERANGE,
			"SCTE35_ENHANCED": AdMarkerHls_SCTE35_ENHANCED,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.AudioCodec",
		reflect.TypeOf((*AudioCodec)(nil)).Elem(),
		map[string]interface{}{
			"AACL": AudioCodec_AACL,
			"AACH": AudioCodec_AACH,
			"AC_3": AudioCodec_AC_3,
			"EC_3": AudioCodec_EC_3,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.BitrateExpression",
		reflect.TypeOf((*BitrateExpression)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BitrateExpression{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.BitrateFilterKey",
		reflect.TypeOf((*BitrateFilterKey)(nil)).Elem(),
		map[string]interface{}{
			"AUDIO_BITRATE": BitrateFilterKey_AUDIO_BITRATE,
			"VIDEO_BITRATE": BitrateFilterKey_VIDEO_BITRATE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.CdnAuthConfiguration",
		reflect.TypeOf((*CdnAuthConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.Channel",
		reflect.TypeOf((*Channel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOriginEndpoint", GoMethod: "AddOriginEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "channelArn", GoGetter: "ChannelArn"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroup", GoGetter: "ChannelGroup"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroupName", GoGetter: "ChannelGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "channelName", GoGetter: "ChannelName"},
			_jsii_.MemberProperty{JsiiProperty: "channelRef", GoGetter: "ChannelRef"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "grants", GoGetter: "Grants"},
			_jsii_.MemberProperty{JsiiProperty: "ingestEndpointUrls", GoGetter: "IngestEndpointUrls"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressBytes", GoMethod: "MetricEgressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressRequestCount", GoMethod: "MetricEgressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressResponseTime", GoMethod: "MetricEgressResponseTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressBytes", GoMethod: "MetricIngressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressRequestCount", GoMethod: "MetricIngressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressResponseTime", GoMethod: "MetricIngressResponseTime"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedAt", GoGetter: "ModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Channel{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IChannel)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.ChannelAttributes",
		reflect.TypeOf((*ChannelAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.ChannelGrants",
		reflect.TypeOf((*ChannelGrants)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "actions", GoMethod: "Actions"},
			_jsii_.MemberMethod{JsiiMethod: "ingest", GoMethod: "Ingest"},
			_jsii_.MemberProperty{JsiiProperty: "policyResource", GoGetter: "PolicyResource"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
		},
		func() interface{} {
			return &jsiiProxy_ChannelGrants{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.ChannelGroup",
		reflect.TypeOf((*ChannelGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addChannel", GoMethod: "AddChannel"},
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroupArn", GoGetter: "ChannelGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroupName", GoGetter: "ChannelGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroupRef", GoGetter: "ChannelGroupRef"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "egressDomain", GoGetter: "EgressDomain"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressBytes", GoMethod: "MetricEgressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressRequestCount", GoMethod: "MetricEgressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressResponseTime", GoMethod: "MetricEgressResponseTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressBytes", GoMethod: "MetricIngressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressRequestCount", GoMethod: "MetricIngressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressResponseTime", GoMethod: "MetricIngressResponseTime"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedAt", GoGetter: "ModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ChannelGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IChannelGroup)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.ChannelGroupAttributes",
		reflect.TypeOf((*ChannelGroupAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.ChannelGroupProps",
		reflect.TypeOf((*ChannelGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.ChannelOptions",
		reflect.TypeOf((*ChannelOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.ChannelPolicy",
		reflect.TypeOf((*ChannelPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "document", GoGetter: "Document"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_ChannelPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.ChannelPolicyProps",
		reflect.TypeOf((*ChannelPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.ChannelProps",
		reflect.TypeOf((*ChannelProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.CmafDrmSystem",
		reflect.TypeOf((*CmafDrmSystem)(nil)).Elem(),
		map[string]interface{}{
			"FAIRPLAY": CmafDrmSystem_FAIRPLAY,
			"PLAYREADY": CmafDrmSystem_PLAYREADY,
			"WIDEVINE": CmafDrmSystem_WIDEVINE,
			"IRDETO": CmafDrmSystem_IRDETO,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.CmafEncryption",
		reflect.TypeOf((*CmafEncryption)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			j := jsiiProxy_CmafEncryption{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EncryptionConfiguration)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.CmafEncryptionMethod",
		reflect.TypeOf((*CmafEncryptionMethod)(nil)).Elem(),
		map[string]interface{}{
			"CENC": CmafEncryptionMethod_CENC,
			"CBCS": CmafEncryptionMethod_CBCS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.CmafInputProps",
		reflect.TypeOf((*CmafInputProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.CmafSegmentProps",
		reflect.TypeOf((*CmafSegmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.CmafSpekeEncryptionProps",
		reflect.TypeOf((*CmafSpekeEncryptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.ContainerType",
		reflect.TypeOf((*ContainerType)(nil)).Elem(),
		map[string]interface{}{
			"TS": ContainerType_TS,
			"CMAF": ContainerType_CMAF,
			"ISM": ContainerType_ISM,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.DashBaseUrlProperty",
		reflect.TypeOf((*DashBaseUrlProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.DashDvbFontDownload",
		reflect.TypeOf((*DashDvbFontDownload)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.DashDvbMetricsReporting",
		reflect.TypeOf((*DashDvbMetricsReporting)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.DashDvbSettings",
		reflect.TypeOf((*DashDvbSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.DashManifestCompactness",
		reflect.TypeOf((*DashManifestCompactness)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": DashManifestCompactness_STANDARD,
			"NONE": DashManifestCompactness_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.DashManifestConfiguration",
		reflect.TypeOf((*DashManifestConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.DashPeriodTriggers",
		reflect.TypeOf((*DashPeriodTriggers)(nil)).Elem(),
		map[string]interface{}{
			"AVAILS": DashPeriodTriggers_AVAILS,
			"DRM_KEY_ROTATION": DashPeriodTriggers_DRM_KEY_ROTATION,
			"SOURCE_CHANGES": DashPeriodTriggers_SOURCE_CHANGES,
			"SOURCE_DISRUPTIONS": DashPeriodTriggers_SOURCE_DISRUPTIONS,
			"NONE": DashPeriodTriggers_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.DashProgramInformation",
		reflect.TypeOf((*DashProgramInformation)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.DashSubtitleConfiguration",
		reflect.TypeOf((*DashSubtitleConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.DashTtmlConfiguration",
		reflect.TypeOf((*DashTtmlConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.DashUtcTimingMode",
		reflect.TypeOf((*DashUtcTimingMode)(nil)).Elem(),
		map[string]interface{}{
			"HTTP_HEAD": DashUtcTimingMode_HTTP_HEAD,
			"HTTP_ISO": DashUtcTimingMode_HTTP_ISO,
			"HTTP_XSDATE": DashUtcTimingMode_HTTP_XSDATE,
			"UTC_DIRECT": DashUtcTimingMode_UTC_DIRECT,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.DrmSettingsKey",
		reflect.TypeOf((*DrmSettingsKey)(nil)).Elem(),
		map[string]interface{}{
			"EXCLUDE_SESSION_KEYS": DrmSettingsKey_EXCLUDE_SESSION_KEYS,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.DrmSignalling",
		reflect.TypeOf((*DrmSignalling)(nil)).Elem(),
		map[string]interface{}{
			"INDIVIDUAL": DrmSignalling_INDIVIDUAL,
			"REFERENCED": DrmSignalling_REFERENCED,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.EncryptionConfiguration",
		reflect.TypeOf((*EncryptionConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_EncryptionConfiguration{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.EndpointErrorConfiguration",
		reflect.TypeOf((*EndpointErrorConfiguration)(nil)).Elem(),
		map[string]interface{}{
			"STALE_MANIFEST": EndpointErrorConfiguration_STALE_MANIFEST,
			"INCOMPLETE_MANIFEST": EndpointErrorConfiguration_INCOMPLETE_MANIFEST,
			"MISSING_DRM_KEY": EndpointErrorConfiguration_MISSING_DRM_KEY,
			"SLATE_INPUT": EndpointErrorConfiguration_SLATE_INPUT,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.FilterConfiguration",
		reflect.TypeOf((*FilterConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.HeadersCMSD",
		reflect.TypeOf((*HeadersCMSD)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_HeadersCMSD{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.HlsManifestConfiguration",
		reflect.TypeOf((*HlsManifestConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediapackagev2-alpha.IChannel",
		reflect.TypeOf((*IChannel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOriginEndpoint", GoMethod: "AddOriginEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "channelArn", GoGetter: "ChannelArn"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroup", GoGetter: "ChannelGroup"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroupName", GoGetter: "ChannelGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "channelName", GoGetter: "ChannelName"},
			_jsii_.MemberProperty{JsiiProperty: "channelRef", GoGetter: "ChannelRef"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "grants", GoGetter: "Grants"},
			_jsii_.MemberProperty{JsiiProperty: "ingestEndpointUrls", GoGetter: "IngestEndpointUrls"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressBytes", GoMethod: "MetricEgressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressRequestCount", GoMethod: "MetricEgressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressResponseTime", GoMethod: "MetricEgressResponseTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressBytes", GoMethod: "MetricIngressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressRequestCount", GoMethod: "MetricIngressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressResponseTime", GoMethod: "MetricIngressResponseTime"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedAt", GoGetter: "ModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IChannel{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediapackagev2IChannelRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediapackagev2-alpha.IChannelGroup",
		reflect.TypeOf((*IChannelGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addChannel", GoMethod: "AddChannel"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroupArn", GoGetter: "ChannelGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroupName", GoGetter: "ChannelGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroupRef", GoGetter: "ChannelGroupRef"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "egressDomain", GoGetter: "EgressDomain"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressBytes", GoMethod: "MetricEgressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressRequestCount", GoMethod: "MetricEgressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressResponseTime", GoMethod: "MetricEgressResponseTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressBytes", GoMethod: "MetricIngressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressRequestCount", GoMethod: "MetricIngressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressResponseTime", GoMethod: "MetricIngressResponseTime"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedAt", GoGetter: "ModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IChannelGroup{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediapackagev2IChannelGroupRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk/aws-mediapackagev2-alpha.IOriginEndpoint",
		reflect.TypeOf((*IOriginEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroupName", GoGetter: "ChannelGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "channelName", GoGetter: "ChannelName"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "dashManifestUrls", GoGetter: "DashManifestUrls"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "hlsManifestUrls", GoGetter: "HlsManifestUrls"},
			_jsii_.MemberProperty{JsiiProperty: "lowLatencyHlsManifestUrls", GoGetter: "LowLatencyHlsManifestUrls"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressBytes", GoMethod: "MetricEgressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressRequestCount", GoMethod: "MetricEgressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressResponseTime", GoMethod: "MetricEgressResponseTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressBytes", GoMethod: "MetricIngressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressRequestCount", GoMethod: "MetricIngressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressResponseTime", GoMethod: "MetricIngressResponseTime"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedAt", GoGetter: "ModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "mssManifestUrls", GoGetter: "MssManifestUrls"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originEndpointArn", GoGetter: "OriginEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "originEndpointName", GoGetter: "OriginEndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "originEndpointRef", GoGetter: "OriginEndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IOriginEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmediapackagev2IOriginEndpointRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.IngestEndpoint",
		reflect.TypeOf((*IngestEndpoint)(nil)).Elem(),
		map[string]interface{}{
			"ENDPOINT_1": IngestEndpoint_ENDPOINT_1,
			"ENDPOINT_2": IngestEndpoint_ENDPOINT_2,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.InputConfiguration",
		reflect.TypeOf((*InputConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "inputSwitchConfiguration", GoGetter: "InputSwitchConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "inputType", GoGetter: "InputType"},
			_jsii_.MemberProperty{JsiiProperty: "outputHeaders", GoGetter: "OutputHeaders"},
		},
		func() interface{} {
			return &jsiiProxy_InputConfiguration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.InputSwitchConfiguration",
		reflect.TypeOf((*InputSwitchConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.InputType",
		reflect.TypeOf((*InputType)(nil)).Elem(),
		map[string]interface{}{
			"CMAF": InputType_CMAF,
			"HLS": InputType_HLS,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.IsmDrmSystem",
		reflect.TypeOf((*IsmDrmSystem)(nil)).Elem(),
		map[string]interface{}{
			"PLAYREADY": IsmDrmSystem_PLAYREADY,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.IsmEncryption",
		reflect.TypeOf((*IsmEncryption)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			j := jsiiProxy_IsmEncryption{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EncryptionConfiguration)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.IsmSegmentProps",
		reflect.TypeOf((*IsmSegmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.IsmSpekeEncryptionProps",
		reflect.TypeOf((*IsmSpekeEncryptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.LowLatencyHlsManifestConfiguration",
		reflect.TypeOf((*LowLatencyHlsManifestConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.Manifest",
		reflect.TypeOf((*Manifest)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Manifest{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.ManifestFilter",
		reflect.TypeOf((*ManifestFilter)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "filterString", GoGetter: "FilterString"},
		},
		func() interface{} {
			return &jsiiProxy_ManifestFilter{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.MediaPackageV2Origin",
		reflect.TypeOf((*MediaPackageV2Origin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "renderCustomOriginConfig", GoMethod: "RenderCustomOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderS3OriginConfig", GoMethod: "RenderS3OriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "renderVpcOriginConfig", GoMethod: "RenderVpcOriginConfig"},
			_jsii_.MemberMethod{JsiiMethod: "validateResponseCompletionTimeoutWithReadTimeout", GoMethod: "ValidateResponseCompletionTimeoutWithReadTimeout"},
		},
		func() interface{} {
			j := jsiiProxy_MediaPackageV2Origin{}
			_jsii_.InitJsiiProxy(&j.Type__awscloudfrontOriginBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.MediaPackageV2OriginProps",
		reflect.TypeOf((*MediaPackageV2OriginProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.MssManifestConfiguration",
		reflect.TypeOf((*MssManifestConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.MssManifestLayout",
		reflect.TypeOf((*MssManifestLayout)(nil)).Elem(),
		map[string]interface{}{
			"FULL": MssManifestLayout_FULL,
			"COMPACT": MssManifestLayout_COMPACT,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.NumericExpression",
		reflect.TypeOf((*NumericExpression)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_NumericExpression{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.NumericFilterKey",
		reflect.TypeOf((*NumericFilterKey)(nil)).Elem(),
		map[string]interface{}{
			"AUDIO_CHANNELS": NumericFilterKey_AUDIO_CHANNELS,
			"AUDIO_SAMPLE_RATE": NumericFilterKey_AUDIO_SAMPLE_RATE,
			"TRICKPLAY_HEIGHT": NumericFilterKey_TRICKPLAY_HEIGHT,
			"VIDEO_FRAMERATE": NumericFilterKey_VIDEO_FRAMERATE,
			"VIDEO_HEIGHT": NumericFilterKey_VIDEO_HEIGHT,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpoint",
		reflect.TypeOf((*OriginEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreatePolicy", GoGetter: "AutoCreatePolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cdnAuthConfig", GoGetter: "CdnAuthConfig"},
			_jsii_.MemberProperty{JsiiProperty: "channelGroupName", GoGetter: "ChannelGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "channelName", GoGetter: "ChannelName"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "dashManifests", GoGetter: "DashManifests"},
			_jsii_.MemberProperty{JsiiProperty: "dashManifestUrls", GoGetter: "DashManifestUrls"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "hlsManifests", GoGetter: "HlsManifests"},
			_jsii_.MemberProperty{JsiiProperty: "hlsManifestUrls", GoGetter: "HlsManifestUrls"},
			_jsii_.MemberProperty{JsiiProperty: "llHlsManifests", GoGetter: "LlHlsManifests"},
			_jsii_.MemberProperty{JsiiProperty: "lowLatencyHlsManifestUrls", GoGetter: "LowLatencyHlsManifestUrls"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressBytes", GoMethod: "MetricEgressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressRequestCount", GoMethod: "MetricEgressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricEgressResponseTime", GoMethod: "MetricEgressResponseTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressBytes", GoMethod: "MetricIngressBytes"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressRequestCount", GoMethod: "MetricIngressRequestCount"},
			_jsii_.MemberMethod{JsiiMethod: "metricIngressResponseTime", GoMethod: "MetricIngressResponseTime"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedAt", GoGetter: "ModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "mssManifests", GoGetter: "MssManifests"},
			_jsii_.MemberProperty{JsiiProperty: "mssManifestUrls", GoGetter: "MssManifestUrls"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "originEndpointArn", GoGetter: "OriginEndpointArn"},
			_jsii_.MemberProperty{JsiiProperty: "originEndpointName", GoGetter: "OriginEndpointName"},
			_jsii_.MemberProperty{JsiiProperty: "originEndpointRef", GoGetter: "OriginEndpointRef"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "segment", GoGetter: "Segment"},
			_jsii_.MemberMethod{JsiiMethod: "segmentValidation", GoMethod: "SegmentValidation"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_OriginEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IOriginEndpoint)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpointAttributes",
		reflect.TypeOf((*OriginEndpointAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpointOptions",
		reflect.TypeOf((*OriginEndpointOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpointPolicy",
		reflect.TypeOf((*OriginEndpointPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "document", GoGetter: "Document"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_OriginEndpointPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpointPolicyProps",
		reflect.TypeOf((*OriginEndpointPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.OriginEndpointProps",
		reflect.TypeOf((*OriginEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.PresetSpeke20Audio",
		reflect.TypeOf((*PresetSpeke20Audio)(nil)).Elem(),
		map[string]interface{}{
			"PRESET_AUDIO_1": PresetSpeke20Audio_PRESET_AUDIO_1,
			"PRESET_AUDIO_2": PresetSpeke20Audio_PRESET_AUDIO_2,
			"PRESET_AUDIO_3": PresetSpeke20Audio_PRESET_AUDIO_3,
			"SHARED": PresetSpeke20Audio_SHARED,
			"UNENCRYPTED": PresetSpeke20Audio_UNENCRYPTED,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.PresetSpeke20Video",
		reflect.TypeOf((*PresetSpeke20Video)(nil)).Elem(),
		map[string]interface{}{
			"PRESET_VIDEO_1": PresetSpeke20Video_PRESET_VIDEO_1,
			"PRESET_VIDEO_2": PresetSpeke20Video_PRESET_VIDEO_2,
			"PRESET_VIDEO_3": PresetSpeke20Video_PRESET_VIDEO_3,
			"PRESET_VIDEO_4": PresetSpeke20Video_PRESET_VIDEO_4,
			"PRESET_VIDEO_5": PresetSpeke20Video_PRESET_VIDEO_5,
			"PRESET_VIDEO_6": PresetSpeke20Video_PRESET_VIDEO_6,
			"PRESET_VIDEO_7": PresetSpeke20Video_PRESET_VIDEO_7,
			"PRESET_VIDEO_8": PresetSpeke20Video_PRESET_VIDEO_8,
			"SHARED": PresetSpeke20Video_SHARED,
			"UNENCRYPTED": PresetSpeke20Video_UNENCRYPTED,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.ScteInSegments",
		reflect.TypeOf((*ScteInSegments)(nil)).Elem(),
		map[string]interface{}{
			"NONE": ScteInSegments_NONE,
			"ALL": ScteInSegments_ALL,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.ScteMessageType",
		reflect.TypeOf((*ScteMessageType)(nil)).Elem(),
		map[string]interface{}{
			"SPLICE_INSERT": ScteMessageType_SPLICE_INSERT,
			"BREAK": ScteMessageType_BREAK,
			"PROVIDER_ADVERTISEMENT": ScteMessageType_PROVIDER_ADVERTISEMENT,
			"DISTRIBUTOR_ADVERTISEMENT": ScteMessageType_DISTRIBUTOR_ADVERTISEMENT,
			"PROVIDER_PLACEMENT_OPPORTUNITY": ScteMessageType_PROVIDER_PLACEMENT_OPPORTUNITY,
			"DISTRIBUTOR_PLACEMENT_OPPORTUNITY": ScteMessageType_DISTRIBUTOR_PLACEMENT_OPPORTUNITY,
			"PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY": ScteMessageType_PROVIDER_OVERLAY_PLACEMENT_OPPORTUNITY,
			"DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY": ScteMessageType_DISTRIBUTOR_OVERLAY_PLACEMENT_OPPORTUNITY,
			"PROGRAM": ScteMessageType_PROGRAM,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.Segment",
		reflect.TypeOf((*Segment)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Segment{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.SegmentConfiguration",
		reflect.TypeOf((*SegmentConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.SegmentPropsBase",
		reflect.TypeOf((*SegmentPropsBase)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.SegmentTemplateFormat",
		reflect.TypeOf((*SegmentTemplateFormat)(nil)).Elem(),
		map[string]interface{}{
			"NUMBER_WITH_TIMELINE": SegmentTemplateFormat_NUMBER_WITH_TIMELINE,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.StartTag",
		reflect.TypeOf((*StartTag)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "precise", GoGetter: "Precise"},
			_jsii_.MemberProperty{JsiiProperty: "timeOffset", GoGetter: "TimeOffset"},
		},
		func() interface{} {
			return &jsiiProxy_StartTag{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.StartTagOptions",
		reflect.TypeOf((*StartTagOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.TextFilterKey",
		reflect.TypeOf((*TextFilterKey)(nil)).Elem(),
		map[string]interface{}{
			"AUDIO_LANGUAGE": TextFilterKey_AUDIO_LANGUAGE,
			"SUBTITLE_LANGUAGE": TextFilterKey_SUBTITLE_LANGUAGE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.TrickplayType",
		reflect.TypeOf((*TrickplayType)(nil)).Elem(),
		map[string]interface{}{
			"IFRAME": TrickplayType_IFRAME,
			"IMAGE": TrickplayType_IMAGE,
			"NONE": TrickplayType_NONE,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.TsDrmSystem",
		reflect.TypeOf((*TsDrmSystem)(nil)).Elem(),
		map[string]interface{}{
			"FAIRPLAY": TsDrmSystem_FAIRPLAY,
			"CLEAR_KEY_AES_128": TsDrmSystem_CLEAR_KEY_AES_128,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-mediapackagev2-alpha.TsEncryption",
		reflect.TypeOf((*TsEncryption)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			j := jsiiProxy_TsEncryption{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_EncryptionConfiguration)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.TsEncryptionMethod",
		reflect.TypeOf((*TsEncryptionMethod)(nil)).Elem(),
		map[string]interface{}{
			"AES_128": TsEncryptionMethod_AES_128,
			"SAMPLE_AES": TsEncryptionMethod_SAMPLE_AES,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.TsSegmentProps",
		reflect.TypeOf((*TsSegmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-mediapackagev2-alpha.TsSpekeEncryptionProps",
		reflect.TypeOf((*TsSpekeEncryptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.TtmlProfile",
		reflect.TypeOf((*TtmlProfile)(nil)).Elem(),
		map[string]interface{}{
			"IMSC_1": TtmlProfile_IMSC_1,
			"EBU_TT_D_101": TtmlProfile_EBU_TT_D_101,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.VideoCodec",
		reflect.TypeOf((*VideoCodec)(nil)).Elem(),
		map[string]interface{}{
			"H264": VideoCodec_H264,
			"H265": VideoCodec_H265,
			"AV1": VideoCodec_AV1,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-mediapackagev2-alpha.VideoDynamicRange",
		reflect.TypeOf((*VideoDynamicRange)(nil)).Elem(),
		map[string]interface{}{
			"DV": VideoDynamicRange_DV,
			"HDR10": VideoDynamicRange_HDR10,
			"HLG": VideoDynamicRange_HLG,
			"SDR": VideoDynamicRange_SDR,
		},
	)
}

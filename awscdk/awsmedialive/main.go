package awsmedialive

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnChannel",
		reflect.TypeOf((*CfnChannel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "anywhereSettings", GoGetter: "AnywhereSettings"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrInputs", GoGetter: "AttrInputs"},
			_jsii_.MemberProperty{JsiiProperty: "cdiInputSpecification", GoGetter: "CdiInputSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "channelClass", GoGetter: "ChannelClass"},
			_jsii_.MemberProperty{JsiiProperty: "channelEngineVersion", GoGetter: "ChannelEngineVersion"},
			_jsii_.MemberProperty{JsiiProperty: "channelRef", GoGetter: "ChannelRef"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinations", GoGetter: "Destinations"},
			_jsii_.MemberProperty{JsiiProperty: "dryRun", GoGetter: "DryRun"},
			_jsii_.MemberProperty{JsiiProperty: "encoderSettings", GoGetter: "EncoderSettings"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "inputAttachments", GoGetter: "InputAttachments"},
			_jsii_.MemberProperty{JsiiProperty: "inputSpecification", GoGetter: "InputSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "logLevel", GoGetter: "LogLevel"},
			_jsii_.MemberProperty{JsiiProperty: "maintenance", GoGetter: "Maintenance"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannel{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveIChannelRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AacSettingsProperty",
		reflect.TypeOf((*CfnChannel_AacSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Ac3SettingsProperty",
		reflect.TypeOf((*CfnChannel_Ac3SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AdditionalDestinationsProperty",
		reflect.TypeOf((*CfnChannel_AdditionalDestinationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AncillarySourceSettingsProperty",
		reflect.TypeOf((*CfnChannel_AncillarySourceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AnywhereSettingsProperty",
		reflect.TypeOf((*CfnChannel_AnywhereSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.ArchiveCdnSettingsProperty",
		reflect.TypeOf((*CfnChannel_ArchiveCdnSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.ArchiveContainerSettingsProperty",
		reflect.TypeOf((*CfnChannel_ArchiveContainerSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.ArchiveGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_ArchiveGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.ArchiveOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_ArchiveOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.ArchiveS3SettingsProperty",
		reflect.TypeOf((*CfnChannel_ArchiveS3SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AribDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_AribDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AribSourceSettingsProperty",
		reflect.TypeOf((*CfnChannel_AribSourceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioChannelMappingProperty",
		reflect.TypeOf((*CfnChannel_AudioChannelMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioCodecSettingsProperty",
		reflect.TypeOf((*CfnChannel_AudioCodecSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioDescriptionProperty",
		reflect.TypeOf((*CfnChannel_AudioDescriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioDolbyEDecodeProperty",
		reflect.TypeOf((*CfnChannel_AudioDolbyEDecodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioHlsRenditionSelectionProperty",
		reflect.TypeOf((*CfnChannel_AudioHlsRenditionSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioLanguageSelectionProperty",
		reflect.TypeOf((*CfnChannel_AudioLanguageSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioNormalizationSettingsProperty",
		reflect.TypeOf((*CfnChannel_AudioNormalizationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioOnlyHlsSettingsProperty",
		reflect.TypeOf((*CfnChannel_AudioOnlyHlsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioPidSelectionProperty",
		reflect.TypeOf((*CfnChannel_AudioPidSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioSelectorProperty",
		reflect.TypeOf((*CfnChannel_AudioSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioSelectorSettingsProperty",
		reflect.TypeOf((*CfnChannel_AudioSelectorSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioSilenceFailoverSettingsProperty",
		reflect.TypeOf((*CfnChannel_AudioSilenceFailoverSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioTrackProperty",
		reflect.TypeOf((*CfnChannel_AudioTrackProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioTrackSelectionProperty",
		reflect.TypeOf((*CfnChannel_AudioTrackSelectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AudioWatermarkSettingsProperty",
		reflect.TypeOf((*CfnChannel_AudioWatermarkSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AutomaticInputFailoverSettingsProperty",
		reflect.TypeOf((*CfnChannel_AutomaticInputFailoverSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Av1ColorSpaceSettingsProperty",
		reflect.TypeOf((*CfnChannel_Av1ColorSpaceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Av1SettingsProperty",
		reflect.TypeOf((*CfnChannel_Av1SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AvailBlankingProperty",
		reflect.TypeOf((*CfnChannel_AvailBlankingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AvailConfigurationProperty",
		reflect.TypeOf((*CfnChannel_AvailConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.AvailSettingsProperty",
		reflect.TypeOf((*CfnChannel_AvailSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.BandwidthReductionFilterSettingsProperty",
		reflect.TypeOf((*CfnChannel_BandwidthReductionFilterSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.BlackoutSlateProperty",
		reflect.TypeOf((*CfnChannel_BlackoutSlateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.BurnInDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_BurnInDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.CaptionDescriptionProperty",
		reflect.TypeOf((*CfnChannel_CaptionDescriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.CaptionDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_CaptionDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.CaptionLanguageMappingProperty",
		reflect.TypeOf((*CfnChannel_CaptionLanguageMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.CaptionRectangleProperty",
		reflect.TypeOf((*CfnChannel_CaptionRectangleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.CaptionSelectorProperty",
		reflect.TypeOf((*CfnChannel_CaptionSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.CaptionSelectorSettingsProperty",
		reflect.TypeOf((*CfnChannel_CaptionSelectorSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.CdiInputSpecificationProperty",
		reflect.TypeOf((*CfnChannel_CdiInputSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.ChannelEngineVersionRequestProperty",
		reflect.TypeOf((*CfnChannel_ChannelEngineVersionRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.CmafIngestCaptionLanguageMappingProperty",
		reflect.TypeOf((*CfnChannel_CmafIngestCaptionLanguageMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.CmafIngestGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_CmafIngestGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.CmafIngestOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_CmafIngestOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.ColorCorrectionProperty",
		reflect.TypeOf((*CfnChannel_ColorCorrectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.ColorCorrectionSettingsProperty",
		reflect.TypeOf((*CfnChannel_ColorCorrectionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.ColorSpacePassthroughSettingsProperty",
		reflect.TypeOf((*CfnChannel_ColorSpacePassthroughSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.DolbyVision81SettingsProperty",
		reflect.TypeOf((*CfnChannel_DolbyVision81SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.DvbNitSettingsProperty",
		reflect.TypeOf((*CfnChannel_DvbNitSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.DvbSdtSettingsProperty",
		reflect.TypeOf((*CfnChannel_DvbSdtSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.DvbSubDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_DvbSubDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.DvbSubSourceSettingsProperty",
		reflect.TypeOf((*CfnChannel_DvbSubSourceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.DvbTdtSettingsProperty",
		reflect.TypeOf((*CfnChannel_DvbTdtSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Eac3AtmosSettingsProperty",
		reflect.TypeOf((*CfnChannel_Eac3AtmosSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Eac3SettingsProperty",
		reflect.TypeOf((*CfnChannel_Eac3SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.EbuTtDDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_EbuTtDDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.EmbeddedDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_EmbeddedDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.EmbeddedPlusScte20DestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_EmbeddedPlusScte20DestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.EmbeddedSourceSettingsProperty",
		reflect.TypeOf((*CfnChannel_EmbeddedSourceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.EncoderSettingsProperty",
		reflect.TypeOf((*CfnChannel_EncoderSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.EpochLockingSettingsProperty",
		reflect.TypeOf((*CfnChannel_EpochLockingSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.EsamProperty",
		reflect.TypeOf((*CfnChannel_EsamProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.FailoverConditionProperty",
		reflect.TypeOf((*CfnChannel_FailoverConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.FailoverConditionSettingsProperty",
		reflect.TypeOf((*CfnChannel_FailoverConditionSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.FeatureActivationsProperty",
		reflect.TypeOf((*CfnChannel_FeatureActivationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.FecOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_FecOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Fmp4HlsSettingsProperty",
		reflect.TypeOf((*CfnChannel_Fmp4HlsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.FrameCaptureCdnSettingsProperty",
		reflect.TypeOf((*CfnChannel_FrameCaptureCdnSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.FrameCaptureGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_FrameCaptureGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.FrameCaptureHlsSettingsProperty",
		reflect.TypeOf((*CfnChannel_FrameCaptureHlsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.FrameCaptureOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_FrameCaptureOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.FrameCaptureS3SettingsProperty",
		reflect.TypeOf((*CfnChannel_FrameCaptureS3SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.FrameCaptureSettingsProperty",
		reflect.TypeOf((*CfnChannel_FrameCaptureSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.GlobalConfigurationProperty",
		reflect.TypeOf((*CfnChannel_GlobalConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.H264ColorSpaceSettingsProperty",
		reflect.TypeOf((*CfnChannel_H264ColorSpaceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.H264FilterSettingsProperty",
		reflect.TypeOf((*CfnChannel_H264FilterSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.H264SettingsProperty",
		reflect.TypeOf((*CfnChannel_H264SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.H265ColorSpaceSettingsProperty",
		reflect.TypeOf((*CfnChannel_H265ColorSpaceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.H265FilterSettingsProperty",
		reflect.TypeOf((*CfnChannel_H265FilterSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.H265SettingsProperty",
		reflect.TypeOf((*CfnChannel_H265SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Hdr10SettingsProperty",
		reflect.TypeOf((*CfnChannel_Hdr10SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HlsAkamaiSettingsProperty",
		reflect.TypeOf((*CfnChannel_HlsAkamaiSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HlsBasicPutSettingsProperty",
		reflect.TypeOf((*CfnChannel_HlsBasicPutSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HlsCdnSettingsProperty",
		reflect.TypeOf((*CfnChannel_HlsCdnSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HlsGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_HlsGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HlsInputSettingsProperty",
		reflect.TypeOf((*CfnChannel_HlsInputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HlsMediaStoreSettingsProperty",
		reflect.TypeOf((*CfnChannel_HlsMediaStoreSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HlsOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_HlsOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HlsS3SettingsProperty",
		reflect.TypeOf((*CfnChannel_HlsS3SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HlsSettingsProperty",
		reflect.TypeOf((*CfnChannel_HlsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HlsWebdavSettingsProperty",
		reflect.TypeOf((*CfnChannel_HlsWebdavSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.HtmlMotionGraphicsSettingsProperty",
		reflect.TypeOf((*CfnChannel_HtmlMotionGraphicsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.InputAttachmentProperty",
		reflect.TypeOf((*CfnChannel_InputAttachmentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.InputChannelLevelProperty",
		reflect.TypeOf((*CfnChannel_InputChannelLevelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.InputLocationProperty",
		reflect.TypeOf((*CfnChannel_InputLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.InputLossBehaviorProperty",
		reflect.TypeOf((*CfnChannel_InputLossBehaviorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.InputLossFailoverSettingsProperty",
		reflect.TypeOf((*CfnChannel_InputLossFailoverSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.InputSettingsProperty",
		reflect.TypeOf((*CfnChannel_InputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.InputSpecificationProperty",
		reflect.TypeOf((*CfnChannel_InputSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.KeyProviderSettingsProperty",
		reflect.TypeOf((*CfnChannel_KeyProviderSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.M2tsSettingsProperty",
		reflect.TypeOf((*CfnChannel_M2tsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.M3u8SettingsProperty",
		reflect.TypeOf((*CfnChannel_M3u8SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MaintenanceCreateSettingsProperty",
		reflect.TypeOf((*CfnChannel_MaintenanceCreateSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MaintenanceUpdateSettingsProperty",
		reflect.TypeOf((*CfnChannel_MaintenanceUpdateSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MediaPackageGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_MediaPackageGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MediaPackageOutputDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_MediaPackageOutputDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MediaPackageOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_MediaPackageOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MediaPackageV2DestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_MediaPackageV2DestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MediaPackageV2GroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_MediaPackageV2GroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MotionGraphicsConfigurationProperty",
		reflect.TypeOf((*CfnChannel_MotionGraphicsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MotionGraphicsSettingsProperty",
		reflect.TypeOf((*CfnChannel_MotionGraphicsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Mp2SettingsProperty",
		reflect.TypeOf((*CfnChannel_Mp2SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Mpeg2FilterSettingsProperty",
		reflect.TypeOf((*CfnChannel_Mpeg2FilterSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Mpeg2SettingsProperty",
		reflect.TypeOf((*CfnChannel_Mpeg2SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MsSmoothGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_MsSmoothGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MsSmoothOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_MsSmoothOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MulticastInputSettingsProperty",
		reflect.TypeOf((*CfnChannel_MulticastInputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MultiplexContainerSettingsProperty",
		reflect.TypeOf((*CfnChannel_MultiplexContainerSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MultiplexGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_MultiplexGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MultiplexM2tsSettingsProperty",
		reflect.TypeOf((*CfnChannel_MultiplexM2tsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MultiplexOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_MultiplexOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.MultiplexProgramChannelDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_MultiplexProgramChannelDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.NetworkInputSettingsProperty",
		reflect.TypeOf((*CfnChannel_NetworkInputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.NielsenCBETProperty",
		reflect.TypeOf((*CfnChannel_NielsenCBETProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.NielsenConfigurationProperty",
		reflect.TypeOf((*CfnChannel_NielsenConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.NielsenNaesIiNwProperty",
		reflect.TypeOf((*CfnChannel_NielsenNaesIiNwProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.NielsenWatermarksSettingsProperty",
		reflect.TypeOf((*CfnChannel_NielsenWatermarksSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.OutputDestinationProperty",
		reflect.TypeOf((*CfnChannel_OutputDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.OutputDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_OutputDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.OutputGroupProperty",
		reflect.TypeOf((*CfnChannel_OutputGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.OutputGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_OutputGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.OutputLocationRefProperty",
		reflect.TypeOf((*CfnChannel_OutputLocationRefProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.OutputLockingSettingsProperty",
		reflect.TypeOf((*CfnChannel_OutputLockingSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.OutputProperty",
		reflect.TypeOf((*CfnChannel_OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.OutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_OutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.PassThroughSettingsProperty",
		reflect.TypeOf((*CfnChannel_PassThroughSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.RawSettingsProperty",
		reflect.TypeOf((*CfnChannel_RawSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Rec601SettingsProperty",
		reflect.TypeOf((*CfnChannel_Rec601SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Rec709SettingsProperty",
		reflect.TypeOf((*CfnChannel_Rec709SettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.RemixSettingsProperty",
		reflect.TypeOf((*CfnChannel_RemixSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.RtmpCaptionInfoDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_RtmpCaptionInfoDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.RtmpGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_RtmpGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.RtmpOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_RtmpOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Scte20PlusEmbeddedDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_Scte20PlusEmbeddedDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Scte20SourceSettingsProperty",
		reflect.TypeOf((*CfnChannel_Scte20SourceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Scte27DestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_Scte27DestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Scte27SourceSettingsProperty",
		reflect.TypeOf((*CfnChannel_Scte27SourceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Scte35SpliceInsertProperty",
		reflect.TypeOf((*CfnChannel_Scte35SpliceInsertProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.Scte35TimeSignalAposProperty",
		reflect.TypeOf((*CfnChannel_Scte35TimeSignalAposProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.SmpteTtDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_SmpteTtDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.SrtGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_SrtGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.SrtOutputDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_SrtOutputDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.SrtOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_SrtOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.StandardHlsSettingsProperty",
		reflect.TypeOf((*CfnChannel_StandardHlsSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.StaticKeySettingsProperty",
		reflect.TypeOf((*CfnChannel_StaticKeySettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.TeletextDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_TeletextDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.TeletextSourceSettingsProperty",
		reflect.TypeOf((*CfnChannel_TeletextSourceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.TemporalFilterSettingsProperty",
		reflect.TypeOf((*CfnChannel_TemporalFilterSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.ThumbnailConfigurationProperty",
		reflect.TypeOf((*CfnChannel_ThumbnailConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.TimecodeBurninSettingsProperty",
		reflect.TypeOf((*CfnChannel_TimecodeBurninSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.TimecodeConfigProperty",
		reflect.TypeOf((*CfnChannel_TimecodeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.TtmlDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_TtmlDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.UdpContainerSettingsProperty",
		reflect.TypeOf((*CfnChannel_UdpContainerSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.UdpGroupSettingsProperty",
		reflect.TypeOf((*CfnChannel_UdpGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.UdpOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_UdpOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.VideoBlackFailoverSettingsProperty",
		reflect.TypeOf((*CfnChannel_VideoBlackFailoverSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.VideoCodecSettingsProperty",
		reflect.TypeOf((*CfnChannel_VideoCodecSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.VideoDescriptionProperty",
		reflect.TypeOf((*CfnChannel_VideoDescriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.VideoSelectorColorSpaceSettingsProperty",
		reflect.TypeOf((*CfnChannel_VideoSelectorColorSpaceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.VideoSelectorPidProperty",
		reflect.TypeOf((*CfnChannel_VideoSelectorPidProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.VideoSelectorProgramIdProperty",
		reflect.TypeOf((*CfnChannel_VideoSelectorProgramIdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.VideoSelectorProperty",
		reflect.TypeOf((*CfnChannel_VideoSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.VideoSelectorSettingsProperty",
		reflect.TypeOf((*CfnChannel_VideoSelectorSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.VpcOutputSettingsProperty",
		reflect.TypeOf((*CfnChannel_VpcOutputSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.WavSettingsProperty",
		reflect.TypeOf((*CfnChannel_WavSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannel.WebvttDestinationSettingsProperty",
		reflect.TypeOf((*CfnChannel_WebvttDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnChannelPlacementGroup",
		reflect.TypeOf((*CfnChannelPlacementGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrChannels", GoGetter: "AttrChannels"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "channelPlacementGroupRef", GoGetter: "ChannelPlacementGroupRef"},
			_jsii_.MemberProperty{JsiiProperty: "clusterId", GoGetter: "ClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodes", GoGetter: "Nodes"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnChannelPlacementGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveIChannelPlacementGroupRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannelPlacementGroupProps",
		reflect.TypeOf((*CfnChannelPlacementGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnChannelProps",
		reflect.TypeOf((*CfnChannelProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnCloudWatchAlarmTemplate",
		reflect.TypeOf((*CfnCloudWatchAlarmTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroupId", GoGetter: "AttrGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentifier", GoGetter: "AttrIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "attrModifiedAt", GoGetter: "AttrModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cloudWatchAlarmTemplateRef", GoGetter: "CloudWatchAlarmTemplateRef"},
			_jsii_.MemberProperty{JsiiProperty: "comparisonOperator", GoGetter: "ComparisonOperator"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "datapointsToAlarm", GoGetter: "DatapointsToAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriods", GoGetter: "EvaluationPeriods"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupIdentifier", GoGetter: "GroupIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "metricName", GoGetter: "MetricName"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "targetResourceType", GoGetter: "TargetResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "treatMissingData", GoGetter: "TreatMissingData"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudWatchAlarmTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveICloudWatchAlarmTemplateRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnCloudWatchAlarmTemplateGroup",
		reflect.TypeOf((*CfnCloudWatchAlarmTemplateGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentifier", GoGetter: "AttrIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "attrModifiedAt", GoGetter: "AttrModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cloudWatchAlarmTemplateGroupRef", GoGetter: "CloudWatchAlarmTemplateGroupRef"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCloudWatchAlarmTemplateGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveICloudWatchAlarmTemplateGroupRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnCloudWatchAlarmTemplateGroupProps",
		reflect.TypeOf((*CfnCloudWatchAlarmTemplateGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnCloudWatchAlarmTemplateProps",
		reflect.TypeOf((*CfnCloudWatchAlarmTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnCluster",
		reflect.TypeOf((*CfnCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrChannelIds", GoGetter: "AttrChannelIds"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterRef", GoGetter: "ClusterRef"},
			_jsii_.MemberProperty{JsiiProperty: "clusterType", GoGetter: "ClusterType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceRoleArn", GoGetter: "InstanceRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkSettings", GoGetter: "NetworkSettings"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveIClusterRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnCluster.ClusterNetworkSettingsProperty",
		reflect.TypeOf((*CfnCluster_ClusterNetworkSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnCluster.InterfaceMappingProperty",
		reflect.TypeOf((*CfnCluster_InterfaceMappingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnClusterProps",
		reflect.TypeOf((*CfnClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnEventBridgeRuleTemplate",
		reflect.TypeOf((*CfnEventBridgeRuleTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrGroupId", GoGetter: "AttrGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentifier", GoGetter: "AttrIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "attrModifiedAt", GoGetter: "AttrModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "eventBridgeRuleTemplateRef", GoGetter: "EventBridgeRuleTemplateRef"},
			_jsii_.MemberProperty{JsiiProperty: "eventTargets", GoGetter: "EventTargets"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "groupIdentifier", GoGetter: "GroupIdentifier"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventBridgeRuleTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveIEventBridgeRuleTemplateRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnEventBridgeRuleTemplate.EventBridgeRuleTemplateTargetProperty",
		reflect.TypeOf((*CfnEventBridgeRuleTemplate_EventBridgeRuleTemplateTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnEventBridgeRuleTemplateGroup",
		reflect.TypeOf((*CfnEventBridgeRuleTemplateGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentifier", GoGetter: "AttrIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "attrModifiedAt", GoGetter: "AttrModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "eventBridgeRuleTemplateGroupRef", GoGetter: "EventBridgeRuleTemplateGroupRef"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEventBridgeRuleTemplateGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveIEventBridgeRuleTemplateGroupRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnEventBridgeRuleTemplateGroupProps",
		reflect.TypeOf((*CfnEventBridgeRuleTemplateGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnEventBridgeRuleTemplateProps",
		reflect.TypeOf((*CfnEventBridgeRuleTemplateProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnInput",
		reflect.TypeOf((*CfnInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrDestinations", GoGetter: "AttrDestinations"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrSources", GoGetter: "AttrSources"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinations", GoGetter: "Destinations"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "inputDevices", GoGetter: "InputDevices"},
			_jsii_.MemberProperty{JsiiProperty: "inputNetworkLocation", GoGetter: "InputNetworkLocation"},
			_jsii_.MemberProperty{JsiiProperty: "inputRef", GoGetter: "InputRef"},
			_jsii_.MemberProperty{JsiiProperty: "inputSecurityGroups", GoGetter: "InputSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "mediaConnectFlows", GoGetter: "MediaConnectFlows"},
			_jsii_.MemberProperty{JsiiProperty: "multicastSettings", GoGetter: "MulticastSettings"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "routerSettings", GoGetter: "RouterSettings"},
			_jsii_.MemberProperty{JsiiProperty: "sdiSources", GoGetter: "SdiSources"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "smpte2110ReceiverGroupSettings", GoGetter: "Smpte2110ReceiverGroupSettings"},
			_jsii_.MemberProperty{JsiiProperty: "sources", GoGetter: "Sources"},
			_jsii_.MemberProperty{JsiiProperty: "srtSettings", GoGetter: "SrtSettings"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInput{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveIInputRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.InputDestinationRequestProperty",
		reflect.TypeOf((*CfnInput_InputDestinationRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.InputDeviceRequestProperty",
		reflect.TypeOf((*CfnInput_InputDeviceRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.InputDeviceSettingsProperty",
		reflect.TypeOf((*CfnInput_InputDeviceSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.InputRequestDestinationRouteProperty",
		reflect.TypeOf((*CfnInput_InputRequestDestinationRouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.InputSdpLocationProperty",
		reflect.TypeOf((*CfnInput_InputSdpLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.InputSourceRequestProperty",
		reflect.TypeOf((*CfnInput_InputSourceRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.InputVpcRequestProperty",
		reflect.TypeOf((*CfnInput_InputVpcRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.MediaConnectFlowRequestProperty",
		reflect.TypeOf((*CfnInput_MediaConnectFlowRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.MulticastSettingsCreateRequestProperty",
		reflect.TypeOf((*CfnInput_MulticastSettingsCreateRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.MulticastSourceCreateRequestProperty",
		reflect.TypeOf((*CfnInput_MulticastSourceCreateRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.RouterDestinationSettingsProperty",
		reflect.TypeOf((*CfnInput_RouterDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.RouterSettingsProperty",
		reflect.TypeOf((*CfnInput_RouterSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.Smpte2110ReceiverGroupProperty",
		reflect.TypeOf((*CfnInput_Smpte2110ReceiverGroupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.Smpte2110ReceiverGroupSdpSettingsProperty",
		reflect.TypeOf((*CfnInput_Smpte2110ReceiverGroupSdpSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.Smpte2110ReceiverGroupSettingsProperty",
		reflect.TypeOf((*CfnInput_Smpte2110ReceiverGroupSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.SrtCallerDecryptionRequestProperty",
		reflect.TypeOf((*CfnInput_SrtCallerDecryptionRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.SrtCallerSourceRequestProperty",
		reflect.TypeOf((*CfnInput_SrtCallerSourceRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInput.SrtSettingsRequestProperty",
		reflect.TypeOf((*CfnInput_SrtSettingsRequestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInputProps",
		reflect.TypeOf((*CfnInputProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup",
		reflect.TypeOf((*CfnInputSecurityGroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "inputSecurityGroupRef", GoGetter: "InputSecurityGroupRef"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRaw", GoGetter: "TagsRaw"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "whitelistRules", GoGetter: "WhitelistRules"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInputSecurityGroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveIInputSecurityGroupRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroup.InputWhitelistRuleCidrProperty",
		reflect.TypeOf((*CfnInputSecurityGroup_InputWhitelistRuleCidrProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnInputSecurityGroupProps",
		reflect.TypeOf((*CfnInputSecurityGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnMultiplex",
		reflect.TypeOf((*CfnMultiplex)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrPipelinesRunningCount", GoGetter: "AttrPipelinesRunningCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrProgramCount", GoGetter: "AttrProgramCount"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinations", GoGetter: "Destinations"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "multiplexRef", GoGetter: "MultiplexRef"},
			_jsii_.MemberProperty{JsiiProperty: "multiplexSettings", GoGetter: "MultiplexSettings"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMultiplex{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveIMultiplexRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplex.MultiplexMediaConnectOutputDestinationSettingsProperty",
		reflect.TypeOf((*CfnMultiplex_MultiplexMediaConnectOutputDestinationSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplex.MultiplexOutputDestinationProperty",
		reflect.TypeOf((*CfnMultiplex_MultiplexOutputDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplex.MultiplexSettingsProperty",
		reflect.TypeOf((*CfnMultiplex_MultiplexSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplexProps",
		reflect.TypeOf((*CfnMultiplexProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram",
		reflect.TypeOf((*CfnMultiplexprogram)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrChannelId", GoGetter: "AttrChannelId"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "multiplexId", GoGetter: "MultiplexId"},
			_jsii_.MemberProperty{JsiiProperty: "multiplexprogramRef", GoGetter: "MultiplexprogramRef"},
			_jsii_.MemberProperty{JsiiProperty: "multiplexProgramSettings", GoGetter: "MultiplexProgramSettings"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "packetIdentifiersMap", GoGetter: "PacketIdentifiersMap"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineDetails", GoGetter: "PipelineDetails"},
			_jsii_.MemberProperty{JsiiProperty: "preferredChannelPipeline", GoGetter: "PreferredChannelPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "programName", GoGetter: "ProgramName"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMultiplexprogram{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveIMultiplexprogramRef)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram.MultiplexProgramPacketIdentifiersMapProperty",
		reflect.TypeOf((*CfnMultiplexprogram_MultiplexProgramPacketIdentifiersMapProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram.MultiplexProgramPipelineDetailProperty",
		reflect.TypeOf((*CfnMultiplexprogram_MultiplexProgramPipelineDetailProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram.MultiplexProgramServiceDescriptorProperty",
		reflect.TypeOf((*CfnMultiplexprogram_MultiplexProgramServiceDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram.MultiplexProgramSettingsProperty",
		reflect.TypeOf((*CfnMultiplexprogram_MultiplexProgramSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram.MultiplexStatmuxVideoSettingsProperty",
		reflect.TypeOf((*CfnMultiplexprogram_MultiplexStatmuxVideoSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogram.MultiplexVideoSettingsProperty",
		reflect.TypeOf((*CfnMultiplexprogram_MultiplexVideoSettingsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnMultiplexprogramProps",
		reflect.TypeOf((*CfnMultiplexprogramProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnNetwork",
		reflect.TypeOf((*CfnNetwork)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrAssociatedClusterIds", GoGetter: "AttrAssociatedClusterIds"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "ipPools", GoGetter: "IpPools"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "networkRef", GoGetter: "NetworkRef"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "routes", GoGetter: "Routes"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNetwork{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveINetworkRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnNetwork.IpPoolProperty",
		reflect.TypeOf((*CfnNetwork_IpPoolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnNetwork.RouteProperty",
		reflect.TypeOf((*CfnNetwork_RouteProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnNetworkProps",
		reflect.TypeOf((*CfnNetworkProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnSdiSource",
		reflect.TypeOf((*CfnSdiSource)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrInputs", GoGetter: "AttrInputs"},
			_jsii_.MemberProperty{JsiiProperty: "attrState", GoGetter: "AttrState"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "mode", GoGetter: "Mode"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "sdiSourceRef", GoGetter: "SdiSourceRef"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSdiSource{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveISdiSourceRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnSdiSourceProps",
		reflect.TypeOf((*CfnSdiSourceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_medialive.CfnSignalMap",
		reflect.TypeOf((*CfnSignalMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCloudWatchAlarmTemplateGroupIds", GoGetter: "AttrCloudWatchAlarmTemplateGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreatedAt", GoGetter: "AttrCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrErrorMessage", GoGetter: "AttrErrorMessage"},
			_jsii_.MemberProperty{JsiiProperty: "attrEventBridgeRuleTemplateGroupIds", GoGetter: "AttrEventBridgeRuleTemplateGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "attrFailedMediaResourceMap", GoGetter: "AttrFailedMediaResourceMap"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrIdentifier", GoGetter: "AttrIdentifier"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastDiscoveredAt", GoGetter: "AttrLastDiscoveredAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrLastSuccessfulMonitorDeployment", GoGetter: "AttrLastSuccessfulMonitorDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "attrMediaResourceMap", GoGetter: "AttrMediaResourceMap"},
			_jsii_.MemberProperty{JsiiProperty: "attrModifiedAt", GoGetter: "AttrModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "attrMonitorChangesPendingDeployment", GoGetter: "AttrMonitorChangesPendingDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "attrMonitorDeployment", GoGetter: "AttrMonitorDeployment"},
			_jsii_.MemberProperty{JsiiProperty: "attrStatus", GoGetter: "AttrStatus"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "cloudWatchAlarmTemplateGroupIdentifiers", GoGetter: "CloudWatchAlarmTemplateGroupIdentifiers"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "discoveryEntryPointArn", GoGetter: "DiscoveryEntryPointArn"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "eventBridgeRuleTemplateGroupIdentifiers", GoGetter: "EventBridgeRuleTemplateGroupIdentifiers"},
			_jsii_.MemberProperty{JsiiProperty: "forceRediscovery", GoGetter: "ForceRediscovery"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "signalMapRef", GoGetter: "SignalMapRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSignalMap{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsmedialiveISignalMapRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnSignalMap.MediaResourceNeighborProperty",
		reflect.TypeOf((*CfnSignalMap_MediaResourceNeighborProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnSignalMap.MediaResourceProperty",
		reflect.TypeOf((*CfnSignalMap_MediaResourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnSignalMap.MonitorDeploymentProperty",
		reflect.TypeOf((*CfnSignalMap_MonitorDeploymentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnSignalMap.SuccessfulMonitorDeploymentProperty",
		reflect.TypeOf((*CfnSignalMap_SuccessfulMonitorDeploymentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_medialive.CfnSignalMapProps",
		reflect.TypeOf((*CfnSignalMapProps)(nil)).Elem(),
	)
}

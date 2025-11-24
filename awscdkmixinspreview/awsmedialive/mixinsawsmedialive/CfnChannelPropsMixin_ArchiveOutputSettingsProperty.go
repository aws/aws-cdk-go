package mixinsawsmedialive


// The archive output settings.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   archiveOutputSettingsProperty := &ArchiveOutputSettingsProperty{
//   	ContainerSettings: &ArchiveContainerSettingsProperty{
//   		M2TsSettings: &M2tsSettingsProperty{
//   			AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   			Arib: jsii.String("arib"),
//   			AribCaptionsPid: jsii.String("aribCaptionsPid"),
//   			AribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   			AudioBufferModel: jsii.String("audioBufferModel"),
//   			AudioFramesPerPes: jsii.Number(123),
//   			AudioPids: jsii.String("audioPids"),
//   			AudioStreamType: jsii.String("audioStreamType"),
//   			Bitrate: jsii.Number(123),
//   			BufferModel: jsii.String("bufferModel"),
//   			CcDescriptor: jsii.String("ccDescriptor"),
//   			DvbNitSettings: &DvbNitSettingsProperty{
//   				NetworkId: jsii.Number(123),
//   				NetworkName: jsii.String("networkName"),
//   				RepInterval: jsii.Number(123),
//   			},
//   			DvbSdtSettings: &DvbSdtSettingsProperty{
//   				OutputSdt: jsii.String("outputSdt"),
//   				RepInterval: jsii.Number(123),
//   				ServiceName: jsii.String("serviceName"),
//   				ServiceProviderName: jsii.String("serviceProviderName"),
//   			},
//   			DvbSubPids: jsii.String("dvbSubPids"),
//   			DvbTdtSettings: &DvbTdtSettingsProperty{
//   				RepInterval: jsii.Number(123),
//   			},
//   			DvbTeletextPid: jsii.String("dvbTeletextPid"),
//   			Ebif: jsii.String("ebif"),
//   			EbpAudioInterval: jsii.String("ebpAudioInterval"),
//   			EbpLookaheadMs: jsii.Number(123),
//   			EbpPlacement: jsii.String("ebpPlacement"),
//   			EcmPid: jsii.String("ecmPid"),
//   			EsRateInPes: jsii.String("esRateInPes"),
//   			EtvPlatformPid: jsii.String("etvPlatformPid"),
//   			EtvSignalPid: jsii.String("etvSignalPid"),
//   			FragmentTime: jsii.Number(123),
//   			Klv: jsii.String("klv"),
//   			KlvDataPids: jsii.String("klvDataPids"),
//   			NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   			NullPacketBitrate: jsii.Number(123),
//   			PatInterval: jsii.Number(123),
//   			PcrControl: jsii.String("pcrControl"),
//   			PcrPeriod: jsii.Number(123),
//   			PcrPid: jsii.String("pcrPid"),
//   			PmtInterval: jsii.Number(123),
//   			PmtPid: jsii.String("pmtPid"),
//   			ProgramNum: jsii.Number(123),
//   			RateMode: jsii.String("rateMode"),
//   			Scte27Pids: jsii.String("scte27Pids"),
//   			Scte35Control: jsii.String("scte35Control"),
//   			Scte35Pid: jsii.String("scte35Pid"),
//   			Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   			SegmentationMarkers: jsii.String("segmentationMarkers"),
//   			SegmentationStyle: jsii.String("segmentationStyle"),
//   			SegmentationTime: jsii.Number(123),
//   			TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   			TimedMetadataPid: jsii.String("timedMetadataPid"),
//   			TransportStreamId: jsii.Number(123),
//   			VideoPid: jsii.String("videoPid"),
//   		},
//   		RawSettings: &RawSettingsProperty{
//   		},
//   	},
//   	Extension: jsii.String("extension"),
//   	NameModifier: jsii.String("nameModifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archiveoutputsettings.html
//
type CfnChannelPropsMixin_ArchiveOutputSettingsProperty struct {
	// The settings that are specific to the container type of the file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archiveoutputsettings.html#cfn-medialive-channel-archiveoutputsettings-containersettings
	//
	ContainerSettings interface{} `field:"optional" json:"containerSettings" yaml:"containerSettings"`
	// The output file extension.
	//
	// If excluded, this is auto-selected from the container type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archiveoutputsettings.html#cfn-medialive-channel-archiveoutputsettings-extension
	//
	Extension *string `field:"optional" json:"extension" yaml:"extension"`
	// A string that is concatenated to the end of the destination file name.
	//
	// The string is required for multiple outputs of the same type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archiveoutputsettings.html#cfn-medialive-channel-archiveoutputsettings-namemodifier
	//
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}


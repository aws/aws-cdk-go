package awsmedialive


// The archive container settings.
//
// The parent of this entity is ArchiveOutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   archiveContainerSettingsProperty := &ArchiveContainerSettingsProperty{
//   	M2TsSettings: &M2tsSettingsProperty{
//   		AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   		Arib: jsii.String("arib"),
//   		AribCaptionsPid: jsii.String("aribCaptionsPid"),
//   		AribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   		AudioBufferModel: jsii.String("audioBufferModel"),
//   		AudioFramesPerPes: jsii.Number(123),
//   		AudioPids: jsii.String("audioPids"),
//   		AudioStreamType: jsii.String("audioStreamType"),
//   		Bitrate: jsii.Number(123),
//   		BufferModel: jsii.String("bufferModel"),
//   		CcDescriptor: jsii.String("ccDescriptor"),
//   		DvbNitSettings: &DvbNitSettingsProperty{
//   			NetworkId: jsii.Number(123),
//   			NetworkName: jsii.String("networkName"),
//   			RepInterval: jsii.Number(123),
//   		},
//   		DvbSdtSettings: &DvbSdtSettingsProperty{
//   			OutputSdt: jsii.String("outputSdt"),
//   			RepInterval: jsii.Number(123),
//   			ServiceName: jsii.String("serviceName"),
//   			ServiceProviderName: jsii.String("serviceProviderName"),
//   		},
//   		DvbSubPids: jsii.String("dvbSubPids"),
//   		DvbTdtSettings: &DvbTdtSettingsProperty{
//   			RepInterval: jsii.Number(123),
//   		},
//   		DvbTeletextPid: jsii.String("dvbTeletextPid"),
//   		Ebif: jsii.String("ebif"),
//   		EbpAudioInterval: jsii.String("ebpAudioInterval"),
//   		EbpLookaheadMs: jsii.Number(123),
//   		EbpPlacement: jsii.String("ebpPlacement"),
//   		EcmPid: jsii.String("ecmPid"),
//   		EsRateInPes: jsii.String("esRateInPes"),
//   		EtvPlatformPid: jsii.String("etvPlatformPid"),
//   		EtvSignalPid: jsii.String("etvSignalPid"),
//   		FragmentTime: jsii.Number(123),
//   		Klv: jsii.String("klv"),
//   		KlvDataPids: jsii.String("klvDataPids"),
//   		NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   		NullPacketBitrate: jsii.Number(123),
//   		PatInterval: jsii.Number(123),
//   		PcrControl: jsii.String("pcrControl"),
//   		PcrPeriod: jsii.Number(123),
//   		PcrPid: jsii.String("pcrPid"),
//   		PmtInterval: jsii.Number(123),
//   		PmtPid: jsii.String("pmtPid"),
//   		ProgramNum: jsii.Number(123),
//   		RateMode: jsii.String("rateMode"),
//   		Scte27Pids: jsii.String("scte27Pids"),
//   		Scte35Control: jsii.String("scte35Control"),
//   		Scte35Pid: jsii.String("scte35Pid"),
//   		Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   		SegmentationMarkers: jsii.String("segmentationMarkers"),
//   		SegmentationStyle: jsii.String("segmentationStyle"),
//   		SegmentationTime: jsii.Number(123),
//   		TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   		TimedMetadataPid: jsii.String("timedMetadataPid"),
//   		TransportStreamId: jsii.Number(123),
//   		VideoPid: jsii.String("videoPid"),
//   	},
//   	RawSettings: &RawSettingsProperty{
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecontainersettings.html
//
type CfnChannel_ArchiveContainerSettingsProperty struct {
	// The settings for the M2TS in the archive output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecontainersettings.html#cfn-medialive-channel-archivecontainersettings-m2tssettings
	//
	M2TsSettings interface{} `field:"optional" json:"m2TsSettings" yaml:"m2TsSettings"`
	// The settings for Raw archive output type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-archivecontainersettings.html#cfn-medialive-channel-archivecontainersettings-rawsettings
	//
	RawSettings interface{} `field:"optional" json:"rawSettings" yaml:"rawSettings"`
}


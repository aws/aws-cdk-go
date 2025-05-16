package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtOutputSettingsProperty := &SrtOutputSettingsProperty{
//   	BufferMsec: jsii.Number(123),
//   	ContainerSettings: &UdpContainerSettingsProperty{
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
//   	},
//   	Destination: &OutputLocationRefProperty{
//   		DestinationRefId: jsii.String("destinationRefId"),
//   	},
//   	EncryptionType: jsii.String("encryptionType"),
//   	Latency: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputsettings.html
//
type CfnChannel_SrtOutputSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputsettings.html#cfn-medialive-channel-srtoutputsettings-buffermsec
	//
	BufferMsec *float64 `field:"optional" json:"bufferMsec" yaml:"bufferMsec"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputsettings.html#cfn-medialive-channel-srtoutputsettings-containersettings
	//
	ContainerSettings interface{} `field:"optional" json:"containerSettings" yaml:"containerSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputsettings.html#cfn-medialive-channel-srtoutputsettings-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputsettings.html#cfn-medialive-channel-srtoutputsettings-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-srtoutputsettings.html#cfn-medialive-channel-srtoutputsettings-latency
	//
	Latency *float64 `field:"optional" json:"latency" yaml:"latency"`
}


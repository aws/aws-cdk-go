package awsmedialive


// Settings for the M3U8 container.
//
// The parent of this entity is StandardHlsSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   m3u8SettingsProperty := &M3u8SettingsProperty{
//   	AudioFramesPerPes: jsii.Number(123),
//   	AudioPids: jsii.String("audioPids"),
//   	EcmPid: jsii.String("ecmPid"),
//   	NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   	PatInterval: jsii.Number(123),
//   	PcrControl: jsii.String("pcrControl"),
//   	PcrPeriod: jsii.Number(123),
//   	PcrPid: jsii.String("pcrPid"),
//   	PmtInterval: jsii.Number(123),
//   	PmtPid: jsii.String("pmtPid"),
//   	ProgramNum: jsii.Number(123),
//   	Scte35Behavior: jsii.String("scte35Behavior"),
//   	Scte35Pid: jsii.String("scte35Pid"),
//   	TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   	TimedMetadataPid: jsii.String("timedMetadataPid"),
//   	TransportStreamId: jsii.Number(123),
//   	VideoPid: jsii.String("videoPid"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html
//
type CfnChannel_M3u8SettingsProperty struct {
	// The number of audio frames to insert for each PES packet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-audioframesperpes
	//
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// The PID of the elementary audio streams in the transport stream.
	//
	// Multiple values are accepted, and can be entered in ranges or by comma separation. You can enter the value as a decimal or hexadecimal value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-audiopids
	//
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// This parameter is unused and deprecated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-ecmpid
	//
	EcmPid *string `field:"optional" json:"ecmPid" yaml:"ecmPid"`
	// If set to passthrough, Nielsen inaudible tones for media tracking will be detected in the input audio and an equivalent ID3 tag will be inserted in the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-nielsenid3behavior
	//
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// The number of milliseconds between instances of this table in the output transport stream.
	//
	// A value of \"0\" writes out the PMT once per segment file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-patinterval
	//
	PatInterval *float64 `field:"optional" json:"patInterval" yaml:"patInterval"`
	// When set to pcrEveryPesPacket, a Program Clock Reference value is inserted for every Packetized Elementary Stream (PES) header.
	//
	// This parameter is effective only when the PCR PID is the same as the video or audio elementary stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-pcrcontrol
	//
	PcrControl *string `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// The maximum time, in milliseconds, between Program Clock References (PCRs) inserted into the transport stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-pcrperiod
	//
	PcrPeriod *float64 `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// The PID of the Program Clock Reference (PCR) in the transport stream.
	//
	// When no value is given, MediaLive assigns the same value as the video PID. You can enter the value as a decimal or hexadecimal value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-pcrpid
	//
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// The number of milliseconds between instances of this table in the output transport stream.
	//
	// A value of \"0\" writes out the PMT once per segment file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-pmtinterval
	//
	PmtInterval *float64 `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// The PID for the Program Map Table (PMT) in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-pmtpid
	//
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// The value of the program number field in the Program Map Table (PMT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-programnum
	//
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// If set to passthrough, passes any SCTE-35 signals from the input source to this output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-scte35behavior
	//
	Scte35Behavior *string `field:"optional" json:"scte35Behavior" yaml:"scte35Behavior"`
	// The PID of the SCTE-35 stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-scte35pid
	//
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// When set to passthrough, timed metadata is passed through from input to output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-timedmetadatabehavior
	//
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
	// The PID of the timed metadata stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-timedmetadatapid
	//
	TimedMetadataPid *string `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// The value of the transport stream ID field in the Program Map Table (PMT).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-transportstreamid
	//
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// The PID of the elementary video stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-m3u8settings.html#cfn-medialive-channel-m3u8settings-videopid
	//
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}


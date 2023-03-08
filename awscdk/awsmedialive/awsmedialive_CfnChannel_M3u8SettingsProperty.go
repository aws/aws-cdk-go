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
//   m3u8SettingsProperty := &m3u8SettingsProperty{
//   	audioFramesPerPes: jsii.Number(123),
//   	audioPids: jsii.String("audioPids"),
//   	ecmPid: jsii.String("ecmPid"),
//   	nielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   	patInterval: jsii.Number(123),
//   	pcrControl: jsii.String("pcrControl"),
//   	pcrPeriod: jsii.Number(123),
//   	pcrPid: jsii.String("pcrPid"),
//   	pmtInterval: jsii.Number(123),
//   	pmtPid: jsii.String("pmtPid"),
//   	programNum: jsii.Number(123),
//   	scte35Behavior: jsii.String("scte35Behavior"),
//   	scte35Pid: jsii.String("scte35Pid"),
//   	timedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   	timedMetadataPid: jsii.String("timedMetadataPid"),
//   	transportStreamId: jsii.Number(123),
//   	videoPid: jsii.String("videoPid"),
//   }
//
type CfnChannel_M3u8SettingsProperty struct {
	// The number of audio frames to insert for each PES packet.
	AudioFramesPerPes *float64 `field:"optional" json:"audioFramesPerPes" yaml:"audioFramesPerPes"`
	// The PID of the elementary audio streams in the transport stream.
	//
	// Multiple values are accepted, and can be entered in ranges or by comma separation. You can enter the value as a decimal or hexadecimal value.
	AudioPids *string `field:"optional" json:"audioPids" yaml:"audioPids"`
	// This parameter is unused and deprecated.
	EcmPid *string `field:"optional" json:"ecmPid" yaml:"ecmPid"`
	// If set to passthrough, Nielsen inaudible tones for media tracking will be detected in the input audio and an equivalent ID3 tag will be inserted in the output.
	NielsenId3Behavior *string `field:"optional" json:"nielsenId3Behavior" yaml:"nielsenId3Behavior"`
	// The number of milliseconds between instances of this table in the output transport stream.
	//
	// A value of \"0\" writes out the PMT once per segment file.
	PatInterval *float64 `field:"optional" json:"patInterval" yaml:"patInterval"`
	// When set to pcrEveryPesPacket, a Program Clock Reference value is inserted for every Packetized Elementary Stream (PES) header.
	//
	// This parameter is effective only when the PCR PID is the same as the video or audio elementary stream.
	PcrControl *string `field:"optional" json:"pcrControl" yaml:"pcrControl"`
	// The maximum time, in milliseconds, between Program Clock References (PCRs) inserted into the transport stream.
	PcrPeriod *float64 `field:"optional" json:"pcrPeriod" yaml:"pcrPeriod"`
	// The PID of the Program Clock Reference (PCR) in the transport stream.
	//
	// When no value is given, MediaLive assigns the same value as the video PID. You can enter the value as a decimal or hexadecimal value.
	PcrPid *string `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// The number of milliseconds between instances of this table in the output transport stream.
	//
	// A value of \"0\" writes out the PMT once per segment file.
	PmtInterval *float64 `field:"optional" json:"pmtInterval" yaml:"pmtInterval"`
	// The PID for the Program Map Table (PMT) in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value.
	PmtPid *string `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// The value of the program number field in the Program Map Table (PMT).
	ProgramNum *float64 `field:"optional" json:"programNum" yaml:"programNum"`
	// If set to passthrough, passes any SCTE-35 signals from the input source to this output.
	Scte35Behavior *string `field:"optional" json:"scte35Behavior" yaml:"scte35Behavior"`
	// The PID of the SCTE-35 stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value.
	Scte35Pid *string `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// When set to passthrough, timed metadata is passed through from input to output.
	TimedMetadataBehavior *string `field:"optional" json:"timedMetadataBehavior" yaml:"timedMetadataBehavior"`
	// The PID of the timed metadata stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value. Valid values are 32 (or 0x20)..8182 (or 0x1ff6).
	TimedMetadataPid *string `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// The value of the transport stream ID field in the Program Map Table (PMT).
	TransportStreamId *float64 `field:"optional" json:"transportStreamId" yaml:"transportStreamId"`
	// The PID of the elementary video stream in the transport stream.
	//
	// You can enter the value as a decimal or hexadecimal value.
	VideoPid *string `field:"optional" json:"videoPid" yaml:"videoPid"`
}


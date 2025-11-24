package mixinsawsmedialive


// Packet identifiers map for a given Multiplex program.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiplexProgramPacketIdentifiersMapProperty := &MultiplexProgramPacketIdentifiersMapProperty{
//   	AudioPids: []interface{}{
//   		jsii.Number(123),
//   	},
//   	DvbSubPids: []interface{}{
//   		jsii.Number(123),
//   	},
//   	DvbTeletextPid: jsii.Number(123),
//   	EtvPlatformPid: jsii.Number(123),
//   	EtvSignalPid: jsii.Number(123),
//   	KlvDataPids: []interface{}{
//   		jsii.Number(123),
//   	},
//   	PcrPid: jsii.Number(123),
//   	PmtPid: jsii.Number(123),
//   	PrivateMetadataPid: jsii.Number(123),
//   	Scte27Pids: []interface{}{
//   		jsii.Number(123),
//   	},
//   	Scte35Pid: jsii.Number(123),
//   	TimedMetadataPid: jsii.Number(123),
//   	VideoPid: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html
//
type CfnMultiplexprogramPropsMixin_MultiplexProgramPacketIdentifiersMapProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-audiopids
	//
	AudioPids interface{} `field:"optional" json:"audioPids" yaml:"audioPids"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-dvbsubpids
	//
	DvbSubPids interface{} `field:"optional" json:"dvbSubPids" yaml:"dvbSubPids"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-dvbteletextpid
	//
	DvbTeletextPid *float64 `field:"optional" json:"dvbTeletextPid" yaml:"dvbTeletextPid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-etvplatformpid
	//
	EtvPlatformPid *float64 `field:"optional" json:"etvPlatformPid" yaml:"etvPlatformPid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-etvsignalpid
	//
	EtvSignalPid *float64 `field:"optional" json:"etvSignalPid" yaml:"etvSignalPid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-klvdatapids
	//
	KlvDataPids interface{} `field:"optional" json:"klvDataPids" yaml:"klvDataPids"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-pcrpid
	//
	PcrPid *float64 `field:"optional" json:"pcrPid" yaml:"pcrPid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-pmtpid
	//
	PmtPid *float64 `field:"optional" json:"pmtPid" yaml:"pmtPid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-privatemetadatapid
	//
	PrivateMetadataPid *float64 `field:"optional" json:"privateMetadataPid" yaml:"privateMetadataPid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-scte27pids
	//
	Scte27Pids interface{} `field:"optional" json:"scte27Pids" yaml:"scte27Pids"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-scte35pid
	//
	Scte35Pid *float64 `field:"optional" json:"scte35Pid" yaml:"scte35Pid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-timedmetadatapid
	//
	TimedMetadataPid *float64 `field:"optional" json:"timedMetadataPid" yaml:"timedMetadataPid"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap.html#cfn-medialive-multiplexprogram-multiplexprogrampacketidentifiersmap-videopid
	//
	VideoPid *float64 `field:"optional" json:"videoPid" yaml:"videoPid"`
}


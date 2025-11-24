package mixinsawsmedialive


// Properties for CfnMultiplexprogramPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMultiplexprogramMixinProps := &CfnMultiplexprogramMixinProps{
//   	MultiplexId: jsii.String("multiplexId"),
//   	MultiplexProgramSettings: &MultiplexProgramSettingsProperty{
//   		PreferredChannelPipeline: jsii.String("preferredChannelPipeline"),
//   		ProgramNumber: jsii.Number(123),
//   		ServiceDescriptor: &MultiplexProgramServiceDescriptorProperty{
//   			ProviderName: jsii.String("providerName"),
//   			ServiceName: jsii.String("serviceName"),
//   		},
//   		VideoSettings: &MultiplexVideoSettingsProperty{
//   			ConstantBitrate: jsii.Number(123),
//   			StatmuxSettings: &MultiplexStatmuxVideoSettingsProperty{
//   				MaximumBitrate: jsii.Number(123),
//   				MinimumBitrate: jsii.Number(123),
//   				Priority: jsii.Number(123),
//   			},
//   		},
//   	},
//   	PacketIdentifiersMap: &MultiplexProgramPacketIdentifiersMapProperty{
//   		AudioPids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		DvbSubPids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		DvbTeletextPid: jsii.Number(123),
//   		EtvPlatformPid: jsii.Number(123),
//   		EtvSignalPid: jsii.Number(123),
//   		KlvDataPids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		PcrPid: jsii.Number(123),
//   		PmtPid: jsii.Number(123),
//   		PrivateMetadataPid: jsii.Number(123),
//   		Scte27Pids: []interface{}{
//   			jsii.Number(123),
//   		},
//   		Scte35Pid: jsii.Number(123),
//   		TimedMetadataPid: jsii.Number(123),
//   		VideoPid: jsii.Number(123),
//   	},
//   	PipelineDetails: []interface{}{
//   		&MultiplexProgramPipelineDetailProperty{
//   			ActiveChannelPipeline: jsii.String("activeChannelPipeline"),
//   			PipelineId: jsii.String("pipelineId"),
//   		},
//   	},
//   	PreferredChannelPipeline: jsii.String("preferredChannelPipeline"),
//   	ProgramName: jsii.String("programName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplexprogram.html
//
type CfnMultiplexprogramMixinProps struct {
	// The unique id of the multiplex.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplexprogram.html#cfn-medialive-multiplexprogram-multiplexid
	//
	MultiplexId *string `field:"optional" json:"multiplexId" yaml:"multiplexId"`
	// Multiplex Program settings configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplexprogram.html#cfn-medialive-multiplexprogram-multiplexprogramsettings
	//
	MultiplexProgramSettings interface{} `field:"optional" json:"multiplexProgramSettings" yaml:"multiplexProgramSettings"`
	// Packet identifiers map for a given Multiplex program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplexprogram.html#cfn-medialive-multiplexprogram-packetidentifiersmap
	//
	PacketIdentifiersMap interface{} `field:"optional" json:"packetIdentifiersMap" yaml:"packetIdentifiersMap"`
	// Contains information about the current sources for the specified program in the specified multiplex.
	//
	// Keep in mind that each multiplex pipeline connects to both pipelines in a given source channel (the channel identified by the program). But only one of those channel pipelines is ever active at one time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplexprogram.html#cfn-medialive-multiplexprogram-pipelinedetails
	//
	PipelineDetails interface{} `field:"optional" json:"pipelineDetails" yaml:"pipelineDetails"`
	// Indicates which pipeline is preferred by the multiplex for program ingest.
	//
	// If set to \"PIPELINE_0\" or \"PIPELINE_1\" and an unhealthy ingest causes the multiplex to switch to the non-preferred pipeline,
	// it will switch back once that ingest is healthy again. If set to \"CURRENTLY_ACTIVE\",
	// it will not switch back to the other pipeline based on it recovering to a healthy state,
	// it will only switch if the active pipeline becomes unhealthy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplexprogram.html#cfn-medialive-multiplexprogram-preferredchannelpipeline
	//
	PreferredChannelPipeline *string `field:"optional" json:"preferredChannelPipeline" yaml:"preferredChannelPipeline"`
	// The name of the multiplex program.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-multiplexprogram.html#cfn-medialive-multiplexprogram-programname
	//
	ProgramName *string `field:"optional" json:"programName" yaml:"programName"`
}


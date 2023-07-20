package awsmedialive


// Destination settings for a Multiplex output.
//
// The parent of this entity is OutputDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiplexProgramChannelDestinationSettingsProperty := &MultiplexProgramChannelDestinationSettingsProperty{
//   	MultiplexId: jsii.String("multiplexId"),
//   	ProgramName: jsii.String("programName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexprogramchanneldestinationsettings.html
//
type CfnChannel_MultiplexProgramChannelDestinationSettingsProperty struct {
	// The ID of the Multiplex that the encoder is providing output to.
	//
	// You do not need to specify the individual inputs to the Multiplex; MediaLive will handle the connection of the two MediaLive pipelines to the two Multiplex instances.
	// The Multiplex must be in the same region as the Channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexprogramchanneldestinationsettings.html#cfn-medialive-channel-multiplexprogramchanneldestinationsettings-multiplexid
	//
	MultiplexId *string `field:"optional" json:"multiplexId" yaml:"multiplexId"`
	// The program name of the Multiplex program that the encoder is providing output to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-multiplexprogramchanneldestinationsettings.html#cfn-medialive-channel-multiplexprogramchanneldestinationsettings-programname
	//
	ProgramName *string `field:"optional" json:"programName" yaml:"programName"`
}


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
//   multiplexProgramChannelDestinationSettingsProperty := &multiplexProgramChannelDestinationSettingsProperty{
//   	multiplexId: jsii.String("multiplexId"),
//   	programName: jsii.String("programName"),
//   }
//
type CfnChannel_MultiplexProgramChannelDestinationSettingsProperty struct {
	// The ID of the Multiplex that the encoder is providing output to.
	//
	// You do not need to specify the individual inputs to the Multiplex; MediaLive will handle the connection of the two MediaLive pipelines to the two Multiplex instances.
	// The Multiplex must be in the same region as the Channel.
	MultiplexId *string `field:"optional" json:"multiplexId" yaml:"multiplexId"`
	// The program name of the Multiplex program that the encoder is providing output to.
	ProgramName *string `field:"optional" json:"programName" yaml:"programName"`
}


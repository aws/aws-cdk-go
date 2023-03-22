package awsmedialive


// Configuration information for an output.
//
// This entity is at the top level in the channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputDestinationProperty := &outputDestinationProperty{
//   	id: jsii.String("id"),
//   	mediaPackageSettings: []interface{}{
//   		&mediaPackageOutputDestinationSettingsProperty{
//   			channelId: jsii.String("channelId"),
//   		},
//   	},
//   	multiplexSettings: &multiplexProgramChannelDestinationSettingsProperty{
//   		multiplexId: jsii.String("multiplexId"),
//   		programName: jsii.String("programName"),
//   	},
//   	settings: []interface{}{
//   		&outputDestinationSettingsProperty{
//   			passwordParam: jsii.String("passwordParam"),
//   			streamName: jsii.String("streamName"),
//   			url: jsii.String("url"),
//   			username: jsii.String("username"),
//   		},
//   	},
//   }
//
type CfnChannel_OutputDestinationProperty struct {
	// The ID for this destination.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The destination settings for a MediaPackage output.
	MediaPackageSettings interface{} `field:"optional" json:"mediaPackageSettings" yaml:"mediaPackageSettings"`
	// Destination settings for a Multiplex output;
	//
	// one destination for both encoders.
	MultiplexSettings interface{} `field:"optional" json:"multiplexSettings" yaml:"multiplexSettings"`
	// The destination settings for an output.
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}


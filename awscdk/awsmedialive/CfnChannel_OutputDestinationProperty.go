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
//   outputDestinationProperty := &OutputDestinationProperty{
//   	Id: jsii.String("id"),
//   	LogicalInterfaceNames: []*string{
//   		jsii.String("logicalInterfaceNames"),
//   	},
//   	MediaPackageSettings: []interface{}{
//   		&MediaPackageOutputDestinationSettingsProperty{
//   			ChannelGroup: jsii.String("channelGroup"),
//   			ChannelId: jsii.String("channelId"),
//   			ChannelName: jsii.String("channelName"),
//   		},
//   	},
//   	MultiplexSettings: &MultiplexProgramChannelDestinationSettingsProperty{
//   		MultiplexId: jsii.String("multiplexId"),
//   		ProgramName: jsii.String("programName"),
//   	},
//   	Settings: []interface{}{
//   		&OutputDestinationSettingsProperty{
//   			PasswordParam: jsii.String("passwordParam"),
//   			StreamName: jsii.String("streamName"),
//   			Url: jsii.String("url"),
//   			Username: jsii.String("username"),
//   		},
//   	},
//   	SrtSettings: []interface{}{
//   		&SrtOutputDestinationSettingsProperty{
//   			EncryptionPassphraseSecretArn: jsii.String("encryptionPassphraseSecretArn"),
//   			StreamId: jsii.String("streamId"),
//   			Url: jsii.String("url"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestination.html
//
type CfnChannel_OutputDestinationProperty struct {
	// The ID for this destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestination.html#cfn-medialive-channel-outputdestination-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestination.html#cfn-medialive-channel-outputdestination-logicalinterfacenames
	//
	LogicalInterfaceNames *[]*string `field:"optional" json:"logicalInterfaceNames" yaml:"logicalInterfaceNames"`
	// The destination settings for a MediaPackage output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestination.html#cfn-medialive-channel-outputdestination-mediapackagesettings
	//
	MediaPackageSettings interface{} `field:"optional" json:"mediaPackageSettings" yaml:"mediaPackageSettings"`
	// Destination settings for a Multiplex output;
	//
	// one destination for both encoders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestination.html#cfn-medialive-channel-outputdestination-multiplexsettings
	//
	MultiplexSettings interface{} `field:"optional" json:"multiplexSettings" yaml:"multiplexSettings"`
	// The destination settings for an output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestination.html#cfn-medialive-channel-outputdestination-settings
	//
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestination.html#cfn-medialive-channel-outputdestination-srtsettings
	//
	SrtSettings interface{} `field:"optional" json:"srtSettings" yaml:"srtSettings"`
}


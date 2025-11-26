package previewawsmedialivemixins


// The configuration information for this output.
//
// The parent of this entity is OutputDestination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   outputDestinationSettingsProperty := &OutputDestinationSettingsProperty{
//   	PasswordParam: jsii.String("passwordParam"),
//   	StreamName: jsii.String("streamName"),
//   	Url: jsii.String("url"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestinationsettings.html
//
type CfnChannelPropsMixin_OutputDestinationSettingsProperty struct {
	// The password parameter that holds the password for accessing the downstream system.
	//
	// This password parameter applies only if the downstream system requires credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestinationsettings.html#cfn-medialive-channel-outputdestinationsettings-passwordparam
	//
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// The stream name for the content.
	//
	// This applies only to RTMP outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestinationsettings.html#cfn-medialive-channel-outputdestinationsettings-streamname
	//
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
	// The URL for the destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestinationsettings.html#cfn-medialive-channel-outputdestinationsettings-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The user name to connect to the downstream system.
	//
	// This applies only if the downstream system requires credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-outputdestinationsettings.html#cfn-medialive-channel-outputdestinationsettings-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}


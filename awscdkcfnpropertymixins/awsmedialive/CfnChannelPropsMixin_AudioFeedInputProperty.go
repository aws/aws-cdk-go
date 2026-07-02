package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   audioFeedInputProperty := &AudioFeedInputProperty{
//   	AudioSelectorName: jsii.String("audioSelectorName"),
//   	FeedInput: jsii.String("feedInput"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiofeedinput.html
//
type CfnChannelPropsMixin_AudioFeedInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiofeedinput.html#cfn-medialive-channel-audiofeedinput-audioselectorname
	//
	AudioSelectorName *string `field:"optional" json:"audioSelectorName" yaml:"audioSelectorName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiofeedinput.html#cfn-medialive-channel-audiofeedinput-feedinput
	//
	FeedInput *string `field:"optional" json:"feedInput" yaml:"feedInput"`
}


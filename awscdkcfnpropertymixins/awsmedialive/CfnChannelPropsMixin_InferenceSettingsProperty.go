package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   inferenceSettingsProperty := &InferenceSettingsProperty{
//   	AudioFeedInputs: []interface{}{
//   		&AudioFeedInputProperty{
//   			AudioSelectorName: jsii.String("audioSelectorName"),
//   			FeedInput: jsii.String("feedInput"),
//   		},
//   	},
//   	FeedArn: jsii.String("feedArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inferencesettings.html
//
type CfnChannelPropsMixin_InferenceSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inferencesettings.html#cfn-medialive-channel-inferencesettings-audiofeedinputs
	//
	AudioFeedInputs interface{} `field:"optional" json:"audioFeedInputs" yaml:"audioFeedInputs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inferencesettings.html#cfn-medialive-channel-inferencesettings-feedarn
	//
	FeedArn *string `field:"optional" json:"feedArn" yaml:"feedArn"`
}


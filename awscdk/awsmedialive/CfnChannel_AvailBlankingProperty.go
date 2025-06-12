package awsmedialive


// The configuration of ad avail blanking in the output.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   availBlankingProperty := &AvailBlankingProperty{
//   	AvailBlankingImage: &InputLocationProperty{
//   		PasswordParam: jsii.String("passwordParam"),
//   		Uri: jsii.String("uri"),
//   		Username: jsii.String("username"),
//   	},
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availblanking.html
//
type CfnChannel_AvailBlankingProperty struct {
	// The blanking image to be used.
	//
	// Keep empty for solid black. Only .bmp and .png images are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availblanking.html#cfn-medialive-channel-availblanking-availblankingimage
	//
	AvailBlankingImage interface{} `field:"optional" json:"availBlankingImage" yaml:"availBlankingImage"`
	// When set to enabled, the video, audio, and captions are blanked when insertion metadata is added.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-availblanking.html#cfn-medialive-channel-availblanking-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}


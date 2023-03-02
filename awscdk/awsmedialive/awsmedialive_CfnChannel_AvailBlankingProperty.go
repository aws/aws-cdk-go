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
//   availBlankingProperty := &availBlankingProperty{
//   	availBlankingImage: &inputLocationProperty{
//   		passwordParam: jsii.String("passwordParam"),
//   		uri: jsii.String("uri"),
//   		username: jsii.String("username"),
//   	},
//   	state: jsii.String("state"),
//   }
//
type CfnChannel_AvailBlankingProperty struct {
	// The blanking image to be used.
	//
	// Keep empty for solid black. Only .bmp and .png images are supported.
	AvailBlankingImage interface{} `field:"optional" json:"availBlankingImage" yaml:"availBlankingImage"`
	// When set to enabled, the video, audio, and captions are blanked when insertion metadata is added.
	State *string `field:"optional" json:"state" yaml:"state"`
}


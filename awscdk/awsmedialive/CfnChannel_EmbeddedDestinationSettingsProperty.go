package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   embeddedDestinationSettingsProperty := &EmbeddedDestinationSettingsProperty{
//   	Position: &EmbeddedCaptionPositionSettingsProperty{
//   		YPositionLine: jsii.Number(123),
//   	},
//   	StyleControl: jsii.String("styleControl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddeddestinationsettings.html
//
type CfnChannel_EmbeddedDestinationSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddeddestinationsettings.html#cfn-medialive-channel-embeddeddestinationsettings-position
	//
	Position interface{} `field:"optional" json:"position" yaml:"position"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddeddestinationsettings.html#cfn-medialive-channel-embeddeddestinationsettings-stylecontrol
	//
	StyleControl *string `field:"optional" json:"styleControl" yaml:"styleControl"`
}


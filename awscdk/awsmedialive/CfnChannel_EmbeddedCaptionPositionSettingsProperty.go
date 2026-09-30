package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   embeddedCaptionPositionSettingsProperty := &EmbeddedCaptionPositionSettingsProperty{
//   	YPositionLine: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedcaptionpositionsettings.html
//
type CfnChannel_EmbeddedCaptionPositionSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-embeddedcaptionpositionsettings.html#cfn-medialive-channel-embeddedcaptionpositionsettings-ypositionline
	//
	YPositionLine *float64 `field:"optional" json:"yPositionLine" yaml:"yPositionLine"`
}


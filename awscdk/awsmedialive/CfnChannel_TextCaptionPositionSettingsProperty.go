package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   textCaptionPositionSettingsProperty := &TextCaptionPositionSettingsProperty{
//   	YPositionPercentage: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-textcaptionpositionsettings.html
//
type CfnChannel_TextCaptionPositionSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-textcaptionpositionsettings.html#cfn-medialive-channel-textcaptionpositionsettings-ypositionpercentage
	//
	YPositionPercentage *float64 `field:"optional" json:"yPositionPercentage" yaml:"yPositionPercentage"`
}


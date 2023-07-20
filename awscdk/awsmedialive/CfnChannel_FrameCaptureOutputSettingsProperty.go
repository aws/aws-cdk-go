package awsmedialive


// The frame capture output settings.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   frameCaptureOutputSettingsProperty := &FrameCaptureOutputSettingsProperty{
//   	NameModifier: jsii.String("nameModifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecaptureoutputsettings.html
//
type CfnChannel_FrameCaptureOutputSettingsProperty struct {
	// Required if the output group contains more than one output.
	//
	// This modifier forms part of the output file name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecaptureoutputsettings.html#cfn-medialive-channel-framecaptureoutputsettings-namemodifier
	//
	NameModifier *string `field:"optional" json:"nameModifier" yaml:"nameModifier"`
}


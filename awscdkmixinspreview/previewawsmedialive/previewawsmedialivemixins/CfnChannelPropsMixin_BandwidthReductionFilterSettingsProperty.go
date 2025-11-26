package previewawsmedialivemixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bandwidthReductionFilterSettingsProperty := &BandwidthReductionFilterSettingsProperty{
//   	PostFilterSharpening: jsii.String("postFilterSharpening"),
//   	Strength: jsii.String("strength"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-bandwidthreductionfiltersettings.html
//
type CfnChannelPropsMixin_BandwidthReductionFilterSettingsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-bandwidthreductionfiltersettings.html#cfn-medialive-channel-bandwidthreductionfiltersettings-postfiltersharpening
	//
	PostFilterSharpening *string `field:"optional" json:"postFilterSharpening" yaml:"postFilterSharpening"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-bandwidthreductionfiltersettings.html#cfn-medialive-channel-bandwidthreductionfiltersettings-strength
	//
	Strength *string `field:"optional" json:"strength" yaml:"strength"`
}


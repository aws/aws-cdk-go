package previewawsmedialivemixins


// The video configuration for each program in a multiplex.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   multiplexVideoSettingsProperty := &MultiplexVideoSettingsProperty{
//   	ConstantBitrate: jsii.Number(123),
//   	StatmuxSettings: &MultiplexStatmuxVideoSettingsProperty{
//   		MaximumBitrate: jsii.Number(123),
//   		MinimumBitrate: jsii.Number(123),
//   		Priority: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexvideosettings.html
//
type CfnMultiplexprogramPropsMixin_MultiplexVideoSettingsProperty struct {
	// The constant bitrate configuration for the video encode.
	//
	// When this field is defined, StatmuxSettings must be undefined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexvideosettings.html#cfn-medialive-multiplexprogram-multiplexvideosettings-constantbitrate
	//
	ConstantBitrate *float64 `field:"optional" json:"constantBitrate" yaml:"constantBitrate"`
	// Statmux rate control settings.
	//
	// When this field is defined, ConstantBitrate must be undefined.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-multiplexprogram-multiplexvideosettings.html#cfn-medialive-multiplexprogram-multiplexvideosettings-statmuxsettings
	//
	StatmuxSettings interface{} `field:"optional" json:"statmuxSettings" yaml:"statmuxSettings"`
}


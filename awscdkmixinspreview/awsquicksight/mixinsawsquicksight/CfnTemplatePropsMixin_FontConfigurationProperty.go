package mixinsawsquicksight


// Configures the display properties of the given text.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fontConfigurationProperty := &FontConfigurationProperty{
//   	FontColor: jsii.String("fontColor"),
//   	FontDecoration: jsii.String("fontDecoration"),
//   	FontFamily: jsii.String("fontFamily"),
//   	FontSize: &FontSizeProperty{
//   		Absolute: jsii.String("absolute"),
//   		Relative: jsii.String("relative"),
//   	},
//   	FontStyle: jsii.String("fontStyle"),
//   	FontWeight: &FontWeightProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fontconfiguration.html
//
type CfnTemplatePropsMixin_FontConfigurationProperty struct {
	// Determines the color of the text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fontconfiguration.html#cfn-quicksight-template-fontconfiguration-fontcolor
	//
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// Determines the appearance of decorative lines on the text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fontconfiguration.html#cfn-quicksight-template-fontconfiguration-fontdecoration
	//
	FontDecoration *string `field:"optional" json:"fontDecoration" yaml:"fontDecoration"`
	// The font family that you want to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fontconfiguration.html#cfn-quicksight-template-fontconfiguration-fontfamily
	//
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// The option that determines the text display size.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fontconfiguration.html#cfn-quicksight-template-fontconfiguration-fontsize
	//
	FontSize interface{} `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Determines the text display face that is inherited by the given font family.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fontconfiguration.html#cfn-quicksight-template-fontconfiguration-fontstyle
	//
	FontStyle *string `field:"optional" json:"fontStyle" yaml:"fontStyle"`
	// The option that determines the text display weight, or boldness.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fontconfiguration.html#cfn-quicksight-template-fontconfiguration-fontweight
	//
	FontWeight interface{} `field:"optional" json:"fontWeight" yaml:"fontWeight"`
}


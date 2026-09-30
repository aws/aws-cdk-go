package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   fontConfigurationProperty := &FontConfigurationProperty{
//   	FontColor: jsii.String("fontColor"),
//   	FontDecoration: jsii.String("fontDecoration"),
//   	FontFamily: jsii.String("fontFamily"),
//   	FontSize: &FontSizeProperty{
//   		Absolute: jsii.String("absolute"),
//   	},
//   	FontStyle: jsii.String("fontStyle"),
//   	FontWeight: &FontWeightProperty{
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-fontconfiguration.html
//
type CfnThemePropsMixin_FontConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-fontconfiguration.html#cfn-quicksight-theme-fontconfiguration-fontcolor
	//
	FontColor *string `field:"optional" json:"fontColor" yaml:"fontColor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-fontconfiguration.html#cfn-quicksight-theme-fontconfiguration-fontdecoration
	//
	FontDecoration *string `field:"optional" json:"fontDecoration" yaml:"fontDecoration"`
	// <p>The font family that you want to use.</p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-fontconfiguration.html#cfn-quicksight-theme-fontconfiguration-fontfamily
	//
	FontFamily *string `field:"optional" json:"fontFamily" yaml:"fontFamily"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-fontconfiguration.html#cfn-quicksight-theme-fontconfiguration-fontsize
	//
	FontSize interface{} `field:"optional" json:"fontSize" yaml:"fontSize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-fontconfiguration.html#cfn-quicksight-theme-fontconfiguration-fontstyle
	//
	FontStyle *string `field:"optional" json:"fontStyle" yaml:"fontStyle"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-fontconfiguration.html#cfn-quicksight-theme-fontconfiguration-fontweight
	//
	FontWeight interface{} `field:"optional" json:"fontWeight" yaml:"fontWeight"`
}


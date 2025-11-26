package previewawsquicksightmixins


// A collection of options that configure how each panel displays in a small multiples chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   panelConfigurationProperty := &PanelConfigurationProperty{
//   	BackgroundColor: jsii.String("backgroundColor"),
//   	BackgroundVisibility: jsii.String("backgroundVisibility"),
//   	BorderColor: jsii.String("borderColor"),
//   	BorderStyle: jsii.String("borderStyle"),
//   	BorderThickness: jsii.String("borderThickness"),
//   	BorderVisibility: jsii.String("borderVisibility"),
//   	GutterSpacing: jsii.String("gutterSpacing"),
//   	GutterVisibility: jsii.String("gutterVisibility"),
//   	Title: &PanelTitleOptionsProperty{
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontFamily: jsii.String("fontFamily"),
//   			FontSize: &FontSizeProperty{
//   				Absolute: jsii.String("absolute"),
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		HorizontalTextAlignment: jsii.String("horizontalTextAlignment"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html
//
type CfnTemplatePropsMixin_PanelConfigurationProperty struct {
	// Sets the background color for each panel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-backgroundcolor
	//
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Determines whether or not a background for each small multiples panel is rendered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-backgroundvisibility
	//
	BackgroundVisibility *string `field:"optional" json:"backgroundVisibility" yaml:"backgroundVisibility"`
	// Sets the line color of panel borders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-bordercolor
	//
	BorderColor *string `field:"optional" json:"borderColor" yaml:"borderColor"`
	// Sets the line style of panel borders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-borderstyle
	//
	BorderStyle *string `field:"optional" json:"borderStyle" yaml:"borderStyle"`
	// Sets the line thickness of panel borders.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-borderthickness
	//
	BorderThickness *string `field:"optional" json:"borderThickness" yaml:"borderThickness"`
	// Determines whether or not each panel displays a border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-bordervisibility
	//
	BorderVisibility *string `field:"optional" json:"borderVisibility" yaml:"borderVisibility"`
	// Sets the total amount of negative space to display between sibling panels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-gutterspacing
	//
	GutterSpacing *string `field:"optional" json:"gutterSpacing" yaml:"gutterSpacing"`
	// Determines whether or not negative space between sibling panels is rendered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-guttervisibility
	//
	GutterVisibility *string `field:"optional" json:"gutterVisibility" yaml:"gutterVisibility"`
	// Configures the title display within each small multiples panel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-panelconfiguration.html#cfn-quicksight-template-panelconfiguration-title
	//
	Title interface{} `field:"optional" json:"title" yaml:"title"`
}


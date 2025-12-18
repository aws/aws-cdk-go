package previewawsconnectmixins


// Contains color configuration for navigation elements in a workspace theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   paletteNavigationProperty := &PaletteNavigationProperty{
//   	Background: jsii.String("background"),
//   	InvertActionsColors: jsii.Boolean(false),
//   	Text: jsii.String("text"),
//   	TextActive: jsii.String("textActive"),
//   	TextBackgroundActive: jsii.String("textBackgroundActive"),
//   	TextBackgroundHover: jsii.String("textBackgroundHover"),
//   	TextHover: jsii.String("textHover"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettenavigation.html
//
type CfnWorkspacePropsMixin_PaletteNavigationProperty struct {
	// The background color of the navigation area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettenavigation.html#cfn-connect-workspace-palettenavigation-background
	//
	Background *string `field:"optional" json:"background" yaml:"background"`
	// Whether to invert the colors of action buttons in the navigation area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettenavigation.html#cfn-connect-workspace-palettenavigation-invertactionscolors
	//
	InvertActionsColors interface{} `field:"optional" json:"invertActionsColors" yaml:"invertActionsColors"`
	// The text color in the navigation area.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettenavigation.html#cfn-connect-workspace-palettenavigation-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The text color for active navigation items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettenavigation.html#cfn-connect-workspace-palettenavigation-textactive
	//
	TextActive *string `field:"optional" json:"textActive" yaml:"textActive"`
	// The background color for active navigation items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettenavigation.html#cfn-connect-workspace-palettenavigation-textbackgroundactive
	//
	TextBackgroundActive *string `field:"optional" json:"textBackgroundActive" yaml:"textBackgroundActive"`
	// The background color when hovering over navigation text.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettenavigation.html#cfn-connect-workspace-palettenavigation-textbackgroundhover
	//
	TextBackgroundHover *string `field:"optional" json:"textBackgroundHover" yaml:"textBackgroundHover"`
	// The text color when hovering over navigation items.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-palettenavigation.html#cfn-connect-workspace-palettenavigation-texthover
	//
	TextHover *string `field:"optional" json:"textHover" yaml:"textHover"`
}


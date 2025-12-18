package previewawsconnectmixins


// Contains color configuration for header elements in a workspace theme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   paletteHeaderProperty := &PaletteHeaderProperty{
//   	Background: jsii.String("background"),
//   	InvertActionsColors: jsii.Boolean(false),
//   	Text: jsii.String("text"),
//   	TextHover: jsii.String("textHover"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-paletteheader.html
//
type CfnWorkspacePropsMixin_PaletteHeaderProperty struct {
	// The background color of the header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-paletteheader.html#cfn-connect-workspace-paletteheader-background
	//
	Background *string `field:"optional" json:"background" yaml:"background"`
	// Whether to invert the colors of action buttons in the header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-paletteheader.html#cfn-connect-workspace-paletteheader-invertactionscolors
	//
	InvertActionsColors interface{} `field:"optional" json:"invertActionsColors" yaml:"invertActionsColors"`
	// The text color in the header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-paletteheader.html#cfn-connect-workspace-paletteheader-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
	// The text color when hovering over header elements.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-paletteheader.html#cfn-connect-workspace-paletteheader-texthover
	//
	TextHover *string `field:"optional" json:"textHover" yaml:"textHover"`
}


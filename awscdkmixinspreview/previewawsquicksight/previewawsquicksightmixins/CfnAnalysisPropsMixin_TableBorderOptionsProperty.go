package previewawsquicksightmixins


// The border options for a table border.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableBorderOptionsProperty := &TableBorderOptionsProperty{
//   	Color: jsii.String("color"),
//   	Style: jsii.String("style"),
//   	Thickness: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tableborderoptions.html
//
type CfnAnalysisPropsMixin_TableBorderOptionsProperty struct {
	// The color of a table border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tableborderoptions.html#cfn-quicksight-analysis-tableborderoptions-color
	//
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The style (none, solid) of a table border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tableborderoptions.html#cfn-quicksight-analysis-tableborderoptions-style
	//
	Style *string `field:"optional" json:"style" yaml:"style"`
	// The thickness of a table border.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-tableborderoptions.html#cfn-quicksight-analysis-tableborderoptions-thickness
	//
	Thickness *float64 `field:"optional" json:"thickness" yaml:"thickness"`
}


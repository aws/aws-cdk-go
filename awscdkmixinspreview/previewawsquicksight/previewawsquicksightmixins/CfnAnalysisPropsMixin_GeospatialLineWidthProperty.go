package previewawsquicksightmixins


// The width properties for a line.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialLineWidthProperty := &GeospatialLineWidthProperty{
//   	LineWidth: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallinewidth.html
//
type CfnAnalysisPropsMixin_GeospatialLineWidthProperty struct {
	// The positive value for the width of a line.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallinewidth.html#cfn-quicksight-analysis-geospatiallinewidth-linewidth
	//
	LineWidth *float64 `field:"optional" json:"lineWidth" yaml:"lineWidth"`
}


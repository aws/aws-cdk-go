package previewawsquicksightmixins


// The geospatial radius for a circle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialCircleRadiusProperty := &GeospatialCircleRadiusProperty{
//   	Radius: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcircleradius.html
//
type CfnAnalysisPropsMixin_GeospatialCircleRadiusProperty struct {
	// The positive value for the radius of a circle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialcircleradius.html#cfn-quicksight-analysis-geospatialcircleradius-radius
	//
	Radius *float64 `field:"optional" json:"radius" yaml:"radius"`
}


package previewawsquicksightmixins


// The bound options (north, south, west, east) of the geospatial window options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialCoordinateBoundsProperty := &GeospatialCoordinateBoundsProperty{
//   	East: jsii.Number(123),
//   	North: jsii.Number(123),
//   	South: jsii.Number(123),
//   	West: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcoordinatebounds.html
//
type CfnTemplatePropsMixin_GeospatialCoordinateBoundsProperty struct {
	// The longitude of the east bound of the geospatial coordinate bounds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcoordinatebounds.html#cfn-quicksight-template-geospatialcoordinatebounds-east
	//
	East *float64 `field:"optional" json:"east" yaml:"east"`
	// The latitude of the north bound of the geospatial coordinate bounds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcoordinatebounds.html#cfn-quicksight-template-geospatialcoordinatebounds-north
	//
	North *float64 `field:"optional" json:"north" yaml:"north"`
	// The latitude of the south bound of the geospatial coordinate bounds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcoordinatebounds.html#cfn-quicksight-template-geospatialcoordinatebounds-south
	//
	South *float64 `field:"optional" json:"south" yaml:"south"`
	// The longitude of the west bound of the geospatial coordinate bounds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-geospatialcoordinatebounds.html#cfn-quicksight-template-geospatialcoordinatebounds-west
	//
	West *float64 `field:"optional" json:"west" yaml:"west"`
}


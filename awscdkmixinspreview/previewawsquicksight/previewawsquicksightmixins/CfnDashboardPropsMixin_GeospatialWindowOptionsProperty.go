package previewawsquicksightmixins


// The window options of the geospatial map visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialWindowOptionsProperty := &GeospatialWindowOptionsProperty{
//   	Bounds: &GeospatialCoordinateBoundsProperty{
//   		East: jsii.Number(123),
//   		North: jsii.Number(123),
//   		South: jsii.Number(123),
//   		West: jsii.Number(123),
//   	},
//   	MapZoomMode: jsii.String("mapZoomMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialwindowoptions.html
//
type CfnDashboardPropsMixin_GeospatialWindowOptionsProperty struct {
	// The bounds options (north, south, west, east) of the geospatial window options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialwindowoptions.html#cfn-quicksight-dashboard-geospatialwindowoptions-bounds
	//
	Bounds interface{} `field:"optional" json:"bounds" yaml:"bounds"`
	// The map zoom modes (manual, auto) of the geospatial window options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialwindowoptions.html#cfn-quicksight-dashboard-geospatialwindowoptions-mapzoommode
	//
	MapZoomMode *string `field:"optional" json:"mapZoomMode" yaml:"mapZoomMode"`
}


package previewawsquicksightmixins


// The map style options of the geospatial map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   geospatialMapStyleOptionsProperty := &GeospatialMapStyleOptionsProperty{
//   	BaseMapStyle: jsii.String("baseMapStyle"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapstyleoptions.html
//
type CfnDashboardPropsMixin_GeospatialMapStyleOptionsProperty struct {
	// The base map style of the geospatial map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-geospatialmapstyleoptions.html#cfn-quicksight-dashboard-geospatialmapstyleoptions-basemapstyle
	//
	BaseMapStyle *string `field:"optional" json:"baseMapStyle" yaml:"baseMapStyle"`
}


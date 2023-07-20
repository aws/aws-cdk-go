package awsquicksight


// The map style options of the geospatial map.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialMapStyleOptionsProperty := &GeospatialMapStyleOptionsProperty{
//   	BaseMapStyle: jsii.String("baseMapStyle"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapstyleoptions.html
//
type CfnAnalysis_GeospatialMapStyleOptionsProperty struct {
	// The base map style of the geospatial map.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatialmapstyleoptions.html#cfn-quicksight-analysis-geospatialmapstyleoptions-basemapstyle
	//
	BaseMapStyle *string `field:"optional" json:"baseMapStyle" yaml:"baseMapStyle"`
}


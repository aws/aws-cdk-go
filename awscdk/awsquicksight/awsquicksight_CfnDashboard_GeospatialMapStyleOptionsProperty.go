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
type CfnDashboard_GeospatialMapStyleOptionsProperty struct {
	// The base map style of the geospatial map.
	BaseMapStyle *string `field:"optional" json:"baseMapStyle" yaml:"baseMapStyle"`
}


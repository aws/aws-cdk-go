package awsquicksight


// The bound options (north, south, west, east) of the geospatial window options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   geospatialCoordinateBoundsProperty := &GeospatialCoordinateBoundsProperty{
//   	East: jsii.Number(123),
//   	North: jsii.Number(123),
//   	South: jsii.Number(123),
//   	West: jsii.Number(123),
//   }
//
type CfnTemplate_GeospatialCoordinateBoundsProperty struct {
	// The longitude of the east bound of the geospatial coordinate bounds.
	East *float64 `field:"required" json:"east" yaml:"east"`
	// The latitude of the north bound of the geospatial coordinate bounds.
	North *float64 `field:"required" json:"north" yaml:"north"`
	// The latitude of the south bound of the geospatial coordinate bounds.
	South *float64 `field:"required" json:"south" yaml:"south"`
	// The longitude of the west bound of the geospatial coordinate bounds.
	West *float64 `field:"required" json:"west" yaml:"west"`
}


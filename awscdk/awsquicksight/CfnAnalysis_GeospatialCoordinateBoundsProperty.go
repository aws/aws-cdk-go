package awsquicksight


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
type CfnAnalysis_GeospatialCoordinateBoundsProperty struct {
	// `CfnAnalysis.GeospatialCoordinateBoundsProperty.East`.
	East *float64 `field:"required" json:"east" yaml:"east"`
	// `CfnAnalysis.GeospatialCoordinateBoundsProperty.North`.
	North *float64 `field:"required" json:"north" yaml:"north"`
	// `CfnAnalysis.GeospatialCoordinateBoundsProperty.South`.
	South *float64 `field:"required" json:"south" yaml:"south"`
	// `CfnAnalysis.GeospatialCoordinateBoundsProperty.West`.
	West *float64 `field:"required" json:"west" yaml:"west"`
}


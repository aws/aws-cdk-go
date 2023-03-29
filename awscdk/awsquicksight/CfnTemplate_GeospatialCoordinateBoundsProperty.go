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
type CfnTemplate_GeospatialCoordinateBoundsProperty struct {
	// `CfnTemplate.GeospatialCoordinateBoundsProperty.East`.
	East *float64 `field:"required" json:"east" yaml:"east"`
	// `CfnTemplate.GeospatialCoordinateBoundsProperty.North`.
	North *float64 `field:"required" json:"north" yaml:"north"`
	// `CfnTemplate.GeospatialCoordinateBoundsProperty.South`.
	South *float64 `field:"required" json:"south" yaml:"south"`
	// `CfnTemplate.GeospatialCoordinateBoundsProperty.West`.
	West *float64 `field:"required" json:"west" yaml:"west"`
}


package awspinpoint


// Specifies the GPS coordinates of a location.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coordinatesProperty := &coordinatesProperty{
//   	latitude: jsii.Number(123),
//   	longitude: jsii.Number(123),
//   }
//
type CfnSegment_CoordinatesProperty struct {
	// The latitude coordinate of the location.
	Latitude *float64 `field:"required" json:"latitude" yaml:"latitude"`
	// The longitude coordinate of the location.
	Longitude *float64 `field:"required" json:"longitude" yaml:"longitude"`
}


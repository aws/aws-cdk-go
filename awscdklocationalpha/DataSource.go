// The CDK Construct Library for AWS::Location
package awscdklocationalpha


// Data source for a place index.
//
// Example:
//   location.NewPlaceIndex(this, jsii.String("PlaceIndex"), &PlaceIndexProps{
//   	PlaceIndexName: jsii.String("MyPlaceIndex"),
//   	 // optional, defaults to a generated name
//   	DataSource: location.DataSource_HERE,
//   })
//
// Experimental.
type DataSource string

const (
	// Esri.
	// See: https://docs.aws.amazon.com/location/latest/developerguide/esri.html
	//
	// Experimental.
	DataSource_ESRI DataSource = "ESRI"
	// HERE.
	// See: https://docs.aws.amazon.com/location/latest/developerguide/HERE.html
	//
	// Experimental.
	DataSource_HERE DataSource = "HERE"
)


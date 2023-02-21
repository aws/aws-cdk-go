// The CDK Construct Library for AWS::Location
package awscdklocationalpha


// Properties for a place index.
//
// Example:
//   location.NewPlaceIndex(this, jsii.String("PlaceIndex"), &PlaceIndexProps{
//   	PlaceIndexName: jsii.String("MyPlaceIndex"),
//   	 // optional, defaults to a generated name
//   	DataSource: location.DataSource_HERE,
//   })
//
// Experimental.
type PlaceIndexProps struct {
	// Data source for the place index.
	// Experimental.
	DataSource DataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	// A description for the place index.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Intend use for the results of an operation.
	// Experimental.
	IntendedUse IntendedUse `field:"optional" json:"intendedUse" yaml:"intendedUse"`
	// A name for the place index.
	// Experimental.
	PlaceIndexName *string `field:"optional" json:"placeIndexName" yaml:"placeIndexName"`
}


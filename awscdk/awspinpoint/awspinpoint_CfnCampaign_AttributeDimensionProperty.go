package awspinpoint


// Specifies attribute-based criteria for including or excluding endpoints from a segment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeDimensionProperty := &attributeDimensionProperty{
//   	attributeType: jsii.String("attributeType"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnCampaign_AttributeDimensionProperty struct {
	// The type of segment dimension to use. Valid values are:.
	//
	// - `INCLUSIVE` – endpoints that have attributes matching the values are included in the segment.
	// - `EXCLUSIVE` – endpoints that have attributes matching the values are excluded from the segment.
	// - `CONTAINS` – endpoints that have attributes' substrings match the values are included in the segment.
	// - `BEFORE` – endpoints with attributes read as ISO_INSTANT datetimes before the value are included in the segment.
	// - `AFTER` – endpoints with attributes read as ISO_INSTANT datetimes after the value are included in the segment.
	// - `BETWEEN` – endpoints with attributes read as ISO_INSTANT datetimes between the values are included in the segment.
	// - `ON` – endpoints with attributes read as ISO_INSTANT dates on the value are included in the segment. Time is ignored in this comparison.
	AttributeType *string `field:"optional" json:"attributeType" yaml:"attributeType"`
	// The criteria values to use for the segment dimension.
	//
	// Depending on the value of the `AttributeType` property, endpoints are included or excluded from the segment if their attribute values match the criteria values.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}


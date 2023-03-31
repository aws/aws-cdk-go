package awspinpoint


// Specifies the dimension type and values for a segment dimension.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   setDimensionProperty := &setDimensionProperty{
//   	dimensionType: jsii.String("dimensionType"),
//   	values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
type CfnCampaign_SetDimensionProperty struct {
	// The type of segment dimension to use.
	//
	// Valid values are: `INCLUSIVE` , endpoints that match the criteria are included in the segment; and, `EXCLUSIVE` , endpoints that match the criteria are excluded from the segment.
	DimensionType *string `field:"optional" json:"dimensionType" yaml:"dimensionType"`
	// The criteria values to use for the segment dimension.
	//
	// Depending on the value of the `DimensionType` property, endpoints are included or excluded from the segment if their values match the criteria values.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}


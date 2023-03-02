package awstimestream


// This type is used to map column(s) from the query result to a dimension in the destination table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionMappingProperty := &dimensionMappingProperty{
//   	dimensionValueType: jsii.String("dimensionValueType"),
//   	name: jsii.String("name"),
//   }
//
type CfnScheduledQuery_DimensionMappingProperty struct {
	// Type for the dimension.
	DimensionValueType *string `field:"required" json:"dimensionValueType" yaml:"dimensionValueType"`
	// Column name from query result.
	Name *string `field:"required" json:"name" yaml:"name"`
}


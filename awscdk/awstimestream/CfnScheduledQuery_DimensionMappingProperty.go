package awstimestream


// This type is used to map column(s) from the query result to a dimension in the destination table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionMappingProperty := &DimensionMappingProperty{
//   	DimensionValueType: jsii.String("dimensionValueType"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-dimensionmapping.html
//
type CfnScheduledQuery_DimensionMappingProperty struct {
	// Type for the dimension: VARCHAR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-dimensionmapping.html#cfn-timestream-scheduledquery-dimensionmapping-dimensionvaluetype
	//
	DimensionValueType *string `field:"required" json:"dimensionValueType" yaml:"dimensionValueType"`
	// Column name from query result.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-dimensionmapping.html#cfn-timestream-scheduledquery-dimensionmapping-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}


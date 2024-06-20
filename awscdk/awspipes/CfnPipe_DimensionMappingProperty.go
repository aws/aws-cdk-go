package awspipes


// Maps source data to a dimension in the target Timestream for LiveAnalytics table.
//
// For more information, see [Amazon Timestream for LiveAnalytics concepts](https://docs.aws.amazon.com/timestream/latest/developerguide/concepts.html)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dimensionMappingProperty := &DimensionMappingProperty{
//   	DimensionName: jsii.String("dimensionName"),
//   	DimensionValue: jsii.String("dimensionValue"),
//   	DimensionValueType: jsii.String("dimensionValueType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-dimensionmapping.html
//
type CfnPipe_DimensionMappingProperty struct {
	// The metadata attributes of the time series.
	//
	// For example, the name and Availability Zone of an Amazon EC2 instance or the name of the manufacturer of a wind turbine are dimensions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-dimensionmapping.html#cfn-pipes-pipe-dimensionmapping-dimensionname
	//
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// Dynamic path to the dimension value in the source event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-dimensionmapping.html#cfn-pipes-pipe-dimensionmapping-dimensionvalue
	//
	DimensionValue *string `field:"required" json:"dimensionValue" yaml:"dimensionValue"`
	// The data type of the dimension for the time-series data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-dimensionmapping.html#cfn-pipes-pipe-dimensionmapping-dimensionvaluetype
	//
	DimensionValueType *string `field:"required" json:"dimensionValueType" yaml:"dimensionValueType"`
}


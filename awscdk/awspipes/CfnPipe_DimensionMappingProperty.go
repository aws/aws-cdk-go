package awspipes


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-dimensionmapping.html#cfn-pipes-pipe-dimensionmapping-dimensionname
	//
	DimensionName *string `field:"required" json:"dimensionName" yaml:"dimensionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-dimensionmapping.html#cfn-pipes-pipe-dimensionmapping-dimensionvalue
	//
	DimensionValue *string `field:"required" json:"dimensionValue" yaml:"dimensionValue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-dimensionmapping.html#cfn-pipes-pipe-dimensionmapping-dimensionvaluetype
	//
	DimensionValueType *string `field:"required" json:"dimensionValueType" yaml:"dimensionValueType"`
}


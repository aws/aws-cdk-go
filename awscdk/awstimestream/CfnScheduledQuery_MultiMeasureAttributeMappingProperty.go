package awstimestream


// Attribute mapping for MULTI value measures.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiMeasureAttributeMappingProperty := &MultiMeasureAttributeMappingProperty{
//   	MeasureValueType: jsii.String("measureValueType"),
//   	SourceColumn: jsii.String("sourceColumn"),
//
//   	// the properties below are optional
//   	TargetMultiMeasureAttributeName: jsii.String("targetMultiMeasureAttributeName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-multimeasureattributemapping.html
//
type CfnScheduledQuery_MultiMeasureAttributeMappingProperty struct {
	// Type of the attribute to be read from the source column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-multimeasureattributemapping.html#cfn-timestream-scheduledquery-multimeasureattributemapping-measurevaluetype
	//
	MeasureValueType *string `field:"required" json:"measureValueType" yaml:"measureValueType"`
	// Source column from where the attribute value is to be read.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-multimeasureattributemapping.html#cfn-timestream-scheduledquery-multimeasureattributemapping-sourcecolumn
	//
	SourceColumn *string `field:"required" json:"sourceColumn" yaml:"sourceColumn"`
	// Custom name to be used for attribute name in derived table.
	//
	// If not provided, source column name would be used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-timestream-scheduledquery-multimeasureattributemapping.html#cfn-timestream-scheduledquery-multimeasureattributemapping-targetmultimeasureattributename
	//
	TargetMultiMeasureAttributeName *string `field:"optional" json:"targetMultiMeasureAttributeName" yaml:"targetMultiMeasureAttributeName"`
}


package previewawsquicksightmixins


// Aggregation for attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   attributeAggregationFunctionProperty := &AttributeAggregationFunctionProperty{
//   	SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   	ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-attributeaggregationfunction.html
//
type CfnAnalysisPropsMixin_AttributeAggregationFunctionProperty struct {
	// The built-in aggregation functions for attributes.
	//
	// - `UNIQUE_VALUE` : Returns the unique value for a field, aggregated by the dimension fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-attributeaggregationfunction.html#cfn-quicksight-analysis-attributeaggregationfunction-simpleattributeaggregation
	//
	SimpleAttributeAggregation *string `field:"optional" json:"simpleAttributeAggregation" yaml:"simpleAttributeAggregation"`
	// Used by the `UNIQUE_VALUE` aggregation function.
	//
	// If there are multiple values for the field used by the aggregation, the value for this property will be returned instead. Defaults to '*'.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-attributeaggregationfunction.html#cfn-quicksight-analysis-attributeaggregationfunction-valueformultiplevalues
	//
	ValueForMultipleValues *string `field:"optional" json:"valueForMultipleValues" yaml:"valueForMultipleValues"`
}


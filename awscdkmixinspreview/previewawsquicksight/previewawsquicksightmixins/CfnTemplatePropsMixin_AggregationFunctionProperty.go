package previewawsquicksightmixins


// An aggregation function aggregates values from a dimension or measure.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aggregationFunctionProperty := &AggregationFunctionProperty{
//   	AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   		SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   		ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   	},
//   	CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   	DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   	NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   		PercentileAggregation: &PercentileAggregationProperty{
//   			PercentileValue: jsii.Number(123),
//   		},
//   		SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-aggregationfunction.html
//
type CfnTemplatePropsMixin_AggregationFunctionProperty struct {
	// Aggregation for attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-aggregationfunction.html#cfn-quicksight-template-aggregationfunction-attributeaggregationfunction
	//
	AttributeAggregationFunction interface{} `field:"optional" json:"attributeAggregationFunction" yaml:"attributeAggregationFunction"`
	// Aggregation for categorical values.
	//
	// - `COUNT` : Aggregate by the total number of values, including duplicates.
	// - `DISTINCT_COUNT` : Aggregate by the total number of distinct values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-aggregationfunction.html#cfn-quicksight-template-aggregationfunction-categoricalaggregationfunction
	//
	CategoricalAggregationFunction *string `field:"optional" json:"categoricalAggregationFunction" yaml:"categoricalAggregationFunction"`
	// Aggregation for date values.
	//
	// - `COUNT` : Aggregate by the total number of values, including duplicates.
	// - `DISTINCT_COUNT` : Aggregate by the total number of distinct values.
	// - `MIN` : Select the smallest date value.
	// - `MAX` : Select the largest date value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-aggregationfunction.html#cfn-quicksight-template-aggregationfunction-dateaggregationfunction
	//
	DateAggregationFunction *string `field:"optional" json:"dateAggregationFunction" yaml:"dateAggregationFunction"`
	// Aggregation for numerical values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-aggregationfunction.html#cfn-quicksight-template-aggregationfunction-numericalaggregationfunction
	//
	NumericalAggregationFunction interface{} `field:"optional" json:"numericalAggregationFunction" yaml:"numericalAggregationFunction"`
}


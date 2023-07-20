package awsquicksight


// A `NumericRangeFilter` filters values that are within the value range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericRangeFilterProperty := &NumericRangeFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FilterId: jsii.String("filterId"),
//   	NullOption: jsii.String("nullOption"),
//
//   	// the properties below are optional
//   	AggregationFunction: &AggregationFunctionProperty{
//   		CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   		DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   		NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   			PercentileAggregation: &PercentileAggregationProperty{
//   				PercentileValue: jsii.Number(123),
//   			},
//   			SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   		},
//   	},
//   	IncludeMaximum: jsii.Boolean(false),
//   	IncludeMinimum: jsii.Boolean(false),
//   	RangeMaximum: &NumericRangeFilterValueProperty{
//   		Parameter: jsii.String("parameter"),
//   		StaticValue: jsii.Number(123),
//   	},
//   	RangeMinimum: &NumericRangeFilterValueProperty{
//   		Parameter: jsii.String("parameter"),
//   		StaticValue: jsii.Number(123),
//   	},
//   	SelectAllOptions: jsii.String("selectAllOptions"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericrangefilter.html
//
type CfnTemplate_NumericRangeFilterProperty struct {
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericrangefilter.html#cfn-quicksight-template-numericrangefilter-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericrangefilter.html#cfn-quicksight-template-numericrangefilter-filterid
	//
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// This option determines how null values should be treated when filtering data.
	//
	// - `ALL_VALUES` : Include null values in filtered results.
	// - `NULLS_ONLY` : Only include null values in filtered results.
	// - `NON_NULLS_ONLY` : Exclude null values from filtered results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericrangefilter.html#cfn-quicksight-template-numericrangefilter-nulloption
	//
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// The aggregation function of the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericrangefilter.html#cfn-quicksight-template-numericrangefilter-aggregationfunction
	//
	AggregationFunction interface{} `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
	// Determines whether the maximum value in the filter value range should be included in the filtered results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericrangefilter.html#cfn-quicksight-template-numericrangefilter-includemaximum
	//
	IncludeMaximum interface{} `field:"optional" json:"includeMaximum" yaml:"includeMaximum"`
	// Determines whether the minimum value in the filter value range should be included in the filtered results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericrangefilter.html#cfn-quicksight-template-numericrangefilter-includeminimum
	//
	IncludeMinimum interface{} `field:"optional" json:"includeMinimum" yaml:"includeMinimum"`
	// The maximum value for the filter value range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericrangefilter.html#cfn-quicksight-template-numericrangefilter-rangemaximum
	//
	RangeMaximum interface{} `field:"optional" json:"rangeMaximum" yaml:"rangeMaximum"`
	// The minimum value for the filter value range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericrangefilter.html#cfn-quicksight-template-numericrangefilter-rangeminimum
	//
	RangeMinimum interface{} `field:"optional" json:"rangeMinimum" yaml:"rangeMinimum"`
	// Select all of the values. Null is not the assigned value of select all.
	//
	// - `FILTER_ALL_VALUES`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-numericrangefilter.html#cfn-quicksight-template-numericrangefilter-selectalloptions
	//
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}


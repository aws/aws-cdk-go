package awsquicksight


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
type CfnTemplate_NumericRangeFilterProperty struct {
	// `CfnTemplate.NumericRangeFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnTemplate.NumericRangeFilterProperty.FilterId`.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// `CfnTemplate.NumericRangeFilterProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnTemplate.NumericRangeFilterProperty.AggregationFunction`.
	AggregationFunction interface{} `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
	// `CfnTemplate.NumericRangeFilterProperty.IncludeMaximum`.
	IncludeMaximum interface{} `field:"optional" json:"includeMaximum" yaml:"includeMaximum"`
	// `CfnTemplate.NumericRangeFilterProperty.IncludeMinimum`.
	IncludeMinimum interface{} `field:"optional" json:"includeMinimum" yaml:"includeMinimum"`
	// `CfnTemplate.NumericRangeFilterProperty.RangeMaximum`.
	RangeMaximum interface{} `field:"optional" json:"rangeMaximum" yaml:"rangeMaximum"`
	// `CfnTemplate.NumericRangeFilterProperty.RangeMinimum`.
	RangeMinimum interface{} `field:"optional" json:"rangeMinimum" yaml:"rangeMinimum"`
	// `CfnTemplate.NumericRangeFilterProperty.SelectAllOptions`.
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}


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
type CfnDashboard_NumericRangeFilterProperty struct {
	// `CfnDashboard.NumericRangeFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnDashboard.NumericRangeFilterProperty.FilterId`.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// `CfnDashboard.NumericRangeFilterProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnDashboard.NumericRangeFilterProperty.AggregationFunction`.
	AggregationFunction interface{} `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
	// `CfnDashboard.NumericRangeFilterProperty.IncludeMaximum`.
	IncludeMaximum interface{} `field:"optional" json:"includeMaximum" yaml:"includeMaximum"`
	// `CfnDashboard.NumericRangeFilterProperty.IncludeMinimum`.
	IncludeMinimum interface{} `field:"optional" json:"includeMinimum" yaml:"includeMinimum"`
	// `CfnDashboard.NumericRangeFilterProperty.RangeMaximum`.
	RangeMaximum interface{} `field:"optional" json:"rangeMaximum" yaml:"rangeMaximum"`
	// `CfnDashboard.NumericRangeFilterProperty.RangeMinimum`.
	RangeMinimum interface{} `field:"optional" json:"rangeMinimum" yaml:"rangeMinimum"`
	// `CfnDashboard.NumericRangeFilterProperty.SelectAllOptions`.
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}


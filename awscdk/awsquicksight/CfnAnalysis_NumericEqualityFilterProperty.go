package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericEqualityFilterProperty := &NumericEqualityFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FilterId: jsii.String("filterId"),
//   	MatchOperator: jsii.String("matchOperator"),
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
//   	ParameterName: jsii.String("parameterName"),
//   	SelectAllOptions: jsii.String("selectAllOptions"),
//   	Value: jsii.Number(123),
//   }
//
type CfnAnalysis_NumericEqualityFilterProperty struct {
	// `CfnAnalysis.NumericEqualityFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnAnalysis.NumericEqualityFilterProperty.FilterId`.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// `CfnAnalysis.NumericEqualityFilterProperty.MatchOperator`.
	MatchOperator *string `field:"required" json:"matchOperator" yaml:"matchOperator"`
	// `CfnAnalysis.NumericEqualityFilterProperty.NullOption`.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// `CfnAnalysis.NumericEqualityFilterProperty.AggregationFunction`.
	AggregationFunction interface{} `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
	// `CfnAnalysis.NumericEqualityFilterProperty.ParameterName`.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// `CfnAnalysis.NumericEqualityFilterProperty.SelectAllOptions`.
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
	// `CfnAnalysis.NumericEqualityFilterProperty.Value`.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}


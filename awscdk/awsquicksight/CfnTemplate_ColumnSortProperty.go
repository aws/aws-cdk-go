package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnSortProperty := &ColumnSortProperty{
//   	Direction: jsii.String("direction"),
//   	SortBy: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
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
//   }
//
type CfnTemplate_ColumnSortProperty struct {
	// `CfnTemplate.ColumnSortProperty.Direction`.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// `CfnTemplate.ColumnSortProperty.SortBy`.
	SortBy interface{} `field:"required" json:"sortBy" yaml:"sortBy"`
	// `CfnTemplate.ColumnSortProperty.AggregationFunction`.
	AggregationFunction interface{} `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
}


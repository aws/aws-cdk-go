package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregationSortConfigurationProperty := &AggregationSortConfigurationProperty{
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
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	SortDirection: jsii.String("sortDirection"),
//   }
//
type CfnTemplate_AggregationSortConfigurationProperty struct {
	// `CfnTemplate.AggregationSortConfigurationProperty.AggregationFunction`.
	AggregationFunction interface{} `field:"required" json:"aggregationFunction" yaml:"aggregationFunction"`
	// `CfnTemplate.AggregationSortConfigurationProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnTemplate.AggregationSortConfigurationProperty.SortDirection`.
	SortDirection *string `field:"required" json:"sortDirection" yaml:"sortDirection"`
}


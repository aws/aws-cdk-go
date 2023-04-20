package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceLineDynamicDataConfigurationProperty := &ReferenceLineDynamicDataConfigurationProperty{
//   	Calculation: &NumericalAggregationFunctionProperty{
//   		PercentileAggregation: &PercentileAggregationProperty{
//   			PercentileValue: jsii.Number(123),
//   		},
//   		SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   	},
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	MeasureAggregationFunction: &AggregationFunctionProperty{
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
type CfnAnalysis_ReferenceLineDynamicDataConfigurationProperty struct {
	// `CfnAnalysis.ReferenceLineDynamicDataConfigurationProperty.Calculation`.
	Calculation interface{} `field:"required" json:"calculation" yaml:"calculation"`
	// `CfnAnalysis.ReferenceLineDynamicDataConfigurationProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnAnalysis.ReferenceLineDynamicDataConfigurationProperty.MeasureAggregationFunction`.
	MeasureAggregationFunction interface{} `field:"required" json:"measureAggregationFunction" yaml:"measureAggregationFunction"`
}


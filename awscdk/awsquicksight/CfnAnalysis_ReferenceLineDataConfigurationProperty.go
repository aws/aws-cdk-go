package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceLineDataConfigurationProperty := &ReferenceLineDataConfigurationProperty{
//   	AxisBinding: jsii.String("axisBinding"),
//   	DynamicConfiguration: &ReferenceLineDynamicDataConfigurationProperty{
//   		Calculation: &NumericalAggregationFunctionProperty{
//   			PercentileAggregation: &PercentileAggregationProperty{
//   				PercentileValue: jsii.Number(123),
//   			},
//   			SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   		},
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		MeasureAggregationFunction: &AggregationFunctionProperty{
//   			CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   			DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   			NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   				PercentileAggregation: &PercentileAggregationProperty{
//   					PercentileValue: jsii.Number(123),
//   				},
//   				SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   			},
//   		},
//   	},
//   	StaticConfiguration: &ReferenceLineStaticDataConfigurationProperty{
//   		Value: jsii.Number(123),
//   	},
//   }
//
type CfnAnalysis_ReferenceLineDataConfigurationProperty struct {
	// `CfnAnalysis.ReferenceLineDataConfigurationProperty.AxisBinding`.
	AxisBinding *string `field:"optional" json:"axisBinding" yaml:"axisBinding"`
	// `CfnAnalysis.ReferenceLineDataConfigurationProperty.DynamicConfiguration`.
	DynamicConfiguration interface{} `field:"optional" json:"dynamicConfiguration" yaml:"dynamicConfiguration"`
	// `CfnAnalysis.ReferenceLineDataConfigurationProperty.StaticConfiguration`.
	StaticConfiguration interface{} `field:"optional" json:"staticConfiguration" yaml:"staticConfiguration"`
}


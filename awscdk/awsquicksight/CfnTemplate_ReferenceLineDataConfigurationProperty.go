package awsquicksight


// The data configuration of the reference line.
//
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
//
//   		// the properties below are optional
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
type CfnTemplate_ReferenceLineDataConfigurationProperty struct {
	// The axis binding type of the reference line. Choose one of the following options:.
	//
	// - PrimaryY
	// - SecondaryY.
	AxisBinding *string `field:"optional" json:"axisBinding" yaml:"axisBinding"`
	// The dynamic configuration of the reference line data configuration.
	DynamicConfiguration interface{} `field:"optional" json:"dynamicConfiguration" yaml:"dynamicConfiguration"`
	// The static data configuration of the reference line data configuration.
	StaticConfiguration interface{} `field:"optional" json:"staticConfiguration" yaml:"staticConfiguration"`
}


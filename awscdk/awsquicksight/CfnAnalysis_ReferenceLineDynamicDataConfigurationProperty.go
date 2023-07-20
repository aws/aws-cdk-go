package awsquicksight


// The dynamic configuration of the reference line data configuration.
//
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
//
//   	// the properties below are optional
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-referencelinedynamicdataconfiguration.html
//
type CfnAnalysis_ReferenceLineDynamicDataConfigurationProperty struct {
	// The calculation that is used in the dynamic data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-referencelinedynamicdataconfiguration.html#cfn-quicksight-analysis-referencelinedynamicdataconfiguration-calculation
	//
	Calculation interface{} `field:"required" json:"calculation" yaml:"calculation"`
	// The column that the dynamic data targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-referencelinedynamicdataconfiguration.html#cfn-quicksight-analysis-referencelinedynamicdataconfiguration-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The aggregation function that is used in the dynamic data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-referencelinedynamicdataconfiguration.html#cfn-quicksight-analysis-referencelinedynamicdataconfiguration-measureaggregationfunction
	//
	MeasureAggregationFunction interface{} `field:"optional" json:"measureAggregationFunction" yaml:"measureAggregationFunction"`
}


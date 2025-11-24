package mixinsawsquicksight


// Describes the *Numeric* dataset column and constraints for the dynamic values used to repeat the contents of a section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bodySectionDynamicNumericDimensionConfigurationProperty := &BodySectionDynamicNumericDimensionConfigurationProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	Limit: jsii.Number(123),
//   	SortByMetrics: []interface{}{
//   		&ColumnSortProperty{
//   			AggregationFunction: &AggregationFunctionProperty{
//   				AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   					SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   					ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   				},
//   				CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   				DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   				NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   					PercentileAggregation: &PercentileAggregationProperty{
//   						PercentileValue: jsii.Number(123),
//   					},
//   					SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   				},
//   			},
//   			Direction: jsii.String("direction"),
//   			SortBy: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-bodysectiondynamicnumericdimensionconfiguration.html
//
type CfnDashboardPropsMixin_BodySectionDynamicNumericDimensionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-bodysectiondynamicnumericdimensionconfiguration.html#cfn-quicksight-dashboard-bodysectiondynamicnumericdimensionconfiguration-column
	//
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// Number of values to use from the column for repetition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-bodysectiondynamicnumericdimensionconfiguration.html#cfn-quicksight-dashboard-bodysectiondynamicnumericdimensionconfiguration-limit
	//
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Sort criteria on the column values that you use for repetition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-bodysectiondynamicnumericdimensionconfiguration.html#cfn-quicksight-dashboard-bodysectiondynamicnumericdimensionconfiguration-sortbymetrics
	//
	SortByMetrics interface{} `field:"optional" json:"sortByMetrics" yaml:"sortByMetrics"`
}


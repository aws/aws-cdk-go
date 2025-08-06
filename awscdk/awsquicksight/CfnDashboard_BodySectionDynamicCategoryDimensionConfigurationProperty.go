package awsquicksight


// Describes the *Category* dataset column and constraints for the dynamic values used to repeat the contents of a section.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bodySectionDynamicCategoryDimensionConfigurationProperty := &BodySectionDynamicCategoryDimensionConfigurationProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	Limit: jsii.Number(123),
//   	SortByMetrics: []interface{}{
//   		&ColumnSortProperty{
//   			Direction: jsii.String("direction"),
//   			SortBy: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//
//   			// the properties below are optional
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
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-bodysectiondynamiccategorydimensionconfiguration.html
//
type CfnDashboard_BodySectionDynamicCategoryDimensionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-bodysectiondynamiccategorydimensionconfiguration.html#cfn-quicksight-dashboard-bodysectiondynamiccategorydimensionconfiguration-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// Number of values to use from the column for repetition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-bodysectiondynamiccategorydimensionconfiguration.html#cfn-quicksight-dashboard-bodysectiondynamiccategorydimensionconfiguration-limit
	//
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Sort criteria on the column values that you use for repetition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-bodysectiondynamiccategorydimensionconfiguration.html#cfn-quicksight-dashboard-bodysectiondynamiccategorydimensionconfiguration-sortbymetrics
	//
	SortByMetrics interface{} `field:"optional" json:"sortByMetrics" yaml:"sortByMetrics"`
}


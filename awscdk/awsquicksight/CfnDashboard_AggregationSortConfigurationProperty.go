package awsquicksight


// The configuration options to sort aggregated values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aggregationSortConfigurationProperty := &AggregationSortConfigurationProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	SortDirection: jsii.String("sortDirection"),
//
//   	// the properties below are optional
//   	AggregationFunction: &AggregationFunctionProperty{
//   		AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   			SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   			ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   		},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-aggregationsortconfiguration.html
//
type CfnDashboard_AggregationSortConfigurationProperty struct {
	// The column that determines the sort order of aggregated values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-aggregationsortconfiguration.html#cfn-quicksight-dashboard-aggregationsortconfiguration-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The sort direction of values.
	//
	// - `ASC` : Sort in ascending order.
	// - `DESC` : Sort in descending order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-aggregationsortconfiguration.html#cfn-quicksight-dashboard-aggregationsortconfiguration-sortdirection
	//
	SortDirection *string `field:"required" json:"sortDirection" yaml:"sortDirection"`
	// The function that aggregates the values in `Column` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-aggregationsortconfiguration.html#cfn-quicksight-dashboard-aggregationsortconfiguration-aggregationfunction
	//
	AggregationFunction interface{} `field:"optional" json:"aggregationFunction" yaml:"aggregationFunction"`
}


package awsquicksight


// The table query sorting options for the plugin visual.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pluginVisualTableQuerySortProperty := &PluginVisualTableQuerySortProperty{
//   	ItemsLimitConfiguration: &PluginVisualItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   	},
//   	RowSort: []interface{}{
//   		&FieldSortOptionsProperty{
//   			ColumnSort: &ColumnSortProperty{
//   				Direction: jsii.String("direction"),
//   				SortBy: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//
//   				// the properties below are optional
//   				AggregationFunction: &AggregationFunctionProperty{
//   					AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   						SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   						ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   					},
//   					CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   					DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   					NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   						PercentileAggregation: &PercentileAggregationProperty{
//   							PercentileValue: jsii.Number(123),
//   						},
//   						SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   					},
//   				},
//   			},
//   			FieldSort: &FieldSortProperty{
//   				Direction: jsii.String("direction"),
//   				FieldId: jsii.String("fieldId"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualtablequerysort.html
//
type CfnDashboard_PluginVisualTableQuerySortProperty struct {
	// The maximum amount of data to be returned by a query.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualtablequerysort.html#cfn-quicksight-dashboard-pluginvisualtablequerysort-itemslimitconfiguration
	//
	ItemsLimitConfiguration interface{} `field:"optional" json:"itemsLimitConfiguration" yaml:"itemsLimitConfiguration"`
	// Determines how data is sorted in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pluginvisualtablequerysort.html#cfn-quicksight-dashboard-pluginvisualtablequerysort-rowsort
	//
	RowSort interface{} `field:"optional" json:"rowSort" yaml:"rowSort"`
}


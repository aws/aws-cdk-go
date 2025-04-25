package awsquicksight


// The sort configuration for a `TableVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableSortConfigurationProperty := &TableSortConfigurationProperty{
//   	PaginationConfiguration: &PaginationConfigurationProperty{
//   		PageNumber: jsii.Number(123),
//   		PageSize: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablesortconfiguration.html
//
type CfnDashboard_TableSortConfigurationProperty struct {
	// The pagination configuration (page size, page number) for the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablesortconfiguration.html#cfn-quicksight-dashboard-tablesortconfiguration-paginationconfiguration
	//
	PaginationConfiguration interface{} `field:"optional" json:"paginationConfiguration" yaml:"paginationConfiguration"`
	// The field sort options for rows in the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-tablesortconfiguration.html#cfn-quicksight-dashboard-tablesortconfiguration-rowsort
	//
	RowSort interface{} `field:"optional" json:"rowSort" yaml:"rowSort"`
}


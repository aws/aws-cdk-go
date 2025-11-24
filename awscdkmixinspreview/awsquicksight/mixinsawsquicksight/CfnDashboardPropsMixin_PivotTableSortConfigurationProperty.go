package mixinsawsquicksight


// The sort configuration for a `PivotTableVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pivotTableSortConfigurationProperty := &PivotTableSortConfigurationProperty{
//   	FieldSortOptions: []interface{}{
//   		&PivotFieldSortOptionsProperty{
//   			FieldId: jsii.String("fieldId"),
//   			SortBy: &PivotTableSortByProperty{
//   				Column: &ColumnSortProperty{
//   					AggregationFunction: &AggregationFunctionProperty{
//   						AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   							SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   							ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   						},
//   						CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   						DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   						NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   							PercentileAggregation: &PercentileAggregationProperty{
//   								PercentileValue: jsii.Number(123),
//   							},
//   							SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   						},
//   					},
//   					Direction: jsii.String("direction"),
//   					SortBy: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   				},
//   				DataPath: &DataPathSortProperty{
//   					Direction: jsii.String("direction"),
//   					SortPaths: []interface{}{
//   						&DataPathValueProperty{
//   							DataPathType: &DataPathTypeProperty{
//   								PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//   							FieldValue: jsii.String("fieldValue"),
//   						},
//   					},
//   				},
//   				Field: &FieldSortProperty{
//   					Direction: jsii.String("direction"),
//   					FieldId: jsii.String("fieldId"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablesortconfiguration.html
//
type CfnDashboardPropsMixin_PivotTableSortConfigurationProperty struct {
	// The field sort options for a pivot table sort configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-pivottablesortconfiguration.html#cfn-quicksight-dashboard-pivottablesortconfiguration-fieldsortoptions
	//
	FieldSortOptions interface{} `field:"optional" json:"fieldSortOptions" yaml:"fieldSortOptions"`
}


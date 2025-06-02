package awsquicksight


// The sort configuration for a `PivotTableVisual` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableSortConfigurationProperty := &PivotTableSortConfigurationProperty{
//   	FieldSortOptions: []interface{}{
//   		&PivotFieldSortOptionsProperty{
//   			FieldId: jsii.String("fieldId"),
//   			SortBy: &PivotTableSortByProperty{
//   				Column: &ColumnSortProperty{
//   					Direction: jsii.String("direction"),
//   					SortBy: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//
//   					// the properties below are optional
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablesortconfiguration.html
//
type CfnAnalysis_PivotTableSortConfigurationProperty struct {
	// The field sort options for a pivot table sort configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablesortconfiguration.html#cfn-quicksight-analysis-pivottablesortconfiguration-fieldsortoptions
	//
	FieldSortOptions interface{} `field:"optional" json:"fieldSortOptions" yaml:"fieldSortOptions"`
}


package awsquicksight


// The sort by field for the field sort options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotTableSortByProperty := &PivotTableSortByProperty{
//   	Column: &ColumnSortProperty{
//   		Direction: jsii.String("direction"),
//   		SortBy: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//
//   		// the properties below are optional
//   		AggregationFunction: &AggregationFunctionProperty{
//   			AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   				SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   				ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   			},
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
//   	DataPath: &DataPathSortProperty{
//   		Direction: jsii.String("direction"),
//   		SortPaths: []interface{}{
//   			&DataPathValueProperty{
//   				DataPathType: &DataPathTypeProperty{
//   					PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//   				FieldValue: jsii.String("fieldValue"),
//   			},
//   		},
//   	},
//   	Field: &FieldSortProperty{
//   		Direction: jsii.String("direction"),
//   		FieldId: jsii.String("fieldId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablesortby.html
//
type CfnAnalysis_PivotTableSortByProperty struct {
	// The column sort (field id, direction) for the pivot table sort by options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablesortby.html#cfn-quicksight-analysis-pivottablesortby-column
	//
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// The data path sort (data path value, direction) for the pivot table sort by options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablesortby.html#cfn-quicksight-analysis-pivottablesortby-datapath
	//
	DataPath interface{} `field:"optional" json:"dataPath" yaml:"dataPath"`
	// The field sort (field id, direction) for the pivot table sort by options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivottablesortby.html#cfn-quicksight-analysis-pivottablesortby-field
	//
	Field interface{} `field:"optional" json:"field" yaml:"field"`
}


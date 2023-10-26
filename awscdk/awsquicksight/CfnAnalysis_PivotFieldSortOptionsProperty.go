package awsquicksight


// The field sort options for a pivot table sort configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pivotFieldSortOptionsProperty := &PivotFieldSortOptionsProperty{
//   	FieldId: jsii.String("fieldId"),
//   	SortBy: &PivotTableSortByProperty{
//   		Column: &ColumnSortProperty{
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
//   		DataPath: &DataPathSortProperty{
//   			Direction: jsii.String("direction"),
//   			SortPaths: []interface{}{
//   				&DataPathValueProperty{
//   					FieldId: jsii.String("fieldId"),
//   					FieldValue: jsii.String("fieldValue"),
//
//   					// the properties below are optional
//   					DataPathType: &DataPathTypeProperty{
//   						PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   					},
//   				},
//   			},
//   		},
//   		Field: &FieldSortProperty{
//   			Direction: jsii.String("direction"),
//   			FieldId: jsii.String("fieldId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivotfieldsortoptions.html
//
type CfnAnalysis_PivotFieldSortOptionsProperty struct {
	// The field ID for the field sort options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivotfieldsortoptions.html#cfn-quicksight-analysis-pivotfieldsortoptions-fieldid
	//
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The sort by field for the field sort options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-pivotfieldsortoptions.html#cfn-quicksight-analysis-pivotfieldsortoptions-sortby
	//
	SortBy interface{} `field:"required" json:"sortBy" yaml:"sortBy"`
}


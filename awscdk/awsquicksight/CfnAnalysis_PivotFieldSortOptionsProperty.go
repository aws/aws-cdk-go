package awsquicksight


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
type CfnAnalysis_PivotFieldSortOptionsProperty struct {
	// `CfnAnalysis.PivotFieldSortOptionsProperty.FieldId`.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// `CfnAnalysis.PivotFieldSortOptionsProperty.SortBy`.
	SortBy interface{} `field:"required" json:"sortBy" yaml:"sortBy"`
}


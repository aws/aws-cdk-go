package awsquicksight


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
type CfnDashboard_PivotTableSortByProperty struct {
	// `CfnDashboard.PivotTableSortByProperty.Column`.
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// `CfnDashboard.PivotTableSortByProperty.DataPath`.
	DataPath interface{} `field:"optional" json:"dataPath" yaml:"dataPath"`
	// `CfnDashboard.PivotTableSortByProperty.Field`.
	Field interface{} `field:"optional" json:"field" yaml:"field"`
}


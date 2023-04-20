package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   treeMapSortConfigurationProperty := &TreeMapSortConfigurationProperty{
//   	TreeMapGroupItemsLimitConfiguration: &ItemsLimitConfigurationProperty{
//   		ItemsLimit: jsii.Number(123),
//   		OtherCategories: jsii.String("otherCategories"),
//   	},
//   	TreeMapSort: []interface{}{
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
type CfnDashboard_TreeMapSortConfigurationProperty struct {
	// `CfnDashboard.TreeMapSortConfigurationProperty.TreeMapGroupItemsLimitConfiguration`.
	TreeMapGroupItemsLimitConfiguration interface{} `field:"optional" json:"treeMapGroupItemsLimitConfiguration" yaml:"treeMapGroupItemsLimitConfiguration"`
	// `CfnDashboard.TreeMapSortConfigurationProperty.TreeMapSort`.
	TreeMapSort interface{} `field:"optional" json:"treeMapSort" yaml:"treeMapSort"`
}


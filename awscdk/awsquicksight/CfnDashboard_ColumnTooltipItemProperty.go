package awsquicksight


// The tooltip item for the columns that are not part of a field well.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnTooltipItemProperty := &ColumnTooltipItemProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//
//   	// the properties below are optional
//   	Aggregation: &AggregationFunctionProperty{
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
//   	Label: jsii.String("label"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columntooltipitem.html
//
type CfnDashboard_ColumnTooltipItemProperty struct {
	// The target column of the tooltip item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columntooltipitem.html#cfn-quicksight-dashboard-columntooltipitem-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The aggregation function of the column tooltip item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columntooltipitem.html#cfn-quicksight-dashboard-columntooltipitem-aggregation
	//
	Aggregation interface{} `field:"optional" json:"aggregation" yaml:"aggregation"`
	// The label of the tooltip item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columntooltipitem.html#cfn-quicksight-dashboard-columntooltipitem-label
	//
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The visibility of the tooltip item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columntooltipitem.html#cfn-quicksight-dashboard-columntooltipitem-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}


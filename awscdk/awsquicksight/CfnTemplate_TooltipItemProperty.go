package awsquicksight


// The tooltip.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tooltipItemProperty := &TooltipItemProperty{
//   	ColumnTooltipItem: &ColumnTooltipItemProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//
//   		// the properties below are optional
//   		Aggregation: &AggregationFunctionProperty{
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
//   		Label: jsii.String("label"),
//   		TooltipTarget: jsii.String("tooltipTarget"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   	FieldTooltipItem: &FieldTooltipItemProperty{
//   		FieldId: jsii.String("fieldId"),
//
//   		// the properties below are optional
//   		Label: jsii.String("label"),
//   		TooltipTarget: jsii.String("tooltipTarget"),
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tooltipitem.html
//
type CfnTemplate_TooltipItemProperty struct {
	// The tooltip item for the columns that are not part of a field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tooltipitem.html#cfn-quicksight-template-tooltipitem-columntooltipitem
	//
	ColumnTooltipItem interface{} `field:"optional" json:"columnTooltipItem" yaml:"columnTooltipItem"`
	// The tooltip item for the fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-tooltipitem.html#cfn-quicksight-template-tooltipitem-fieldtooltipitem
	//
	FieldTooltipItem interface{} `field:"optional" json:"fieldTooltipItem" yaml:"fieldTooltipItem"`
}


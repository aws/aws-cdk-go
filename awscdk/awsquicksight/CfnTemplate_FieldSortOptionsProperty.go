package awsquicksight


// The field sort options in a chart configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fieldSortOptionsProperty := &FieldSortOptionsProperty{
//   	ColumnSort: &ColumnSortProperty{
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
//   	FieldSort: &FieldSortProperty{
//   		Direction: jsii.String("direction"),
//   		FieldId: jsii.String("fieldId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fieldsortoptions.html
//
type CfnTemplate_FieldSortOptionsProperty struct {
	// The sort configuration for a column that is not used in a field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fieldsortoptions.html#cfn-quicksight-template-fieldsortoptions-columnsort
	//
	ColumnSort interface{} `field:"optional" json:"columnSort" yaml:"columnSort"`
	// The sort configuration for a field in a field well.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-fieldsortoptions.html#cfn-quicksight-template-fieldsortoptions-fieldsort
	//
	FieldSort interface{} `field:"optional" json:"fieldSort" yaml:"fieldSort"`
}


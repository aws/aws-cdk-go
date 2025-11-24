package mixinsawsquicksight


// Determines how the plugin visual sorts the data during query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pluginVisualSortConfigurationProperty := &PluginVisualSortConfigurationProperty{
//   	PluginVisualTableQuerySort: &PluginVisualTableQuerySortProperty{
//   		ItemsLimitConfiguration: &PluginVisualItemsLimitConfigurationProperty{
//   			ItemsLimit: jsii.Number(123),
//   		},
//   		RowSort: []interface{}{
//   			&FieldSortOptionsProperty{
//   				ColumnSort: &ColumnSortProperty{
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
//   				FieldSort: &FieldSortProperty{
//   					Direction: jsii.String("direction"),
//   					FieldId: jsii.String("fieldId"),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualsortconfiguration.html
//
type CfnTemplatePropsMixin_PluginVisualSortConfigurationProperty struct {
	// The table query sorting options for the plugin visual.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-pluginvisualsortconfiguration.html#cfn-quicksight-template-pluginvisualsortconfiguration-pluginvisualtablequerysort
	//
	PluginVisualTableQuerySort interface{} `field:"optional" json:"pluginVisualTableQuerySort" yaml:"pluginVisualTableQuerySort"`
}


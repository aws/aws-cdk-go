package previewawsquicksightmixins


// Describes the configurations that are required to declare a section as repeating.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bodySectionRepeatConfigurationProperty := &BodySectionRepeatConfigurationProperty{
//   	DimensionConfigurations: []interface{}{
//   		&BodySectionRepeatDimensionConfigurationProperty{
//   			DynamicCategoryDimensionConfiguration: &BodySectionDynamicCategoryDimensionConfigurationProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				Limit: jsii.Number(123),
//   				SortByMetrics: []interface{}{
//   					&ColumnSortProperty{
//   						AggregationFunction: &AggregationFunctionProperty{
//   							AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   								SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   								ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   							},
//   							CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   							DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   							NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   						},
//   						Direction: jsii.String("direction"),
//   						SortBy: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   					},
//   				},
//   			},
//   			DynamicNumericDimensionConfiguration: &BodySectionDynamicNumericDimensionConfigurationProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				Limit: jsii.Number(123),
//   				SortByMetrics: []interface{}{
//   					&ColumnSortProperty{
//   						AggregationFunction: &AggregationFunctionProperty{
//   							AttributeAggregationFunction: &AttributeAggregationFunctionProperty{
//   								SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   								ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   							},
//   							CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   							DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   							NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   						},
//   						Direction: jsii.String("direction"),
//   						SortBy: &ColumnIdentifierProperty{
//   							ColumnName: jsii.String("columnName"),
//   							DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	NonRepeatingVisuals: []*string{
//   		jsii.String("nonRepeatingVisuals"),
//   	},
//   	PageBreakConfiguration: &BodySectionRepeatPageBreakConfigurationProperty{
//   		After: &SectionAfterPageBreakProperty{
//   			Status: jsii.String("status"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionrepeatconfiguration.html
//
type CfnAnalysisPropsMixin_BodySectionRepeatConfigurationProperty struct {
	// List of `BodySectionRepeatDimensionConfiguration` values that describe the dataset column and constraints for the column used to repeat the contents of a section.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionrepeatconfiguration.html#cfn-quicksight-analysis-bodysectionrepeatconfiguration-dimensionconfigurations
	//
	DimensionConfigurations interface{} `field:"optional" json:"dimensionConfigurations" yaml:"dimensionConfigurations"`
	// List of visuals to exclude from repetition in repeating sections.
	//
	// The visuals will render identically, and ignore the repeating configurations in all repeating instances.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionrepeatconfiguration.html#cfn-quicksight-analysis-bodysectionrepeatconfiguration-nonrepeatingvisuals
	//
	NonRepeatingVisuals *[]*string `field:"optional" json:"nonRepeatingVisuals" yaml:"nonRepeatingVisuals"`
	// Page break configuration to apply for each repeating instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-bodysectionrepeatconfiguration.html#cfn-quicksight-analysis-bodysectionrepeatconfiguration-pagebreakconfiguration
	//
	PageBreakConfiguration interface{} `field:"optional" json:"pageBreakConfiguration" yaml:"pageBreakConfiguration"`
}


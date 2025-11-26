package previewawsquicksightmixins


// The data configuration of the reference line.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceLineDataConfigurationProperty := &ReferenceLineDataConfigurationProperty{
//   	AxisBinding: jsii.String("axisBinding"),
//   	DynamicConfiguration: &ReferenceLineDynamicDataConfigurationProperty{
//   		Calculation: &NumericalAggregationFunctionProperty{
//   			PercentileAggregation: &PercentileAggregationProperty{
//   				PercentileValue: jsii.Number(123),
//   			},
//   			SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   		},
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		MeasureAggregationFunction: &AggregationFunctionProperty{
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
//   	SeriesType: jsii.String("seriesType"),
//   	StaticConfiguration: &ReferenceLineStaticDataConfigurationProperty{
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinedataconfiguration.html
//
type CfnDashboardPropsMixin_ReferenceLineDataConfigurationProperty struct {
	// The axis binding type of the reference line. Choose one of the following options:.
	//
	// - `PrimaryY`
	// - `SecondaryY`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinedataconfiguration.html#cfn-quicksight-dashboard-referencelinedataconfiguration-axisbinding
	//
	AxisBinding *string `field:"optional" json:"axisBinding" yaml:"axisBinding"`
	// The dynamic configuration of the reference line data configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinedataconfiguration.html#cfn-quicksight-dashboard-referencelinedataconfiguration-dynamicconfiguration
	//
	DynamicConfiguration interface{} `field:"optional" json:"dynamicConfiguration" yaml:"dynamicConfiguration"`
	// The series type of the reference line data configuration. Choose one of the following options:.
	//
	// - `BAR`
	// - `LINE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinedataconfiguration.html#cfn-quicksight-dashboard-referencelinedataconfiguration-seriestype
	//
	SeriesType *string `field:"optional" json:"seriesType" yaml:"seriesType"`
	// The static data configuration of the reference line data configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-referencelinedataconfiguration.html#cfn-quicksight-dashboard-referencelinedataconfiguration-staticconfiguration
	//
	StaticConfiguration interface{} `field:"optional" json:"staticConfiguration" yaml:"staticConfiguration"`
}


package awsquicksight


// The label options for an axis on a chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   chartAxisLabelOptionsProperty := &ChartAxisLabelOptionsProperty{
//   	AxisLabelOptions: []interface{}{
//   		&AxisLabelOptionsProperty{
//   			ApplyTo: &AxisLabelReferenceOptionsProperty{
//   				Column: &ColumnIdentifierProperty{
//   					ColumnName: jsii.String("columnName"),
//   					DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   				},
//   				FieldId: jsii.String("fieldId"),
//   			},
//   			CustomLabel: jsii.String("customLabel"),
//   			FontConfiguration: &FontConfigurationProperty{
//   				FontColor: jsii.String("fontColor"),
//   				FontDecoration: jsii.String("fontDecoration"),
//   				FontFamily: jsii.String("fontFamily"),
//   				FontSize: &FontSizeProperty{
//   					Absolute: jsii.String("absolute"),
//   					Relative: jsii.String("relative"),
//   				},
//   				FontStyle: jsii.String("fontStyle"),
//   				FontWeight: &FontWeightProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	SortIconVisibility: jsii.String("sortIconVisibility"),
//   	Visibility: jsii.String("visibility"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-chartaxislabeloptions.html
//
type CfnTemplate_ChartAxisLabelOptionsProperty struct {
	// The label options for a chart axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-chartaxislabeloptions.html#cfn-quicksight-template-chartaxislabeloptions-axislabeloptions
	//
	AxisLabelOptions interface{} `field:"optional" json:"axisLabelOptions" yaml:"axisLabelOptions"`
	// The visibility configuration of the sort icon on a chart's axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-chartaxislabeloptions.html#cfn-quicksight-template-chartaxislabeloptions-sorticonvisibility
	//
	SortIconVisibility *string `field:"optional" json:"sortIconVisibility" yaml:"sortIconVisibility"`
	// The visibility of an axis label on a chart. Choose one of the following options:.
	//
	// - `VISIBLE` : Shows the axis.
	// - `HIDDEN` : Hides the axis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-chartaxislabeloptions.html#cfn-quicksight-template-chartaxislabeloptions-visibility
	//
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}


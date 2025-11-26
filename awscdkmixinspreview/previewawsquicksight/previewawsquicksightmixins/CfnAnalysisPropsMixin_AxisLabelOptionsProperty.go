package previewawsquicksightmixins


// The label options for a chart axis.
//
// You must specify the field that the label is targeted to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   axisLabelOptionsProperty := &AxisLabelOptionsProperty{
//   	ApplyTo: &AxisLabelReferenceOptionsProperty{
//   		Column: &ColumnIdentifierProperty{
//   			ColumnName: jsii.String("columnName"),
//   			DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   		},
//   		FieldId: jsii.String("fieldId"),
//   	},
//   	CustomLabel: jsii.String("customLabel"),
//   	FontConfiguration: &FontConfigurationProperty{
//   		FontColor: jsii.String("fontColor"),
//   		FontDecoration: jsii.String("fontDecoration"),
//   		FontFamily: jsii.String("fontFamily"),
//   		FontSize: &FontSizeProperty{
//   			Absolute: jsii.String("absolute"),
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-axislabeloptions.html
//
type CfnAnalysisPropsMixin_AxisLabelOptionsProperty struct {
	// The options that indicate which field the label belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-axislabeloptions.html#cfn-quicksight-analysis-axislabeloptions-applyto
	//
	ApplyTo interface{} `field:"optional" json:"applyTo" yaml:"applyTo"`
	// The text for the axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-axislabeloptions.html#cfn-quicksight-analysis-axislabeloptions-customlabel
	//
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// The font configuration of the axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-axislabeloptions.html#cfn-quicksight-analysis-axislabeloptions-fontconfiguration
	//
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
}


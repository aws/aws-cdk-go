package awsquicksight


// The label options for a chart axis.
//
// You must specify the field that the label is targeted to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   		FontSize: &FontSizeProperty{
//   			Relative: jsii.String("relative"),
//   		},
//   		FontStyle: jsii.String("fontStyle"),
//   		FontWeight: &FontWeightProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislabeloptions.html
//
type CfnDashboard_AxisLabelOptionsProperty struct {
	// The options that indicate which field the label belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislabeloptions.html#cfn-quicksight-dashboard-axislabeloptions-applyto
	//
	ApplyTo interface{} `field:"optional" json:"applyTo" yaml:"applyTo"`
	// The text for the axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislabeloptions.html#cfn-quicksight-dashboard-axislabeloptions-customlabel
	//
	CustomLabel *string `field:"optional" json:"customLabel" yaml:"customLabel"`
	// The font configuration of the axis label.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-axislabeloptions.html#cfn-quicksight-dashboard-axislabeloptions-fontconfiguration
	//
	FontConfiguration interface{} `field:"optional" json:"fontConfiguration" yaml:"fontConfiguration"`
}


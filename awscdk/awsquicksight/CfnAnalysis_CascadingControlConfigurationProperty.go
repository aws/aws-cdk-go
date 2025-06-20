package awsquicksight


// The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cascadingControlConfigurationProperty := &CascadingControlConfigurationProperty{
//   	SourceControls: []interface{}{
//   		&CascadingControlSourceProperty{
//   			ColumnToMatch: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-cascadingcontrolconfiguration.html
//
type CfnAnalysis_CascadingControlConfigurationProperty struct {
	// A list of source controls that determine the values that are used in the current control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-cascadingcontrolconfiguration.html#cfn-quicksight-analysis-cascadingcontrolconfiguration-sourcecontrols
	//
	SourceControls interface{} `field:"optional" json:"sourceControls" yaml:"sourceControls"`
}


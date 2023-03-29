package awsquicksight


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
type CfnDashboard_CascadingControlConfigurationProperty struct {
	// `CfnDashboard.CascadingControlConfigurationProperty.SourceControls`.
	SourceControls interface{} `field:"optional" json:"sourceControls" yaml:"sourceControls"`
}


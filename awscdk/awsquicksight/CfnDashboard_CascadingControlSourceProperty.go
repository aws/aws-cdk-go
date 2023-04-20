package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cascadingControlSourceProperty := &CascadingControlSourceProperty{
//   	ColumnToMatch: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   }
//
type CfnDashboard_CascadingControlSourceProperty struct {
	// `CfnDashboard.CascadingControlSourceProperty.ColumnToMatch`.
	ColumnToMatch interface{} `field:"optional" json:"columnToMatch" yaml:"columnToMatch"`
	// `CfnDashboard.CascadingControlSourceProperty.SourceSheetControlId`.
	SourceSheetControlId *string `field:"optional" json:"sourceSheetControlId" yaml:"sourceSheetControlId"`
}


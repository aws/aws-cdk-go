package awsquicksight


// A column of a data set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnIdentifierProperty := &ColumnIdentifierProperty{
//   	ColumnName: jsii.String("columnName"),
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnidentifier.html
//
type CfnDashboard_ColumnIdentifierProperty struct {
	// The name of the column.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnidentifier.html#cfn-quicksight-dashboard-columnidentifier-columnname
	//
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The data set that the column belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-columnidentifier.html#cfn-quicksight-dashboard-columnidentifier-datasetidentifier
	//
	DataSetIdentifier *string `field:"required" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
}


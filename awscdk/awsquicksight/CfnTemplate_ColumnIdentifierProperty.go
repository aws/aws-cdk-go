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
type CfnTemplate_ColumnIdentifierProperty struct {
	// The name of the column.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The data set that the column belongs to.
	DataSetIdentifier *string `field:"required" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
}


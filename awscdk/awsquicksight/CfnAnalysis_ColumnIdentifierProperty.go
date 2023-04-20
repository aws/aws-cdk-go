package awsquicksight


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
type CfnAnalysis_ColumnIdentifierProperty struct {
	// `CfnAnalysis.ColumnIdentifierProperty.ColumnName`.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// `CfnAnalysis.ColumnIdentifierProperty.DataSetIdentifier`.
	DataSetIdentifier *string `field:"required" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
}


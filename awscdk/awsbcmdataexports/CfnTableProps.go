package awsbcmdataexports


// Properties for defining a `CfnTable`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTableProps := &CfnTableProps{
//   	TableName: jsii.String("tableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmdataexports-table.html
//
type CfnTableProps struct {
	// The name of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmdataexports-table.html#cfn-bcmdataexports-table-tablename
	//
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


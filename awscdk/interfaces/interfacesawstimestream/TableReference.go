package interfacesawstimestream


// A reference to a Table resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableReference := &TableReference{
//   	DatabaseName: jsii.String("databaseName"),
//   	TableArn: jsii.String("tableArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
type TableReference struct {
	// The DatabaseName of the Table resource.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The ARN of the Table resource.
	TableArn *string `field:"required" json:"tableArn" yaml:"tableArn"`
	// The TableName of the Table resource.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


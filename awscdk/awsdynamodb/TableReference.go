package awsdynamodb


// A reference to a Table resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableReference := &TableReference{
//   	TableArn: jsii.String("tableArn"),
//   	TableName: jsii.String("tableName"),
//   }
//
type TableReference struct {
	// The ARN of the Table resource.
	TableArn *string `field:"required" json:"tableArn" yaml:"tableArn"`
	// The TableName of the Table resource.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


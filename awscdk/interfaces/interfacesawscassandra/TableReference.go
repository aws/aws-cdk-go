package interfacesawscassandra


// A reference to a Table resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableReference := &TableReference{
//   	KeyspaceName: jsii.String("keyspaceName"),
//   	TableName: jsii.String("tableName"),
//   }
//
type TableReference struct {
	// The KeyspaceName of the Table resource.
	KeyspaceName *string `field:"required" json:"keyspaceName" yaml:"keyspaceName"`
	// The TableName of the Table resource.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


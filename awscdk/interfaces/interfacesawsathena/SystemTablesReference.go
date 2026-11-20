package interfacesawsathena


// A reference to a SystemTables resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemTablesReference := &SystemTablesReference{
//   	TableName: jsii.String("tableName"),
//   }
//
type SystemTablesReference struct {
	// The TableName of the SystemTables resource.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


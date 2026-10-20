package interfacesawsglue


// A reference to a Table resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableReference := &TableReference{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	TableName: jsii.String("tableName"),
//   }
//
type TableReference struct {
	// The CatalogId of the Table resource.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The DatabaseName of the Table resource.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The Name of the Table resource.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


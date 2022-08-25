package awslakeformation


// A structure for the table object.
//
// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableResourceProperty := &tableResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   	tableWildcard: &tableWildcardProperty{
//   	},
//   }
//
type CfnPermissions_TableResourceProperty struct {
	// `CfnPermissions.TableResourceProperty.CatalogId`.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the database for the table.
	//
	// Unique to a Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the table.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An empty object representing all tables under a database.
	//
	// If this field is specified instead of the `Name` field, all tables under `DatabaseName` will have permission changes applied.
	TableWildcard interface{} `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}


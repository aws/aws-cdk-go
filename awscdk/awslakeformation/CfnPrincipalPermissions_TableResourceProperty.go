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
//   var tableWildcard interface{}
//
//   tableResourceProperty := &TableResourceProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   	TableWildcard: tableWildcard,
//   }
//
type CfnPrincipalPermissions_TableResourceProperty struct {
	// `CfnPrincipalPermissions.TableResourceProperty.CatalogId`.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The name of the database for the table.
	//
	// Unique to a Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name of the table.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A wildcard object representing every table under a database.
	//
	// At least one of `TableResource$Name` or `TableResource$TableWildcard` is required.
	TableWildcard interface{} `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}


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
//   tableResourceProperty := &tableResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	databaseName: jsii.String("databaseName"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tableWildcard: tableWildcard,
//   }
//
type CfnTagAssociation_TableResourceProperty struct {
	// The identifier for the Data Catalog .
	//
	// By default, it is the account ID of the caller.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The name of the database for the table.
	//
	// Unique to a Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name of the table.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A wildcard object representing every table under a database.This is an object with no properties that effectively behaves as a true or false depending on whether not it is passed as a parameter. The valid inputs for a property with this type in either yaml or json is null or {}.
	//
	// At least one of `TableResource$Name` or `TableResource$TableWildcard` is required.
	TableWildcard interface{} `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}


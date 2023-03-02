package awslakeformation


// A structure for a table with columns object. This object is only used when granting a SELECT permission.
//
// This object must take a value for at least one of `ColumnsNames` , `ColumnsIndexes` , or `ColumnsWildcard` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableWithColumnsResourceProperty := &tableWithColumnsResourceProperty{
//   	catalogId: jsii.String("catalogId"),
//   	columnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   }
//
type CfnTagAssociation_TableWithColumnsResourceProperty struct {
	// A wildcard object representing every table under a database.
	//
	// At least one of TableResource$Name or TableResource$TableWildcard is required.
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The list of column names for the table.
	//
	// At least one of `ColumnNames` or `ColumnWildcard` is required.
	ColumnNames *[]*string `field:"required" json:"columnNames" yaml:"columnNames"`
	// The name of the database for the table with columns resource.
	//
	// Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name of the table resource.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	Name *string `field:"required" json:"name" yaml:"name"`
}


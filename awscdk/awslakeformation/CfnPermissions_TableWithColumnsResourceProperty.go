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
//   tableWithColumnsResourceProperty := &TableWithColumnsResourceProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	ColumnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	ColumnWildcard: &ColumnWildcardProperty{
//   		ExcludedColumnNames: []*string{
//   			jsii.String("excludedColumnNames"),
//   		},
//   	},
//   	DatabaseName: jsii.String("databaseName"),
//   	Name: jsii.String("name"),
//   }
//
type CfnPermissions_TableWithColumnsResourceProperty struct {
	// The identifier for the Data Catalog .
	//
	// By default, it is the account ID of the caller.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The list of column names for the table.
	//
	// At least one of `ColumnNames` or `ColumnWildcard` is required.
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// A wildcard specified by a `ColumnWildcard` object.
	//
	// At least one of `ColumnNames` or `ColumnWildcard` is required.
	ColumnWildcard interface{} `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
	// The name of the database for the table with columns resource.
	//
	// Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the table resource.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	Name *string `field:"optional" json:"name" yaml:"name"`
}


package awslakeformation


// A structure that describes certain columns on certain rows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCellsFilterResourceProperty := &DataCellsFilterResourceProperty{
//   	DatabaseName: jsii.String("databaseName"),
//   	Name: jsii.String("name"),
//   	TableCatalogId: jsii.String("tableCatalogId"),
//   	TableName: jsii.String("tableName"),
//   }
//
type CfnPrincipalPermissions_DataCellsFilterResourceProperty struct {
	// A database in the Data Catalog .
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name given by the user to the data filter cell.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the catalog to which the table belongs.
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// The name of the table.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


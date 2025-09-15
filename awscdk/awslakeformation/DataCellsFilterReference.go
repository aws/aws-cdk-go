package awslakeformation


// A reference to a DataCellsFilter resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCellsFilterReference := &DataCellsFilterReference{
//   	DatabaseName: jsii.String("databaseName"),
//   	DataCellsFilterName: jsii.String("dataCellsFilterName"),
//   	TableCatalogId: jsii.String("tableCatalogId"),
//   	TableName: jsii.String("tableName"),
//   }
//
type DataCellsFilterReference struct {
	// The DatabaseName of the DataCellsFilter resource.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The Name of the DataCellsFilter resource.
	DataCellsFilterName *string `field:"required" json:"dataCellsFilterName" yaml:"dataCellsFilterName"`
	// The TableCatalogId of the DataCellsFilter resource.
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// The TableName of the DataCellsFilter resource.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


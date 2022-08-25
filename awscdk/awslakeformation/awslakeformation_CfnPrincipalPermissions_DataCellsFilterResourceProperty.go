package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataCellsFilterResourceProperty := &dataCellsFilterResourceProperty{
//   	databaseName: jsii.String("databaseName"),
//   	name: jsii.String("name"),
//   	tableCatalogId: jsii.String("tableCatalogId"),
//   	tableName: jsii.String("tableName"),
//   }
//
type CfnPrincipalPermissions_DataCellsFilterResourceProperty struct {
	// `CfnPrincipalPermissions.DataCellsFilterResourceProperty.DatabaseName`.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// `CfnPrincipalPermissions.DataCellsFilterResourceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnPrincipalPermissions.DataCellsFilterResourceProperty.TableCatalogId`.
	TableCatalogId *string `field:"required" json:"tableCatalogId" yaml:"tableCatalogId"`
	// `CfnPrincipalPermissions.DataCellsFilterResourceProperty.TableName`.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


package awslakeformation


// A structure for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   resourceProperty := &ResourceProperty{
//   	Catalog: catalog,
//   	Database: &DatabaseResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		Name: jsii.String("name"),
//   	},
//   	DataCellsFilter: &DataCellsFilterResourceProperty{
//   		DatabaseName: jsii.String("databaseName"),
//   		Name: jsii.String("name"),
//   		TableCatalogId: jsii.String("tableCatalogId"),
//   		TableName: jsii.String("tableName"),
//   	},
//   	DataLocation: &DataLocationResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		ResourceArn: jsii.String("resourceArn"),
//   	},
//   	LfTag: &LFTagKeyResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		TagKey: jsii.String("tagKey"),
//   		TagValues: []*string{
//   			jsii.String("tagValues"),
//   		},
//   	},
//   	LfTagPolicy: &LFTagPolicyResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		Expression: []interface{}{
//   			&LFTagProperty{
//   				TagKey: jsii.String("tagKey"),
//   				TagValues: []*string{
//   					jsii.String("tagValues"),
//   				},
//   			},
//   		},
//   		ResourceType: jsii.String("resourceType"),
//   	},
//   	Table: &TableResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		DatabaseName: jsii.String("databaseName"),
//
//   		// the properties below are optional
//   		Name: jsii.String("name"),
//   		TableWildcard: tableWildcard,
//   	},
//   	TableWithColumns: &TableWithColumnsResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		ColumnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		ColumnWildcard: &ColumnWildcardProperty{
//   			ExcludedColumnNames: []*string{
//   				jsii.String("excludedColumnNames"),
//   			},
//   		},
//   	},
//   }
//
type CfnPrincipalPermissions_ResourceProperty struct {
	// The identifier for the Data Catalog.
	//
	// By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.
	Catalog interface{} `field:"optional" json:"catalog" yaml:"catalog"`
	// The database for the resource.
	//
	// Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database permissions to a principal.
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// A data cell filter.
	DataCellsFilter interface{} `field:"optional" json:"dataCellsFilter" yaml:"dataCellsFilter"`
	// The location of an Amazon S3 path where permissions are granted or revoked.
	DataLocation interface{} `field:"optional" json:"dataLocation" yaml:"dataLocation"`
	// The LF-tag key and values attached to a resource.
	LfTag interface{} `field:"optional" json:"lfTag" yaml:"lfTag"`
	// A list of LF-tag conditions that define a resource's LF-tag policy.
	LfTagPolicy interface{} `field:"optional" json:"lfTagPolicy" yaml:"lfTagPolicy"`
	// The table for the resource.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	Table interface{} `field:"optional" json:"table" yaml:"table"`
	// The table with columns for the resource.
	//
	// A principal with permissions to this resource can select metadata from the columns of a table in the Data Catalog and the underlying data in Amazon S3.
	TableWithColumns interface{} `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}


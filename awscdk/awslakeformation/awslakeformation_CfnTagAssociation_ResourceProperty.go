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
//   resourceProperty := &resourceProperty{
//   	catalog: catalog,
//   	database: &databaseResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		name: jsii.String("name"),
//   	},
//   	table: &tableResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   		tableWildcard: tableWildcard,
//   	},
//   	tableWithColumns: &tableWithColumnsResourceProperty{
//   		catalogId: jsii.String("catalogId"),
//   		columnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		databaseName: jsii.String("databaseName"),
//   		name: jsii.String("name"),
//   	},
//   }
//
type CfnTagAssociation_ResourceProperty struct {
	// The identifier for the Data Catalog.
	//
	// By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.
	Catalog interface{} `field:"optional" json:"catalog" yaml:"catalog"`
	// The database for the resource.
	//
	// Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database permissions to a principal.
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// The table for the resource.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	Table interface{} `field:"optional" json:"table" yaml:"table"`
	// The table with columns for the resource.
	//
	// A principal with permissions to this resource can select metadata from the columns of a table in the Data Catalog and the underlying data in Amazon S3.
	TableWithColumns interface{} `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}


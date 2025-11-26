package previewawslakeformationmixins


// A structure for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   	Table: &TableResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Name: jsii.String("name"),
//   		TableWildcard: tableWildcard,
//   	},
//   	TableWithColumns: &TableWithColumnsResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		ColumnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		DatabaseName: jsii.String("databaseName"),
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-tagassociation-resource.html
//
type CfnTagAssociationPropsMixin_ResourceProperty struct {
	// The identifier for the Data Catalog.
	//
	// By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-tagassociation-resource.html#cfn-lakeformation-tagassociation-resource-catalog
	//
	Catalog interface{} `field:"optional" json:"catalog" yaml:"catalog"`
	// The database for the resource.
	//
	// Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database permissions to a principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-tagassociation-resource.html#cfn-lakeformation-tagassociation-resource-database
	//
	Database interface{} `field:"optional" json:"database" yaml:"database"`
	// The table for the resource.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-tagassociation-resource.html#cfn-lakeformation-tagassociation-resource-table
	//
	Table interface{} `field:"optional" json:"table" yaml:"table"`
	// The table with columns for the resource.
	//
	// A principal with permissions to this resource can select metadata from the columns of a table in the Data Catalog and the underlying data in Amazon S3.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-tagassociation-resource.html#cfn-lakeformation-tagassociation-resource-tablewithcolumns
	//
	TableWithColumns interface{} `field:"optional" json:"tableWithColumns" yaml:"tableWithColumns"`
}


package previewawslakeformationmixins


// A structure for the table object.
//
// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tableResourceProperty := &TableResourceProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseName: jsii.String("databaseName"),
//   	Name: jsii.String("name"),
//   	TableWildcard: &TableWildcardProperty{
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-tableresource.html
//
type CfnPermissionsPropsMixin_TableResourceProperty struct {
	// The identifier for the Data Catalog .
	//
	// By default, it is the account ID of the caller.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-tableresource.html#cfn-lakeformation-permissions-tableresource-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the database for the table.
	//
	// Unique to a Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-tableresource.html#cfn-lakeformation-permissions-tableresource-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// The name of the table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-tableresource.html#cfn-lakeformation-permissions-tableresource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An empty object representing all tables under a database.
	//
	// If this field is specified instead of the `Name` field, all tables under `DatabaseName` will have permission changes applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-tableresource.html#cfn-lakeformation-permissions-tableresource-tablewildcard
	//
	TableWildcard interface{} `field:"optional" json:"tableWildcard" yaml:"tableWildcard"`
}


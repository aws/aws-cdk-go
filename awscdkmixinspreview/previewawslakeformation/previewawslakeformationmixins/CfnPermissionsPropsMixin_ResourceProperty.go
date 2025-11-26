package previewawslakeformationmixins


// A structure for the resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceProperty := &ResourceProperty{
//   	DatabaseResource: &DatabaseResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		Name: jsii.String("name"),
//   	},
//   	DataLocationResource: &DataLocationResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		S3Resource: jsii.String("s3Resource"),
//   	},
//   	TableResource: &TableResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		DatabaseName: jsii.String("databaseName"),
//   		Name: jsii.String("name"),
//   		TableWildcard: &TableWildcardProperty{
//   		},
//   	},
//   	TableWithColumnsResource: &TableWithColumnsResourceProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		ColumnNames: []*string{
//   			jsii.String("columnNames"),
//   		},
//   		ColumnWildcard: &ColumnWildcardProperty{
//   			ExcludedColumnNames: []*string{
//   				jsii.String("excludedColumnNames"),
//   			},
//   		},
//   		DatabaseName: jsii.String("databaseName"),
//   		Name: jsii.String("name"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-resource.html
//
type CfnPermissionsPropsMixin_ResourceProperty struct {
	// A structure for the database object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-resource.html#cfn-lakeformation-permissions-resource-databaseresource
	//
	DatabaseResource interface{} `field:"optional" json:"databaseResource" yaml:"databaseResource"`
	// A structure for a data location object where permissions are granted or revoked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-resource.html#cfn-lakeformation-permissions-resource-datalocationresource
	//
	DataLocationResource interface{} `field:"optional" json:"dataLocationResource" yaml:"dataLocationResource"`
	// A structure for the table object.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-resource.html#cfn-lakeformation-permissions-resource-tableresource
	//
	TableResource interface{} `field:"optional" json:"tableResource" yaml:"tableResource"`
	// A structure for a table with columns object.
	//
	// This object is only used when granting a SELECT permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-permissions-resource.html#cfn-lakeformation-permissions-resource-tablewithcolumnsresource
	//
	TableWithColumnsResource interface{} `field:"optional" json:"tableWithColumnsResource" yaml:"tableWithColumnsResource"`
}


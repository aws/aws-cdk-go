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
//   	DatabaseName: jsii.String("databaseName"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ColumnNames: []*string{
//   		jsii.String("columnNames"),
//   	},
//   	ColumnWildcard: &ColumnWildcardProperty{
//   		ExcludedColumnNames: []*string{
//   			jsii.String("excludedColumnNames"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-tablewithcolumnsresource.html
//
type CfnPrincipalPermissions_TableWithColumnsResourceProperty struct {
	// The identifier for the Data Catalog where the location is registered with AWS Lake Formation .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-tablewithcolumnsresource.html#cfn-lakeformation-principalpermissions-tablewithcolumnsresource-catalogid
	//
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The name of the database for the table with columns resource.
	//
	// Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-tablewithcolumnsresource.html#cfn-lakeformation-principalpermissions-tablewithcolumnsresource-databasename
	//
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The name of the table resource.
	//
	// A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-tablewithcolumnsresource.html#cfn-lakeformation-principalpermissions-tablewithcolumnsresource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of column names for the table.
	//
	// At least one of `ColumnNames` or `ColumnWildcard` is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-tablewithcolumnsresource.html#cfn-lakeformation-principalpermissions-tablewithcolumnsresource-columnnames
	//
	ColumnNames *[]*string `field:"optional" json:"columnNames" yaml:"columnNames"`
	// A wildcard specified by a `ColumnWildcard` object.
	//
	// At least one of `ColumnNames` or `ColumnWildcard` is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-tablewithcolumnsresource.html#cfn-lakeformation-principalpermissions-tablewithcolumnsresource-columnwildcard
	//
	ColumnWildcard interface{} `field:"optional" json:"columnWildcard" yaml:"columnWildcard"`
}


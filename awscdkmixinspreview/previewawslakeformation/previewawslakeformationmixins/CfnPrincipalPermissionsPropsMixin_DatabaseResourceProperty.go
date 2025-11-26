package previewawslakeformationmixins


// A structure for the database object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   databaseResourceProperty := &DatabaseResourceProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-databaseresource.html
//
type CfnPrincipalPermissionsPropsMixin_DatabaseResourceProperty struct {
	// The identifier for the Data Catalog.
	//
	// By default, it is the account ID of the caller.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-databaseresource.html#cfn-lakeformation-principalpermissions-databaseresource-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The name of the database resource.
	//
	// Unique to the Data Catalog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-databaseresource.html#cfn-lakeformation-principalpermissions-databaseresource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}


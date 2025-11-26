package previewawslakeformationmixins


// Properties for CfnPermissionsPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPermissionsMixinProps := &CfnPermissionsMixinProps{
//   	DataLakePrincipal: &DataLakePrincipalProperty{
//   		DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	PermissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   	Resource: &ResourceProperty{
//   		DatabaseResource: &DatabaseResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			Name: jsii.String("name"),
//   		},
//   		DataLocationResource: &DataLocationResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			S3Resource: jsii.String("s3Resource"),
//   		},
//   		TableResource: &TableResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//   			TableWildcard: &TableWildcardProperty{
//   			},
//   		},
//   		TableWithColumnsResource: &TableWithColumnsResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			ColumnWildcard: &ColumnWildcardProperty{
//   				ExcludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html
//
type CfnPermissionsMixinProps struct {
	// The AWS Lake Formation principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html#cfn-lakeformation-permissions-datalakeprincipal
	//
	DataLakePrincipal interface{} `field:"optional" json:"dataLakePrincipal" yaml:"dataLakePrincipal"`
	// The permissions granted or revoked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html#cfn-lakeformation-permissions-permissions
	//
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Indicates the ability to grant permissions (as a subset of permissions granted).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html#cfn-lakeformation-permissions-permissionswithgrantoption
	//
	PermissionsWithGrantOption *[]*string `field:"optional" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
	// A structure for the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html#cfn-lakeformation-permissions-resource
	//
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
}


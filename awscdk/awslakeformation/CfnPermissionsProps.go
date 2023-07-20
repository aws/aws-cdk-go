package awslakeformation


// Properties for defining a `CfnPermissions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermissionsProps := &CfnPermissionsProps{
//   	DataLakePrincipal: &DataLakePrincipalProperty{
//   		DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
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
//
//   	// the properties below are optional
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	PermissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html
//
type CfnPermissionsProps struct {
	// The AWS Lake Formation principal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html#cfn-lakeformation-permissions-datalakeprincipal
	//
	DataLakePrincipal interface{} `field:"required" json:"dataLakePrincipal" yaml:"dataLakePrincipal"`
	// A structure for the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html#cfn-lakeformation-permissions-resource
	//
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
	// The permissions granted or revoked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html#cfn-lakeformation-permissions-permissions
	//
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Indicates the ability to grant permissions (as a subset of permissions granted).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-permissions.html#cfn-lakeformation-permissions-permissionswithgrantoption
	//
	PermissionsWithGrantOption *[]*string `field:"optional" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
}


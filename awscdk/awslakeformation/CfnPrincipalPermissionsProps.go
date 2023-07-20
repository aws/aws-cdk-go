package awslakeformation


// Properties for defining a `CfnPrincipalPermissions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var catalog interface{}
//   var tableWildcard interface{}
//
//   cfnPrincipalPermissionsProps := &CfnPrincipalPermissionsProps{
//   	Permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	PermissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   	Principal: &DataLakePrincipalProperty{
//   		DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	Resource: &ResourceProperty{
//   		Catalog: catalog,
//   		Database: &DatabaseResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			Name: jsii.String("name"),
//   		},
//   		DataCellsFilter: &DataCellsFilterResourceProperty{
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//   			TableCatalogId: jsii.String("tableCatalogId"),
//   			TableName: jsii.String("tableName"),
//   		},
//   		DataLocation: &DataLocationResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			ResourceArn: jsii.String("resourceArn"),
//   		},
//   		LfTag: &LFTagKeyResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			TagKey: jsii.String("tagKey"),
//   			TagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   		LfTagPolicy: &LFTagPolicyResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			Expression: []interface{}{
//   				&LFTagProperty{
//   					TagKey: jsii.String("tagKey"),
//   					TagValues: []*string{
//   						jsii.String("tagValues"),
//   					},
//   				},
//   			},
//   			ResourceType: jsii.String("resourceType"),
//   		},
//   		Table: &TableResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//
//   			// the properties below are optional
//   			Name: jsii.String("name"),
//   			TableWildcard: tableWildcard,
//   		},
//   		TableWithColumns: &TableWithColumnsResourceProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			ColumnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			ColumnWildcard: &ColumnWildcardProperty{
//   				ExcludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Catalog: jsii.String("catalog"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-principalpermissions.html
//
type CfnPrincipalPermissionsProps struct {
	// The permissions granted or revoked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-principalpermissions.html#cfn-lakeformation-principalpermissions-permissions
	//
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// Indicates the ability to grant permissions (as a subset of permissions granted).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-principalpermissions.html#cfn-lakeformation-principalpermissions-permissionswithgrantoption
	//
	PermissionsWithGrantOption *[]*string `field:"required" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
	// The principal to be granted a permission.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-principalpermissions.html#cfn-lakeformation-principalpermissions-principal
	//
	Principal interface{} `field:"required" json:"principal" yaml:"principal"`
	// The resource to be granted or revoked permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-principalpermissions.html#cfn-lakeformation-principalpermissions-resource
	//
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
	// The identifier for the Data Catalog .
	//
	// By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lakeformation-principalpermissions.html#cfn-lakeformation-principalpermissions-catalog
	//
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
}


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
//   cfnPrincipalPermissionsProps := &cfnPrincipalPermissionsProps{
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	permissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   	principal: &dataLakePrincipalProperty{
//   		dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	resource: &resourceProperty{
//   		catalog: catalog,
//   		database: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		dataCellsFilter: &dataCellsFilterResourceProperty{
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   			tableCatalogId: jsii.String("tableCatalogId"),
//   			tableName: jsii.String("tableName"),
//   		},
//   		dataLocation: &dataLocationResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			resourceArn: jsii.String("resourceArn"),
//   		},
//   		lfTag: &lFTagKeyResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			tagKey: jsii.String("tagKey"),
//   			tagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   		lfTagPolicy: &lFTagPolicyResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			expression: []interface{}{
//   				&lFTagProperty{
//   					tagKey: jsii.String("tagKey"),
//   					tagValues: []*string{
//   						jsii.String("tagValues"),
//   					},
//   				},
//   			},
//   			resourceType: jsii.String("resourceType"),
//   		},
//   		table: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//
//   			// the properties below are optional
//   			name: jsii.String("name"),
//   			tableWildcard: tableWildcard,
//   		},
//   		tableWithColumns: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//
//   			// the properties below are optional
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			columnWildcard: &columnWildcardProperty{
//   				excludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	catalog: jsii.String("catalog"),
//   }
//
type CfnPrincipalPermissionsProps struct {
	// The permissions granted or revoked.
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// Indicates the ability to grant permissions (as a subset of permissions granted).
	PermissionsWithGrantOption *[]*string `field:"required" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
	// The principal to be granted a permission.
	Principal interface{} `field:"required" json:"principal" yaml:"principal"`
	// The resource to be granted or revoked permissions.
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
	// The identifier for the Data Catalog .
	//
	// By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
}


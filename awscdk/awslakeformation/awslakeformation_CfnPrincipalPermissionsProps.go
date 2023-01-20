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
	// `AWS::LakeFormation::PrincipalPermissions.Permissions`.
	Permissions *[]*string `field:"required" json:"permissions" yaml:"permissions"`
	// `AWS::LakeFormation::PrincipalPermissions.PermissionsWithGrantOption`.
	PermissionsWithGrantOption *[]*string `field:"required" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
	// `AWS::LakeFormation::PrincipalPermissions.Principal`.
	Principal interface{} `field:"required" json:"principal" yaml:"principal"`
	// `AWS::LakeFormation::PrincipalPermissions.Resource`.
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
	// `AWS::LakeFormation::PrincipalPermissions.Catalog`.
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
}


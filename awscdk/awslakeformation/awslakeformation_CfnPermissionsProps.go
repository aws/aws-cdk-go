package awslakeformation


// Properties for defining a `CfnPermissions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPermissionsProps := &cfnPermissionsProps{
//   	dataLakePrincipal: &dataLakePrincipalProperty{
//   		dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   	},
//   	resource: &resourceProperty{
//   		databaseResource: &databaseResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			name: jsii.String("name"),
//   		},
//   		dataLocationResource: &dataLocationResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			s3Resource: jsii.String("s3Resource"),
//   		},
//   		tableResource: &tableResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   			tableWildcard: &tableWildcardProperty{
//   			},
//   		},
//   		tableWithColumnsResource: &tableWithColumnsResourceProperty{
//   			catalogId: jsii.String("catalogId"),
//   			columnNames: []*string{
//   				jsii.String("columnNames"),
//   			},
//   			columnWildcard: &columnWildcardProperty{
//   				excludedColumnNames: []*string{
//   					jsii.String("excludedColumnNames"),
//   				},
//   			},
//   			databaseName: jsii.String("databaseName"),
//   			name: jsii.String("name"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	permissions: []*string{
//   		jsii.String("permissions"),
//   	},
//   	permissionsWithGrantOption: []*string{
//   		jsii.String("permissionsWithGrantOption"),
//   	},
//   }
//
type CfnPermissionsProps struct {
	// The AWS Lake Formation principal.
	DataLakePrincipal interface{} `field:"required" json:"dataLakePrincipal" yaml:"dataLakePrincipal"`
	// A structure for the resource.
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
	// The permissions granted or revoked.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// Indicates whether to grant the ability to grant permissions (as a subset of permissions granted).
	PermissionsWithGrantOption *[]*string `field:"optional" json:"permissionsWithGrantOption" yaml:"permissionsWithGrantOption"`
}


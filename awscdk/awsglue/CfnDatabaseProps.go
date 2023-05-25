package awsglue


// Properties for defining a `CfnDatabase`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   cfnDatabaseProps := &CfnDatabaseProps{
//   	CatalogId: jsii.String("catalogId"),
//   	DatabaseInput: &DatabaseInputProperty{
//   		CreateTableDefaultPermissions: []interface{}{
//   			&PrincipalPrivilegesProperty{
//   				Permissions: []*string{
//   					jsii.String("permissions"),
//   				},
//   				Principal: &DataLakePrincipalProperty{
//   					DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   				},
//   			},
//   		},
//   		Description: jsii.String("description"),
//   		FederatedDatabase: &FederatedDatabaseProperty{
//   			ConnectionName: jsii.String("connectionName"),
//   			Identifier: jsii.String("identifier"),
//   		},
//   		LocationUri: jsii.String("locationUri"),
//   		Name: jsii.String("name"),
//   		Parameters: parameters,
//   		TargetDatabase: &DatabaseIdentifierProperty{
//   			CatalogId: jsii.String("catalogId"),
//   			DatabaseName: jsii.String("databaseName"),
//   		},
//   	},
//   }
//
type CfnDatabaseProps struct {
	// The AWS account ID for the account in which to create the catalog object.
	//
	// > To specify the account ID, you can use the `Ref` intrinsic function with the `AWS::AccountId` pseudo parameter. For example: `!Ref AWS::AccountId`
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The metadata for the database.
	DatabaseInput interface{} `field:"required" json:"databaseInput" yaml:"databaseInput"`
}


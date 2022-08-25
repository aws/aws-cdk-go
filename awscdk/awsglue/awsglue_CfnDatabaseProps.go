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
//   cfnDatabaseProps := &cfnDatabaseProps{
//   	catalogId: jsii.String("catalogId"),
//   	databaseInput: &databaseInputProperty{
//   		createTableDefaultPermissions: []interface{}{
//   			&principalPrivilegesProperty{
//   				permissions: []*string{
//   					jsii.String("permissions"),
//   				},
//   				principal: &dataLakePrincipalProperty{
//   					dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   				},
//   			},
//   		},
//   		description: jsii.String("description"),
//   		locationUri: jsii.String("locationUri"),
//   		name: jsii.String("name"),
//   		parameters: parameters,
//   		targetDatabase: &databaseIdentifierProperty{
//   			catalogId: jsii.String("catalogId"),
//   			databaseName: jsii.String("databaseName"),
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


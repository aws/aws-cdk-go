package previewawsgluemixins


// Properties for CfnDatabasePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var parameters interface{}
//
//   cfnDatabaseMixinProps := &CfnDatabaseMixinProps{
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
//   			Region: jsii.String("region"),
//   		},
//   	},
//   	DatabaseName: jsii.String("databaseName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html
//
type CfnDatabaseMixinProps struct {
	// The AWS account ID for the account in which to create the catalog object.
	//
	// > To specify the account ID, you can use the `Ref` intrinsic function with the `AWS::AccountId` pseudo parameter. For example: `!Ref AWS::AccountId`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html#cfn-glue-database-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// The metadata for the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html#cfn-glue-database-databaseinput
	//
	DatabaseInput interface{} `field:"optional" json:"databaseInput" yaml:"databaseInput"`
	// The name of the catalog database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-database.html#cfn-glue-database-databasename
	//
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
}


package awsglue


// The structure used to create or update a database.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameters interface{}
//
//   databaseInputProperty := &DatabaseInputProperty{
//   	CreateTableDefaultPermissions: []interface{}{
//   		&PrincipalPrivilegesProperty{
//   			Permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   			Principal: &DataLakePrincipalProperty{
//   				DataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	FederatedDatabase: &FederatedDatabaseProperty{
//   		ConnectionName: jsii.String("connectionName"),
//   		Identifier: jsii.String("identifier"),
//   	},
//   	LocationUri: jsii.String("locationUri"),
//   	Name: jsii.String("name"),
//   	Parameters: parameters,
//   	TargetDatabase: &DatabaseIdentifierProperty{
//   		CatalogId: jsii.String("catalogId"),
//   		DatabaseName: jsii.String("databaseName"),
//   	},
//   }
//
type CfnDatabase_DatabaseInputProperty struct {
	// Creates a set of default permissions on the table for principals.
	//
	// Used by AWS Lake Formation . Not used in the normal course of AWS Glue operations.
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// A description of the database.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `CfnDatabase.DatabaseInputProperty.FederatedDatabase`.
	FederatedDatabase interface{} `field:"optional" json:"federatedDatabase" yaml:"federatedDatabase"`
	// The location of the database (for example, an HDFS path).
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
	// The name of the database.
	//
	// For Hive compatibility, this is folded to lowercase when it is stored.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// These key-value pairs define parameters and properties of the database.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A `DatabaseIdentifier` structure that describes a target database for resource linking.
	TargetDatabase interface{} `field:"optional" json:"targetDatabase" yaml:"targetDatabase"`
}


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
//   databaseInputProperty := &databaseInputProperty{
//   	createTableDefaultPermissions: []interface{}{
//   		&principalPrivilegesProperty{
//   			permissions: []*string{
//   				jsii.String("permissions"),
//   			},
//   			principal: &dataLakePrincipalProperty{
//   				dataLakePrincipalIdentifier: jsii.String("dataLakePrincipalIdentifier"),
//   			},
//   		},
//   	},
//   	description: jsii.String("description"),
//   	locationUri: jsii.String("locationUri"),
//   	name: jsii.String("name"),
//   	parameters: parameters,
//   	targetDatabase: &databaseIdentifierProperty{
//   		catalogId: jsii.String("catalogId"),
//   		databaseName: jsii.String("databaseName"),
//   	},
//   }
//
type CfnDatabase_DatabaseInputProperty struct {
	// Creates a set of default permissions on the table for principals.
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// A description of the database.
	Description *string `field:"optional" json:"description" yaml:"description"`
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


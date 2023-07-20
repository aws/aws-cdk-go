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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html
//
type CfnDatabase_DatabaseInputProperty struct {
	// Creates a set of default permissions on the table for principals.
	//
	// Used by AWS Lake Formation . Not used in the normal course of AWS Glue operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html#cfn-glue-database-databaseinput-createtabledefaultpermissions
	//
	CreateTableDefaultPermissions interface{} `field:"optional" json:"createTableDefaultPermissions" yaml:"createTableDefaultPermissions"`
	// A description of the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html#cfn-glue-database-databaseinput-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A `FederatedDatabase` structure that references an entity outside the AWS Glue Data Catalog .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html#cfn-glue-database-databaseinput-federateddatabase
	//
	FederatedDatabase interface{} `field:"optional" json:"federatedDatabase" yaml:"federatedDatabase"`
	// The location of the database (for example, an HDFS path).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html#cfn-glue-database-databaseinput-locationuri
	//
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
	// The name of the database.
	//
	// For Hive compatibility, this is folded to lowercase when it is stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html#cfn-glue-database-databaseinput-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// These key-value pairs define parameters and properties of the database.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html#cfn-glue-database-databaseinput-parameters
	//
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// A `DatabaseIdentifier` structure that describes a target database for resource linking.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-database-databaseinput.html#cfn-glue-database-databaseinput-targetdatabase
	//
	TargetDatabase interface{} `field:"optional" json:"targetDatabase" yaml:"targetDatabase"`
}


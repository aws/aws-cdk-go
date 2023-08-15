package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// A full specification of a Redshift user that can be used to import it fluently into the CDK application.
//
// Example:
//   databaseName := "databaseName"
//   username := "myuser"
//   tableName := "mytable"
//
//   user := awscdkredshiftalpha.User_FromUserAttributes(this, jsii.String("User"), &UserAttributes{
//   	Username: username,
//   	Password: awscdk.SecretValue_UnsafePlainText(jsii.String("NOT_FOR_PRODUCTION")),
//   	Cluster: cluster,
//   	DatabaseName: databaseName,
//   })
//   table := awscdkredshiftalpha.Table_FromTableAttributes(this, jsii.String("Table"), &TableAttributes{
//   	TableName: tableName,
//   	TableColumns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			DataType: jsii.String("varchar(4)"),
//   		},
//   		&column{
//   			Name: jsii.String("col2"),
//   			DataType: jsii.String("float"),
//   		},
//   	},
//   	Cluster: cluster,
//   	DatabaseName: jsii.String("databaseName"),
//   })
//   table.Grant(user, awscdkredshiftalpha.TableAction_INSERT)
//
// Experimental.
type UserAttributes struct {
	// The cluster containing the database.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the database.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The secret containing credentials to a Redshift user with administrator privileges.
	//
	// Secret JSON schema: `{ username: string; password: string }`.
	// Default: - the admin secret is taken from the cluster.
	//
	// Experimental.
	AdminUser awssecretsmanager.ISecret `field:"optional" json:"adminUser" yaml:"adminUser"`
	// The password of the user.
	//
	// Do not put passwords in CDK code directly.
	// Experimental.
	Password awscdk.SecretValue `field:"required" json:"password" yaml:"password"`
	// The name of the user.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}


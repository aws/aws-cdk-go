// The CDK Construct Library for AWS::Redshift
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
//   user := awscdkredshiftalpha.User.fromUserAttributes(this, jsii.String("User"), &userAttributes{
//   	username: username,
//   	password: awscdk.SecretValue.unsafePlainText(jsii.String("NOT_FOR_PRODUCTION")),
//   	cluster: cluster,
//   	databaseName: databaseName,
//   })
//   table := awscdkredshiftalpha.Table.fromTableAttributes(this, jsii.String("Table"), &tableAttributes{
//   	tableName: tableName,
//   	tableColumns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			dataType: jsii.String("varchar(4)"),
//   		},
//   		&column{
//   			name: jsii.String("col2"),
//   			dataType: jsii.String("float"),
//   		},
//   	},
//   	cluster: cluster,
//   	databaseName: jsii.String("databaseName"),
//   })
//   table.grant(user, awscdkredshiftalpha.TableAction_INSERT)
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


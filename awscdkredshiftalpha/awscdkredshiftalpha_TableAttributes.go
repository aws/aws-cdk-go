// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha


// A full specification of a Redshift table that can be used to import it fluently into the CDK application.
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
type TableAttributes struct {
	// The cluster where the table is located.
	// Experimental.
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// The name of the database where the table is located.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// The columns of the table.
	// Experimental.
	TableColumns *[]*Column `field:"required" json:"tableColumns" yaml:"tableColumns"`
	// Name of the table.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}


package awscdkredshiftalpha


// A full specification of a Redshift table that can be used to import it fluently into the CDK application.
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


package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Properties for configuring a Redshift table.
//
// Example:
//   awscdkredshiftalpha.NewTable(this, jsii.String("Table"), &TableProps{
//   	TableColumns: []column{
//   		&column{
//   			Name: jsii.String("col1"),
//   			DataType: jsii.String("varchar(4)"),
//   			DistKey: jsii.Boolean(true),
//   		},
//   		&column{
//   			Name: jsii.String("col2"),
//   			DataType: jsii.String("float"),
//   		},
//   	},
//   	Cluster: cluster,
//   	DatabaseName: jsii.String("databaseName"),
//   	DistStyle: awscdkredshiftalpha.TableDistStyle_KEY,
//   })
//
// Experimental.
type TableProps struct {
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
	// The columns of the table.
	// Experimental.
	TableColumns *[]*Column `field:"required" json:"tableColumns" yaml:"tableColumns"`
	// The distribution style of the table.
	// Default: TableDistStyle.AUTO
	//
	// Experimental.
	DistStyle TableDistStyle `field:"optional" json:"distStyle" yaml:"distStyle"`
	// The policy to apply when this resource is removed from the application.
	// Default: cdk.RemovalPolicy.Retain
	//
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The sort style of the table.
	// Default: TableSortStyle.AUTO if no sort key is specified, TableSortStyle.COMPOUND if a sort key is specified
	//
	// Experimental.
	SortStyle TableSortStyle `field:"optional" json:"sortStyle" yaml:"sortStyle"`
	// A comment to attach to the table.
	// Default: - no comment.
	//
	// Experimental.
	TableComment *string `field:"optional" json:"tableComment" yaml:"tableComment"`
	// The name of the table.
	// Default: - a name is generated.
	//
	// Experimental.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}


package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Properties for configuring a Redshift table.
//
// Example:
//   awscdk.NewTable(this, jsii.String("Table"), &tableProps{
//   	tableColumns: []column{
//   		&column{
//   			name: jsii.String("col1"),
//   			dataType: jsii.String("varchar(4)"),
//   			distKey: jsii.Boolean(true),
//   		},
//   		&column{
//   			name: jsii.String("col2"),
//   			dataType: jsii.String("float"),
//   		},
//   	},
//   	cluster: cluster,
//   	databaseName: jsii.String("databaseName"),
//   	distStyle: awscdk.TableDistStyle_KEY,
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
	// Experimental.
	AdminUser awssecretsmanager.ISecret `field:"optional" json:"adminUser" yaml:"adminUser"`
	// The columns of the table.
	// Experimental.
	TableColumns *[]*Column `field:"required" json:"tableColumns" yaml:"tableColumns"`
	// The distribution style of the table.
	// Experimental.
	DistStyle TableDistStyle `field:"optional" json:"distStyle" yaml:"distStyle"`
	// The policy to apply when this resource is removed from the application.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The sort style of the table.
	// Experimental.
	SortStyle TableSortStyle `field:"optional" json:"sortStyle" yaml:"sortStyle"`
	// The name of the table.
	// Experimental.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}


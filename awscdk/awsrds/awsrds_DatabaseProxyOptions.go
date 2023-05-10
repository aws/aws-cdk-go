package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awssecretsmanager"
)

// Options for a new DatabaseProxy.
//
// Example:
//   var vpc vpc
//   var securityGroup securityGroup
//   var secrets []secret
//   var dbInstance databaseInstance
//
//
//   proxy := dbInstance.AddProxy(jsii.String("proxy"), &DatabaseProxyOptions{
//   	BorrowTimeout: awscdk.Duration_Seconds(jsii.Number(30)),
//   	MaxConnectionsPercent: jsii.Number(50),
//   	Secrets: Secrets,
//   	Vpc: Vpc,
//   })
//
// Experimental.
type DatabaseProxyOptions struct {
	// The secret that the proxy uses to authenticate to the RDS DB instance or Aurora DB cluster.
	//
	// These secrets are stored within Amazon Secrets Manager.
	// One or more secrets are required.
	// Experimental.
	Secrets *[]awssecretsmanager.ISecret `field:"required" json:"secrets" yaml:"secrets"`
	// The VPC to associate with the new proxy.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// The duration for a proxy to wait for a connection to become available in the connection pool.
	//
	// Only applies when the proxy has opened its maximum number of connections and all connections are busy with client
	// sessions.
	//
	// Value must be between 1 second and 1 hour, or `Duration.seconds(0)` to represent unlimited.
	// Experimental.
	BorrowTimeout awscdk.Duration `field:"optional" json:"borrowTimeout" yaml:"borrowTimeout"`
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
	// An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens;
	// it can't end with a hyphen or contain two consecutive hyphens.
	// Experimental.
	DbProxyName *string `field:"optional" json:"dbProxyName" yaml:"dbProxyName"`
	// Whether the proxy includes detailed information about SQL statements in its logs.
	//
	// This information helps you to debug issues involving SQL behavior or the performance and scalability of the proxy connections.
	// The debug information includes the text of SQL statements that you submit through the proxy.
	// Thus, only enable this setting when needed for debugging, and only when you have security measures in place to safeguard any sensitive
	// information that appears in the logs.
	// Experimental.
	DebugLogging *bool `field:"optional" json:"debugLogging" yaml:"debugLogging"`
	// Whether to require or disallow AWS Identity and Access Management (IAM) authentication for connections to the proxy.
	// Experimental.
	IamAuth *bool `field:"optional" json:"iamAuth" yaml:"iamAuth"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
	//
	// You can set this value higher or lower than the connection timeout limit for the associated database.
	// Experimental.
	IdleClientTimeout awscdk.Duration `field:"optional" json:"idleClientTimeout" yaml:"idleClientTimeout"`
	// One or more SQL statements for the proxy to run when opening each new database connection.
	//
	// Typically used with SET statements to make sure that each connection has identical settings such as time zone
	// and character set.
	// For multiple statements, use semicolons as the separator.
	// You can also include multiple variables in a single SET statement, such as SET x=1, y=2.
	//
	// not currently supported for PostgreSQL.
	// Experimental.
	InitQuery *string `field:"optional" json:"initQuery" yaml:"initQuery"`
	// The maximum size of the connection pool for each target in a target group.
	//
	// For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance or Aurora DB
	// cluster used by the target group.
	//
	// 1-100.
	// Experimental.
	MaxConnectionsPercent *float64 `field:"optional" json:"maxConnectionsPercent" yaml:"maxConnectionsPercent"`
	// Controls how actively the proxy closes idle database connections in the connection pool.
	//
	// A high value enables the proxy to leave a high percentage of idle connections open.
	// A low value causes the proxy to close idle client connections and return the underlying database connections
	// to the connection pool.
	// For Aurora MySQL, it is expressed as a percentage of the max_connections setting for the RDS DB instance
	// or Aurora DB cluster used by the target group.
	//
	// between 0 and MaxConnectionsPercent.
	// Experimental.
	MaxIdleConnectionsPercent *float64 `field:"optional" json:"maxIdleConnectionsPercent" yaml:"maxIdleConnectionsPercent"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	//
	// By enabling this setting, you can enforce encrypted TLS connections to the proxy.
	// Experimental.
	RequireTLS *bool `field:"optional" json:"requireTLS" yaml:"requireTLS"`
	// IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// One or more VPC security groups to associate with the new proxy.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.
	//
	// Including an item in the list exempts that class of SQL operations from the pinning behavior.
	// Experimental.
	SessionPinningFilters *[]SessionPinningFilter `field:"optional" json:"sessionPinningFilters" yaml:"sessionPinningFilters"`
	// The subnets used by the proxy.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}


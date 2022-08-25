package awsrds


// Specifies the settings that control the size and behavior of the connection pool associated with a `DBProxyTargetGroup` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionPoolConfigurationInfoFormatProperty := &connectionPoolConfigurationInfoFormatProperty{
//   	connectionBorrowTimeout: jsii.Number(123),
//   	initQuery: jsii.String("initQuery"),
//   	maxConnectionsPercent: jsii.Number(123),
//   	maxIdleConnectionsPercent: jsii.Number(123),
//   	sessionPinningFilters: []*string{
//   		jsii.String("sessionPinningFilters"),
//   	},
//   }
//
type CfnDBProxyTargetGroup_ConnectionPoolConfigurationInfoFormatProperty struct {
	// The number of seconds for a proxy to wait for a connection to become available in the connection pool.
	//
	// Only applies when the proxy has opened its maximum number of connections and all connections are busy with client sessions.
	//
	// Default: 120
	//
	// Constraints: between 1 and 3600, or 0 representing unlimited.
	ConnectionBorrowTimeout *float64 `field:"optional" json:"connectionBorrowTimeout" yaml:"connectionBorrowTimeout"`
	// One or more SQL statements for the proxy to run when opening each new database connection.
	//
	// Typically used with `SET` statements to make sure that each connection has identical settings such as time zone and character set. For multiple statements, use semicolons as the separator. You can also include multiple variables in a single `SET` statement, such as `SET x=1, y=2` .
	//
	// Default: no initialization query.
	InitQuery *string `field:"optional" json:"initQuery" yaml:"initQuery"`
	// The maximum size of the connection pool for each target in a target group.
	//
	// The value is expressed as a percentage of the `max_connections` setting for the RDS DB instance or Aurora DB cluster used by the target group.
	//
	// Default: 100
	//
	// Constraints: between 1 and 100.
	MaxConnectionsPercent *float64 `field:"optional" json:"maxConnectionsPercent" yaml:"maxConnectionsPercent"`
	// Controls how actively the proxy closes idle database connections in the connection pool.
	//
	// The value is expressed as a percentage of the `max_connections` setting for the RDS DB instance or Aurora DB cluster used by the target group. With a high value, the proxy leaves a high percentage of idle database connections open. A low value causes the proxy to close more idle connections and return them to the database.
	//
	// Default: 50
	//
	// Constraints: between 0 and `MaxConnectionsPercent`.
	MaxIdleConnectionsPercent *float64 `field:"optional" json:"maxIdleConnectionsPercent" yaml:"maxIdleConnectionsPercent"`
	// Each item in the list represents a class of SQL operations that normally cause all later statements in a session using a proxy to be pinned to the same underlying database connection.
	//
	// Including an item in the list exempts that class of SQL operations from the pinning behavior.
	//
	// Default: no session pinning filters.
	SessionPinningFilters *[]*string `field:"optional" json:"sessionPinningFilters" yaml:"sessionPinningFilters"`
}


package awsrds


// Properties for PostgreSQL instance engines.
//
// Used in `DatabaseInstanceEngine.postgres`.
//
// Example:
//   var vpc Vpc
//
//   instance := rds.NewDatabaseInstance(this, jsii.String("Database"), &DatabaseInstanceProps{
//   	Engine: rds.DatabaseInstanceEngine_Postgres(&PostgresInstanceEngineProps{
//   		Version: rds.PostgresEngineVersion_VER_17_7(),
//   	}),
//   	Vpc: Vpc,
//   	IamAuthentication: jsii.Boolean(true),
//   })
//
//   proxy := rds.NewDatabaseProxy(this, jsii.String("Proxy"), &DatabaseProxyProps{
//   	ProxyTarget: rds.ProxyTarget_FromInstance(instance),
//   	Vpc: Vpc,
//   	DefaultAuthScheme: rds.DefaultAuthScheme_IAM_AUTH,
//   })
//
//   // Grant IAM permissions for database connection
//   role := iam.NewRole(this, jsii.String("DBRole"), &RoleProps{
//   	AssumedBy: iam.NewAccountPrincipal(this.Account),
//   })
//   proxy.GrantConnect(role, jsii.String("database-user"))
//
type PostgresInstanceEngineProps struct {
	// The exact version of the engine to use.
	Version PostgresEngineVersion `field:"required" json:"version" yaml:"version"`
}


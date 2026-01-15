package awsrds


// The default authentication scheme that the proxy uses for client connections to the proxy and connections from the proxy to the underlying database.
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
type DefaultAuthScheme string

const (
	// IAM authentication.
	DefaultAuthScheme_IAM_AUTH DefaultAuthScheme = "IAM_AUTH"
	// No default authentication.
	DefaultAuthScheme_NONE DefaultAuthScheme = "NONE"
)


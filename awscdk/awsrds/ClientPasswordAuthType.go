package awsrds


// Client password authentication type used by a proxy to log in as a specific database user.
//
// Example:
//   var vpc vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_03_0(),
//   	}),
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer")),
//   	Vpc: Vpc,
//   })
//
//   proxy := rds.NewDatabaseProxy(this, jsii.String("Proxy"), &DatabaseProxyProps{
//   	ProxyTarget: rds.ProxyTarget_FromCluster(cluster),
//   	Secrets: []iSecret{
//   		cluster.Secret,
//   	},
//   	Vpc: Vpc,
//   	ClientPasswordAuthType: rds.ClientPasswordAuthType_MYSQL_NATIVE_PASSWORD,
//   })
//
type ClientPasswordAuthType string

const (
	// MySQL Native Password client authentication type.
	ClientPasswordAuthType_MYSQL_NATIVE_PASSWORD ClientPasswordAuthType = "MYSQL_NATIVE_PASSWORD"
	// SCRAM SHA 256 client authentication type.
	ClientPasswordAuthType_POSTGRES_SCRAM_SHA_256 ClientPasswordAuthType = "POSTGRES_SCRAM_SHA_256"
	// PostgreSQL MD5 client authentication type.
	ClientPasswordAuthType_POSTGRES_MD5 ClientPasswordAuthType = "POSTGRES_MD5"
	// SQL Server Authentication client authentication type.
	ClientPasswordAuthType_SQL_SERVER_AUTHENTICATION ClientPasswordAuthType = "SQL_SERVER_AUTHENTICATION"
)


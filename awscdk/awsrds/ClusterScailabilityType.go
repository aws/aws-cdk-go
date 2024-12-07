package awsrds


// The scalability mode of the Aurora DB cluster.
//
// Example:
//   var vpc iVpc
//
//
//   rds.NewDatabaseCluster(this, jsii.String("LimitlessDatabaseCluster"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraPostgres(&AuroraPostgresClusterEngineProps{
//   		Version: rds.AuroraPostgresEngineVersion_VER_16_4_LIMITLESS(),
//   	}),
//   	Vpc: Vpc,
//   	ClusterScailabilityType: rds.ClusterScailabilityType_LIMITLESS,
//   	// Requires enabling Performance Insights
//   	EnablePerformanceInsights: jsii.Boolean(true),
//   	PerformanceInsightRetention: rds.PerformanceInsightRetention_MONTHS_1,
//   	// Requires enabling Enhanced Monitoring at the cluster level
//   	MonitoringInterval: awscdk.Duration_Minutes(jsii.Number(1)),
//   	EnableClusterLevelEnhancedMonitoring: jsii.Boolean(true),
//   	// Requires I/O optimized storage type
//   	StorageType: rds.DBClusterStorageType_AURORA_IOPT1,
//   	// Requires exporting the PostgreSQL log to Amazon CloudWatch Logs.
//   	CloudwatchLogsExports: []*string{
//   		jsii.String("postgresql"),
//   	},
//   })
//
type ClusterScailabilityType string

const (
	// The cluster uses normal DB instance creation.
	ClusterScailabilityType_STANDARD ClusterScailabilityType = "STANDARD"
	// The cluster operates as an Aurora Limitless Database, allowing you to create a DB shard group for horizontal scaling (sharding) capabilities.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/limitless.html
	//
	ClusterScailabilityType_LIMITLESS ClusterScailabilityType = "LIMITLESS"
)


package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Properties for a new database cluster.
//
// Example:
//   var vpc Vpc
//
//   cluster := rds.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	Engine: rds.DatabaseClusterEngine_AuroraMysql(&AuroraMysqlClusterEngineProps{
//   		Version: rds.AuroraMysqlEngineVersion_VER_3_01_0(),
//   	}),
//   	Credentials: rds.Credentials_FromGeneratedSecret(jsii.String("clusteradmin")),
//   	 // Optional - will default to 'admin' username and generated password
//   	Writer: rds.ClusterInstance_Provisioned(jsii.String("writer"), &ProvisionedClusterInstanceProps{
//   		PubliclyAccessible: jsii.Boolean(false),
//   	}),
//   	Readers: []IClusterInstance{
//   		rds.ClusterInstance_*Provisioned(jsii.String("reader1"), &ProvisionedClusterInstanceProps{
//   			PromotionTier: jsii.Number(1),
//   		}),
//   		rds.ClusterInstance_ServerlessV2(jsii.String("reader2")),
//   	},
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PRIVATE_WITH_EGRESS,
//   	},
//   	Vpc: Vpc,
//   })
//
type DatabaseClusterProps struct {
	// What kind of database to start.
	Engine IClusterEngine `field:"required" json:"engine" yaml:"engine"`
	// Specifies whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window.
	// Default: true.
	//
	AutoMinorVersionUpgrade *bool `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The number of seconds to set a cluster's target backtrack window to.
	//
	// This feature is only supported by the Aurora MySQL database engine and
	// cannot be enabled on existing clusters.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Managing.Backtrack.html
	//
	// Default: 0 seconds (no backtrack).
	//
	BacktrackWindow awscdk.Duration `field:"optional" json:"backtrackWindow" yaml:"backtrackWindow"`
	// Backup settings.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow
	//
	// Default: - Backup retention period for automated backups is 1 day.
	// Backup preferred window is set to a 30-minute window selected at random from an
	// 8-hour block of time for each AWS Region, occurring on a random day of the week.
	//
	Backup *BackupProps `field:"optional" json:"backup" yaml:"backup"`
	// The list of log types that need to be enabled for exporting to CloudWatch Logs.
	// Default: - no log exports.
	//
	CloudwatchLogsExports *[]*string `field:"optional" json:"cloudwatchLogsExports" yaml:"cloudwatchLogsExports"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Default: - logs never expire.
	//
	CloudwatchLogsRetention awslogs.RetentionDays `field:"optional" json:"cloudwatchLogsRetention" yaml:"cloudwatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Default: - a new role is created.
	//
	CloudwatchLogsRetentionRole awsiam.IRole `field:"optional" json:"cloudwatchLogsRetentionRole" yaml:"cloudwatchLogsRetentionRole"`
	// An optional identifier for the cluster.
	// Default: - A name is automatically generated.
	//
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// [Misspelled] Specifies the scalability mode of the Aurora DB cluster.
	//
	// Set LIMITLESS if you want to use a limitless database; otherwise, set it to STANDARD.
	// Default: ClusterScailabilityType.STANDARD
	//
	// Deprecated: Use clusterScalabilityType instead. This will be removed in the next major version.
	ClusterScailabilityType ClusterScailabilityType `field:"optional" json:"clusterScailabilityType" yaml:"clusterScailabilityType"`
	// Specifies the scalability mode of the Aurora DB cluster.
	//
	// Set LIMITLESS if you want to use a limitless database; otherwise, set it to STANDARD.
	// Default: ClusterScalabilityType.STANDARD
	//
	ClusterScalabilityType ClusterScalabilityType `field:"optional" json:"clusterScalabilityType" yaml:"clusterScalabilityType"`
	// Whether to copy tags to the snapshot when a snapshot is created.
	// Default: - true.
	//
	CopyTagsToSnapshot *bool `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Credentials for the administrative user.
	// Default: - A username of 'admin' (or 'postgres' for PostgreSQL) and SecretsManager-generated password.
	//
	Credentials Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// The database insights mode.
	// Default: - DatabaseInsightsMode.STANDARD when performance insights are enabled and Amazon Aurora engine is used, otherwise not set.
	//
	DatabaseInsightsMode DatabaseInsightsMode `field:"optional" json:"databaseInsightsMode" yaml:"databaseInsightsMode"`
	// Name of a database which is automatically created inside the cluster.
	// Default: - Database is not created in cluster.
	//
	DefaultDatabaseName *string `field:"optional" json:"defaultDatabaseName" yaml:"defaultDatabaseName"`
	// Specifies whether to remove automated backups immediately after the DB cluster is deleted.
	// Default: undefined - AWS RDS default is to remove automated backups immediately after the DB cluster is deleted, unless the AWS Backup policy specifies a point-in-time restore rule.
	//
	DeleteAutomatedBackups *bool `field:"optional" json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Indicates whether the DB cluster should have deletion protection enabled.
	// Default: - true if `removalPolicy` is RETAIN, `undefined` otherwise, which will not enable deletion protection.
	// To disable deletion protection after it has been enabled, you must explicitly set this value to `false`.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Directory ID for associating the DB cluster with a specific Active Directory.
	//
	// Necessary for enabling Kerberos authentication. If specified, the DB cluster joins the given Active Directory, enabling Kerberos authentication.
	// If not specified, the DB cluster will not be associated with any Active Directory, and Kerberos authentication will not be enabled.
	// Default: - DB cluster is not associated with an Active Directory; Kerberos authentication is not enabled.
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The IAM role to be used when making API calls to the Directory Service.
	//
	// The role needs the AWS-managed policy
	// `AmazonRDSDirectoryServiceAccess` or equivalent.
	// Default: - If `DatabaseClusterBaseProps.domain` is specified, a role with the `AmazonRDSDirectoryServiceAccess` policy is automatically created.
	//
	DomainRole awsiam.IRole `field:"optional" json:"domainRole" yaml:"domainRole"`
	// Whether to enable enhanced monitoring at the cluster level.
	//
	// If set to true, `monitoringInterval` and `monitoringRole` are applied to not the instances, but the cluster.
	// `monitoringInterval` is required to be set if `enableClusterLevelEnhancedMonitoring` is set to true.
	// Default: - When the `monitoringInterval` is set, enhanced monitoring is enabled for each instance.
	//
	EnableClusterLevelEnhancedMonitoring *bool `field:"optional" json:"enableClusterLevelEnhancedMonitoring" yaml:"enableClusterLevelEnhancedMonitoring"`
	// Whether to enable the Data API for the cluster.
	// Default: - false.
	//
	EnableDataApi *bool `field:"optional" json:"enableDataApi" yaml:"enableDataApi"`
	// Whether read replicas can forward write operations to the writer DB instance in the DB cluster.
	//
	// This setting can only be enabled for Aurora MySQL 3.04 or higher, and for Aurora PostgreSQL 16.4
	// or higher (for version 16), 15.8 or higher (for version 15), and 14.13 or higher (for version 14).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-postgresql-write-forwarding.html
	//
	// Default: false.
	//
	EnableLocalWriteForwarding *bool `field:"optional" json:"enableLocalWriteForwarding" yaml:"enableLocalWriteForwarding"`
	// Whether to enable Performance Insights for the DB cluster.
	// Default: - false, unless `performanceInsightRetention` or `performanceInsightEncryptionKey` is set,
	// or `databaseInsightsMode` is set to `DatabaseInsightsMode.ADVANCED`.
	//
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The life cycle type for this DB cluster.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/extended-support.html
	//
	// Default: undefined - AWS RDS default setting is `EngineLifecycleSupport.OPEN_SOURCE_RDS_EXTENDED_SUPPORT`
	//
	EngineLifecycleSupport EngineLifecycleSupport `field:"optional" json:"engineLifecycleSupport" yaml:"engineLifecycleSupport"`
	// Whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts.
	// Default: false.
	//
	IamAuthentication *bool `field:"optional" json:"iamAuthentication" yaml:"iamAuthentication"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	// Default: - clusterIdentifier is used with the word "Instance" appended.
	// If clusterIdentifier is not provided, the identifier is automatically generated.
	//
	InstanceIdentifierBase *string `field:"optional" json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// Settings for the individual instances that are launched.
	// Deprecated: - use writer and readers instead.
	InstanceProps *InstanceProps `field:"optional" json:"instanceProps" yaml:"instanceProps"`
	// How many replicas/instances to create.
	//
	// Has to be at least 1.
	// Default: 2.
	//
	// Deprecated: - use writer and readers instead.
	Instances *float64 `field:"optional" json:"instances" yaml:"instances"`
	// The ordering of updates for instances.
	// Default: InstanceUpdateBehaviour.BULK
	//
	InstanceUpdateBehaviour InstanceUpdateBehaviour `field:"optional" json:"instanceUpdateBehaviour" yaml:"instanceUpdateBehaviour"`
	// The interval between points when Amazon RDS collects enhanced monitoring metrics.
	//
	// If you enable `enableClusterLevelEnhancedMonitoring`, this property is applied to the cluster,
	// otherwise it is applied to the instances.
	// Default: - no enhanced monitoring.
	//
	MonitoringInterval awscdk.Duration `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// Role that will be used to manage DB monitoring.
	//
	// If you enable `enableClusterLevelEnhancedMonitoring`, this property is applied to the cluster,
	// otherwise it is applied to the instances.
	// Default: - A role is automatically created for you.
	//
	MonitoringRole awsiam.IRole `field:"optional" json:"monitoringRole" yaml:"monitoringRole"`
	// The network type of the DB instance.
	// Default: - IPV4.
	//
	NetworkType NetworkType `field:"optional" json:"networkType" yaml:"networkType"`
	// Additional parameters to pass to the database engine.
	// Default: - No parameter group.
	//
	ParameterGroup IParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The parameters in the DBClusterParameterGroup to create automatically.
	//
	// You can only specify parameterGroup or parameters but not both.
	// You need to use a versioned engine to auto-generate a DBClusterParameterGroup.
	// Default: - None.
	//
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The AWS KMS key for encryption of Performance Insights data.
	// Default: - default master key.
	//
	PerformanceInsightEncryptionKey awskms.IKey `field:"optional" json:"performanceInsightEncryptionKey" yaml:"performanceInsightEncryptionKey"`
	// The amount of time, in days, to retain Performance Insights data.
	//
	// If you set `databaseInsightsMode` to `DatabaseInsightsMode.ADVANCED`, you must set this property to `PerformanceInsightRetention.MONTHS_15`.
	// Default: - 7.
	//
	PerformanceInsightRetention PerformanceInsightRetention `field:"optional" json:"performanceInsightRetention" yaml:"performanceInsightRetention"`
	// What port to listen on.
	// Default: - The default for the engine is used.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A preferred maintenance window day/time range. Should be specified as a range ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC).
	//
	// Example: 'Sun:23:45-Mon:00:15'.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance
	//
	// Default: - 30-minute window selected at random from an 8-hour block of time for
	// each AWS Region, occurring on a random day of the week.
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// A list of instances to create as cluster reader instances.
	// Default: - no readers are created. The cluster will have a single writer/reader
	//
	Readers *[]IClusterInstance `field:"optional" json:"readers" yaml:"readers"`
	// The removal policy to apply when the cluster and its instances are removed from the stack or replaced during an update.
	// Default: - RemovalPolicy.SNAPSHOT (remove the cluster and instances, but retain a snapshot of the data)
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.
	//
	// Cannot be used with credentials.
	// Default: - This DB Cluster is not a read replica.
	//
	ReplicationSourceIdentifier *string `field:"optional" json:"replicationSourceIdentifier" yaml:"replicationSourceIdentifier"`
	// S3 buckets that you want to load data into. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportRole` is used.
	//
	// For MySQL:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	// Default: - None.
	//
	S3ExportBuckets *[]awss3.IBucket `field:"optional" json:"s3ExportBuckets" yaml:"s3ExportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 export.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ExportBuckets` is used.
	// To use this property with Aurora PostgreSQL, it must be configured with the S3 export feature enabled when creating the DatabaseClusterEngine
	// For MySQL:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/postgresql-s3-export.html
	//
	// Default: - New role is created if `s3ExportBuckets` is set, no role is defined otherwise.
	//
	S3ExportRole awsiam.IRole `field:"optional" json:"s3ExportRole" yaml:"s3ExportRole"`
	// S3 buckets that you want to load data from. This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportRole` is used.
	//
	// For MySQL:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	// Default: - None.
	//
	S3ImportBuckets *[]awss3.IBucket `field:"optional" json:"s3ImportBuckets" yaml:"s3ImportBuckets"`
	// Role that will be associated with this DB cluster to enable S3 import.
	//
	// This feature is only supported by the Aurora database engine.
	//
	// This property must not be used if `s3ImportBuckets` is used.
	// To use this property with Aurora PostgreSQL, it must be configured with the S3 import feature enabled when creating the DatabaseClusterEngine
	// For MySQL:.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Migrating.html
	//
	// Default: - New role is created if `s3ImportBuckets` is set, no role is defined otherwise.
	//
	S3ImportRole awsiam.IRole `field:"optional" json:"s3ImportRole" yaml:"s3ImportRole"`
	// Security group.
	// Default: - a new security group is created.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Specifies the duration an Aurora Serverless v2 DB instance must be idle before Aurora attempts to automatically pause it.
	//
	// The duration must be between 300 seconds (5 minutes) and 86,400 seconds (24 hours).
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2-auto-pause.html
	//
	// Default: - The default is 300 seconds (5 minutes).
	//
	ServerlessV2AutoPauseDuration awscdk.Duration `field:"optional" json:"serverlessV2AutoPauseDuration" yaml:"serverlessV2AutoPauseDuration"`
	// The maximum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster.
	//
	// You can specify ACU values in half-step increments, such as 40, 40.5, 41, and so on.
	// The largest value that you can use is 256.
	//
	// The maximum capacity must be higher than 0.5 ACUs.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.setting-capacity.html#aurora-serverless-v2.max_capacity_considerations
	//
	// Default: 2.
	//
	ServerlessV2MaxCapacity *float64 `field:"optional" json:"serverlessV2MaxCapacity" yaml:"serverlessV2MaxCapacity"`
	// The minimum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster.
	//
	// You can specify ACU values in half-step increments, such as 8, 8.5, 9, and so on.
	// The smallest value that you can use is 0.
	//
	// For Aurora versions that support the Aurora Serverless v2 auto-pause feature, the smallest value that you can use is 0.
	// For versions that don't support Aurora Serverless v2 auto-pause, the smallest value that you can use is 0.5.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.setting-capacity.html#aurora-serverless-v2.min_capacity_considerations
	//
	// Default: 0.5
	//
	ServerlessV2MinCapacity *float64 `field:"optional" json:"serverlessV2MinCapacity" yaml:"serverlessV2MinCapacity"`
	// Whether to enable storage encryption.
	// Default: - true if storageEncryptionKey is provided, false otherwise.
	//
	StorageEncrypted *bool `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// The KMS key for storage encryption.
	//
	// If specified, `storageEncrypted` will be set to `true`.
	// Default: - if storageEncrypted is true then the default master key, no key otherwise.
	//
	StorageEncryptionKey interfacesawskms.IKeyRef `field:"optional" json:"storageEncryptionKey" yaml:"storageEncryptionKey"`
	// The storage type to be associated with the DB cluster.
	// Default: - DBClusterStorageType.AURORA
	//
	StorageType DBClusterStorageType `field:"optional" json:"storageType" yaml:"storageType"`
	// Existing subnet group for the cluster.
	// Default: - a new subnet group will be created.
	//
	SubnetGroup ISubnetGroup `field:"optional" json:"subnetGroup" yaml:"subnetGroup"`
	// What subnets to run the RDS instances in.
	//
	// Must be at least 2 subnets in two different AZs.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where to place the instances within the VPC.
	// Default: - the Vpc default strategy if not specified.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
	// The instance to use for the cluster writer.
	// Default: - required if instanceProps is not provided.
	//
	Writer IClusterInstance `field:"optional" json:"writer" yaml:"writer"`
}


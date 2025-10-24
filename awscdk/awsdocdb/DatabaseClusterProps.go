package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
)

// Properties for a new database cluster.
//
// Example:
//   var vpc Vpc
//
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	MasterUser: &Login{
//   		Username: jsii.String("myuser"),
//   	},
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   	Vpc: Vpc,
//   	DeletionProtection: jsii.Boolean(true),
//   })
//
type DatabaseClusterProps struct {
	// Username and password for the administrative user.
	MasterUser *Login `field:"required" json:"masterUser" yaml:"masterUser"`
	// What subnets to run the DocumentDB instances in.
	//
	// Must be at least 2 subnets in two different AZs.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Backup settings.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/backup-restore.db-cluster-snapshots.html#backup-restore.backup-window
	//
	// Default: - Backup retention period for automated backups is 1 day.
	// Backup preferred window is set to a 30-minute window selected at random from an
	// 8-hour block of time for each AWS Region, occurring on a random day of the week.
	//
	Backup *BackupProps `field:"optional" json:"backup" yaml:"backup"`
	// The identifier of the CA certificate used for the instances.
	//
	// Specifying or updating this property triggers a reboot.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html
	//
	// Default: - DocumentDB will choose a certificate authority.
	//
	CaCertificate awsrds.CaCertificate `field:"optional" json:"caCertificate" yaml:"caCertificate"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Default: - logs never expire.
	//
	CloudWatchLogsRetention awslogs.RetentionDays `field:"optional" json:"cloudWatchLogsRetention" yaml:"cloudWatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Default: - a new role is created.
	//
	CloudWatchLogsRetentionRole awsiam.IRole `field:"optional" json:"cloudWatchLogsRetentionRole" yaml:"cloudWatchLogsRetentionRole"`
	// Whether to copy tags to the snapshot when a snapshot is created.
	// Default: - false.
	//
	CopyTagsToSnapshot *bool `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// An optional identifier for the cluster.
	// Default: - A name is automatically generated.
	//
	DbClusterName *string `field:"optional" json:"dbClusterName" yaml:"dbClusterName"`
	// Specifies whether this cluster can be deleted.
	//
	// If deletionProtection is
	// enabled, the cluster cannot be deleted unless it is modified and
	// deletionProtection is disabled. deletionProtection protects clusters from
	// being accidentally deleted.
	// Default: - false.
	//
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// A value that indicates whether to enable Performance Insights for the instances in the DB Cluster.
	// Default: - false.
	//
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// What version of the database to start.
	// Default: -  the latest major version.
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Whether the audit logs should be exported to CloudWatch.
	//
	// Note that you also have to configure the audit log export in the Cluster's Parameter Group.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html#event-auditing-enabling-auditing
	//
	// Default: false.
	//
	ExportAuditLogsToCloudWatch *bool `field:"optional" json:"exportAuditLogsToCloudWatch" yaml:"exportAuditLogsToCloudWatch"`
	// Whether the profiler logs should be exported to CloudWatch.
	//
	// Note that you also have to configure the profiler log export in the Cluster's Parameter Group.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html#profiling.enable-profiling
	//
	// Default: false.
	//
	ExportProfilerLogsToCloudWatch *bool `field:"optional" json:"exportProfilerLogsToCloudWatch" yaml:"exportProfilerLogsToCloudWatch"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	// Only applicable for provisioned clusters.
	// Default: - `dbClusterName` is used with the word "Instance" appended. If `dbClusterName` is not provided, the
	// identifier is automatically generated.
	//
	InstanceIdentifierBase *string `field:"optional" json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// The removal policy to apply to the cluster's instances.
	//
	// Cannot be set to `SNAPSHOT`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
	//
	// Default: - `RemovalPolicy.DESTROY` when `removalPolicy` is set to `SNAPSHOT`, `removalPolicy` otherwise.
	//
	InstanceRemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"instanceRemovalPolicy" yaml:"instanceRemovalPolicy"`
	// Number of DocDB compute instances.
	// Default: 1.
	//
	Instances *float64 `field:"optional" json:"instances" yaml:"instances"`
	// What type of instance to start for the replicas.
	//
	// Required for provisioned clusters, not applicable for serverless clusters.
	// Default: None.
	//
	InstanceType awsec2.InstanceType `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The KMS key for storage encryption.
	// Default: - default master key.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The DB parameter group to associate with the instance.
	// Default: no parameter group.
	//
	ParameterGroup IClusterParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The port the DocumentDB cluster will listen on.
	// Default: DatabaseCluster.DEFAULT_PORT
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A weekly time range in which maintenance should preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: 'tue:04:17-tue:04:47'.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-maintain.html#maintenance-window
	//
	// Default: - 30-minute window selected at random from an 8-hour block of time for
	// each AWS Region, occurring on a random day of the week.
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed or replaced during a stack update, or when the stack is deleted.
	//
	// This
	// removal policy also applies to the implicit security group created for the
	// cluster if one is not supplied as a parameter.
	//
	// When set to `SNAPSHOT`, the removal policy for the instances and the security group
	// will default to `DESTROY` as those resources do not support the policy.
	//
	// Use the `instanceRemovalPolicy` and `securityGroupRemovalPolicy` to change the behavior.
	// Default: - Retain cluster.
	//
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Security group.
	// Default: a new security group is created.
	//
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// The removal policy to apply to the cluster's security group.
	//
	// Cannot be set to `SNAPSHOT`.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
	//
	// Default: - `RemovalPolicy.DESTROY` when `removalPolicy` is set to `SNAPSHOT`, `removalPolicy` otherwise.
	//
	SecurityGroupRemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"securityGroupRemovalPolicy" yaml:"securityGroupRemovalPolicy"`
	// ServerlessV2 scaling configuration.
	//
	// When specified, the cluster will be created as a serverless cluster.
	// Default: None.
	//
	ServerlessV2ScalingConfiguration *ServerlessV2ScalingConfiguration `field:"optional" json:"serverlessV2ScalingConfiguration" yaml:"serverlessV2ScalingConfiguration"`
	// Whether to enable storage encryption.
	// Default: true.
	//
	StorageEncrypted *bool `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// The storage type of the DocDB cluster.
	//
	// I/O-optimized storage is supported starting with engine version 5.0.0.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/release-notes.html#release-notes.11-21-2023
	//
	// Default: StorageType.STANDARD
	//
	StorageType StorageType `field:"optional" json:"storageType" yaml:"storageType"`
	// Where to place the instances within the VPC.
	// Default: private subnets.
	//
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}


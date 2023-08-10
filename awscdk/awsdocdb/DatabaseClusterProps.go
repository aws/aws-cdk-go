package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Properties for a new database cluster.
//
// Example:
//   var vpc vpc
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
	// What type of instance to start for the replicas.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Username and password for the administrative user.
	MasterUser *Login `field:"required" json:"masterUser" yaml:"masterUser"`
	// What subnets to run the DocumentDB instances in.
	//
	// Must be at least 2 subnets in two different AZs.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Backup settings.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/backup-restore.db-cluster-snapshots.html#backup-restore.backup-window
	//
	Backup *BackupProps `field:"optional" json:"backup" yaml:"backup"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	CloudWatchLogsRetention awslogs.RetentionDays `field:"optional" json:"cloudWatchLogsRetention" yaml:"cloudWatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	CloudWatchLogsRetentionRole awsiam.IRole `field:"optional" json:"cloudWatchLogsRetentionRole" yaml:"cloudWatchLogsRetentionRole"`
	// An optional identifier for the cluster.
	DbClusterName *string `field:"optional" json:"dbClusterName" yaml:"dbClusterName"`
	// Specifies whether this cluster can be deleted.
	//
	// If deletionProtection is
	// enabled, the cluster cannot be deleted unless it is modified and
	// deletionProtection is disabled. deletionProtection protects clusters from
	// being accidentally deleted.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// A value that indicates whether to enable Performance Insights for the instances in the DB Cluster.
	EnablePerformanceInsights *bool `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// What version of the database to start.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Whether the audit logs should be exported to CloudWatch.
	//
	// Note that you also have to configure the audit log export in the Cluster's Parameter Group.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html#event-auditing-enabling-auditing
	//
	ExportAuditLogsToCloudWatch *bool `field:"optional" json:"exportAuditLogsToCloudWatch" yaml:"exportAuditLogsToCloudWatch"`
	// Whether the profiler logs should be exported to CloudWatch.
	//
	// Note that you also have to configure the profiler log export in the Cluster's Parameter Group.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html#profiling.enable-profiling
	//
	ExportProfilerLogsToCloudWatch *bool `field:"optional" json:"exportProfilerLogsToCloudWatch" yaml:"exportProfilerLogsToCloudWatch"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	InstanceIdentifierBase *string `field:"optional" json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// Number of DocDB compute instances.
	Instances *float64 `field:"optional" json:"instances" yaml:"instances"`
	// The KMS key for storage encryption.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The DB parameter group to associate with the instance.
	ParameterGroup IClusterParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The port the DocumentDB cluster will listen on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A weekly time range in which maintenance should preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: 'tue:04:17-tue:04:47'.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-maintain.html#maintenance-window
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed or replaced during a stack update, or when the stack is deleted.
	//
	// This
	// removal policy also applies to the implicit security group created for the
	// cluster if one is not supplied as a parameter.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Security group.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Whether to enable storage encryption.
	StorageEncrypted *bool `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Where to place the instances within the VPC.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}


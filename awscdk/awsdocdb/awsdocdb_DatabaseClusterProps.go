package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
)

// Properties for a new database cluster.
//
// Example:
//   var vpc vpc
//
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &databaseClusterProps{
//   	masterUser: &login{
//   		username: jsii.String("myuser"),
//   		 // NOTE: 'admin' is reserved by DocumentDB
//   		excludeCharacters: jsii.String("\"@/:"),
//   		 // optional, defaults to the set "\"@/" and is also used for eventually created rotations
//   		secretName: jsii.String("/myapp/mydocdb/masteruser"),
//   	},
//   	instanceType: ec2.instanceType.of(ec2.instanceClass_R5, ec2.instanceSize_LARGE),
//   	vpcSubnets: &subnetSelection{
//   		subnetType: ec2.subnetType_PUBLIC,
//   	},
//   	vpc: vpc,
//   })
//
// Experimental.
type DatabaseClusterProps struct {
	// What type of instance to start for the replicas.
	// Experimental.
	InstanceType awsec2.InstanceType `field:"required" json:"instanceType" yaml:"instanceType"`
	// Username and password for the administrative user.
	// Experimental.
	MasterUser *Login `field:"required" json:"masterUser" yaml:"masterUser"`
	// What subnets to run the DocumentDB instances in.
	//
	// Must be at least 2 subnets in two different AZs.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Backup settings.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/backup-restore.db-cluster-snapshots.html#backup-restore.backup-window
	//
	// Experimental.
	Backup *BackupProps `field:"optional" json:"backup" yaml:"backup"`
	// The number of days log events are kept in CloudWatch Logs.
	//
	// When updating
	// this property, unsetting it doesn't remove the log retention policy. To
	// remove the retention policy, set the value to `Infinity`.
	// Experimental.
	CloudWatchLogsRetention awslogs.RetentionDays `field:"optional" json:"cloudWatchLogsRetention" yaml:"cloudWatchLogsRetention"`
	// The IAM role for the Lambda function associated with the custom resource that sets the retention policy.
	// Experimental.
	CloudWatchLogsRetentionRole awsiam.IRole `field:"optional" json:"cloudWatchLogsRetentionRole" yaml:"cloudWatchLogsRetentionRole"`
	// An optional identifier for the cluster.
	// Experimental.
	DbClusterName *string `field:"optional" json:"dbClusterName" yaml:"dbClusterName"`
	// Specifies whether this cluster can be deleted.
	//
	// If deletionProtection is
	// enabled, the cluster cannot be deleted unless it is modified and
	// deletionProtection is disabled. deletionProtection protects clusters from
	// being accidentally deleted.
	// Experimental.
	DeletionProtection *bool `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// What version of the database to start.
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Whether the audit logs should be exported to CloudWatch.
	//
	// Note that you also have to configure the audit log export in the Cluster's Parameter Group.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html#event-auditing-enabling-auditing
	//
	// Experimental.
	ExportAuditLogsToCloudWatch *bool `field:"optional" json:"exportAuditLogsToCloudWatch" yaml:"exportAuditLogsToCloudWatch"`
	// Whether the profiler logs should be exported to CloudWatch.
	//
	// Note that you also have to configure the profiler log export in the Cluster's Parameter Group.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html#profiling.enable-profiling
	//
	// Experimental.
	ExportProfilerLogsToCloudWatch *bool `field:"optional" json:"exportProfilerLogsToCloudWatch" yaml:"exportProfilerLogsToCloudWatch"`
	// Base identifier for instances.
	//
	// Every replica is named by appending the replica number to this string, 1-based.
	// Experimental.
	InstanceIdentifierBase *string `field:"optional" json:"instanceIdentifierBase" yaml:"instanceIdentifierBase"`
	// Number of DocDB compute instances.
	// Experimental.
	Instances *float64 `field:"optional" json:"instances" yaml:"instances"`
	// The KMS key for storage encryption.
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The DB parameter group to associate with the instance.
	// Experimental.
	ParameterGroup IClusterParameterGroup `field:"optional" json:"parameterGroup" yaml:"parameterGroup"`
	// The port the DocumentDB cluster will listen on.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A weekly time range in which maintenance should preferably execute.
	//
	// Must be at least 30 minutes long.
	//
	// Example: 'tue:04:17-tue:04:47'.
	// See: https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-maintain.html#maintenance-window
	//
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The removal policy to apply when the cluster and its instances are removed or replaced during a stack update, or when the stack is deleted.
	//
	// This
	// removal policy also applies to the implicit security group created for the
	// cluster if one is not supplied as a parameter.
	// Experimental.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// Security group.
	// Experimental.
	SecurityGroup awsec2.ISecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// Whether to enable storage encryption.
	// Experimental.
	StorageEncrypted *bool `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Where to place the instances within the VPC.
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}


package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDBCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBClusterProps := &cfnDBClusterProps{
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	backupRetentionPeriod: jsii.Number(123),
//   	copyTagsToSnapshot: jsii.Boolean(false),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	dbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	deletionProtection: jsii.Boolean(false),
//   	enableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	engineVersion: jsii.String("engineVersion"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	masterUsername: jsii.String("masterUsername"),
//   	masterUserPassword: jsii.String("masterUserPassword"),
//   	port: jsii.Number(123),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	restoreToTime: jsii.String("restoreToTime"),
//   	restoreType: jsii.String("restoreType"),
//   	snapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	sourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	storageEncrypted: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	useLatestRestorableTime: jsii.Boolean(false),
//   	vpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
type CfnDBClusterProps struct {
	// A list of Amazon EC2 Availability Zones that instances in the cluster can be created in.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The number of days for which automated backups are retained. You must specify a minimum value of 1.
	//
	// Default: 1
	//
	// Constraints:
	//
	// - Must be a value from 1 to 35.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// `AWS::DocDB::DBCluster.CopyTagsToSnapshot`.
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// The cluster identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - The first character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `my-cluster`.
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The name of the cluster parameter group to associate with this cluster.
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// A subnet group to associate with this cluster.
	//
	// Constraints: Must match the name of an existing `DBSubnetGroup` . Must not be default.
	//
	// Example: `mySubnetgroup`.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Protects clusters from being accidentally deleted.
	//
	// If enabled, the cluster cannot be deleted unless it is modified and `DeletionProtection` is disabled.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The list of log types that need to be enabled for exporting to Amazon CloudWatch Logs.
	//
	// You can enable audit logs or profiler logs. For more information, see [Auditing Amazon DocumentDB Events](https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html) and [Profiling Amazon DocumentDB Operations](https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html) .
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// The version number of the database engine to use.
	//
	// The `--engine-version` will default to the latest major engine version. For production workloads, we recommend explicitly declaring this parameter with the intended major engine version.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The AWS KMS key identifier for an encrypted cluster.
	//
	// The AWS KMS key identifier is the Amazon Resource Name (ARN) for the AWS KMS encryption key. If you are creating a cluster using the same AWS account that owns the AWS KMS encryption key that is used to encrypt the new cluster, you can use the AWS KMS key alias instead of the ARN for the AWS KMS encryption key.
	//
	// If an encryption key is not specified in `KmsKeyId` :
	//
	// - If the `StorageEncrypted` parameter is `true` , Amazon DocumentDB uses your default encryption key.
	//
	// AWS KMS creates the default encryption key for your AWS account . Your AWS account has a different default encryption key for each AWS Regions .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of the master user for the cluster.
	//
	// Constraints:
	//
	// - Must be from 1 to 63 letters or numbers.
	// - The first character must be a letter.
	// - Cannot be a reserved word for the chosen database engine.
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// The password for the master database user.
	//
	// This password can contain any printable ASCII character except forward slash (/), double quote ("), or the "at" symbol (@).
	//
	// Constraints: Must contain from 8 to 100 characters.
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// Specifies the port that the database engine is listening on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The daily time range during which automated backups are created if automated backups are enabled using the `BackupRetentionPeriod` parameter.
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region .
	//
	// Constraints:
	//
	// - Must be in the format `hh24:mi-hh24:mi` .
	// - Must be in Universal Coordinated Time (UTC).
	// - Must not conflict with the preferred maintenance window.
	// - Must be at least 30 minutes.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region , occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// `AWS::DocDB::DBCluster.RestoreToTime`.
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// `AWS::DocDB::DBCluster.RestoreType`.
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// The identifier for the snapshot or cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a cluster snapshot. However, you can use only the ARN to specify a snapshot.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing snapshot.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// `AWS::DocDB::DBCluster.SourceDBClusterIdentifier`.
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// Specifies whether the cluster is encrypted.
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// The tags to be assigned to the cluster.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::DocDB::DBCluster.UseLatestRestorableTime`.
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// A list of EC2 VPC security groups to associate with this cluster.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}


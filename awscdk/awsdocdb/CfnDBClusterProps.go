package awsdocdb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDBCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBClusterProps := &CfnDBClusterProps{
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	BackupRetentionPeriod: jsii.Number(123),
//   	CopyTagsToSnapshot: jsii.Boolean(false),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	DeletionProtection: jsii.Boolean(false),
//   	EnableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	EngineVersion: jsii.String("engineVersion"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ManageMasterUserPassword: jsii.Boolean(false),
//   	MasterUsername: jsii.String("masterUsername"),
//   	MasterUserPassword: jsii.String("masterUserPassword"),
//   	MasterUserSecretKmsKeyId: jsii.String("masterUserSecretKmsKeyId"),
//   	Port: jsii.Number(123),
//   	PreferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	RestoreToTime: jsii.String("restoreToTime"),
//   	RestoreType: jsii.String("restoreType"),
//   	RotateMasterUserPassword: jsii.Boolean(false),
//   	ServerlessV2ScalingConfiguration: &ServerlessV2ScalingConfigurationProperty{
//   		MaxCapacity: jsii.Number(123),
//   		MinCapacity: jsii.Number(123),
//   	},
//   	SnapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	SourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	StorageEncrypted: jsii.Boolean(false),
//   	StorageType: jsii.String("storageType"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UseLatestRestorableTime: jsii.Boolean(false),
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html
//
type CfnDBClusterProps struct {
	// A list of Amazon EC2 Availability Zones that instances in the cluster can be created in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-availabilityzones
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// The number of days for which automated backups are retained. You must specify a minimum value of 1.
	//
	// Default: 1
	//
	// Constraints:
	//
	// - Must be a value from 1 to 35.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-backupretentionperiod
	//
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// Set to `true` to copy all tags from the source cluster snapshot to the target cluster snapshot, and otherwise `false` .
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-copytagstosnapshot
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-dbclusteridentifier
	//
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The name of the cluster parameter group to associate with this cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-dbclusterparametergroupname
	//
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// A subnet group to associate with this cluster.
	//
	// Constraints: Must match the name of an existing `DBSubnetGroup` . Must not be default.
	//
	// Example: `mySubnetgroup`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-dbsubnetgroupname
	//
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Protects clusters from being accidentally deleted.
	//
	// If enabled, the cluster cannot be deleted unless it is modified and `DeletionProtection` is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-deletionprotection
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// The list of log types that need to be enabled for exporting to Amazon CloudWatch Logs.
	//
	// You can enable audit logs or profiler logs. For more information, see [Auditing Amazon DocumentDB Events](https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html) and [Profiling Amazon DocumentDB Operations](https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-enablecloudwatchlogsexports
	//
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// The version number of the database engine to use.
	//
	// The `--engine-version` will default to the latest major engine version. For production workloads, we recommend explicitly declaring this parameter with the intended major engine version.
	//
	// Changing the `EngineVersion` will start an in-place engine version upgrade. Note that in-place engine version upgrade will cause downtime in the cluster. See [Amazon DocumentDB in-place major version upgrade](https://docs.aws.amazon.com/documentdb/latest/developerguide/docdb-mvu.html) before starting an in-place engine version upgrade.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-engineversion
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-managemasteruserpassword
	//
	ManageMasterUserPassword interface{} `field:"optional" json:"manageMasterUserPassword" yaml:"manageMasterUserPassword"`
	// The name of the master user for the cluster.
	//
	// Constraints:
	//
	// - Must be from 1 to 63 letters or numbers.
	// - The first character must be a letter.
	// - Cannot be a reserved word for the chosen database engine.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-masterusername
	//
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// The password for the master database user.
	//
	// This password can contain any printable ASCII character except forward slash (/), double quote ("), or the "at" symbol (@).
	//
	// Constraints: Must contain from 8 to 100 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-masteruserpassword
	//
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-masterusersecretkmskeyid
	//
	MasterUserSecretKmsKeyId *string `field:"optional" json:"masterUserSecretKmsKeyId" yaml:"masterUserSecretKmsKeyId"`
	// Specifies the port that the database engine is listening on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-port
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-preferredbackupwindow
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The date and time to restore the cluster to.
	//
	// Valid values: A time in Universal Coordinated Time (UTC) format.
	//
	// Constraints:
	//
	// - Must be before the latest restorable time for the instance.
	// - Must be specified if the `UseLatestRestorableTime` parameter is not provided.
	// - Cannot be specified if the `UseLatestRestorableTime` parameter is `true` .
	// - Cannot be specified if the `RestoreType` parameter is `copy-on-write` .
	//
	// Example: `2015-03-07T23:45:00Z`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-restoretotime
	//
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// The type of restore to be performed. You can specify one of the following values:.
	//
	// - `full-copy` - The new DB cluster is restored as a full copy of the source DB cluster.
	// - `copy-on-write` - The new DB cluster is restored as a clone of the source DB cluster.
	//
	// Constraints: You can't specify `copy-on-write` if the engine version of the source DB cluster is earlier than 1.11.
	//
	// If you don't specify a `RestoreType` value, then the new DB cluster is restored as a full copy of the source DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-restoretype
	//
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-rotatemasteruserpassword
	//
	RotateMasterUserPassword interface{} `field:"optional" json:"rotateMasterUserPassword" yaml:"rotateMasterUserPassword"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-serverlessv2scalingconfiguration
	//
	ServerlessV2ScalingConfiguration interface{} `field:"optional" json:"serverlessV2ScalingConfiguration" yaml:"serverlessV2ScalingConfiguration"`
	// The identifier for the snapshot or cluster snapshot to restore from.
	//
	// You can use either the name or the Amazon Resource Name (ARN) to specify a cluster snapshot. However, you can use only the ARN to specify a snapshot.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-snapshotidentifier
	//
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// The identifier of the source cluster from which to restore.
	//
	// Constraints:
	//
	// - Must match the identifier of an existing `DBCluster` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-sourcedbclusteridentifier
	//
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// Specifies whether the cluster is encrypted.
	//
	// If you specify `SourceDBClusterIdentifier` or `SnapshotIdentifier` and don’t specify `StorageEncrypted` , the encryption property is inherited from the source cluster or snapshot (unless `KMSKeyId` is specified, in which case the restored cluster will be encrypted with that KMS key). If the source is encrypted and `StorageEncrypted` is specified to be true, the restored cluster will be encrypted (if you want to use a different KMS key, specify the `KMSKeyId` property as well). If the source is unencrypted and `StorageEncrypted` is specified to be true, then the `KMSKeyId` property must be specified. If the source is encrypted, don’t specify `StorageEncrypted` to be false as opting out of encryption is not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-storageencrypted
	//
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// The storage type to associate with the DB cluster.
	//
	// For information on storage types for Amazon DocumentDB clusters, see Cluster storage configurations in the *Amazon DocumentDB Developer Guide* .
	//
	// Valid values for storage type - `standard | iopt1`
	//
	// Default value is `standard`
	//
	// > When you create a DocumentDB DB cluster with the storage type set to `iopt1` , the storage type is returned in the response. The storage type isn't returned when you set it to `standard` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// The tags to be assigned to the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A value that is set to `true` to restore the cluster to the latest restorable backup time, and `false` otherwise.
	//
	// Default: `false`
	//
	// Constraints: Cannot be specified if the `RestoreToTime` parameter is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-uselatestrestorabletime
	//
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// A list of EC2 VPC security groups to associate with this cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}


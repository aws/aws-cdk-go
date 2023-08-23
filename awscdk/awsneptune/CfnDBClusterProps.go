package awsneptune

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
//   	AssociatedRoles: []interface{}{
//   		&DBClusterRoleProperty{
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			FeatureName: jsii.String("featureName"),
//   		},
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	BackupRetentionPeriod: jsii.Number(123),
//   	CopyTagsToSnapshot: jsii.Boolean(false),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	DbInstanceParameterGroupName: jsii.String("dbInstanceParameterGroupName"),
//   	DbPort: jsii.Number(123),
//   	DbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	DeletionProtection: jsii.Boolean(false),
//   	EnableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	EngineVersion: jsii.String("engineVersion"),
//   	IamAuthEnabled: jsii.Boolean(false),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	PreferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	RestoreToTime: jsii.String("restoreToTime"),
//   	RestoreType: jsii.String("restoreType"),
//   	ServerlessScalingConfiguration: &ServerlessScalingConfigurationProperty{
//   		MaxCapacity: jsii.Number(123),
//   		MinCapacity: jsii.Number(123),
//   	},
//   	SnapshotIdentifier: jsii.String("snapshotIdentifier"),
//   	SourceDbClusterIdentifier: jsii.String("sourceDbClusterIdentifier"),
//   	StorageEncrypted: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html
//
type CfnDBClusterProps struct {
	// Provides a list of the Amazon Identity and Access Management (IAM) roles that are associated with the DB cluster.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon services on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-associatedroles
	//
	AssociatedRoles interface{} `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-availabilityzones
	//
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Specifies the number of days for which automatic DB snapshots are retained.
	//
	// An update may require some interruption. See [ModifyDBInstance](https://docs.aws.amazon.com/neptune/latest/userguide/api-instances.html#ModifyDBInstance) in the Amazon Neptune User Guide for more information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-backupretentionperiod
	//
	// Default: - 1.
	//
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// *If set to `true` , tags are copied to any snapshot of the DB cluster that is created.*.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-copytagstosnapshot
	//
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Contains a user-supplied DB cluster identifier.
	//
	// This identifier is the unique key that identifies a DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-dbclusteridentifier
	//
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Provides the name of the DB cluster parameter group.
	//
	// An update may require some interruption. See [ModifyDBInstance](https://docs.aws.amazon.com/neptune/latest/userguide/api-instances.html#ModifyDBInstance) in the Amazon Neptune User Guide for more information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-dbclusterparametergroupname
	//
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// The name of the DB parameter group to apply to all instances of the DB cluster.
	//
	// Used only in case of a major engine version upgrade request
	//
	// Note that when you apply a parameter group using `DBInstanceParameterGroupName` , parameter changes are applied immediately, not during the next maintenance window.
	//
	// **Constraints** - The DB parameter group must be in the same DB parameter group family as the target DB cluster version.
	// - The `DBInstanceParameterGroupName` parameter is only valid for major engine version upgrades.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-dbinstanceparametergroupname
	//
	DbInstanceParameterGroupName *string `field:"optional" json:"dbInstanceParameterGroupName" yaml:"dbInstanceParameterGroupName"`
	// The port number on which the DB instances in the DB cluster accept connections.
	//
	// If not specified, the default port used is `8182` .
	//
	// > The `Port` property will soon be deprecated. Please update existing templates to use the new `DBPort` property that has the same functionality.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-dbport
	//
	DbPort *float64 `field:"optional" json:"dbPort" yaml:"dbPort"`
	// Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-dbsubnetgroupname
	//
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Indicates whether or not the DB cluster has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-deletionprotection
	//
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Specifies a list of log types that are enabled for export to CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-enablecloudwatchlogsexports
	//
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// Indicates the database engine version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-iamauthenabled
	//
	IamAuthEnabled interface{} `field:"optional" json:"iamAuthEnabled" yaml:"iamAuthEnabled"`
	// If `StorageEncrypted` is true, the Amazon KMS key identifier for the encrypted DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the `BackupRetentionPeriod` .
	//
	// An update may require some interruption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-preferredbackupwindow
	//
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-restoretotime
	//
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-restoretype
	//
	// Default: - "full-copy".
	//
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// Contains the scaling configuration of an Neptune Serverless DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-serverlessscalingconfiguration
	//
	ServerlessScalingConfiguration interface{} `field:"optional" json:"serverlessScalingConfiguration" yaml:"serverlessScalingConfiguration"`
	// Specifies the identifier for a DB cluster snapshot. Must match the identifier of an existing snapshot.
	//
	// After you restore a DB cluster using a `SnapshotIdentifier` , you must specify the same `SnapshotIdentifier` for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed.
	//
	// However, if you don't specify the `SnapshotIdentifier` , an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, the DB cluster is restored from the snapshot specified by the `SnapshotIdentifier` , and the original DB cluster is deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-snapshotidentifier
	//
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-sourcedbclusteridentifier
	//
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// Indicates whether the DB cluster is encrypted.
	//
	// If you specify the `DBClusterIdentifier` , `DBSnapshotIdentifier` , or `SourceDBInstanceIdentifier` property, don't specify this property. The value is inherited from the cluster, snapshot, or source DB instance. If you specify the `KmsKeyId` property, you must enable encryption.
	//
	// If you specify the `KmsKeyId` , you must enable encryption by setting `StorageEncrypted` to true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-storageencrypted
	//
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// The tags assigned to this cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-uselatestrestorabletime
	//
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// Provides a list of VPC security groups that the DB cluster belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html#cfn-neptune-dbcluster-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}


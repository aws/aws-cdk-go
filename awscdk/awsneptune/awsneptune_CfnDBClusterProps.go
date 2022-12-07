package awsneptune

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
//   	associatedRoles: []interface{}{
//   		&dBClusterRoleProperty{
//   			roleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			featureName: jsii.String("featureName"),
//   		},
//   	},
//   	availabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	backupRetentionPeriod: jsii.Number(123),
//   	dbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	dbClusterParameterGroupName: jsii.String("dbClusterParameterGroupName"),
//   	dbSubnetGroupName: jsii.String("dbSubnetGroupName"),
//   	deletionProtection: jsii.Boolean(false),
//   	enableCloudwatchLogsExports: []*string{
//   		jsii.String("enableCloudwatchLogsExports"),
//   	},
//   	engineVersion: jsii.String("engineVersion"),
//   	iamAuthEnabled: jsii.Boolean(false),
//   	kmsKeyId: jsii.String("kmsKeyId"),
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
	// Provides a list of the Amazon Identity and Access Management (IAM) roles that are associated with the DB cluster.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon services on your behalf.
	AssociatedRoles interface{} `field:"optional" json:"associatedRoles" yaml:"associatedRoles"`
	// Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Specifies the number of days for which automatic DB snapshots are retained.
	//
	// An update may require some interruption. See [ModifyDBInstance](https://docs.aws.amazon.com/neptune/latest/userguide/api-instances.html#ModifyDBInstance) in the Amazon Neptune User Guide for more information.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// Contains a user-supplied DB cluster identifier.
	//
	// This identifier is the unique key that identifies a DB cluster.
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// Provides the name of the DB cluster parameter group.
	//
	// An update may require some interruption. See [ModifyDBInstance](https://docs.aws.amazon.com/neptune/latest/userguide/api-instances.html#ModifyDBInstance) in the Amazon Neptune User Guide for more information.
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Indicates whether or not the DB cluster has deletion protection enabled.
	//
	// The database can't be deleted when deletion protection is enabled.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Specifies a list of log types that are enabled for export to CloudWatch Logs.
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// Indicates the database engine version.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.
	IamAuthEnabled interface{} `field:"optional" json:"iamAuthEnabled" yaml:"iamAuthEnabled"`
	// If `StorageEncrypted` is true, the Amazon KMS key identifier for the encrypted DB cluster.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the port that the database engine is listening on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the `BackupRetentionPeriod` .
	//
	// An update may require some interruption.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// Specifies the identifier for a DB cluster snapshot. Must match the identifier of an existing snapshot.
	//
	// After you restore a DB cluster using a `SnapshotIdentifier` , you must specify the same `SnapshotIdentifier` for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed.
	//
	// However, if you don't specify the `SnapshotIdentifier` , an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, the DB cluster is restored from the snapshot specified by the `SnapshotIdentifier` , and the original DB cluster is deleted.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	SourceDbClusterIdentifier *string `field:"optional" json:"sourceDbClusterIdentifier" yaml:"sourceDbClusterIdentifier"`
	// Indicates whether the DB cluster is encrypted.
	//
	// If you specify the `DBClusterIdentifier` , `DBSnapshotIdentifier` , or `SourceDBInstanceIdentifier` property, don't specify this property. The value is inherited from the cluster, snapshot, or source DB instance. If you specify the `KmsKeyId` property, you must enable encryption.
	//
	// If you specify the `KmsKeyId` , you must enable encryption by setting `StorageEncrypted` to true.
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// The tags assigned to this cluster.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
	//
	// If a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.
	//
	// If a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
	// Provides a list of VPC security groups that the DB cluster belongs to.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}


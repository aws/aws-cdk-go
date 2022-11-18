package awsmemorydb

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &cfnClusterProps{
//   	aclName: jsii.String("aclName"),
//   	clusterName: jsii.String("clusterName"),
//   	nodeType: jsii.String("nodeType"),
//
//   	// the properties below are optional
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	engineVersion: jsii.String("engineVersion"),
//   	finalSnapshotName: jsii.String("finalSnapshotName"),
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   	maintenanceWindow: jsii.String("maintenanceWindow"),
//   	numReplicasPerShard: jsii.Number(123),
//   	numShards: jsii.Number(123),
//   	parameterGroupName: jsii.String("parameterGroupName"),
//   	port: jsii.Number(123),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	snapshotArns: []*string{
//   		jsii.String("snapshotArns"),
//   	},
//   	snapshotName: jsii.String("snapshotName"),
//   	snapshotRetentionLimit: jsii.Number(123),
//   	snapshotWindow: jsii.String("snapshotWindow"),
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   	snsTopicStatus: jsii.String("snsTopicStatus"),
//   	subnetGroupName: jsii.String("subnetGroupName"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	tlsEnabled: jsii.Boolean(false),
//   }
//
type CfnClusterProps struct {
	// The name of the Access Control List to associate with the cluster .
	AclName *string `field:"required" json:"aclName" yaml:"aclName"`
	// The name of the cluster .
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The cluster 's node type.
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// When set to true, the cluster will automatically receive minor engine version upgrades after launch.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// A description of the cluster .
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Redis engine version used by the cluster .
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The user-supplied name of a final cluster snapshot.
	//
	// This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.
	FinalSnapshotName *string `field:"optional" json:"finalSnapshotName" yaml:"finalSnapshotName"`
	// The ID of the KMS key used to encrypt the cluster .
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the weekly time range during which maintenance on the cluster is performed.
	//
	// It is specified as a range in the format `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	//
	// *Pattern* : `ddd:hh24:mi-ddd:hh24:mi`.
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The number of replicas to apply to each shard.
	//
	// *Default value* : `1`
	//
	// *Maximum value* : `5`.
	NumReplicasPerShard *float64 `field:"optional" json:"numReplicasPerShard" yaml:"numReplicasPerShard"`
	// The number of shards in the cluster .
	NumShards *float64 `field:"optional" json:"numShards" yaml:"numShards"`
	// The name of the parameter group used by the cluster .
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// The port used by the cluster .
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A list of security group names to associate with this cluster .
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3.
	//
	// The snapshot files are used to populate the new cluster . The Amazon S3 object name in the ARN cannot contain any commas.
	SnapshotArns *[]*string `field:"optional" json:"snapshotArns" yaml:"snapshotArns"`
	// The name of a snapshot from which to restore data into the new cluster .
	//
	// The snapshot status changes to restoring while the new cluster is being created.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// The number of days for which MemoryDB retains automatic snapshots before deleting them.
	//
	// For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.
	SnapshotRetentionLimit *float64 `field:"optional" json:"snapshotRetentionLimit" yaml:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your shard.
	//
	// Example: 05:00-09:00 If you do not specify this parameter, MemoryDB automatically chooses an appropriate time range.
	SnapshotWindow *string `field:"optional" json:"snapshotWindow" yaml:"snapshotWindow"`
	// When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ARN of the SNS topic, such as `arn:aws:memorydb:us-east-1:123456789012:mySNSTopic`.
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
	// The SNS topic must be in Active status to receive notifications.
	SnsTopicStatus *string `field:"optional" json:"snsTopicStatus" yaml:"snsTopicStatus"`
	// The name of the subnet group used by the cluster .
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// A flag to indicate if In-transit encryption is enabled.
	TlsEnabled interface{} `field:"optional" json:"tlsEnabled" yaml:"tlsEnabled"`
}


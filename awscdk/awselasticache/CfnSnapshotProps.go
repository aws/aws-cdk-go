package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSnapshotProps := &CfnSnapshotProps{
//   	SnapshotName: jsii.String("snapshotName"),
//
//   	// the properties below are optional
//   	CacheClusterId: jsii.String("cacheClusterId"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	ReplicationGroupId: jsii.String("replicationGroupId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-snapshot.html
//
type CfnSnapshotProps struct {
	// The name of a snapshot.
	//
	// Must be unique within the customer account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-snapshot.html#cfn-elasticache-snapshot-snapshotname
	//
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// The identifier of an existing cluster to create a snapshot from.
	//
	// The snapshot is created from this cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-snapshot.html#cfn-elasticache-snapshot-cacheclusterid
	//
	CacheClusterId *string `field:"optional" json:"cacheClusterId" yaml:"cacheClusterId"`
	// The ID of the KMS key used to encrypt the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-snapshot.html#cfn-elasticache-snapshot-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The identifier of an existing replication group to create a snapshot from.
	//
	// The snapshot is created from this replication group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-snapshot.html#cfn-elasticache-snapshot-replicationgroupid
	//
	ReplicationGroupId *string `field:"optional" json:"replicationGroupId" yaml:"replicationGroupId"`
	// A list of tags to be added to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-snapshot.html#cfn-elasticache-snapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


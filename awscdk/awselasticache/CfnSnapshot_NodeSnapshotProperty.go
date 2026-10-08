package awselasticache


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeSnapshotProperty := &NodeSnapshotProperty{
//   	CacheClusterId: jsii.String("cacheClusterId"),
//   	CacheNodeCreateTime: jsii.String("cacheNodeCreateTime"),
//   	CacheNodeId: jsii.String("cacheNodeId"),
//   	CacheSize: jsii.String("cacheSize"),
//   	NodeGroupId: jsii.String("nodeGroupId"),
//   	SnapshotCreateTime: jsii.String("snapshotCreateTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-snapshot-nodesnapshot.html
//
type CfnSnapshot_NodeSnapshotProperty struct {
	// A unique identifier for the source cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-snapshot-nodesnapshot.html#cfn-elasticache-snapshot-nodesnapshot-cacheclusterid
	//
	CacheClusterId *string `field:"optional" json:"cacheClusterId" yaml:"cacheClusterId"`
	// The date and time when the cache node was created in the source cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-snapshot-nodesnapshot.html#cfn-elasticache-snapshot-nodesnapshot-cachenodecreatetime
	//
	CacheNodeCreateTime *string `field:"optional" json:"cacheNodeCreateTime" yaml:"cacheNodeCreateTime"`
	// The cache node identifier for the node in the source cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-snapshot-nodesnapshot.html#cfn-elasticache-snapshot-nodesnapshot-cachenodeid
	//
	CacheNodeId *string `field:"optional" json:"cacheNodeId" yaml:"cacheNodeId"`
	// The size of the cache on the source cache node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-snapshot-nodesnapshot.html#cfn-elasticache-snapshot-nodesnapshot-cachesize
	//
	CacheSize *string `field:"optional" json:"cacheSize" yaml:"cacheSize"`
	// A unique identifier for the source node group (shard).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-snapshot-nodesnapshot.html#cfn-elasticache-snapshot-nodesnapshot-nodegroupid
	//
	NodeGroupId *string `field:"optional" json:"nodeGroupId" yaml:"nodeGroupId"`
	// The date and time when the source node's metadata and cache data set was obtained for the snapshot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-snapshot-nodesnapshot.html#cfn-elasticache-snapshot-nodesnapshot-snapshotcreatetime
	//
	SnapshotCreateTime *string `field:"optional" json:"snapshotCreateTime" yaml:"snapshotCreateTime"`
}


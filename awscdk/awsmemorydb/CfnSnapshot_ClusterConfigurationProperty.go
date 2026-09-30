package awsmemorydb


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigurationProperty := &ClusterConfigurationProperty{
//   	Description: jsii.String("description"),
//   	Engine: jsii.String("engine"),
//   	EngineVersion: jsii.String("engineVersion"),
//   	MaintenanceWindow: jsii.String("maintenanceWindow"),
//   	Name: jsii.String("name"),
//   	NodeType: jsii.String("nodeType"),
//   	NumShards: jsii.Number(123),
//   	ParameterGroupName: jsii.String("parameterGroupName"),
//   	Port: jsii.Number(123),
//   	SnapshotRetentionLimit: jsii.Number(123),
//   	SnapshotWindow: jsii.String("snapshotWindow"),
//   	SubnetGroupName: jsii.String("subnetGroupName"),
//   	TopicArn: jsii.String("topicArn"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html
//
type CfnSnapshot_ClusterConfigurationProperty struct {
	// The description of the cluster configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the engine used by the cluster configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-engine
	//
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// The Redis OSS engine version used by the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The specified maintenance window for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-maintenancewindow
	//
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The name of the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The node type used for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-nodetype
	//
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// The number of shards in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-numshards
	//
	NumShards *float64 `field:"optional" json:"numShards" yaml:"numShards"`
	// The name of parameter group used by the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-parametergroupname
	//
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// The port used by the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The snapshot retention limit set by the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-snapshotretentionlimit
	//
	SnapshotRetentionLimit *float64 `field:"optional" json:"snapshotRetentionLimit" yaml:"snapshotRetentionLimit"`
	// The snapshot window set by the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-snapshotwindow
	//
	SnapshotWindow *string `field:"optional" json:"snapshotWindow" yaml:"snapshotWindow"`
	// The name of the subnet group used by the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-subnetgroupname
	//
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// The Amazon Resource Name (ARN) of the SNS notification topic for the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-topicarn
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
	// The ID of the VPC the cluster belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-memorydb-snapshot-clusterconfiguration.html#cfn-memorydb-snapshot-clusterconfiguration-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}


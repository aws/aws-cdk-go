package awsopensearchservice


// The cluster configuration for the OpenSearch Service domain.
//
// You can specify options such as the instance type and the number of instances. For more information, see [Creating and managing Amazon OpenSearch Service domains](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html) in the *Amazon OpenSearch Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clusterConfigProperty := &clusterConfigProperty{
//   	dedicatedMasterCount: jsii.Number(123),
//   	dedicatedMasterEnabled: jsii.Boolean(false),
//   	dedicatedMasterType: jsii.String("dedicatedMasterType"),
//   	instanceCount: jsii.Number(123),
//   	instanceType: jsii.String("instanceType"),
//   	warmCount: jsii.Number(123),
//   	warmEnabled: jsii.Boolean(false),
//   	warmType: jsii.String("warmType"),
//   	zoneAwarenessConfig: &zoneAwarenessConfigProperty{
//   		availabilityZoneCount: jsii.Number(123),
//   	},
//   	zoneAwarenessEnabled: jsii.Boolean(false),
//   }
//
type CfnDomain_ClusterConfigProperty struct {
	// The number of instances to use for the master node.
	//
	// If you specify this property, you must specify `true` for the `DedicatedMasterEnabled` property.
	DedicatedMasterCount *float64 `field:"optional" json:"dedicatedMasterCount" yaml:"dedicatedMasterCount"`
	// Indicates whether to use a dedicated master node for the OpenSearch Service domain.
	//
	// A dedicated master node is a cluster node that performs cluster management tasks, but doesn't hold data or respond to data upload requests. Dedicated master nodes offload cluster management tasks to increase the stability of your search clusters. See [Dedicated master nodes in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-dedicatedmasternodes.html) .
	DedicatedMasterEnabled interface{} `field:"optional" json:"dedicatedMasterEnabled" yaml:"dedicatedMasterEnabled"`
	// The hardware configuration of the computer that hosts the dedicated master node, such as `m3.medium.search` . If you specify this property, you must specify `true` for the `DedicatedMasterEnabled` property. For valid values, see [Supported instance types in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) .
	DedicatedMasterType *string `field:"optional" json:"dedicatedMasterType" yaml:"dedicatedMasterType"`
	// The number of data nodes (instances) to use in the OpenSearch Service domain.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The instance type for your data nodes, such as `m3.medium.search` . For valid values, see [Supported instance types in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) .
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The number of warm nodes in the cluster.
	WarmCount *float64 `field:"optional" json:"warmCount" yaml:"warmCount"`
	// Whether to enable UltraWarm storage for the cluster.
	//
	// See [UltraWarm storage for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ultrawarm.html) .
	WarmEnabled interface{} `field:"optional" json:"warmEnabled" yaml:"warmEnabled"`
	// The instance type for the cluster's warm nodes.
	WarmType *string `field:"optional" json:"warmType" yaml:"warmType"`
	// Specifies zone awareness configuration options.
	//
	// Only use if `ZoneAwarenessEnabled` is `true` .
	ZoneAwarenessConfig interface{} `field:"optional" json:"zoneAwarenessConfig" yaml:"zoneAwarenessConfig"`
	// Indicates whether to enable zone awareness for the OpenSearch Service domain.
	//
	// When you enable zone awareness, OpenSearch Service allocates the nodes and replica index shards that belong to a cluster across two Availability Zones (AZs) in the same region to prevent data loss and minimize downtime in the event of node or data center failure. Don't enable zone awareness if your cluster has no replica index shards or is a single-node cluster. For more information, see [Configuring a multi-AZ domain in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html) .
	ZoneAwarenessEnabled interface{} `field:"optional" json:"zoneAwarenessEnabled" yaml:"zoneAwarenessEnabled"`
}


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
//   clusterConfigProperty := &ClusterConfigProperty{
//   	ColdStorageOptions: &ColdStorageOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   	DedicatedMasterCount: jsii.Number(123),
//   	DedicatedMasterEnabled: jsii.Boolean(false),
//   	DedicatedMasterType: jsii.String("dedicatedMasterType"),
//   	InstanceCount: jsii.Number(123),
//   	InstanceType: jsii.String("instanceType"),
//   	MultiAzWithStandbyEnabled: jsii.Boolean(false),
//   	WarmCount: jsii.Number(123),
//   	WarmEnabled: jsii.Boolean(false),
//   	WarmType: jsii.String("warmType"),
//   	ZoneAwarenessConfig: &ZoneAwarenessConfigProperty{
//   		AvailabilityZoneCount: jsii.Number(123),
//   	},
//   	ZoneAwarenessEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html
//
type CfnDomain_ClusterConfigProperty struct {
	// Specifies cold storage options for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-coldstorageoptions
	//
	ColdStorageOptions interface{} `field:"optional" json:"coldStorageOptions" yaml:"coldStorageOptions"`
	// The number of instances to use for the master node.
	//
	// If you specify this property, you must specify `true` for the `DedicatedMasterEnabled` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-dedicatedmastercount
	//
	DedicatedMasterCount *float64 `field:"optional" json:"dedicatedMasterCount" yaml:"dedicatedMasterCount"`
	// Indicates whether to use a dedicated master node for the OpenSearch Service domain.
	//
	// A dedicated master node is a cluster node that performs cluster management tasks, but doesn't hold data or respond to data upload requests. Dedicated master nodes offload cluster management tasks to increase the stability of your search clusters. See [Dedicated master nodes in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-dedicatedmasternodes.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-dedicatedmasterenabled
	//
	DedicatedMasterEnabled interface{} `field:"optional" json:"dedicatedMasterEnabled" yaml:"dedicatedMasterEnabled"`
	// The hardware configuration of the computer that hosts the dedicated master node, such as `m3.medium.search` . If you specify this property, you must specify `true` for the `DedicatedMasterEnabled` property. For valid values, see [Supported instance types in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-dedicatedmastertype
	//
	DedicatedMasterType *string `field:"optional" json:"dedicatedMasterType" yaml:"dedicatedMasterType"`
	// The number of data nodes (instances) to use in the OpenSearch Service domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-instancecount
	//
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// The instance type for your data nodes, such as `m3.medium.search` . For valid values, see [Supported instance types in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-instancetype
	//
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Indicates whether Multi-AZ with Standby deployment option is enabled.
	//
	// For more information, see [Multi-AZ with Standby](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html#managedomains-za-standby) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-multiazwithstandbyenabled
	//
	MultiAzWithStandbyEnabled interface{} `field:"optional" json:"multiAzWithStandbyEnabled" yaml:"multiAzWithStandbyEnabled"`
	// The number of warm nodes in the cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-warmcount
	//
	WarmCount *float64 `field:"optional" json:"warmCount" yaml:"warmCount"`
	// Whether to enable UltraWarm storage for the cluster.
	//
	// See [UltraWarm storage for Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ultrawarm.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-warmenabled
	//
	WarmEnabled interface{} `field:"optional" json:"warmEnabled" yaml:"warmEnabled"`
	// The instance type for the cluster's warm nodes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-warmtype
	//
	WarmType *string `field:"optional" json:"warmType" yaml:"warmType"`
	// Specifies zone awareness configuration options.
	//
	// Only use if `ZoneAwarenessEnabled` is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-zoneawarenessconfig
	//
	ZoneAwarenessConfig interface{} `field:"optional" json:"zoneAwarenessConfig" yaml:"zoneAwarenessConfig"`
	// Indicates whether to enable zone awareness for the OpenSearch Service domain.
	//
	// When you enable zone awareness, OpenSearch Service allocates the nodes and replica index shards that belong to a cluster across two Availability Zones (AZs) in the same region to prevent data loss and minimize downtime in the event of node or data center failure. Don't enable zone awareness if your cluster has no replica index shards or is a single-node cluster. For more information, see [Configuring a multi-AZ domain in Amazon OpenSearch Service](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-clusterconfig.html#cfn-opensearchservice-domain-clusterconfig-zoneawarenessenabled
	//
	ZoneAwarenessEnabled interface{} `field:"optional" json:"zoneAwarenessEnabled" yaml:"zoneAwarenessEnabled"`
}


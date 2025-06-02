package awselasticsearch


// Configures the capacity of the cluster such as the instance type and the number of instances.
//
// Example:
//   domain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: es.ElasticsearchVersion_V7_10(),
//   	Capacity: &CapacityConfig{
//   		MasterNodes: jsii.Number(2),
//   		WarmNodes: jsii.Number(2),
//   		WarmInstanceType: jsii.String("ultrawarm1.medium.elasticsearch"),
//   	},
//   })
//
// Deprecated: use opensearchservice module instead.
type CapacityConfig struct {
	// The instance type for your data nodes, such as `m3.medium.elasticsearch`. For valid values, see [Supported Instance Types](https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html) in the Amazon Elasticsearch Service Developer Guide.
	// Default: - r5.large.elasticsearch
	//
	// Deprecated: use opensearchservice module instead.
	DataNodeInstanceType *string `field:"optional" json:"dataNodeInstanceType" yaml:"dataNodeInstanceType"`
	// The number of data nodes (instances) to use in the Amazon ES domain.
	// Default: - 1.
	//
	// Deprecated: use opensearchservice module instead.
	DataNodes *float64 `field:"optional" json:"dataNodes" yaml:"dataNodes"`
	// The hardware configuration of the computer that hosts the dedicated master node, such as `m3.medium.elasticsearch`. For valid values, see [Supported Instance Types] (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-supported-instance-types.html) in the Amazon Elasticsearch Service Developer Guide.
	// Default: - r5.large.elasticsearch
	//
	// Deprecated: use opensearchservice module instead.
	MasterNodeInstanceType *string `field:"optional" json:"masterNodeInstanceType" yaml:"masterNodeInstanceType"`
	// The number of instances to use for the master node.
	// Default: - no dedicated master nodes.
	//
	// Deprecated: use opensearchservice module instead.
	MasterNodes *float64 `field:"optional" json:"masterNodes" yaml:"masterNodes"`
	// The instance type for your UltraWarm node, such as `ultrawarm1.medium.elasticsearch`. For valid values, see [UltraWarm Storage Limits] (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/aes-limits.html#limits-ultrawarm) in the Amazon Elasticsearch Service Developer Guide.
	// Default: - ultrawarm1.medium.elasticsearch
	//
	// Deprecated: use opensearchservice module instead.
	WarmInstanceType *string `field:"optional" json:"warmInstanceType" yaml:"warmInstanceType"`
	// The number of UltraWarm nodes (instances) to use in the Amazon ES domain.
	// Default: - no UltraWarm nodes.
	//
	// Deprecated: use opensearchservice module instead.
	WarmNodes *float64 `field:"optional" json:"warmNodes" yaml:"warmNodes"`
}


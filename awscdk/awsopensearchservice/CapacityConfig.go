package awsopensearchservice


// Configures the capacity of the cluster such as the instance type and the number of instances.
//
// Example:
//   domain := opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: opensearch.EngineVersion_OPENSEARCH_1_0(),
//   	Capacity: &CapacityConfig{
//   		MasterNodes: jsii.Number(2),
//   		WarmNodes: jsii.Number(2),
//   		WarmInstanceType: jsii.String("ultrawarm1.medium.search"),
//   	},
//   })
//
// Experimental.
type CapacityConfig struct {
	// The instance type for your data nodes, such as `m3.medium.search`. For valid values, see [Supported Instance Types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) in the Amazon OpenSearch Service Developer Guide.
	// Experimental.
	DataNodeInstanceType *string `field:"optional" json:"dataNodeInstanceType" yaml:"dataNodeInstanceType"`
	// The number of data nodes (instances) to use in the Amazon OpenSearch Service domain.
	// Experimental.
	DataNodes *float64 `field:"optional" json:"dataNodes" yaml:"dataNodes"`
	// The hardware configuration of the computer that hosts the dedicated master node, such as `m3.medium.search`. For valid values, see [Supported Instance Types] (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) in the Amazon OpenSearch Service Developer Guide.
	// Experimental.
	MasterNodeInstanceType *string `field:"optional" json:"masterNodeInstanceType" yaml:"masterNodeInstanceType"`
	// The number of instances to use for the master node.
	// Experimental.
	MasterNodes *float64 `field:"optional" json:"masterNodes" yaml:"masterNodes"`
	// The instance type for your UltraWarm node, such as `ultrawarm1.medium.search`. For valid values, see [UltraWarm Storage Limits] (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#limits-ultrawarm) in the Amazon OpenSearch Service Developer Guide.
	// Experimental.
	WarmInstanceType *string `field:"optional" json:"warmInstanceType" yaml:"warmInstanceType"`
	// The number of UltraWarm nodes (instances) to use in the Amazon OpenSearch Service domain.
	// Experimental.
	WarmNodes *float64 `field:"optional" json:"warmNodes" yaml:"warmNodes"`
}


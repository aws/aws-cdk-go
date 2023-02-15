package awsopensearchservice


// Configures the capacity of the cluster such as the instance type and the number of instances.
//
// Example:
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &domainProps{
//   	version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	capacity: &capacityConfig{
//   		masterNodes: jsii.Number(2),
//   		warmNodes: jsii.Number(2),
//   		warmInstanceType: jsii.String("ultrawarm1.medium.search"),
//   	},
//   })
//
type CapacityConfig struct {
	// The instance type for your data nodes, such as `m3.medium.search`. For valid values, see [Supported Instance Types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) in the Amazon OpenSearch Service Developer Guide.
	DataNodeInstanceType *string `field:"optional" json:"dataNodeInstanceType" yaml:"dataNodeInstanceType"`
	// The number of data nodes (instances) to use in the Amazon OpenSearch Service domain.
	DataNodes *float64 `field:"optional" json:"dataNodes" yaml:"dataNodes"`
	// The hardware configuration of the computer that hosts the dedicated master node, such as `m3.medium.search`. For valid values, see [Supported Instance Types] (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) in the Amazon OpenSearch Service Developer Guide.
	MasterNodeInstanceType *string `field:"optional" json:"masterNodeInstanceType" yaml:"masterNodeInstanceType"`
	// The number of instances to use for the master node.
	MasterNodes *float64 `field:"optional" json:"masterNodes" yaml:"masterNodes"`
	// The instance type for your UltraWarm node, such as `ultrawarm1.medium.search`. For valid values, see [UltraWarm Storage Limits] (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#limits-ultrawarm) in the Amazon OpenSearch Service Developer Guide.
	WarmInstanceType *string `field:"optional" json:"warmInstanceType" yaml:"warmInstanceType"`
	// The number of UltraWarm nodes (instances) to use in the Amazon OpenSearch Service domain.
	WarmNodes *float64 `field:"optional" json:"warmNodes" yaml:"warmNodes"`
}


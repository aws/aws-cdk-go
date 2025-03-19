package awsopensearchservice


// Configures the capacity of the cluster such as the instance type and the number of instances.
//
// Example:
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	Capacity: &CapacityConfig{
//   		MasterNodes: jsii.Number(2),
//   		WarmNodes: jsii.Number(2),
//   		WarmInstanceType: jsii.String("ultrawarm1.medium.search"),
//   	},
//   	ColdStorageEnabled: jsii.Boolean(true),
//   })
//
type CapacityConfig struct {
	// The instance type for your data nodes, such as `m3.medium.search`. For valid values, see [Supported Instance Types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) in the Amazon OpenSearch Service Developer Guide.
	// Default: - r5.large.search
	//
	DataNodeInstanceType *string `field:"optional" json:"dataNodeInstanceType" yaml:"dataNodeInstanceType"`
	// The number of data nodes (instances) to use in the Amazon OpenSearch Service domain.
	// Default: - 1.
	//
	DataNodes *float64 `field:"optional" json:"dataNodes" yaml:"dataNodes"`
	// The hardware configuration of the computer that hosts the dedicated master node, such as `m3.medium.search`. For valid values, see [Supported Instance Types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/supported-instance-types.html) in the Amazon OpenSearch Service Developer Guide.
	// Default: - r5.large.search
	//
	MasterNodeInstanceType *string `field:"optional" json:"masterNodeInstanceType" yaml:"masterNodeInstanceType"`
	// The number of instances to use for the master node.
	// Default: - no dedicated master nodes.
	//
	MasterNodes *float64 `field:"optional" json:"masterNodes" yaml:"masterNodes"`
	// Indicates whether Multi-AZ with Standby deployment option is enabled.
	//
	// For more information, see [Multi-AZ with
	// Standby](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html#managedomains-za-standby)
	// Default: - multi-az with standby if the feature flag `ENABLE_OPENSEARCH_MULTIAZ_WITH_STANDBY`
	// is true, no multi-az with standby otherwise.
	//
	MultiAzWithStandbyEnabled *bool `field:"optional" json:"multiAzWithStandbyEnabled" yaml:"multiAzWithStandbyEnabled"`
	// Additional node options for the domain.
	// Default: - no additional node options.
	//
	NodeOptions *[]*NodeOptions `field:"optional" json:"nodeOptions" yaml:"nodeOptions"`
	// The instance type for your UltraWarm node, such as `ultrawarm1.medium.search`. For valid values, see [UltraWarm Storage Limits](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/limits.html#limits-ultrawarm) in the Amazon OpenSearch Service Developer Guide.
	// Default: - ultrawarm1.medium.search
	//
	WarmInstanceType *string `field:"optional" json:"warmInstanceType" yaml:"warmInstanceType"`
	// The number of UltraWarm nodes (instances) to use in the Amazon OpenSearch Service domain.
	// Default: - no UltraWarm nodes.
	//
	WarmNodes *float64 `field:"optional" json:"warmNodes" yaml:"warmNodes"`
}


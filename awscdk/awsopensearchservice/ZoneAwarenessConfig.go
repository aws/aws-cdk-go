package awsopensearchservice


// Specifies zone awareness configuration options.
//
// Example:
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_3(),
//   	Ebs: &EbsOptions{
//   		VolumeSize: jsii.Number(10),
//   		VolumeType: ec2.EbsDeviceVolumeType_GENERAL_PURPOSE_SSD_GP3,
//   	},
//   	ZoneAwareness: &ZoneAwarenessConfig{
//   		Enabled: jsii.Boolean(true),
//   		AvailabilityZoneCount: jsii.Number(3),
//   	},
//   	Capacity: &CapacityConfig{
//   		MultiAzWithStandbyEnabled: jsii.Boolean(true),
//   		MasterNodes: jsii.Number(3),
//   		DataNodes: jsii.Number(3),
//   	},
//   })
//
type ZoneAwarenessConfig struct {
	// If you enabled multiple Availability Zones (AZs), the number of AZs that you want the domain to use.
	//
	// Valid values are 2 and 3.
	AvailabilityZoneCount *float64 `field:"optional" json:"availabilityZoneCount" yaml:"availabilityZoneCount"`
	// Indicates whether to enable zone awareness for the Amazon OpenSearch Service domain.
	//
	// When you enable zone awareness, Amazon OpenSearch Service allocates the nodes and replica
	// index shards that belong to a cluster across two Availability Zones (AZs)
	// in the same region to prevent data loss and minimize downtime in the event
	// of node or data center failure. Don't enable zone awareness if your cluster
	// has no replica index shards or is a single-node cluster. For more information,
	// see [Configuring a Multi-AZ Domain]
	// (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-multiaz.html)
	// in the Amazon OpenSearch Service Developer Guide.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}


package awselasticsearch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// The configurations of Amazon Elastic Block Store (Amazon EBS) volumes that are attached to data nodes in the Amazon ES domain.
//
// For more information, see
// [Configuring EBS-based Storage]
// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
// in the Amazon Elasticsearch Service Developer Guide.
//
// Example:
//   prodDomain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: es.ElasticsearchVersion_V7_1(),
//   	Capacity: &CapacityConfig{
//   		MasterNodes: jsii.Number(5),
//   		DataNodes: jsii.Number(20),
//   	},
//   	Ebs: &EbsOptions{
//   		VolumeSize: jsii.Number(20),
//   	},
//   	ZoneAwareness: &ZoneAwarenessConfig{
//   		AvailabilityZoneCount: jsii.Number(3),
//   	},
//   	Logging: &LoggingOptions{
//   		SlowSearchLogEnabled: jsii.Boolean(true),
//   		AppLogEnabled: jsii.Boolean(true),
//   		SlowIndexLogEnabled: jsii.Boolean(true),
//   	},
//   })
//
// Deprecated: use opensearchservice module instead.
type EbsOptions struct {
	// Specifies whether Amazon EBS volumes are attached to data nodes in the Amazon ES domain.
	// Default: - true.
	//
	// Deprecated: use opensearchservice module instead.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	//
	// This property applies only to the Provisioned IOPS (SSD) EBS
	// volume type.
	// Default: - iops are not set.
	//
	// Deprecated: use opensearchservice module instead.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The size (in GiB) of the EBS volume for each data node.
	//
	// The minimum and
	// maximum size of an EBS volume depends on the EBS volume type and the
	// instance type to which it is attached.  For more information, see
	// [Configuring EBS-based Storage]
	// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
	// in the Amazon Elasticsearch Service Developer Guide.
	// Default: 10.
	//
	// Deprecated: use opensearchservice module instead.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// The EBS volume type to use with the Amazon ES domain, such as standard, gp2, io1.
	//
	// For more information, see[Configuring EBS-based Storage]
	// (https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)
	// in the Amazon Elasticsearch Service Developer Guide.
	// Default: gp2.
	//
	// Deprecated: use opensearchservice module instead.
	VolumeType awsec2.EbsDeviceVolumeType `field:"optional" json:"volumeType" yaml:"volumeType"`
}


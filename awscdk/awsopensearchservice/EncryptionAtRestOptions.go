package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service (KMS) key to use.
//
// Can only be used to create a new domain,
// not update an existing one. Requires Elasticsearch version 5.1 or later or OpenSearch version 1.0 or later.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   domain := opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: opensearch.EngineVersion_OPENSEARCH_2_17(),
//   	EncryptionAtRest: &EncryptionAtRestOptions{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	NodeToNodeEncryption: jsii.Boolean(true),
//   	EnforceHttps: jsii.Boolean(true),
//   	Capacity: &CapacityConfig{
//   		MultiAzWithStandbyEnabled: jsii.Boolean(false),
//   	},
//   	Ebs: &EbsOptions{
//   		Enabled: jsii.Boolean(true),
//   		VolumeSize: jsii.Number(10),
//   	},
//   })
//   api := appsync.NewEventApi(this, jsii.String("EventApiOpenSearch"), &EventApiProps{
//   	ApiName: jsii.String("OpenSearchEventApi"),
//   })
//
//   dataSource := api.AddOpenSearchDataSource(jsii.String("opensearchds"), domain)
//
type EncryptionAtRestOptions struct {
	// Specify true to enable encryption at rest.
	// Default: - encryption at rest is disabled.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Supply if using KMS key for encryption at rest.
	// Default: - uses default aws/es KMS key.
	//
	KmsKey awskms.IKeyRef `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}


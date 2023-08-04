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
//   domain := awscdk.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: awscdk.EngineVersion_OPENSEARCH_1_0(),
//   	EnforceHttps: jsii.Boolean(true),
//   	NodeToNodeEncryption: jsii.Boolean(true),
//   	EncryptionAtRest: &EncryptionAtRestOptions{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	FineGrainedAccessControl: &AdvancedSecurityOptions{
//   		MasterUserName: jsii.String("master-user"),
//   	},
//   	Logging: &LoggingOptions{
//   		AuditLogEnabled: jsii.Boolean(true),
//   		SlowSearchLogEnabled: jsii.Boolean(true),
//   		AppLogEnabled: jsii.Boolean(true),
//   		SlowIndexLogEnabled: jsii.Boolean(true),
//   	},
//   })
//
type EncryptionAtRestOptions struct {
	// Specify true to enable encryption at rest.
	// Default: - encryption at rest is disabled.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Supply if using KMS key for encryption at rest.
	// Default: - uses default aws/es KMS key.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}


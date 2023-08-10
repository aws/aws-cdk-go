package awselasticsearch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Whether the domain should encrypt data at rest, and if so, the AWS Key Management Service (KMS) key to use.
//
// Can only be used to create a new domain,
// not update an existing one. Requires Elasticsearch version 5.1 or later.
//
// Example:
//   domain := es.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	Version: es.ElasticsearchVersion_V7_1(),
//   	EnforceHttps: jsii.Boolean(true),
//   	NodeToNodeEncryption: jsii.Boolean(true),
//   	EncryptionAtRest: &EncryptionAtRestOptions{
//   		Enabled: jsii.Boolean(true),
//   	},
//   	FineGrainedAccessControl: &AdvancedSecurityOptions{
//   		MasterUserName: jsii.String("master-user"),
//   	},
//   })
//
//   masterUserPassword := domain.MasterUserPassword
//
// Deprecated: use opensearchservice module instead.
type EncryptionAtRestOptions struct {
	// Specify true to enable encryption at rest.
	// Deprecated: use opensearchservice module instead.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Supply if using KMS key for encryption at rest.
	// Deprecated: use opensearchservice module instead.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}


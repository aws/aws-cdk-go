package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Specifies options for fine-grained access control.
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
//   		SamlAuthenticationEnabled: jsii.Boolean(true),
//   		SamlAuthenticationOptions: &SAMLOptionsProperty{
//   			IdpEntityId: jsii.String("entity-id"),
//   			IdpMetadataContent: jsii.String("metadata-content-with-quotes-escaped"),
//   		},
//   	},
//   })
//
type AdvancedSecurityOptions struct {
	// ARN for the master user.
	//
	// Only specify this or masterUserName, but not both.
	// Default: - fine-grained access control is disabled.
	//
	MasterUserArn *string `field:"optional" json:"masterUserArn" yaml:"masterUserArn"`
	// Username for the master user.
	//
	// Only specify this or masterUserArn, but not both.
	// Default: - fine-grained access control is disabled.
	//
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Password for the master user.
	//
	// You can use `SecretValue.unsafePlainText` to specify a password in plain text or
	// use `secretsmanager.Secret.fromSecretAttributes` to reference a secret in
	// Secrets Manager.
	// Default: - A Secrets Manager generated password.
	//
	MasterUserPassword awscdk.SecretValue `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
	// True to enable SAML authentication for a domain.
	// See: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/saml.html
	//
	// Default: - SAML authentication is disabled. Enabled if `samlAuthenticationOptions` is set.
	//
	SamlAuthenticationEnabled *bool `field:"optional" json:"samlAuthenticationEnabled" yaml:"samlAuthenticationEnabled"`
	// Container for information about the SAML configuration for OpenSearch Dashboards.
	//
	// If set, `samlAuthenticationEnabled` will be enabled.
	// Default: - no SAML authentication options.
	//
	SamlAuthenticationOptions *SAMLOptionsProperty `field:"optional" json:"samlAuthenticationOptions" yaml:"samlAuthenticationOptions"`
}


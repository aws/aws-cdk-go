package awselasticsearch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Specifies options for fine-grained access control.
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
type AdvancedSecurityOptions struct {
	// ARN for the master user.
	//
	// Only specify this or masterUserName, but not both.
	// Deprecated: use opensearchservice module instead.
	MasterUserArn *string `field:"optional" json:"masterUserArn" yaml:"masterUserArn"`
	// Username for the master user.
	//
	// Only specify this or masterUserArn, but not both.
	// Deprecated: use opensearchservice module instead.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Password for the master user.
	//
	// You can use `SecretValue.unsafePlainText` to specify a password in plain text or
	// use `secretsmanager.Secret.fromSecretAttributes` to reference a secret in
	// Secrets Manager.
	// Deprecated: use opensearchservice module instead.
	MasterUserPassword awscdk.SecretValue `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
}


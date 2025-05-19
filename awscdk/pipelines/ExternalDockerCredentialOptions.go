package pipelines

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options for defining credentials for a Docker Credential.
//
// Example:
//   dockerHubSecret := secretsmanager.Secret_FromSecretCompleteArn(this, jsii.String("DHSecret"), jsii.String("arn:aws:..."))
//   // Only the image asset publishing actions will be granted read access to the secret.
//   creds := pipelines.DockerCredential_DockerHub(dockerHubSecret, &ExternalDockerCredentialOptions{
//   	Usages: []dockerCredentialUsage{
//   		pipelines.*dockerCredentialUsage_ASSET_PUBLISHING,
//   	},
//   })
//
type ExternalDockerCredentialOptions struct {
	// An IAM role to assume prior to accessing the secret.
	// Default: - none. The current execution role will be used.
	//
	AssumeRole awsiam.IRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// The name of the JSON field of the secret which contains the secret/password.
	// Default: 'secret'.
	//
	SecretPasswordField *string `field:"optional" json:"secretPasswordField" yaml:"secretPasswordField"`
	// The name of the JSON field of the secret which contains the user/login name.
	// Default: 'username'.
	//
	SecretUsernameField *string `field:"optional" json:"secretUsernameField" yaml:"secretUsernameField"`
	// Defines which stages of the pipeline should be granted access to these credentials.
	// Default: - all relevant stages (synth, self-update, asset publishing) are granted access.
	//
	Usages *[]DockerCredentialUsage `field:"optional" json:"usages" yaml:"usages"`
}


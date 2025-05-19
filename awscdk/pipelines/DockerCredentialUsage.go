package pipelines


// Defines which stages of a pipeline require the specified credentials.
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
type DockerCredentialUsage string

const (
	// Synth/Build.
	DockerCredentialUsage_SYNTH DockerCredentialUsage = "SYNTH"
	// Self-update.
	DockerCredentialUsage_SELF_UPDATE DockerCredentialUsage = "SELF_UPDATE"
	// Asset publishing.
	DockerCredentialUsage_ASSET_PUBLISHING DockerCredentialUsage = "ASSET_PUBLISHING"
)


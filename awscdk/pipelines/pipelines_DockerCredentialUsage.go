package pipelines


// Defines which stages of a pipeline require the specified credentials.
//
// Example:
//   dockerHubSecret := secretsmanager.secret.fromSecretCompleteArn(this, jsii.String("DHSecret"), jsii.String("arn:aws:..."))
//   // Only the image asset publishing actions will be granted read access to the secret.
//   creds := pipelines.dockerCredential.dockerHub(dockerHubSecret, &externalDockerCredentialOptions{
//   	usages: []dockerCredentialUsage{
//   		pipelines.*dockerCredentialUsage_ASSET_PUBLISHING,
//   	},
//   })
//
// Experimental.
type DockerCredentialUsage string

const (
	// Synth/Build.
	// Experimental.
	DockerCredentialUsage_SYNTH DockerCredentialUsage = "SYNTH"
	// Self-update.
	// Experimental.
	DockerCredentialUsage_SELF_UPDATE DockerCredentialUsage = "SELF_UPDATE"
	// Asset publishing.
	// Experimental.
	DockerCredentialUsage_ASSET_PUBLISHING DockerCredentialUsage = "ASSET_PUBLISHING"
)


package awscdk


// Options for referencing a secret value from Secrets Manager.
//
// Example:
//   codebuild.NewBitBucketSourceCredentials(this, jsii.String("CodeBuildBitBucketCreds"), &BitBucketSourceCredentialsProps{
//   	Username: awscdk.SecretValue_SecretsManager(jsii.String("my-bitbucket-creds"), &SecretsManagerSecretOptions{
//   		JsonField: jsii.String("username"),
//   	}),
//   	Password: awscdk.SecretValue_*SecretsManager(jsii.String("my-bitbucket-creds"), &SecretsManagerSecretOptions{
//   		JsonField: jsii.String("password"),
//   	}),
//   })
//
type SecretsManagerSecretOptions struct {
	// The key of a JSON field to retrieve.
	//
	// This can only be used if the secret
	// stores a JSON object.
	// Default: - returns all the content stored in the Secrets Manager secret.
	//
	JsonField *string `field:"optional" json:"jsonField" yaml:"jsonField"`
	// Specifies the unique identifier of the version of the secret you want to use.
	//
	// Can specify at most one of `versionId` and `versionStage`.
	// Default: AWSCURRENT.
	//
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// Specifies the secret version that you want to retrieve by the staging label attached to the version.
	//
	// Can specify at most one of `versionId` and `versionStage`.
	// Default: AWSCURRENT.
	//
	VersionStage *string `field:"optional" json:"versionStage" yaml:"versionStage"`
}


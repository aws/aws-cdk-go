// Version 2 of the AWS Cloud Development Kit library
package awscdk


// Options for referencing a secret value from Secrets Manager.
//
// Example:
//   codebuild.NewBitBucketSourceCredentials(this, jsii.String("CodeBuildBitBucketCreds"), &bitBucketSourceCredentialsProps{
//   	username: awscdk.SecretValue.secretsManager(jsii.String("my-bitbucket-creds"), &secretsManagerSecretOptions{
//   		jsonField: jsii.String("username"),
//   	}),
//   	password: awscdk.SecretValue.secretsManager(jsii.String("my-bitbucket-creds"), &secretsManagerSecretOptions{
//   		jsonField: jsii.String("password"),
//   	}),
//   })
//
type SecretsManagerSecretOptions struct {
	// The key of a JSON field to retrieve.
	//
	// This can only be used if the secret
	// stores a JSON object.
	JsonField *string `field:"optional" json:"jsonField" yaml:"jsonField"`
	// Specifies the unique identifier of the version of the secret you want to use.
	//
	// Can specify at most one of `versionId` and `versionStage`.
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
	// Specifies the secret version that you want to retrieve by the staging label attached to the version.
	//
	// Can specify at most one of `versionId` and `versionStage`.
	VersionStage *string `field:"optional" json:"versionStage" yaml:"versionStage"`
}


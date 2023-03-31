package awscodebuild


// `RegistryCredential` is a property of the [AWS CodeBuild Project Environment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html) property type that specifies information about credentials that provide access to a private Docker registry. When this is set:.
//
// - `imagePullCredentialsType` must be set to `SERVICE_ROLE` .
// - images cannot be curated or an Amazon ECR image.
//
// For more information, see [Private Registry with AWS Secrets Manager Sample for AWS CodeBuild](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-private-registry.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registryCredentialProperty := &registryCredentialProperty{
//   	credential: jsii.String("credential"),
//   	credentialProvider: jsii.String("credentialProvider"),
//   }
//
type CfnProject_RegistryCredentialProperty struct {
	// The Amazon Resource Name (ARN) or name of credentials created using AWS Secrets Manager .
	//
	// > The `credential` can use the name of the credentials only if they exist in your current AWS Region .
	Credential *string `field:"required" json:"credential" yaml:"credential"`
	// The service that created the credentials to access a private Docker registry.
	//
	// The valid value, SECRETS_MANAGER, is for AWS Secrets Manager .
	CredentialProvider *string `field:"required" json:"credentialProvider" yaml:"credentialProvider"`
}


package awsecs


// The `RepositoryCredentials` property specifies the repository credentials for private registry authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryCredentialsProperty := &repositoryCredentialsProperty{
//   	credentialsParameter: jsii.String("credentialsParameter"),
//   }
//
type CfnTaskDefinition_RepositoryCredentialsProperty struct {
	// The Amazon Resource Name (ARN) of the secret containing the private repository credentials.
	//
	// > When you use the Amazon ECS API, AWS CLI , or AWS SDK, if the secret exists in the same Region as the task that you're launching then you can use either the full ARN or the name of the secret. When you use the AWS Management Console, you must specify the full ARN of the secret.
	CredentialsParameter *string `field:"optional" json:"credentialsParameter" yaml:"credentialsParameter"`
}


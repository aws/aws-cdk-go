package awsecs


// The repository credentials for private registry authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryCredentialsProperty := &RepositoryCredentialsProperty{
//   	CredentialsParameter: jsii.String("credentialsParameter"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-repositorycredentials.html
//
type CfnTaskDefinition_RepositoryCredentialsProperty struct {
	// The Amazon Resource Name (ARN) of the secret containing the private repository credentials.
	//
	// > When you use the Amazon ECS API, AWS CLI , or AWS SDK, if the secret exists in the same Region as the task that you're launching then you can use either the full ARN or the name of the secret. When you use the AWS Management Console, you must specify the full ARN of the secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-repositorycredentials.html#cfn-ecs-taskdefinition-repositorycredentials-credentialsparameter
	//
	CredentialsParameter *string `field:"optional" json:"credentialsParameter" yaml:"credentialsParameter"`
}


package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryCredentialsProperty := &RepositoryCredentialsProperty{
//   	CredentialsParameter: jsii.String("credentialsParameter"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-repositorycredentials.html
//
type CfnDaemonTaskDefinition_RepositoryCredentialsProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-daemontaskdefinition-repositorycredentials.html#cfn-ecs-daemontaskdefinition-repositorycredentials-credentialsparameter
	//
	CredentialsParameter *string `field:"optional" json:"credentialsParameter" yaml:"credentialsParameter"`
}


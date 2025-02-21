package awsbatch


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-repositorycredentials.html
//
type CfnJobDefinition_RepositoryCredentialsProperty struct {
	// The Amazon Resource Name (ARN) of the secret containing the private repository credentials.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-repositorycredentials.html#cfn-batch-jobdefinition-repositorycredentials-credentialsparameter
	//
	CredentialsParameter *string `field:"required" json:"credentialsParameter" yaml:"credentialsParameter"`
}


package awssagemaker


// Specifies an authentication configuration for the private docker registry where your model image is hosted.
//
// Specify a value for this property only if you specified `Vpc` as the value for the `RepositoryAccessMode` field of the `ImageConfig` object that you passed to a call to `CreateModel` and the private Docker registry where the model image is hosted requires authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   repositoryAuthConfigProperty := &repositoryAuthConfigProperty{
//   	repositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   }
//
type CfnModel_RepositoryAuthConfigProperty struct {
	// The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted.
	//
	// For information about how to create an AWS Lambda function, see [Create a Lambda function with the console](https://docs.aws.amazon.com/lambda/latest/dg/getting-started-create-function.html) in the *AWS Lambda Developer Guide* .
	RepositoryCredentialsProviderArn *string `field:"required" json:"repositoryCredentialsProviderArn" yaml:"repositoryCredentialsProviderArn"`
}


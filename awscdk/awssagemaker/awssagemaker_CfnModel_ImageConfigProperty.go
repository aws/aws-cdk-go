package awssagemaker


// Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageConfigProperty := &imageConfigProperty{
//   	repositoryAccessMode: jsii.String("repositoryAccessMode"),
//
//   	// the properties below are optional
//   	repositoryAuthConfig: &repositoryAuthConfigProperty{
//   		repositoryCredentialsProviderArn: jsii.String("repositoryCredentialsProviderArn"),
//   	},
//   }
//
type CfnModel_ImageConfigProperty struct {
	// Set this to one of the following values:.
	//
	// - `Platform` - The model image is hosted in Amazon ECR.
	// - `Vpc` - The model image is hosted in a private Docker registry in your VPC.
	RepositoryAccessMode *string `field:"required" json:"repositoryAccessMode" yaml:"repositoryAccessMode"`
	// (Optional) Specifies an authentication configuration for the private docker registry where your model image is hosted.
	//
	// Specify a value for this property only if you specified `Vpc` as the value for the `RepositoryAccessMode` field, and the private Docker registry where the model image is hosted requires authentication.
	RepositoryAuthConfig interface{} `field:"optional" json:"repositoryAuthConfig" yaml:"repositoryAuthConfig"`
}


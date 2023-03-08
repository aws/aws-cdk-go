package awstransfer


// Required when `IdentityProviderType` is set to `AWS_DIRECTORY_SERVICE` or `API_GATEWAY` .
//
// Accepts an array containing all of the information required to use a directory in `AWS_DIRECTORY_SERVICE` or invoke a customer-supplied authentication API, including the API Gateway URL. Not required when `IdentityProviderType` is set to `SERVICE_MANAGED` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityProviderDetailsProperty := &identityProviderDetailsProperty{
//   	directoryId: jsii.String("directoryId"),
//   	function: jsii.String("function"),
//   	invocationRole: jsii.String("invocationRole"),
//   	url: jsii.String("url"),
//   }
//
type CfnServer_IdentityProviderDetailsProperty struct {
	// The identifier of the AWS Directory Service directory that you want to stop sharing.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// The ARN for a lambda function to use for the Identity provider.
	Function *string `field:"optional" json:"function" yaml:"function"`
	// Provides the type of `InvocationRole` used to authenticate the user account.
	InvocationRole *string `field:"optional" json:"invocationRole" yaml:"invocationRole"`
	// Provides the location of the service endpoint used to authenticate users.
	Url *string `field:"optional" json:"url" yaml:"url"`
}


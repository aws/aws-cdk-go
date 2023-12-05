package awstransfer


// Required when `IdentityProviderType` is set to `AWS_DIRECTORY_SERVICE` , `AWS _LAMBDA` or `API_GATEWAY` .
//
// Accepts an array containing all of the information required to use a directory in `AWS_DIRECTORY_SERVICE` or invoke a customer-supplied authentication API, including the API Gateway URL. Not required when `IdentityProviderType` is set to `SERVICE_MANAGED` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   identityProviderDetailsProperty := &IdentityProviderDetailsProperty{
//   	DirectoryId: jsii.String("directoryId"),
//   	Function: jsii.String("function"),
//   	InvocationRole: jsii.String("invocationRole"),
//   	SftpAuthenticationMethods: jsii.String("sftpAuthenticationMethods"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html
//
type CfnServer_IdentityProviderDetailsProperty struct {
	// The identifier of the AWS Directory Service directory that you want to use as your identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html#cfn-transfer-server-identityproviderdetails-directoryid
	//
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// The ARN for a Lambda function to use for the Identity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html#cfn-transfer-server-identityproviderdetails-function
	//
	Function *string `field:"optional" json:"function" yaml:"function"`
	// This parameter is only applicable if your `IdentityProviderType` is `API_GATEWAY` .
	//
	// Provides the type of `InvocationRole` used to authenticate the user account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html#cfn-transfer-server-identityproviderdetails-invocationrole
	//
	InvocationRole *string `field:"optional" json:"invocationRole" yaml:"invocationRole"`
	// For SFTP-enabled servers, and for custom identity providers *only* , you can specify whether to authenticate using a password, SSH key pair, or both.
	//
	// - `PASSWORD` - users must provide their password to connect.
	// - `PUBLIC_KEY` - users must provide their private key to connect.
	// - `PUBLIC_KEY_OR_PASSWORD` - users can authenticate with either their password or their key. This is the default value.
	// - `PUBLIC_KEY_AND_PASSWORD` - users must provide both their private key and their password to connect. The server checks the key first, and then if the key is valid, the system prompts for a password. If the private key provided does not match the public key that is stored, authentication fails.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html#cfn-transfer-server-identityproviderdetails-sftpauthenticationmethods
	//
	SftpAuthenticationMethods *string `field:"optional" json:"sftpAuthenticationMethods" yaml:"sftpAuthenticationMethods"`
	// Provides the location of the service endpoint used to authenticate users.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-identityproviderdetails.html#cfn-transfer-server-identityproviderdetails-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}


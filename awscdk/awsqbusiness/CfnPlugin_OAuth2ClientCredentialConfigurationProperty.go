package awsqbusiness


// Information about the OAuth 2.0 authentication credential/token used to configure a plugin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oAuth2ClientCredentialConfigurationProperty := &OAuth2ClientCredentialConfigurationProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	SecretArn: jsii.String("secretArn"),
//
//   	// the properties below are optional
//   	AuthorizationUrl: jsii.String("authorizationUrl"),
//   	TokenUrl: jsii.String("tokenUrl"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-oauth2clientcredentialconfiguration.html
//
type CfnPlugin_OAuth2ClientCredentialConfigurationProperty struct {
	// The ARN of an IAM role used by Amazon Q Business to access the OAuth 2.0 authentication credentials stored in a Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-oauth2clientcredentialconfiguration.html#cfn-qbusiness-plugin-oauth2clientcredentialconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The ARN of the Secrets Manager secret that stores the OAuth 2.0 credentials/token used for plugin configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-oauth2clientcredentialconfiguration.html#cfn-qbusiness-plugin-oauth2clientcredentialconfiguration-secretarn
	//
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// The redirect URL required by the OAuth 2.0 protocol for Amazon Q Business to authenticate a plugin user through a third party authentication server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-oauth2clientcredentialconfiguration.html#cfn-qbusiness-plugin-oauth2clientcredentialconfiguration-authorizationurl
	//
	AuthorizationUrl *string `field:"optional" json:"authorizationUrl" yaml:"authorizationUrl"`
	// The URL required by the OAuth 2.0 protocol to exchange an end user authorization code for an access token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-oauth2clientcredentialconfiguration.html#cfn-qbusiness-plugin-oauth2clientcredentialconfiguration-tokenurl
	//
	TokenUrl *string `field:"optional" json:"tokenUrl" yaml:"tokenUrl"`
}


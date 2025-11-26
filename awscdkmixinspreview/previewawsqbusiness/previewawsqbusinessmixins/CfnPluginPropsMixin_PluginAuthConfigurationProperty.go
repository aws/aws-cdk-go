package previewawsqbusinessmixins


// Authentication configuration information for an Amazon Q Business plugin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var noAuthConfiguration interface{}
//
//   pluginAuthConfigurationProperty := &PluginAuthConfigurationProperty{
//   	BasicAuthConfiguration: &BasicAuthConfigurationProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	NoAuthConfiguration: noAuthConfiguration,
//   	OAuth2ClientCredentialConfiguration: &OAuth2ClientCredentialConfigurationProperty{
//   		AuthorizationUrl: jsii.String("authorizationUrl"),
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   		TokenUrl: jsii.String("tokenUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html
//
type CfnPluginPropsMixin_PluginAuthConfigurationProperty struct {
	// Information about the basic authentication credentials used to configure a plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html#cfn-qbusiness-plugin-pluginauthconfiguration-basicauthconfiguration
	//
	BasicAuthConfiguration interface{} `field:"optional" json:"basicAuthConfiguration" yaml:"basicAuthConfiguration"`
	// Information about invoking a custom plugin without any authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html#cfn-qbusiness-plugin-pluginauthconfiguration-noauthconfiguration
	//
	NoAuthConfiguration interface{} `field:"optional" json:"noAuthConfiguration" yaml:"noAuthConfiguration"`
	// Information about the OAuth 2.0 authentication credential/token used to configure a plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html#cfn-qbusiness-plugin-pluginauthconfiguration-oauth2clientcredentialconfiguration
	//
	OAuth2ClientCredentialConfiguration interface{} `field:"optional" json:"oAuth2ClientCredentialConfiguration" yaml:"oAuth2ClientCredentialConfiguration"`
}


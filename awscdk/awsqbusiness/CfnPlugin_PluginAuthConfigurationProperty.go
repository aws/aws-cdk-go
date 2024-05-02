package awsqbusiness


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html
//
type CfnPlugin_PluginAuthConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html#cfn-qbusiness-plugin-pluginauthconfiguration-basicauthconfiguration
	//
	BasicAuthConfiguration interface{} `field:"optional" json:"basicAuthConfiguration" yaml:"basicAuthConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html#cfn-qbusiness-plugin-pluginauthconfiguration-noauthconfiguration
	//
	NoAuthConfiguration interface{} `field:"optional" json:"noAuthConfiguration" yaml:"noAuthConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-pluginauthconfiguration.html#cfn-qbusiness-plugin-pluginauthconfiguration-oauth2clientcredentialconfiguration
	//
	OAuth2ClientCredentialConfiguration interface{} `field:"optional" json:"oAuth2ClientCredentialConfiguration" yaml:"oAuth2ClientCredentialConfiguration"`
}


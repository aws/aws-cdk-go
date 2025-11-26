package previewawsqbusinessmixins


// Information about the basic authentication credentials used to configure a plugin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   basicAuthConfigurationProperty := &BasicAuthConfigurationProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-basicauthconfiguration.html
//
type CfnPluginPropsMixin_BasicAuthConfigurationProperty struct {
	// The ARN of an IAM role used by Amazon Q Business to access the basic authentication credentials stored in a Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-basicauthconfiguration.html#cfn-qbusiness-plugin-basicauthconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ARN of the Secrets Manager secret that stores the basic authentication credentials used for plugin configuration..
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-basicauthconfiguration.html#cfn-qbusiness-plugin-basicauthconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}


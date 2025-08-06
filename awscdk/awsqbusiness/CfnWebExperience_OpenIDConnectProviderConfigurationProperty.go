package awsqbusiness


// Information about the OIDC-compliant identity provider (IdP) used to authenticate end users of an Amazon Q Business web experience.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openIDConnectProviderConfigurationProperty := &OpenIDConnectProviderConfigurationProperty{
//   	SecretsArn: jsii.String("secretsArn"),
//   	SecretsRole: jsii.String("secretsRole"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-openidconnectproviderconfiguration.html
//
type CfnWebExperience_OpenIDConnectProviderConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of a Secrets Manager secret containing the OIDC client secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-openidconnectproviderconfiguration.html#cfn-qbusiness-webexperience-openidconnectproviderconfiguration-secretsarn
	//
	SecretsArn *string `field:"required" json:"secretsArn" yaml:"secretsArn"`
	// An IAM role with permissions to access AWS KMS to decrypt the Secrets Manager secret containing your OIDC client secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-webexperience-openidconnectproviderconfiguration.html#cfn-qbusiness-webexperience-openidconnectproviderconfiguration-secretsrole
	//
	SecretsRole *string `field:"required" json:"secretsRole" yaml:"secretsRole"`
}


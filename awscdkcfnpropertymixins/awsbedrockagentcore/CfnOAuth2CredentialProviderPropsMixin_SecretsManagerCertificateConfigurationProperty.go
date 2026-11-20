package awsbedrockagentcore


// A reference to a PEM-encoded private CA certificate stored in an AWS Secrets Manager secret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   secretsManagerCertificateConfigurationProperty := &SecretsManagerCertificateConfigurationProperty{
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-secretsmanagercertificateconfiguration.html
//
type CfnOAuth2CredentialProviderPropsMixin_SecretsManagerCertificateConfigurationProperty struct {
	// The ARN of the AWS Secrets Manager secret that contains the PEM-encoded certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-secretsmanagercertificateconfiguration.html#cfn-bedrockagentcore-oauth2credentialprovider-secretsmanagercertificateconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}


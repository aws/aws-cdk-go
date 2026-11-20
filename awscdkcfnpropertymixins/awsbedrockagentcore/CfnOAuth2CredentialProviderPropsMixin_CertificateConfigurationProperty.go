package awsbedrockagentcore


// A reference to a private CA certificate.
//
// Specify exactly one certificate source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   certificateConfigurationProperty := &CertificateConfigurationProperty{
//   	S3: &S3CertificateConfigurationProperty{
//   		BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   		Uri: jsii.String("uri"),
//   	},
//   	SecretsManager: &SecretsManagerCertificateConfigurationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-certificateconfiguration.html
//
type CfnOAuth2CredentialProviderPropsMixin_CertificateConfigurationProperty struct {
	// A reference to a PEM-encoded private CA certificate stored as an Amazon S3 object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-certificateconfiguration.html#cfn-bedrockagentcore-oauth2credentialprovider-certificateconfiguration-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// A reference to a PEM-encoded private CA certificate stored in an AWS Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-certificateconfiguration.html#cfn-bedrockagentcore-oauth2credentialprovider-certificateconfiguration-secretsmanager
	//
	SecretsManager interface{} `field:"optional" json:"secretsManager" yaml:"secretsManager"`
}


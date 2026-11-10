package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateConfigurationProperty := &CertificateConfigurationProperty{
//   	S3: &S3CertificateConfigurationProperty{
//   		Uri: jsii.String("uri"),
//
//   		// the properties below are optional
//   		BucketOwnerAccountId: jsii.String("bucketOwnerAccountId"),
//   	},
//   	SecretsManager: &SecretsManagerCertificateConfigurationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-certificateconfiguration.html
//
type CfnGatewayTarget_CertificateConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-certificateconfiguration.html#cfn-bedrockagentcore-gatewaytarget-certificateconfiguration-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-certificateconfiguration.html#cfn-bedrockagentcore-gatewaytarget-certificateconfiguration-secretsmanager
	//
	SecretsManager interface{} `field:"optional" json:"secretsManager" yaml:"secretsManager"`
}


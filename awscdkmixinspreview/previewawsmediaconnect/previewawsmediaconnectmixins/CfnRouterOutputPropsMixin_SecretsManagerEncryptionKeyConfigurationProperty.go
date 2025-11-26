package previewawsmediaconnectmixins


// The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   secretsManagerEncryptionKeyConfigurationProperty := &SecretsManagerEncryptionKeyConfigurationProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-secretsmanagerencryptionkeyconfiguration.html
//
type CfnRouterOutputPropsMixin_SecretsManagerEncryptionKeyConfigurationProperty struct {
	// The ARN of the IAM role assumed by MediaConnect to access the AWS Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-secretsmanagerencryptionkeyconfiguration.html#cfn-mediaconnect-routeroutput-secretsmanagerencryptionkeyconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ARN of the AWS Secrets Manager secret used for transit encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-secretsmanagerencryptionkeyconfiguration.html#cfn-mediaconnect-routeroutput-secretsmanagerencryptionkeyconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}


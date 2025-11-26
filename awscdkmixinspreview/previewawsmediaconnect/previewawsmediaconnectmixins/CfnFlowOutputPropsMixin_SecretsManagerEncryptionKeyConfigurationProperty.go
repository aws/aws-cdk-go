package previewawsmediaconnectmixins


// The configuration settings for transit encryption of a flow output using AWS Secrets Manager, including the secret ARN and role ARN.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-secretsmanagerencryptionkeyconfiguration.html
//
type CfnFlowOutputPropsMixin_SecretsManagerEncryptionKeyConfigurationProperty struct {
	// The ARN of the IAM role used for transit encryption to the router input using AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-secretsmanagerencryptionkeyconfiguration.html#cfn-mediaconnect-flowoutput-secretsmanagerencryptionkeyconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ARN of the AWS Secrets Manager secret used for transit encryption to the router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowoutput-secretsmanagerencryptionkeyconfiguration.html#cfn-mediaconnect-flowoutput-secretsmanagerencryptionkeyconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}


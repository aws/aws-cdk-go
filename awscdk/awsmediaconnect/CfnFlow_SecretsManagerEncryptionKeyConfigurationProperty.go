package awsmediaconnect


// The configuration settings for transit encryption of a flow source using AWS Secrets Manager, including the secret ARN and role ARN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretsManagerEncryptionKeyConfigurationProperty := &SecretsManagerEncryptionKeyConfigurationProperty{
//   	RoleArn: jsii.String("roleArn"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-secretsmanagerencryptionkeyconfiguration.html
//
type CfnFlow_SecretsManagerEncryptionKeyConfigurationProperty struct {
	// The ARN of the IAM role used for transit encryption from the router output using AWS Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-secretsmanagerencryptionkeyconfiguration.html#cfn-mediaconnect-flow-secretsmanagerencryptionkeyconfiguration-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The ARN of the AWS Secrets Manager secret used for transit encryption from the router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-secretsmanagerencryptionkeyconfiguration.html#cfn-mediaconnect-flow-secretsmanagerencryptionkeyconfiguration-secretarn
	//
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
}


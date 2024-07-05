package awskinesisfirehose


// The structure that defines how Firehose accesses the secret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretsManagerConfigurationProperty := &SecretsManagerConfigurationProperty{
//   	Enabled: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	RoleArn: jsii.String("roleArn"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-secretsmanagerconfiguration.html
//
type CfnDeliveryStream_SecretsManagerConfigurationProperty struct {
	// Specifies whether you want to use the the secrets manager feature.
	//
	// When set as `True` the secrets manager configuration overwrites the existing secrets in the destination configuration. When it's set to `False` Firehose falls back to the credentials in the destination configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-secretsmanagerconfiguration.html#cfn-kinesisfirehose-deliverystream-secretsmanagerconfiguration-enabled
	//
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Specifies the role that Firehose assumes when calling the Secrets Manager API operation.
	//
	// When you provide the role, it overrides any destination specific role defined in the destination configuration. If you do not provide the then we use the destination specific role. This parameter is required for Splunk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-secretsmanagerconfiguration.html#cfn-kinesisfirehose-deliverystream-secretsmanagerconfiguration-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ARN of the secret that stores your credentials.
	//
	// It must be in the same region as the Firehose stream and the role. The secret ARN can reside in a different account than the delivery stream and role as Firehose supports cross-account secret access. This parameter is required when *Enabled* is set to `True` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-secretsmanagerconfiguration.html#cfn-kinesisfirehose-deliverystream-secretsmanagerconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}


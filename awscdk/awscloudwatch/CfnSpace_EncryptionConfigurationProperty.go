package awscloudwatch


// The encryption configuration for the space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	EncryptionStrategy: jsii.String("encryptionStrategy"),
//
//   	// the properties below are optional
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-space-encryptionconfiguration.html
//
type CfnSpace_EncryptionConfigurationProperty struct {
	// Which kind of key to use.
	//
	// Required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-space-encryptionconfiguration.html#cfn-cloudwatch-space-encryptionconfiguration-encryptionstrategy
	//
	EncryptionStrategy *string `field:"required" json:"encryptionStrategy" yaml:"encryptionStrategy"`
	// Customer managed KMS key ARN.
	//
	// Required when encryptionStrategy is CUSTOMER_MANAGED, and must be omitted when it is AWS_OWNED. Must be a symmetric ENCRYPT_DECRYPT key in the caller's account and region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-space-encryptionconfiguration.html#cfn-cloudwatch-space-encryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}


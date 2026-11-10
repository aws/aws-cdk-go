package awscloudwatch


// The encryption configuration for the space.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	EncryptionStrategy: jsii.String("encryptionStrategy"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-space-encryptionconfiguration.html
//
type CfnSpacePropsMixin_EncryptionConfigurationProperty struct {
	// Whether the space is encrypted with an AWS owned key or a customer managed key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-space-encryptionconfiguration.html#cfn-cloudwatch-space-encryptionconfiguration-encryptionstrategy
	//
	EncryptionStrategy *string `field:"optional" json:"encryptionStrategy" yaml:"encryptionStrategy"`
	// The Amazon Resource Name (ARN) of the customer managed KMS key.
	//
	// Required when encryptionStrategy is CUSTOMER_MANAGED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudwatch-space-encryptionconfiguration.html#cfn-cloudwatch-space-encryptionconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}


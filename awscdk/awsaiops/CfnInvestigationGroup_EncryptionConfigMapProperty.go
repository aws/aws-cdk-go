package awsaiops


// Use this structure if you want to use a customer managed AWS KMS key to encrypt your investigation data.
//
// If you omit this parameter, CloudWatch investigations will use an AWS key to encrypt the data. For more information, see [Encryption of investigation data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Investigations-Security.html#Investigations-KMS) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigMapProperty := &EncryptionConfigMapProperty{
//   	EncryptionConfigurationType: jsii.String("encryptionConfigurationType"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-encryptionconfigmap.html
//
type CfnInvestigationGroup_EncryptionConfigMapProperty struct {
	// Displays whether investigation data is encrypted by a customer managed key or an AWS owned key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-encryptionconfigmap.html#cfn-aiops-investigationgroup-encryptionconfigmap-encryptionconfigurationtype
	//
	EncryptionConfigurationType *string `field:"optional" json:"encryptionConfigurationType" yaml:"encryptionConfigurationType"`
	// If the investigation group uses a customer managed key for encryption, this field displays the ID of that key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-encryptionconfigmap.html#cfn-aiops-investigationgroup-encryptionconfigmap-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}


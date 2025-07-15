package awsaiops


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-encryptionconfigmap.html#cfn-aiops-investigationgroup-encryptionconfigmap-encryptionconfigurationtype
	//
	EncryptionConfigurationType *string `field:"optional" json:"encryptionConfigurationType" yaml:"encryptionConfigurationType"`
	// If the investigation group uses a customer managed key for encryption, this field displays the ID of that key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-encryptionconfigmap.html#cfn-aiops-investigationgroup-encryptionconfigmap-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}


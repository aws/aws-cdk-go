package awsbedrock


// KMS Encryption Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//
//   	// the properties below are optional
//   	KmsEncryptionContext: map[string]*string{
//   		"kmsEncryptionContextKey": jsii.String("kmsEncryptionContext"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationlibrary-encryptionconfiguration.html
//
type CfnDataAutomationLibrary_EncryptionConfigurationProperty struct {
	// KMS Key Identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationlibrary-encryptionconfiguration.html#cfn-bedrock-dataautomationlibrary-encryptionconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"required" json:"kmsKeyId" yaml:"kmsKeyId"`
	// KMS Encryption Context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationlibrary-encryptionconfiguration.html#cfn-bedrock-dataautomationlibrary-encryptionconfiguration-kmsencryptioncontext
	//
	KmsEncryptionContext interface{} `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
}


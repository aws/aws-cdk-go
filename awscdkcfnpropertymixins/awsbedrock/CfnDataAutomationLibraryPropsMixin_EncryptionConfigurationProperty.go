package awsbedrock


// KMS Encryption Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   encryptionConfigurationProperty := &EncryptionConfigurationProperty{
//   	KmsEncryptionContext: map[string]*string{
//   		"kmsEncryptionContextKey": jsii.String("kmsEncryptionContext"),
//   	},
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationlibrary-encryptionconfiguration.html
//
type CfnDataAutomationLibraryPropsMixin_EncryptionConfigurationProperty struct {
	// KMS Encryption Context.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationlibrary-encryptionconfiguration.html#cfn-bedrock-dataautomationlibrary-encryptionconfiguration-kmsencryptioncontext
	//
	KmsEncryptionContext interface{} `field:"optional" json:"kmsEncryptionContext" yaml:"kmsEncryptionContext"`
	// KMS Key Identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationlibrary-encryptionconfiguration.html#cfn-bedrock-dataautomationlibrary-encryptionconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}


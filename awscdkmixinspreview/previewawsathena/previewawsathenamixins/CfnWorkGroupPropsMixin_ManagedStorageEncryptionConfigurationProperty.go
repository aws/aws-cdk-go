package previewawsathenamixins


// Indicates the encryption configuration for Athena Managed Storage.
//
// If not setting this field, Managed Storage will encrypt the query results with Athena's encryption key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedStorageEncryptionConfigurationProperty := &ManagedStorageEncryptionConfigurationProperty{
//   	KmsKey: jsii.String("kmsKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-managedstorageencryptionconfiguration.html
//
type CfnWorkGroupPropsMixin_ManagedStorageEncryptionConfigurationProperty struct {
	// For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-managedstorageencryptionconfiguration.html#cfn-athena-workgroup-managedstorageencryptionconfiguration-kmskey
	//
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}


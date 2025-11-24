package mixinsawshealthlake


// The customer-managed-key(CMK) used when creating a Data Store.
//
// If a customer owned key is not specified, an Amazon owned key will be used for encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kmsEncryptionConfigProperty := &KmsEncryptionConfigProperty{
//   	CmkType: jsii.String("cmkType"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-kmsencryptionconfig.html
//
type CfnFHIRDatastorePropsMixin_KmsEncryptionConfigProperty struct {
	// The type of customer-managed-key(CMK) used for encryption.
	//
	// The two types of supported CMKs are customer owned CMKs and Amazon owned CMKs. For more information on CMK types, see [KmsEncryptionConfig](https://docs.aws.amazon.com/healthlake/latest/APIReference/API_KmsEncryptionConfig.html#HealthLake-Type-KmsEncryptionConfig-CmkType) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-kmsencryptionconfig.html#cfn-healthlake-fhirdatastore-kmsencryptionconfig-cmktype
	//
	CmkType *string `field:"optional" json:"cmkType" yaml:"cmkType"`
	// The Key Management Service (KMS) encryption key id/alias used to encrypt the data store contents at rest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-kmsencryptionconfig.html#cfn-healthlake-fhirdatastore-kmsencryptionconfig-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}


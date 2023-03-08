package awshealthlake


// The customer-managed-key(CMK) used when creating a Data Store.
//
// If a customer owned key is not specified, an Amazon owned key will be used for encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kmsEncryptionConfigProperty := &kmsEncryptionConfigProperty{
//   	cmkType: jsii.String("cmkType"),
//
//   	// the properties below are optional
//   	kmsKeyId: jsii.String("kmsKeyId"),
//   }
//
type CfnFHIRDatastore_KmsEncryptionConfigProperty struct {
	// The type of customer-managed-key(CMK) used for encryption.
	//
	// The two types of supported CMKs are customer owned CMKs and Amazon owned CMKs. For more information on CMK types, see [KmsEncryptionConfig](https://docs.aws.amazon.com/healthlake/latest/APIReference/API_KmsEncryptionConfig.html#HealthLake-Type-KmsEncryptionConfig-CmkType) .
	CmkType *string `field:"required" json:"cmkType" yaml:"cmkType"`
	// The KMS encryption key id/alias used to encrypt the Data Store contents at rest.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}


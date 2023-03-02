package awshealthlake


// The server-side encryption key configuration for a customer provided encryption key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sseConfigurationProperty := &sseConfigurationProperty{
//   	kmsEncryptionConfig: &kmsEncryptionConfigProperty{
//   		cmkType: jsii.String("cmkType"),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   }
//
type CfnFHIRDatastore_SseConfigurationProperty struct {
	// The server-side encryption key configuration for a customer provided encryption key (CMK).
	KmsEncryptionConfig interface{} `field:"required" json:"kmsEncryptionConfig" yaml:"kmsEncryptionConfig"`
}


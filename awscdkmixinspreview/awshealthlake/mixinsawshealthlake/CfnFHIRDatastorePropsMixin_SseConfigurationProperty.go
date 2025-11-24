package mixinsawshealthlake


// The server-side encryption key configuration for a customer-provided encryption key.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sseConfigurationProperty := &SseConfigurationProperty{
//   	KmsEncryptionConfig: &KmsEncryptionConfigProperty{
//   		CmkType: jsii.String("cmkType"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-sseconfiguration.html
//
type CfnFHIRDatastorePropsMixin_SseConfigurationProperty struct {
	// The server-side encryption key configuration for a customer provided encryption key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-healthlake-fhirdatastore-sseconfiguration.html#cfn-healthlake-fhirdatastore-sseconfiguration-kmsencryptionconfig
	//
	KmsEncryptionConfig interface{} `field:"optional" json:"kmsEncryptionConfig" yaml:"kmsEncryptionConfig"`
}


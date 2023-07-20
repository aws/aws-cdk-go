package awsathena


// Specifies the KMS key that is used to encrypt the user's data stores in Athena.
//
// This setting does not apply to Athena SQL workgroups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customerContentEncryptionConfigurationProperty := &CustomerContentEncryptionConfigurationProperty{
//   	KmsKey: jsii.String("kmsKey"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-customercontentencryptionconfiguration.html
//
type CfnWorkGroup_CustomerContentEncryptionConfigurationProperty struct {
	// The KMS key that is used to encrypt the user's data stores in Athena.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-customercontentencryptionconfiguration.html#cfn-athena-workgroup-customercontentencryptionconfiguration-kmskey
	//
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}


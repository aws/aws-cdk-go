package previewawsvoiceidmixins


// > End of support notice: On May 20, 2026, AWS will end support for Amazon Connect Voice ID.
//
// After May 20, 2026, you will no longer be able to access Voice ID on the Amazon Connect console, access Voice ID features on the Amazon Connect admin website or Contact Control Panel, or access Voice ID resources. For more information, visit [Amazon Connect Voice ID end of support](https://docs.aws.amazon.com/connect/latest/adminguide/amazonconnect-voiceid-end-of-support.html) .
//
// The configuration containing information about the customer managed key used for encrypting customer data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   serverSideEncryptionConfigurationProperty := &ServerSideEncryptionConfigurationProperty{
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-voiceid-domain-serversideencryptionconfiguration.html
//
type CfnDomainPropsMixin_ServerSideEncryptionConfigurationProperty struct {
	// The identifier of the KMS key to use to encrypt data stored by Voice ID.
	//
	// Voice ID doesn't support asymmetric customer managed keys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-voiceid-domain-serversideencryptionconfiguration.html#cfn-voiceid-domain-serversideencryptionconfiguration-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}


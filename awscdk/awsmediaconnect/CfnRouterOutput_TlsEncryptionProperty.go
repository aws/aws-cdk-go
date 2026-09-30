package awsmediaconnect


// The Transport Layer Security (TLS) encryption settings used to establish a secure connection to a destination.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var public_ interface{}
//
//   tlsEncryptionProperty := &TlsEncryptionProperty{
//   	EncryptionConfiguration: &TlsEncryptionConfigurationProperty{
//   		Public: public_,
//   	},
//
//   	// the properties below are optional
//   	EncryptionType: jsii.String("encryptionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-tlsencryption.html
//
type CfnRouterOutput_TlsEncryptionProperty struct {
	// The configuration settings for TLS encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-tlsencryption.html#cfn-mediaconnect-routeroutput-tlsencryption-encryptionconfiguration
	//
	EncryptionConfiguration interface{} `field:"required" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-tlsencryption.html#cfn-mediaconnect-routeroutput-tlsencryption-encryptiontype
	//
	EncryptionType *string `field:"optional" json:"encryptionType" yaml:"encryptionType"`
}


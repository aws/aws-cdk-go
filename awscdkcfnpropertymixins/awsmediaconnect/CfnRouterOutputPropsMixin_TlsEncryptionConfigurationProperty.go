package awsmediaconnect


// The configuration settings for TLS encryption.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var public_ interface{}
//
//   tlsEncryptionConfigurationProperty := &TlsEncryptionConfigurationProperty{
//   	Public: public_,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-tlsencryptionconfiguration.html
//
type CfnRouterOutputPropsMixin_TlsEncryptionConfigurationProperty struct {
	// The TLS encryption configuration for destinations that present a certificate from a publicly trusted certificate authority.
	//
	// This type does not require any additional settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-tlsencryptionconfiguration.html#cfn-mediaconnect-routeroutput-tlsencryptionconfiguration-public
	//
	Public interface{} `field:"optional" json:"public" yaml:"public"`
}


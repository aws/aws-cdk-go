package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var automatic interface{}
//
//   endpointEncryptionConfigurationProperty := &EndpointEncryptionConfigurationProperty{
//   	Automatic: automatic,
//   	SecretsManager: &SecretsManagerEndpointEncryptionConfigurationProperty{
//   		EncryptionSecret: &EndpointEncryptionSecretConfigurationProperty{
//   			SecretArn: jsii.String("secretArn"),
//   		},
//   		EncryptionSecretType: jsii.String("encryptionSecretType"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-endpointencryptionconfiguration.html
//
type CfnRouterOutputPropsMixin_EndpointEncryptionConfigurationProperty struct {
	// The configuration for automatic endpoint encryption.
	//
	// In this mode, MediaConnect determines the encryption key when the endpoint connects. This type does not require any additional settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-endpointencryptionconfiguration.html#cfn-mediaconnect-routeroutput-endpointencryptionconfiguration-automatic
	//
	Automatic interface{} `field:"optional" json:"automatic" yaml:"automatic"`
	// The endpoint encryption configuration for connections that use a key or passphrase that you store in Secrets Manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-endpointencryptionconfiguration.html#cfn-mediaconnect-routeroutput-endpointencryptionconfiguration-secretsmanager
	//
	SecretsManager interface{} `field:"optional" json:"secretsManager" yaml:"secretsManager"`
}


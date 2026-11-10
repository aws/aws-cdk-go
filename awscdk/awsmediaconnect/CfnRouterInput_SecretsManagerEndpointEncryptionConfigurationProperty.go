package awsmediaconnect


// The endpoint encryption configuration for connections that use a key or passphrase that you store in Secrets Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   secretsManagerEndpointEncryptionConfigurationProperty := &SecretsManagerEndpointEncryptionConfigurationProperty{
//   	EncryptionSecret: &EndpointEncryptionSecretConfigurationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//
//   	// the properties below are optional
//   	EncryptionSecretType: jsii.String("encryptionSecretType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-secretsmanagerendpointencryptionconfiguration.html
//
type CfnRouterInput_SecretsManagerEndpointEncryptionConfigurationProperty struct {
	// The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-secretsmanagerendpointencryptionconfiguration.html#cfn-mediaconnect-routerinput-secretsmanagerendpointencryptionconfiguration-encryptionsecret
	//
	EncryptionSecret interface{} `field:"required" json:"encryptionSecret" yaml:"encryptionSecret"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-secretsmanagerendpointencryptionconfiguration.html#cfn-mediaconnect-routerinput-secretsmanagerendpointencryptionconfiguration-encryptionsecrettype
	//
	EncryptionSecretType *string `field:"optional" json:"encryptionSecretType" yaml:"encryptionSecretType"`
}


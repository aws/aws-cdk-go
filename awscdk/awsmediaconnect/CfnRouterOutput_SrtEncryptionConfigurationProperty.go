package awsmediaconnect


// Contains the configuration settings for encrypting SRT streams, including the encryption key details and encryption parameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   srtEncryptionConfigurationProperty := &SrtEncryptionConfigurationProperty{
//   	EncryptionKey: &SecretsManagerEncryptionKeyConfigurationProperty{
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtencryptionconfiguration.html
//
type CfnRouterOutput_SrtEncryptionConfigurationProperty struct {
	// The configuration settings for transit encryption using AWS Secrets Manager, including the secret ARN and role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-srtencryptionconfiguration.html#cfn-mediaconnect-routeroutput-srtencryptionconfiguration-encryptionkey
	//
	EncryptionKey interface{} `field:"required" json:"encryptionKey" yaml:"encryptionKey"`
}


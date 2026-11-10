package awsmediaconnect


// The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   endpointEncryptionSecretConfigurationProperty := &EndpointEncryptionSecretConfigurationProperty{
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-endpointencryptionsecretconfiguration.html
//
type CfnDestinationEndpointPropsMixin_EndpointEncryptionSecretConfigurationProperty struct {
	// The ARN of the Secrets Manager secret that holds the key or passphrase.
	//
	// The secret must be encrypted with a customer managed Amazon Web Services KMS key that is enabled and symmetric; it cannot use the Amazon Web Services managed key aws/secretsmanager. The secret must be in the same Amazon Web Services account and Region as the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-endpointencryptionsecretconfiguration.html#cfn-mediaconnect-destinationendpoint-endpointencryptionsecretconfiguration-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}


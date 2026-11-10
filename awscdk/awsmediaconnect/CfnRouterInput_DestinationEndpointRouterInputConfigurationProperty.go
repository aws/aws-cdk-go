package awsmediaconnect


// Configuration settings for a router input that is represented by a destination endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var automatic interface{}
//
//   destinationEndpointRouterInputConfigurationProperty := &DestinationEndpointRouterInputConfigurationProperty{
//   	EndpointLifecycleManagement: jsii.String("endpointLifecycleManagement"),
//   	SourceTransitDecryption: &EndpointEncryptionConfigurationProperty{
//   		Automatic: automatic,
//   		SecretsManager: &SecretsManagerEndpointEncryptionConfigurationProperty{
//   			EncryptionSecret: &EndpointEncryptionSecretConfigurationProperty{
//   				SecretArn: jsii.String("secretArn"),
//   			},
//
//   			// the properties below are optional
//   			EncryptionSecretType: jsii.String("encryptionSecretType"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	SourceTransitDecryptionType: jsii.String("sourceTransitDecryptionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-destinationendpointrouterinputconfiguration.html
//
type CfnRouterInput_DestinationEndpointRouterInputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-destinationendpointrouterinputconfiguration.html#cfn-mediaconnect-routerinput-destinationendpointrouterinputconfiguration-endpointlifecyclemanagement
	//
	EndpointLifecycleManagement *string `field:"required" json:"endpointLifecycleManagement" yaml:"endpointLifecycleManagement"`
	// The encryption settings for an endpoint connection.
	//
	// This is a union type, so exactly one of its members is set: secretsManager when you store the encryption key or passphrase in Secrets Manager, or automatic when you let MediaConnect determine the encryption key when the endpoint connects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-destinationendpointrouterinputconfiguration.html#cfn-mediaconnect-routerinput-destinationendpointrouterinputconfiguration-sourcetransitdecryption
	//
	SourceTransitDecryption interface{} `field:"required" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routerinput-destinationendpointrouterinputconfiguration.html#cfn-mediaconnect-routerinput-destinationendpointrouterinputconfiguration-sourcetransitdecryptiontype
	//
	SourceTransitDecryptionType *string `field:"optional" json:"sourceTransitDecryptionType" yaml:"sourceTransitDecryptionType"`
}


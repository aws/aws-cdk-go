package awsmediaconnect


// Configuration settings for a router output that is represented by a source endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var automatic interface{}
//
//   sourceEndpointRouterOutputConfigurationProperty := &SourceEndpointRouterOutputConfigurationProperty{
//   	DestinationTransitEncryption: &EndpointEncryptionConfigurationProperty{
//   		Automatic: automatic,
//   		SecretsManager: &SecretsManagerEndpointEncryptionConfigurationProperty{
//   			EncryptionSecret: &EndpointEncryptionSecretConfigurationProperty{
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   			EncryptionSecretType: jsii.String("encryptionSecretType"),
//   		},
//   	},
//   	DestinationTransitEncryptionType: jsii.String("destinationTransitEncryptionType"),
//   	EndpointLifecycleManagement: jsii.String("endpointLifecycleManagement"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-sourceendpointrouteroutputconfiguration.html
//
type CfnRouterOutputPropsMixin_SourceEndpointRouterOutputConfigurationProperty struct {
	// The encryption settings for an endpoint connection.
	//
	// This is a union type, so exactly one of its members is set: secretsManager when you store the encryption key or passphrase in Secrets Manager, or automatic when you let MediaConnect determine the encryption key when the endpoint connects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-sourceendpointrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-sourceendpointrouteroutputconfiguration-destinationtransitencryption
	//
	DestinationTransitEncryption interface{} `field:"optional" json:"destinationTransitEncryption" yaml:"destinationTransitEncryption"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-sourceendpointrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-sourceendpointrouteroutputconfiguration-destinationtransitencryptiontype
	//
	DestinationTransitEncryptionType *string `field:"optional" json:"destinationTransitEncryptionType" yaml:"destinationTransitEncryptionType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routeroutput-sourceendpointrouteroutputconfiguration.html#cfn-mediaconnect-routeroutput-sourceendpointrouteroutputconfiguration-endpointlifecyclemanagement
	//
	EndpointLifecycleManagement *string `field:"optional" json:"endpointLifecycleManagement" yaml:"endpointLifecycleManagement"`
}


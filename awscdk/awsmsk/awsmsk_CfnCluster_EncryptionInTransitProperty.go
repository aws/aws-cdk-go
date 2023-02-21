package awsmsk


// The settings for encrypting data in transit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptionInTransitProperty := &encryptionInTransitProperty{
//   	clientBroker: jsii.String("clientBroker"),
//   	inCluster: jsii.Boolean(false),
//   }
//
type CfnCluster_EncryptionInTransitProperty struct {
	// Indicates the encryption setting for data in transit between clients and brokers. The following are the possible values.
	//
	// - `TLS` means that client-broker communication is enabled with TLS only.
	// - `TLS_PLAINTEXT` means that client-broker communication is enabled for both TLS-encrypted, as well as plain text data.
	// - `PLAINTEXT` means that client-broker communication is enabled in plain text only.
	//
	// The default value is `TLS` .
	ClientBroker *string `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// When set to true, it indicates that data communication among the broker nodes of the cluster is encrypted.
	//
	// When set to false, the communication happens in plain text. The default value is true.
	InCluster interface{} `field:"optional" json:"inCluster" yaml:"inCluster"`
}


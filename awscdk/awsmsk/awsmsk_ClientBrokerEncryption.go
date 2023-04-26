package awsmsk


// Indicates the encryption setting for data in transit between clients and brokers.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &ClusterProps{
//   	ClusterName: jsii.String("myCluster"),
//   	KafkaVersion: msk.KafkaVersion_V2_8_1(),
//   	Vpc: Vpc,
//   	EncryptionInTransit: &EncryptionInTransitConfig{
//   		ClientBroker: msk.ClientBrokerEncryption_TLS,
//   	},
//   	ClientAuthentication: msk.ClientAuthentication_Sasl(&SaslAuthProps{
//   		Scram: jsii.Boolean(true),
//   	}),
//   })
//
// Experimental.
type ClientBrokerEncryption string

const (
	// TLS means that client-broker communication is enabled with TLS only.
	// Experimental.
	ClientBrokerEncryption_TLS ClientBrokerEncryption = "TLS"
	// TLS_PLAINTEXT means that client-broker communication is enabled for both TLS-encrypted, as well as plaintext data.
	// Experimental.
	ClientBrokerEncryption_TLS_PLAINTEXT ClientBrokerEncryption = "TLS_PLAINTEXT"
	// PLAINTEXT means that client-broker communication is enabled in plaintext only.
	// Experimental.
	ClientBrokerEncryption_PLAINTEXT ClientBrokerEncryption = "PLAINTEXT"
)


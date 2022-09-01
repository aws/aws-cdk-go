// The CDK Construct Library for AWS::MSK
package awscdkmskalpha


// Indicates the encryption setting for data in transit between clients and brokers.
//
// Example:
//   var vpc vpc
//
//   cluster := msk.NewCluster(this, jsii.String("cluster"), &clusterProps{
//   	clusterName: jsii.String("myCluster"),
//   	kafkaVersion: msk.kafkaVersion_V2_8_1(),
//   	vpc: vpc,
//   	encryptionInTransit: &encryptionInTransitConfig{
//   		clientBroker: msk.clientBrokerEncryption_TLS,
//   	},
//   	clientAuthentication: msk.clientAuthentication.sasl(&saslAuthProps{
//   		scram: jsii.Boolean(true),
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


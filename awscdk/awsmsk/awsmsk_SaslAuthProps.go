package awsmsk

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// SASL authentication properties.
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
type SaslAuthProps struct {
	// Enable IAM access control.
	// Experimental.
	Iam *bool `field:"optional" json:"iam" yaml:"iam"`
	// KMS Key to encrypt SASL/SCRAM secrets.
	//
	// You must use a customer master key (CMK) when creating users in secrets manager.
	// You cannot use a Secret with Amazon MSK that uses the default Secrets Manager encryption key.
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Enable SASL/SCRAM authentication.
	// Experimental.
	Scram *bool `field:"optional" json:"scram" yaml:"scram"`
}


package awscdkmskalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// SASL authentication properties.
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
type SaslAuthProps struct {
	// Enable IAM access control.
	// Default: false.
	//
	// Experimental.
	Iam *bool `field:"optional" json:"iam" yaml:"iam"`
	// KMS Key to encrypt SASL/SCRAM secrets.
	//
	// You must use a customer master key (CMK) when creating users in secrets manager.
	// You cannot use a Secret with Amazon MSK that uses the default Secrets Manager encryption key.
	// Default: - CMK will be created with alias msk/{clusterName}/sasl/scram.
	//
	// Experimental.
	Key awskms.IKey `field:"optional" json:"key" yaml:"key"`
	// Enable SASL/SCRAM authentication.
	// Default: false.
	//
	// Experimental.
	Scram *bool `field:"optional" json:"scram" yaml:"scram"`
}


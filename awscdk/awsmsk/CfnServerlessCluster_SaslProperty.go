package awsmsk


// Details for client authentication using SASL.
//
// To turn on SASL, you must also turn on `EncryptionInTransit` by setting `inCluster` to true. You must set `clientBroker` to either `TLS` or `TLS_PLAINTEXT` . If you choose `TLS_PLAINTEXT` , then you must also set `unauthenticated` to true.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   saslProperty := &SaslProperty{
//   	Iam: &IamProperty{
//   		Enabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnServerlessCluster_SaslProperty struct {
	// Details for client authentication using IAM.
	Iam interface{} `field:"required" json:"iam" yaml:"iam"`
}


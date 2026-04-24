package awsmsk


// Details of the client authentication used by the Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterClientAuthenticationProperty := &KafkaClusterClientAuthenticationProperty{
//   	SaslScram: &KafkaClusterSaslScramAuthenticationProperty{
//   		Mechanism: jsii.String("mechanism"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusterclientauthentication.html
//
type CfnReplicator_KafkaClusterClientAuthenticationProperty struct {
	// Details for SASL/SCRAM client authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusterclientauthentication.html#cfn-msk-replicator-kafkaclusterclientauthentication-saslscram
	//
	SaslScram interface{} `field:"required" json:"saslScram" yaml:"saslScram"`
}


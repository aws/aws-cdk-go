package awsmsk


// Details for SASL/SCRAM client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   kafkaClusterSaslScramAuthenticationProperty := &KafkaClusterSaslScramAuthenticationProperty{
//   	Mechanism: jsii.String("mechanism"),
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersaslscramauthentication.html
//
type CfnReplicatorPropsMixin_KafkaClusterSaslScramAuthenticationProperty struct {
	// The SASL/SCRAM authentication mechanism.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersaslscramauthentication.html#cfn-msk-replicator-kafkaclustersaslscramauthentication-mechanism
	//
	Mechanism *string `field:"optional" json:"mechanism" yaml:"mechanism"`
	// The Amazon Resource Name (ARN) of the Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersaslscramauthentication.html#cfn-msk-replicator-kafkaclustersaslscramauthentication-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}


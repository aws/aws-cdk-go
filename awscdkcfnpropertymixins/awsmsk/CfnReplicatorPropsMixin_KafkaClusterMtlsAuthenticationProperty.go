package awsmsk


// Details for mTLS client authentication.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   kafkaClusterMtlsAuthenticationProperty := &KafkaClusterMtlsAuthenticationProperty{
//   	SecretArn: jsii.String("secretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustermtlsauthentication.html
//
type CfnReplicatorPropsMixin_KafkaClusterMtlsAuthenticationProperty struct {
	// The Amazon Resource Name (ARN) of the Secrets Manager secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustermtlsauthentication.html#cfn-msk-replicator-kafkaclustermtlsauthentication-secretarn
	//
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}


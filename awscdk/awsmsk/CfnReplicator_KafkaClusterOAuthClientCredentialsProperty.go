package awsmsk


// Details for SASL/OAUTHBEARER using the standard client_credentials grant.
//
// The referenced secret must contain client_id and client_secret.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterOAuthClientCredentialsProperty := &KafkaClusterOAuthClientCredentialsProperty{
//   	TokenRequestSecretArn: jsii.String("tokenRequestSecretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusteroauthclientcredentials.html
//
type CfnReplicator_KafkaClusterOAuthClientCredentialsProperty struct {
	// Secrets Manager ARN of the secret containing the client_id and client_secret used to obtain an OAuth Bearer token via the client_credentials grant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusteroauthclientcredentials.html#cfn-msk-replicator-kafkaclusteroauthclientcredentials-tokenrequestsecretarn
	//
	TokenRequestSecretArn *string `field:"required" json:"tokenRequestSecretArn" yaml:"tokenRequestSecretArn"`
}


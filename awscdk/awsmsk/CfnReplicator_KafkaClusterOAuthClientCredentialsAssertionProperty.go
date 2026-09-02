package awsmsk


// Details for SASL/OAUTHBEARER using the client credentials grant with a JWT client assertion (RFC 7521/7523 Section 2.2). An STS-vended JWT is used as the client_assertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterOAuthClientCredentialsAssertionProperty := &KafkaClusterOAuthClientCredentialsAssertionProperty{
//   	Audience: jsii.String("audience"),
//   	SigningAlgorithm: jsii.String("signingAlgorithm"),
//
//   	// the properties below are optional
//   	TokenRequestSecretArn: jsii.String("tokenRequestSecretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusteroauthclientcredentialsassertion.html
//
type CfnReplicator_KafkaClusterOAuthClientCredentialsAssertionProperty struct {
	// The audience (aud claim) set in the STS JWT client assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusteroauthclientcredentialsassertion.html#cfn-msk-replicator-kafkaclusteroauthclientcredentialsassertion-audience
	//
	Audience *string `field:"required" json:"audience" yaml:"audience"`
	// The algorithm used to sign the STS JWT assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusteroauthclientcredentialsassertion.html#cfn-msk-replicator-kafkaclusteroauthclientcredentialsassertion-signingalgorithm
	//
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// Optional Secrets Manager ARN for identity providers that require client_id as a form parameter alongside the JWT client assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusteroauthclientcredentialsassertion.html#cfn-msk-replicator-kafkaclusteroauthclientcredentialsassertion-tokenrequestsecretarn
	//
	TokenRequestSecretArn *string `field:"optional" json:"tokenRequestSecretArn" yaml:"tokenRequestSecretArn"`
}


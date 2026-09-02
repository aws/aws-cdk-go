package awsmsk


// Details for SASL/OAUTHBEARER using the JWT Bearer assertion grant (RFC 7523).
//
// An STS-vended JWT is used as the assertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   kafkaClusterOAuthIamJwtBearerProperty := &KafkaClusterOAuthIamJwtBearerProperty{
//   	Audience: jsii.String("audience"),
//   	SigningAlgorithm: jsii.String("signingAlgorithm"),
//   	TokenRequestSecretArn: jsii.String("tokenRequestSecretArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusteroauthiamjwtbearer.html
//
type CfnReplicatorPropsMixin_KafkaClusterOAuthIamJwtBearerProperty struct {
	// The audience (aud claim) set in the STS JWT assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusteroauthiamjwtbearer.html#cfn-msk-replicator-kafkaclusteroauthiamjwtbearer-audience
	//
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// The algorithm used to sign the STS JWT assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusteroauthiamjwtbearer.html#cfn-msk-replicator-kafkaclusteroauthiamjwtbearer-signingalgorithm
	//
	SigningAlgorithm *string `field:"optional" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// Optional Secrets Manager ARN for identity providers that require client authentication alongside the JWT Bearer assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusteroauthiamjwtbearer.html#cfn-msk-replicator-kafkaclusteroauthiamjwtbearer-tokenrequestsecretarn
	//
	TokenRequestSecretArn *string `field:"optional" json:"tokenRequestSecretArn" yaml:"tokenRequestSecretArn"`
}


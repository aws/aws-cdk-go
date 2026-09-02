package awsmsk


// Details of the client authentication used by the Apache Kafka cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterClientAuthenticationProperty := &KafkaClusterClientAuthenticationProperty{
//   	Mtls: &KafkaClusterMtlsAuthenticationProperty{
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	SaslOAuthBearer: &KafkaClusterSaslOAuthBearerAuthenticationProperty{
//   		TokenEndpointAuthenticationMethod: jsii.String("tokenEndpointAuthenticationMethod"),
//   		TokenEndpointUrl: jsii.String("tokenEndpointUrl"),
//
//   		// the properties below are optional
//   		ClientCredentials: &KafkaClusterOAuthClientCredentialsProperty{
//   			TokenRequestSecretArn: jsii.String("tokenRequestSecretArn"),
//   		},
//   		ClientCredentialsAssertion: &KafkaClusterOAuthClientCredentialsAssertionProperty{
//   			Audience: jsii.String("audience"),
//   			SigningAlgorithm: jsii.String("signingAlgorithm"),
//
//   			// the properties below are optional
//   			TokenRequestSecretArn: jsii.String("tokenRequestSecretArn"),
//   		},
//   		IamJwtBearer: &KafkaClusterOAuthIamJwtBearerProperty{
//   			Audience: jsii.String("audience"),
//   			SigningAlgorithm: jsii.String("signingAlgorithm"),
//
//   			// the properties below are optional
//   			TokenRequestSecretArn: jsii.String("tokenRequestSecretArn"),
//   		},
//   		Scope: jsii.String("scope"),
//   		TokenEndpointTlsCertificateArn: jsii.String("tokenEndpointTlsCertificateArn"),
//   	},
//   	SaslScram: &KafkaClusterSaslScramAuthenticationProperty{
//   		Mechanism: jsii.String("mechanism"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusterclientauthentication.html
//
type CfnReplicator_KafkaClusterClientAuthenticationProperty struct {
	// Details for mTLS client authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusterclientauthentication.html#cfn-msk-replicator-kafkaclusterclientauthentication-mtls
	//
	Mtls interface{} `field:"optional" json:"mtls" yaml:"mtls"`
	// Details for client authentication using SASL/OAUTHBEARER.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusterclientauthentication.html#cfn-msk-replicator-kafkaclusterclientauthentication-sasloauthbearer
	//
	SaslOAuthBearer interface{} `field:"optional" json:"saslOAuthBearer" yaml:"saslOAuthBearer"`
	// Details for SASL/SCRAM client authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclusterclientauthentication.html#cfn-msk-replicator-kafkaclusterclientauthentication-saslscram
	//
	SaslScram interface{} `field:"optional" json:"saslScram" yaml:"saslScram"`
}


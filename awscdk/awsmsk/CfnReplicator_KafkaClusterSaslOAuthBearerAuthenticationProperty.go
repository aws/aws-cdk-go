package awsmsk


// Details for client authentication using SASL/OAUTHBEARER.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   kafkaClusterSaslOAuthBearerAuthenticationProperty := &KafkaClusterSaslOAuthBearerAuthenticationProperty{
//   	TokenEndpointAuthenticationMethod: jsii.String("tokenEndpointAuthenticationMethod"),
//   	TokenEndpointUrl: jsii.String("tokenEndpointUrl"),
//
//   	// the properties below are optional
//   	ClientCredentials: &KafkaClusterOAuthClientCredentialsProperty{
//   		TokenRequestSecretArn: jsii.String("tokenRequestSecretArn"),
//   	},
//   	ClientCredentialsAssertion: &KafkaClusterOAuthClientCredentialsAssertionProperty{
//   		Audience: jsii.String("audience"),
//   		SigningAlgorithm: jsii.String("signingAlgorithm"),
//
//   		// the properties below are optional
//   		TokenRequestSecretArn: jsii.String("tokenRequestSecretArn"),
//   	},
//   	IamJwtBearer: &KafkaClusterOAuthIamJwtBearerProperty{
//   		Audience: jsii.String("audience"),
//   		SigningAlgorithm: jsii.String("signingAlgorithm"),
//
//   		// the properties below are optional
//   		TokenRequestSecretArn: jsii.String("tokenRequestSecretArn"),
//   	},
//   	Scope: jsii.String("scope"),
//   	TokenEndpointTlsCertificateArn: jsii.String("tokenEndpointTlsCertificateArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersasloauthbearerauthentication.html
//
type CfnReplicator_KafkaClusterSaslOAuthBearerAuthenticationProperty struct {
	// How client credentials are sent to the identity provider's token endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersasloauthbearerauthentication.html#cfn-msk-replicator-kafkaclustersasloauthbearerauthentication-tokenendpointauthenticationmethod
	//
	TokenEndpointAuthenticationMethod *string `field:"required" json:"tokenEndpointAuthenticationMethod" yaml:"tokenEndpointAuthenticationMethod"`
	// The HTTPS URL of the OAuth token endpoint that vends OAuth Bearer tokens per RFC 6749.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersasloauthbearerauthentication.html#cfn-msk-replicator-kafkaclustersasloauthbearerauthentication-tokenendpointurl
	//
	TokenEndpointUrl *string `field:"required" json:"tokenEndpointUrl" yaml:"tokenEndpointUrl"`
	// Details for SASL/OAUTHBEARER using the standard client_credentials grant.
	//
	// The referenced secret must contain client_id and client_secret.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersasloauthbearerauthentication.html#cfn-msk-replicator-kafkaclustersasloauthbearerauthentication-clientcredentials
	//
	ClientCredentials interface{} `field:"optional" json:"clientCredentials" yaml:"clientCredentials"`
	// Details for SASL/OAUTHBEARER using the client credentials grant with a JWT client assertion (RFC 7521/7523 Section 2.2). An STS-vended JWT is used as the client_assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersasloauthbearerauthentication.html#cfn-msk-replicator-kafkaclustersasloauthbearerauthentication-clientcredentialsassertion
	//
	ClientCredentialsAssertion interface{} `field:"optional" json:"clientCredentialsAssertion" yaml:"clientCredentialsAssertion"`
	// Details for SASL/OAUTHBEARER using the JWT Bearer assertion grant (RFC 7523).
	//
	// An STS-vended JWT is used as the assertion.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersasloauthbearerauthentication.html#cfn-msk-replicator-kafkaclustersasloauthbearerauthentication-iamjwtbearer
	//
	IamJwtBearer interface{} `field:"optional" json:"iamJwtBearer" yaml:"iamJwtBearer"`
	// OAuth scope to request.
	//
	// Included in the token request if provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersasloauthbearerauthentication.html#cfn-msk-replicator-kafkaclustersasloauthbearerauthentication-scope
	//
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Secrets Manager ARN containing a custom CA certificate for the identity provider.
	//
	// Required only if the identity provider uses a private CA.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-replicator-kafkaclustersasloauthbearerauthentication.html#cfn-msk-replicator-kafkaclustersasloauthbearerauthentication-tokenendpointtlscertificatearn
	//
	TokenEndpointTlsCertificateArn *string `field:"optional" json:"tokenEndpointTlsCertificateArn" yaml:"tokenEndpointTlsCertificateArn"`
}


package awselasticsearch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configures Amazon ES to use Amazon Cognito authentication for Kibana.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   es.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	CognitoKibanaAuth: &CognitoOptions{
//   		IdentityPoolId: jsii.String("test-identity-pool-id"),
//   		UserPoolId: jsii.String("test-user-pool-id"),
//   		Role: role,
//   	},
//   	Version: elasticsearchVersion,
//   })
//
// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html
//
// Deprecated: use opensearchservice module instead.
type CognitoOptions struct {
	// The Amazon Cognito identity pool ID that you want Amazon ES to use for Kibana authentication.
	// Deprecated: use opensearchservice module instead.
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// A role that allows Amazon ES to configure your user pool and identity pool.
	//
	// It must have the `AmazonESCognitoAccess` policy attached to it.
	// See: https://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html#es-cognito-auth-prereq
	//
	// Deprecated: use opensearchservice module instead.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// The Amazon Cognito user pool ID that you want Amazon ES to use for Kibana authentication.
	// Deprecated: use opensearchservice module instead.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}


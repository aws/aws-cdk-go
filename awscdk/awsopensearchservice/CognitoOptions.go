package awsopensearchservice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Configures Amazon OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   opensearch.NewDomain(this, jsii.String("Domain"), &DomainProps{
//   	CognitoDashboardsAuth: &CognitoOptions{
//   		IdentityPoolId: jsii.String("test-identity-pool-id"),
//   		UserPoolId: jsii.String("test-user-pool-id"),
//   		Role: role,
//   	},
//   	Version: openSearchVersion,
//   })
//
// See: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html
//
type CognitoOptions struct {
	// The Amazon Cognito identity pool ID that you want Amazon OpenSearch Service to use for OpenSearch Dashboards authentication.
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// A role that allows Amazon OpenSearch Service to configure your user pool and identity pool.
	//
	// It must have the `AmazonESCognitoAccess` policy attached to it.
	// See: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html#cognito-auth-prereq
	//
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// The Amazon Cognito user pool ID that you want Amazon OpenSearch Service to use for OpenSearch Dashboards authentication.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
}


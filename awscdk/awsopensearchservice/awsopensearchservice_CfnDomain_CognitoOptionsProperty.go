package awsopensearchservice


// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoOptionsProperty := &cognitoOptionsProperty{
//   	enabled: jsii.Boolean(false),
//   	identityPoolId: jsii.String("identityPoolId"),
//   	roleArn: jsii.String("roleArn"),
//   	userPoolId: jsii.String("userPoolId"),
//   }
//
type CfnDomain_CognitoOptionsProperty struct {
	// Whether to enable or disable Amazon Cognito authentication for OpenSearch Dashboards.
	//
	// See [Amazon Cognito authentication for OpenSearch Dashboards](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html) .
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The Amazon Cognito identity pool ID that you want OpenSearch Service to use for OpenSearch Dashboards authentication.
	//
	// Required if you enabled Cognito Authentication for OpenSearch Dashboards.
	IdentityPoolId *string `field:"optional" json:"identityPoolId" yaml:"identityPoolId"`
	// The `AmazonOpenSearchServiceCognitoAccess` role that allows OpenSearch Service to configure your user pool and identity pool.
	//
	// Required if you enabled Cognito Authentication for OpenSearch Dashboards.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The Amazon Cognito user pool ID that you want OpenSearch Service to use for OpenSearch Dashboards authentication.
	//
	// Required if you enabled Cognito Authentication for OpenSearch Dashboards.
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}


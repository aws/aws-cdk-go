package awsopensearchservice


// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cognitoOptionsProperty := &CognitoOptionsProperty{
//   	Enabled: jsii.Boolean(false),
//   	IdentityPoolId: jsii.String("identityPoolId"),
//   	RoleArn: jsii.String("roleArn"),
//   	UserPoolId: jsii.String("userPoolId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-cognitooptions.html
//
type CfnDomain_CognitoOptionsProperty struct {
	// Whether to enable or disable Amazon Cognito authentication for OpenSearch Dashboards.
	//
	// See [Amazon Cognito authentication for OpenSearch Dashboards](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-cognitooptions.html#cfn-opensearchservice-domain-cognitooptions-enabled
	//
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The Amazon Cognito identity pool ID that you want OpenSearch Service to use for OpenSearch Dashboards authentication.
	//
	// Required if you enabled Cognito Authentication for OpenSearch Dashboards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-cognitooptions.html#cfn-opensearchservice-domain-cognitooptions-identitypoolid
	//
	IdentityPoolId *string `field:"optional" json:"identityPoolId" yaml:"identityPoolId"`
	// The `AmazonOpenSearchServiceCognitoAccess` role that allows OpenSearch Service to configure your user pool and identity pool.
	//
	// Required if you enabled Cognito Authentication for OpenSearch Dashboards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-cognitooptions.html#cfn-opensearchservice-domain-cognitooptions-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The Amazon Cognito user pool ID that you want OpenSearch Service to use for OpenSearch Dashboards authentication.
	//
	// Required if you enabled Cognito Authentication for OpenSearch Dashboards.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opensearchservice-domain-cognitooptions.html#cfn-opensearchservice-domain-cognitooptions-userpoolid
	//
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}


package awselasticsearch


// Configures OpenSearch Service to use Amazon Cognito authentication for OpenSearch Dashboards.
//
// > The `AWS::Elasticsearch::Domain` resource is being replaced by the [AWS::OpenSearchService::Domain](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opensearchservice-domain.html) resource. While the legacy Elasticsearch resource and options are still supported, we recommend modifying your existing Cloudformation templates to use the new OpenSearch Service resource, which supports both OpenSearch and Elasticsearch. For more information about the service rename, see [New resource types](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/rename.html#rename-resource) in the *Amazon OpenSearch Service Developer Guide* .
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
	// Required if you enable Cognito authentication.
	IdentityPoolId *string `field:"optional" json:"identityPoolId" yaml:"identityPoolId"`
	// The `AmazonESCognitoAccess` role that allows OpenSearch Service to configure your user pool and identity pool.
	//
	// Required if you enable Cognito authentication.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The Amazon Cognito user pool ID that you want OpenSearch Service to use for OpenSearch Dashboards authentication.
	//
	// Required if you enable Cognito authentication.
	UserPoolId *string `field:"optional" json:"userPoolId" yaml:"userPoolId"`
}


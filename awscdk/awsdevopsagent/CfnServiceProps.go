package awsdevopsagent


// Properties for defining a `CfnService`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var exchangeParameters interface{}
//
//   cfnServiceProps := &CfnServiceProps{
//   	ServiceType: jsii.String("serviceType"),
//
//   	// the properties below are optional
//   	ServiceDetails: &ServiceDetailsProperty{
//   		Dynatrace: &DynatraceServiceDetailsProperty{
//   			AccountUrn: jsii.String("accountUrn"),
//
//   			// the properties below are optional
//   			AuthorizationConfig: &DynatraceAuthorizationConfigProperty{
//   				OAuthClientCredentials: &OAuthClientDetailsProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//
//   					// the properties below are optional
//   					ClientName: jsii.String("clientName"),
//   					ExchangeParameters: exchangeParameters,
//   				},
//   			},
//   		},
//   		GitLab: &GitLabDetailsProperty{
//   			TargetUrl: jsii.String("targetUrl"),
//   			TokenType: jsii.String("tokenType"),
//   			TokenValue: jsii.String("tokenValue"),
//
//   			// the properties below are optional
//   			GroupId: jsii.String("groupId"),
//   		},
//   		McpServer: &MCPServerDetailsProperty{
//   			AuthorizationConfig: &MCPServerAuthorizationConfigProperty{
//   				ApiKey: &ApiKeyDetailsProperty{
//   					ApiKeyHeader: jsii.String("apiKeyHeader"),
//   					ApiKeyName: jsii.String("apiKeyName"),
//   					ApiKeyValue: jsii.String("apiKeyValue"),
//   				},
//   				OAuthClientCredentials: &MCPServerOAuthClientCredentialsConfigProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ExchangeUrl: jsii.String("exchangeUrl"),
//
//   					// the properties below are optional
//   					ClientName: jsii.String("clientName"),
//   					ExchangeParameters: exchangeParameters,
//   					Scopes: []*string{
//   						jsii.String("scopes"),
//   					},
//   				},
//   			},
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   		McpServerNewRelic: &NewRelicServiceDetailsProperty{
//   			AuthorizationConfig: &NewRelicAuthorizationConfigProperty{
//   				ApiKey: &NewRelicApiKeyConfigProperty{
//   					AccountId: jsii.String("accountId"),
//   					ApiKey: jsii.String("apiKey"),
//   					Region: jsii.String("region"),
//
//   					// the properties below are optional
//   					AlertPolicyIds: []*string{
//   						jsii.String("alertPolicyIds"),
//   					},
//   					ApplicationIds: []*string{
//   						jsii.String("applicationIds"),
//   					},
//   					EntityGuids: []*string{
//   						jsii.String("entityGuids"),
//   					},
//   				},
//   			},
//   		},
//   		McpServerSplunk: &MCPServerSplunkDetailsProperty{
//   			AuthorizationConfig: &MCPServerSplunkAuthorizationConfigProperty{
//   				BearerToken: &BearerTokenDetailsProperty{
//   					TokenName: jsii.String("tokenName"),
//   					TokenValue: jsii.String("tokenValue"),
//
//   					// the properties below are optional
//   					AuthorizationHeader: jsii.String("authorizationHeader"),
//   				},
//   			},
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   		ServiceNow: &ServiceNowServiceDetailsProperty{
//   			InstanceUrl: jsii.String("instanceUrl"),
//
//   			// the properties below are optional
//   			AuthorizationConfig: &ServiceNowAuthorizationConfigProperty{
//   				OAuthClientCredentials: &OAuthClientDetailsProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//
//   					// the properties below are optional
//   					ClientName: jsii.String("clientName"),
//   					ExchangeParameters: exchangeParameters,
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html
//
type CfnServiceProps struct {
	// The type of service being registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-servicetype
	//
	ServiceType *string `field:"required" json:"serviceType" yaml:"serviceType"`
	// Service-specific configuration details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-servicedetails
	//
	ServiceDetails interface{} `field:"optional" json:"serviceDetails" yaml:"serviceDetails"`
}


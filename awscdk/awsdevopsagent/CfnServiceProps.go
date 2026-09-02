package awsdevopsagent

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

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
//   	ExchangeUrlPrivateConnectionName: jsii.String("exchangeUrlPrivateConnectionName"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	PrivateConnectionName: jsii.String("privateConnectionName"),
//   	ServiceDetails: &ServiceDetailsProperty{
//   		AzureIdentity: &AzureIdentityServiceDetailsProperty{
//   			ClientId: jsii.String("clientId"),
//   			TenantId: jsii.String("tenantId"),
//   			WebIdentityRoleArn: jsii.String("webIdentityRoleArn"),
//   			WebIdentityTokenAudiences: []*string{
//   				jsii.String("webIdentityTokenAudiences"),
//   			},
//   		},
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
//   				BearerToken: &BearerTokenDetailsProperty{
//   					TokenName: jsii.String("tokenName"),
//   					TokenValue: jsii.String("tokenValue"),
//
//   					// the properties below are optional
//   					AuthorizationHeader: jsii.String("authorizationHeader"),
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
//   		McpServerGrafana: &MCPServerGrafanaDetailsProperty{
//   			AuthorizationConfig: &MCPServerGrafanaAuthorizationConfigProperty{
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
//   		McpServerSigV4: &MCPServerSigV4DetailsProperty{
//   			AuthorizationConfig: &MCPServerSigV4AuthorizationConfigProperty{
//   				Region: jsii.String("region"),
//   				Service: jsii.String("service"),
//
//   				// the properties below are optional
//   				CustomHeaders: map[string]*string{
//   					"customHeadersKey": jsii.String("customHeaders"),
//   				},
//   				McpRoleArn: jsii.String("mcpRoleArn"),
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
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
//   		PagerDuty: &PagerDutyDetailsProperty{
//   			AuthorizationConfig: &PagerDutyAuthorizationConfigProperty{
//   				OAuthClientCredentials: &OAuthClientDetailsProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientSecret: jsii.String("clientSecret"),
//
//   					// the properties below are optional
//   					ClientName: jsii.String("clientName"),
//   					ExchangeParameters: exchangeParameters,
//   				},
//   			},
//   			Scopes: []*string{
//   				jsii.String("scopes"),
//   			},
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetUrlPrivateConnectionName: jsii.String("targetUrlPrivateConnectionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html
//
type CfnServiceProps struct {
	// The type of service being registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-servicetype
	//
	ServiceType *string `field:"required" json:"serviceType" yaml:"serviceType"`
	// The name of the private connection to use for OAuth token exchange requests only.
	//
	// Cannot be specified when PrivateConnectionName is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-exchangeurlprivateconnectionname
	//
	ExchangeUrlPrivateConnectionName *string `field:"optional" json:"exchangeUrlPrivateConnectionName" yaml:"exchangeUrlPrivateConnectionName"`
	// The ARN of the KMS key to use for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The name of the private connection to use for VPC connectivity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-privateconnectionname
	//
	PrivateConnectionName *string `field:"optional" json:"privateConnectionName" yaml:"privateConnectionName"`
	// Service-specific configuration details - MCPServerSigV4 supports in-place updates;
	//
	// GitLab (TokenValue), MCPServer (ApiKey or BearerToken credential rotation), MCPServerNewRelic (ApiKey rotation), and MCPServerGrafana (BearerToken rotation) support in-place credential rotation; all other service types and fields require replacement when modified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-servicedetails
	//
	ServiceDetails interface{} `field:"optional" json:"serviceDetails" yaml:"serviceDetails"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the private connection to use for API calls (target URL) only.
	//
	// Cannot be specified when PrivateConnectionName is provided.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-targeturlprivateconnectionname
	//
	TargetUrlPrivateConnectionName *string `field:"optional" json:"targetUrlPrivateConnectionName" yaml:"targetUrlPrivateConnectionName"`
}


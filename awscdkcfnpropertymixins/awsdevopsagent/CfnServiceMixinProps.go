package awsdevopsagent

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnServicePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var exchangeParameters interface{}
//
//   cfnServiceMixinProps := &CfnServiceMixinProps{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
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
//   			AuthorizationConfig: &DynatraceAuthorizationConfigProperty{
//   				OAuthClientCredentials: &OAuthClientDetailsProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientName: jsii.String("clientName"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ExchangeParameters: exchangeParameters,
//   				},
//   			},
//   		},
//   		GitLab: &GitLabDetailsProperty{
//   			GroupId: jsii.String("groupId"),
//   			TargetUrl: jsii.String("targetUrl"),
//   			TokenType: jsii.String("tokenType"),
//   			TokenValue: jsii.String("tokenValue"),
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
//   					ClientName: jsii.String("clientName"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ExchangeParameters: exchangeParameters,
//   					ExchangeUrl: jsii.String("exchangeUrl"),
//   					Scopes: []*string{
//   						jsii.String("scopes"),
//   					},
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   		},
//   		McpServerNewRelic: &NewRelicServiceDetailsProperty{
//   			AuthorizationConfig: &NewRelicAuthorizationConfigProperty{
//   				ApiKey: &NewRelicApiKeyConfigProperty{
//   					AccountId: jsii.String("accountId"),
//   					AlertPolicyIds: []*string{
//   						jsii.String("alertPolicyIds"),
//   					},
//   					ApiKey: jsii.String("apiKey"),
//   					ApplicationIds: []*string{
//   						jsii.String("applicationIds"),
//   					},
//   					EntityGuids: []*string{
//   						jsii.String("entityGuids"),
//   					},
//   					Region: jsii.String("region"),
//   				},
//   			},
//   		},
//   		McpServerSigV4: &MCPServerSigV4DetailsProperty{
//   			AuthorizationConfig: &MCPServerSigV4AuthorizationConfigProperty{
//   				CustomHeaders: map[string]*string{
//   					"customHeadersKey": jsii.String("customHeaders"),
//   				},
//   				Region: jsii.String("region"),
//   				RoleArn: jsii.String("roleArn"),
//   				Service: jsii.String("service"),
//   			},
//   			Description: jsii.String("description"),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   		},
//   		McpServerSplunk: &MCPServerSplunkDetailsProperty{
//   			AuthorizationConfig: &MCPServerSplunkAuthorizationConfigProperty{
//   				BearerToken: &BearerTokenDetailsProperty{
//   					AuthorizationHeader: jsii.String("authorizationHeader"),
//   					TokenName: jsii.String("tokenName"),
//   					TokenValue: jsii.String("tokenValue"),
//   				},
//   			},
//   			Description: jsii.String("description"),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   		},
//   		PagerDuty: &PagerDutyDetailsProperty{
//   			AuthorizationConfig: &PagerDutyAuthorizationConfigProperty{
//   				OAuthClientCredentials: &OAuthClientDetailsProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientName: jsii.String("clientName"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ExchangeParameters: exchangeParameters,
//   				},
//   			},
//   			Scopes: []*string{
//   				jsii.String("scopes"),
//   			},
//   		},
//   		ServiceNow: &ServiceNowServiceDetailsProperty{
//   			AuthorizationConfig: &ServiceNowAuthorizationConfigProperty{
//   				OAuthClientCredentials: &OAuthClientDetailsProperty{
//   					ClientId: jsii.String("clientId"),
//   					ClientName: jsii.String("clientName"),
//   					ClientSecret: jsii.String("clientSecret"),
//   					ExchangeParameters: exchangeParameters,
//   				},
//   			},
//   			InstanceUrl: jsii.String("instanceUrl"),
//   		},
//   	},
//   	ServiceType: jsii.String("serviceType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html
//
type CfnServiceMixinProps struct {
	// The ARN of the KMS key to use for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Service-specific configuration details.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-servicedetails
	//
	ServiceDetails interface{} `field:"optional" json:"serviceDetails" yaml:"serviceDetails"`
	// The type of service being registered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-servicetype
	//
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-service.html#cfn-devopsagent-service-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


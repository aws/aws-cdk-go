package awsdevopsagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var exchangeParameters interface{}
//
//   serviceDetailsProperty := &ServiceDetailsProperty{
//   	AzureIdentity: &AzureIdentityServiceDetailsProperty{
//   		ClientId: jsii.String("clientId"),
//   		TenantId: jsii.String("tenantId"),
//   		WebIdentityRoleArn: jsii.String("webIdentityRoleArn"),
//   		WebIdentityTokenAudiences: []*string{
//   			jsii.String("webIdentityTokenAudiences"),
//   		},
//   	},
//   	Dynatrace: &DynatraceServiceDetailsProperty{
//   		AccountUrn: jsii.String("accountUrn"),
//   		AuthorizationConfig: &DynatraceAuthorizationConfigProperty{
//   			OAuthClientCredentials: &OAuthClientDetailsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientName: jsii.String("clientName"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				ExchangeParameters: exchangeParameters,
//   			},
//   		},
//   	},
//   	GitLab: &GitLabDetailsProperty{
//   		GroupId: jsii.String("groupId"),
//   		TargetUrl: jsii.String("targetUrl"),
//   		TokenType: jsii.String("tokenType"),
//   		TokenValue: jsii.String("tokenValue"),
//   	},
//   	McpServer: &MCPServerDetailsProperty{
//   		AuthorizationConfig: &MCPServerAuthorizationConfigProperty{
//   			ApiKey: &ApiKeyDetailsProperty{
//   				ApiKeyHeader: jsii.String("apiKeyHeader"),
//   				ApiKeyName: jsii.String("apiKeyName"),
//   				ApiKeyValue: jsii.String("apiKeyValue"),
//   			},
//   			OAuthClientCredentials: &MCPServerOAuthClientCredentialsConfigProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientName: jsii.String("clientName"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				ExchangeParameters: exchangeParameters,
//   				ExchangeUrl: jsii.String("exchangeUrl"),
//   				Scopes: []*string{
//   					jsii.String("scopes"),
//   				},
//   			},
//   		},
//   		Description: jsii.String("description"),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   	},
//   	McpServerNewRelic: &NewRelicServiceDetailsProperty{
//   		AuthorizationConfig: &NewRelicAuthorizationConfigProperty{
//   			ApiKey: &NewRelicApiKeyConfigProperty{
//   				AccountId: jsii.String("accountId"),
//   				AlertPolicyIds: []*string{
//   					jsii.String("alertPolicyIds"),
//   				},
//   				ApiKey: jsii.String("apiKey"),
//   				ApplicationIds: []*string{
//   					jsii.String("applicationIds"),
//   				},
//   				EntityGuids: []*string{
//   					jsii.String("entityGuids"),
//   				},
//   				Region: jsii.String("region"),
//   			},
//   		},
//   	},
//   	McpServerSigV4: &MCPServerSigV4DetailsProperty{
//   		AuthorizationConfig: &MCPServerSigV4AuthorizationConfigProperty{
//   			CustomHeaders: map[string]*string{
//   				"customHeadersKey": jsii.String("customHeaders"),
//   			},
//   			Region: jsii.String("region"),
//   			RoleArn: jsii.String("roleArn"),
//   			Service: jsii.String("service"),
//   		},
//   		Description: jsii.String("description"),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   	},
//   	McpServerSplunk: &MCPServerSplunkDetailsProperty{
//   		AuthorizationConfig: &MCPServerSplunkAuthorizationConfigProperty{
//   			BearerToken: &BearerTokenDetailsProperty{
//   				AuthorizationHeader: jsii.String("authorizationHeader"),
//   				TokenName: jsii.String("tokenName"),
//   				TokenValue: jsii.String("tokenValue"),
//   			},
//   		},
//   		Description: jsii.String("description"),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   	},
//   	PagerDuty: &PagerDutyDetailsProperty{
//   		AuthorizationConfig: &PagerDutyAuthorizationConfigProperty{
//   			OAuthClientCredentials: &OAuthClientDetailsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientName: jsii.String("clientName"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				ExchangeParameters: exchangeParameters,
//   			},
//   		},
//   		Scopes: []*string{
//   			jsii.String("scopes"),
//   		},
//   	},
//   	ServiceNow: &ServiceNowServiceDetailsProperty{
//   		AuthorizationConfig: &ServiceNowAuthorizationConfigProperty{
//   			OAuthClientCredentials: &OAuthClientDetailsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientName: jsii.String("clientName"),
//   				ClientSecret: jsii.String("clientSecret"),
//   				ExchangeParameters: exchangeParameters,
//   			},
//   		},
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html
//
type CfnServicePropsMixin_ServiceDetailsProperty struct {
	// Azure Identity service configuration for federated identity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-azureidentity
	//
	AzureIdentity interface{} `field:"optional" json:"azureIdentity" yaml:"azureIdentity"`
	// Dynatrace service configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-dynatrace
	//
	Dynatrace interface{} `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// GitLab service configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-gitlab
	//
	GitLab interface{} `field:"optional" json:"gitLab" yaml:"gitLab"`
	// MCP server configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-mcpserver
	//
	McpServer interface{} `field:"optional" json:"mcpServer" yaml:"mcpServer"`
	// New Relic service configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-mcpservernewrelic
	//
	McpServerNewRelic interface{} `field:"optional" json:"mcpServerNewRelic" yaml:"mcpServerNewRelic"`
	// SigV4-authenticated MCP server configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-mcpserversigv4
	//
	McpServerSigV4 interface{} `field:"optional" json:"mcpServerSigV4" yaml:"mcpServerSigV4"`
	// Splunk MCP server configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-mcpserversplunk
	//
	McpServerSplunk interface{} `field:"optional" json:"mcpServerSplunk" yaml:"mcpServerSplunk"`
	// PagerDuty service configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-pagerduty
	//
	PagerDuty interface{} `field:"optional" json:"pagerDuty" yaml:"pagerDuty"`
	// ServiceNow service configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
}


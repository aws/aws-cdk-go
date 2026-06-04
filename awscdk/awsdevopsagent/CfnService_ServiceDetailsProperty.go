package awsdevopsagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//
//   		// the properties below are optional
//   		AuthorizationConfig: &DynatraceAuthorizationConfigProperty{
//   			OAuthClientCredentials: &OAuthClientDetailsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				ClientName: jsii.String("clientName"),
//   				ExchangeParameters: exchangeParameters,
//   			},
//   		},
//   	},
//   	GitLab: &GitLabDetailsProperty{
//   		TargetUrl: jsii.String("targetUrl"),
//   		TokenType: jsii.String("tokenType"),
//   		TokenValue: jsii.String("tokenValue"),
//
//   		// the properties below are optional
//   		GroupId: jsii.String("groupId"),
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
//   				ClientSecret: jsii.String("clientSecret"),
//   				ExchangeUrl: jsii.String("exchangeUrl"),
//
//   				// the properties below are optional
//   				ClientName: jsii.String("clientName"),
//   				ExchangeParameters: exchangeParameters,
//   				Scopes: []*string{
//   					jsii.String("scopes"),
//   				},
//   			},
//   		},
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   	},
//   	McpServerGrafana: &MCPServerGrafanaDetailsProperty{
//   		AuthorizationConfig: &MCPServerGrafanaAuthorizationConfigProperty{
//   			BearerToken: &BearerTokenDetailsProperty{
//   				TokenName: jsii.String("tokenName"),
//   				TokenValue: jsii.String("tokenValue"),
//
//   				// the properties below are optional
//   				AuthorizationHeader: jsii.String("authorizationHeader"),
//   			},
//   		},
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   	},
//   	McpServerNewRelic: &NewRelicServiceDetailsProperty{
//   		AuthorizationConfig: &NewRelicAuthorizationConfigProperty{
//   			ApiKey: &NewRelicApiKeyConfigProperty{
//   				AccountId: jsii.String("accountId"),
//   				ApiKey: jsii.String("apiKey"),
//   				Region: jsii.String("region"),
//
//   				// the properties below are optional
//   				AlertPolicyIds: []*string{
//   					jsii.String("alertPolicyIds"),
//   				},
//   				ApplicationIds: []*string{
//   					jsii.String("applicationIds"),
//   				},
//   				EntityGuids: []*string{
//   					jsii.String("entityGuids"),
//   				},
//   			},
//   		},
//   	},
//   	McpServerSigV4: &MCPServerSigV4DetailsProperty{
//   		AuthorizationConfig: &MCPServerSigV4AuthorizationConfigProperty{
//   			Region: jsii.String("region"),
//   			RoleArn: jsii.String("roleArn"),
//   			Service: jsii.String("service"),
//
//   			// the properties below are optional
//   			CustomHeaders: map[string]*string{
//   				"customHeadersKey": jsii.String("customHeaders"),
//   			},
//   		},
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   	},
//   	McpServerSplunk: &MCPServerSplunkDetailsProperty{
//   		AuthorizationConfig: &MCPServerSplunkAuthorizationConfigProperty{
//   			BearerToken: &BearerTokenDetailsProperty{
//   				TokenName: jsii.String("tokenName"),
//   				TokenValue: jsii.String("tokenValue"),
//
//   				// the properties below are optional
//   				AuthorizationHeader: jsii.String("authorizationHeader"),
//   			},
//   		},
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   	},
//   	PagerDuty: &PagerDutyDetailsProperty{
//   		AuthorizationConfig: &PagerDutyAuthorizationConfigProperty{
//   			OAuthClientCredentials: &OAuthClientDetailsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				ClientName: jsii.String("clientName"),
//   				ExchangeParameters: exchangeParameters,
//   			},
//   		},
//   		Scopes: []*string{
//   			jsii.String("scopes"),
//   		},
//   	},
//   	ServiceNow: &ServiceNowServiceDetailsProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//
//   		// the properties below are optional
//   		AuthorizationConfig: &ServiceNowAuthorizationConfigProperty{
//   			OAuthClientCredentials: &OAuthClientDetailsProperty{
//   				ClientId: jsii.String("clientId"),
//   				ClientSecret: jsii.String("clientSecret"),
//
//   				// the properties below are optional
//   				ClientName: jsii.String("clientName"),
//   				ExchangeParameters: exchangeParameters,
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html
//
type CfnService_ServiceDetailsProperty struct {
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
	// Grafana MCP server configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-mcpservergrafana
	//
	McpServerGrafana interface{} `field:"optional" json:"mcpServerGrafana" yaml:"mcpServerGrafana"`
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


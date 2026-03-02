package previewawsdevopsagentmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var exchangeParameters interface{}
//
//   serviceDetailsProperty := &ServiceDetailsProperty{
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
	// Splunk MCP server configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-mcpserversplunk
	//
	McpServerSplunk interface{} `field:"optional" json:"mcpServerSplunk" yaml:"mcpServerSplunk"`
	// ServiceNow service configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-servicedetails.html#cfn-devopsagent-service-servicedetails-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
}


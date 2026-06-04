package awsdevopsagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   additionalServiceDetailsProperty := &AdditionalServiceDetailsProperty{
//   	AzureIdentity: &RegisteredAzureIdentityDetailsProperty{
//   		ClientId: jsii.String("clientId"),
//   		TenantId: jsii.String("tenantId"),
//   		WebIdentityRoleArn: jsii.String("webIdentityRoleArn"),
//   		WebIdentityTokenAudiences: []*string{
//   			jsii.String("webIdentityTokenAudiences"),
//   		},
//   	},
//   	Dynatrace: &RegisteredDynatraceDetailsProperty{
//   		AccountUrn: jsii.String("accountUrn"),
//   	},
//   	GitLab: &RegisteredGitLabServiceDetailsProperty{
//   		GroupId: jsii.String("groupId"),
//   		TargetUrl: jsii.String("targetUrl"),
//   		TokenType: jsii.String("tokenType"),
//   	},
//   	McpServer: &RegisteredMCPServerDetailsProperty{
//   		ApiKeyHeader: jsii.String("apiKeyHeader"),
//   		AuthorizationMethod: jsii.String("authorizationMethod"),
//   		Description: jsii.String("description"),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   	},
//   	McpServerGrafana: &RegisteredMCPServerGrafanaDetailsProperty{
//   		AuthorizationMethod: jsii.String("authorizationMethod"),
//   		Description: jsii.String("description"),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   	},
//   	McpServerNewRelic: &RegisteredNewRelicDetailsProperty{
//   		AccountId: jsii.String("accountId"),
//   		Description: jsii.String("description"),
//   		Region: jsii.String("region"),
//   	},
//   	McpServerSigV4: &RegisteredMCPServerSigV4DetailsProperty{
//   		CustomHeaders: map[string]*string{
//   			"customHeadersKey": jsii.String("customHeaders"),
//   		},
//   		Description: jsii.String("description"),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   		Region: jsii.String("region"),
//   		RoleArn: jsii.String("roleArn"),
//   		Service: jsii.String("service"),
//   	},
//   	McpServerSplunk: &RegisteredMCPServerDetailsProperty{
//   		ApiKeyHeader: jsii.String("apiKeyHeader"),
//   		AuthorizationMethod: jsii.String("authorizationMethod"),
//   		Description: jsii.String("description"),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   	},
//   	PagerDuty: &RegisteredPagerDutyDetailsProperty{
//   		Scopes: []*string{
//   			jsii.String("scopes"),
//   		},
//   	},
//   	ServiceNow: &RegisteredServiceNowDetailsProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html
//
type CfnServicePropsMixin_AdditionalServiceDetailsProperty struct {
	// Azure Identity service details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-azureidentity
	//
	AzureIdentity interface{} `field:"optional" json:"azureIdentity" yaml:"azureIdentity"`
	// Dynatrace service details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-dynatrace
	//
	Dynatrace interface{} `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// GitLab service details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-gitlab
	//
	GitLab interface{} `field:"optional" json:"gitLab" yaml:"gitLab"`
	// MCP server details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-mcpserver
	//
	McpServer interface{} `field:"optional" json:"mcpServer" yaml:"mcpServer"`
	// Grafana MCP server details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-mcpservergrafana
	//
	McpServerGrafana interface{} `field:"optional" json:"mcpServerGrafana" yaml:"mcpServerGrafana"`
	// New Relic service details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-mcpservernewrelic
	//
	McpServerNewRelic interface{} `field:"optional" json:"mcpServerNewRelic" yaml:"mcpServerNewRelic"`
	// SigV4-authenticated MCP server details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-mcpserversigv4
	//
	McpServerSigV4 interface{} `field:"optional" json:"mcpServerSigV4" yaml:"mcpServerSigV4"`
	// MCP server details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-mcpserversplunk
	//
	McpServerSplunk interface{} `field:"optional" json:"mcpServerSplunk" yaml:"mcpServerSplunk"`
	// PagerDuty service details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-pagerduty
	//
	PagerDuty interface{} `field:"optional" json:"pagerDuty" yaml:"pagerDuty"`
	// ServiceNow service details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
}


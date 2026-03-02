package previewawsdevopsagentmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   additionalServiceDetailsProperty := &AdditionalServiceDetailsProperty{
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
//   	McpServerNewRelic: &RegisteredNewRelicDetailsProperty{
//   		AccountId: jsii.String("accountId"),
//   		Description: jsii.String("description"),
//   		Region: jsii.String("region"),
//   	},
//   	McpServerSplunk: &RegisteredMCPServerDetailsProperty{
//   		ApiKeyHeader: jsii.String("apiKeyHeader"),
//   		AuthorizationMethod: jsii.String("authorizationMethod"),
//   		Description: jsii.String("description"),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   	},
//   	ServiceNow: &RegisteredServiceNowDetailsProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html
//
type CfnServicePropsMixin_AdditionalServiceDetailsProperty struct {
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
	// New Relic service details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-mcpservernewrelic
	//
	McpServerNewRelic interface{} `field:"optional" json:"mcpServerNewRelic" yaml:"mcpServerNewRelic"`
	// MCP server details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-mcpserversplunk
	//
	McpServerSplunk interface{} `field:"optional" json:"mcpServerSplunk" yaml:"mcpServerSplunk"`
	// ServiceNow service details returned after registration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html#cfn-devopsagent-service-additionalservicedetails-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
}


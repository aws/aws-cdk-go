package awsdevopsagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   additionalServiceDetailsProperty := &AdditionalServiceDetailsProperty{
//   	Dynatrace: &RegisteredDynatraceDetailsProperty{
//   		AccountUrn: jsii.String("accountUrn"),
//   	},
//   	GitLab: &RegisteredGitLabServiceDetailsProperty{
//   		TargetUrl: jsii.String("targetUrl"),
//   		TokenType: jsii.String("tokenType"),
//
//   		// the properties below are optional
//   		GroupId: jsii.String("groupId"),
//   	},
//   	McpServer: &RegisteredMCPServerDetailsProperty{
//   		AuthorizationMethod: jsii.String("authorizationMethod"),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		ApiKeyHeader: jsii.String("apiKeyHeader"),
//   		Description: jsii.String("description"),
//   	},
//   	McpServerNewRelic: &RegisteredNewRelicDetailsProperty{
//   		AccountId: jsii.String("accountId"),
//   		Region: jsii.String("region"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   	},
//   	McpServerSplunk: &RegisteredMCPServerDetailsProperty{
//   		AuthorizationMethod: jsii.String("authorizationMethod"),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		ApiKeyHeader: jsii.String("apiKeyHeader"),
//   		Description: jsii.String("description"),
//   	},
//   	ServiceNow: &RegisteredServiceNowDetailsProperty{
//   		InstanceUrl: jsii.String("instanceUrl"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-additionalservicedetails.html
//
type CfnService_AdditionalServiceDetailsProperty struct {
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


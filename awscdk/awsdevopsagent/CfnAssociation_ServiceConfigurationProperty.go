package awsdevopsagent


// The configuration that directs how Agent Space interacts with the given service.
//
// You can specify only one configuration type per association.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var resourceMetadata interface{}
//
//   serviceConfigurationProperty := &ServiceConfigurationProperty{
//   	Aws: &AWSConfigurationProperty{
//   		AccountId: jsii.String("accountId"),
//   		AccountType: jsii.String("accountType"),
//   		AssumableRoleArn: jsii.String("assumableRoleArn"),
//
//   		// the properties below are optional
//   		Resources: []interface{}{
//   			&AWSResourceProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//
//   				// the properties below are optional
//   				ResourceMetadata: resourceMetadata,
//   				ResourceType: jsii.String("resourceType"),
//   			},
//   		},
//   		Tags: []KeyValuePairProperty{
//   			&KeyValuePairProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Dynatrace: &DynatraceConfigurationProperty{
//   		EnvId: jsii.String("envId"),
//
//   		// the properties below are optional
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   	},
//   	EventChannel: &EventChannelConfigurationProperty{
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   	},
//   	GitHub: &GitHubConfigurationProperty{
//   		Owner: jsii.String("owner"),
//   		OwnerType: jsii.String("ownerType"),
//   		RepoId: jsii.String("repoId"),
//   		RepoName: jsii.String("repoName"),
//   	},
//   	GitLab: &GitLabConfigurationProperty{
//   		ProjectId: jsii.String("projectId"),
//   		ProjectPath: jsii.String("projectPath"),
//
//   		// the properties below are optional
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   		InstanceIdentifier: jsii.String("instanceIdentifier"),
//   	},
//   	McpServer: &MCPServerConfigurationProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   		Tools: []*string{
//   			jsii.String("tools"),
//   		},
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   	},
//   	McpServerDatadog: &MCPServerDatadogConfigurationProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   	},
//   	McpServerNewRelic: &MCPServerNewRelicConfigurationProperty{
//   		AccountId: jsii.String("accountId"),
//   		Endpoint: jsii.String("endpoint"),
//   	},
//   	McpServerSplunk: &MCPServerSplunkConfigurationProperty{
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   	},
//   	ServiceNow: &ServiceNowConfigurationProperty{
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   		InstanceId: jsii.String("instanceId"),
//   	},
//   	Slack: &SlackConfigurationProperty{
//   		TransmissionTarget: &SlackTransmissionTargetProperty{
//   			IncidentResponseTarget: &SlackChannelProperty{
//   				ChannelId: jsii.String("channelId"),
//
//   				// the properties below are optional
//   				ChannelName: jsii.String("channelName"),
//   			},
//   		},
//   		WorkspaceId: jsii.String("workspaceId"),
//   		WorkspaceName: jsii.String("workspaceName"),
//   	},
//   	SourceAws: &SourceAwsConfigurationProperty{
//   		AccountId: jsii.String("accountId"),
//   		AccountType: jsii.String("accountType"),
//   		AssumableRoleArn: jsii.String("assumableRoleArn"),
//
//   		// the properties below are optional
//   		Resources: []interface{}{
//   			&AWSResourceProperty{
//   				ResourceArn: jsii.String("resourceArn"),
//
//   				// the properties below are optional
//   				ResourceMetadata: resourceMetadata,
//   				ResourceType: jsii.String("resourceType"),
//   			},
//   		},
//   		Tags: []KeyValuePairProperty{
//   			&KeyValuePairProperty{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html
//
type CfnAssociation_ServiceConfigurationProperty struct {
	// Configuration for AWS monitor account integration.
	//
	// Specifies the account ID, assumable role ARN, and resources to be monitored in the primary monitoring account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-aws
	//
	Aws interface{} `field:"optional" json:"aws" yaml:"aws"`
	// Configuration for Dynatrace monitoring integration.
	//
	// Specifies the environment ID, resources to monitor, and webhook settings to enable the Agent Space to access Dynatrace metrics, traces, and logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-dynatrace
	//
	Dynatrace interface{} `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// Configuration for Event Channel integration.
	//
	// Specifies webhook settings to enable the Agent Space to receive and process real-time events from external systems.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-eventchannel
	//
	EventChannel interface{} `field:"optional" json:"eventChannel" yaml:"eventChannel"`
	// Configuration for GitHub repository integration.
	//
	// Specifies the repository name, repository ID, owner, and owner type to enable the Agent Space to access code, pull requests, and issues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-github
	//
	GitHub interface{} `field:"optional" json:"gitHub" yaml:"gitHub"`
	// Configuration for GitLab project integration.
	//
	// Specifies the project ID, project path, instance identifier, and webhook settings to enable the Agent Space to access code, merge requests, and issues.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-gitlab
	//
	GitLab interface{} `field:"optional" json:"gitLab" yaml:"gitLab"`
	// Configuration for custom MCP (Model Context Protocol) server integration.
	//
	// Specifies the server name, endpoint URL, available tools, description, and webhook settings to enable the Agent Space to interact with custom MCP servers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-mcpserver
	//
	McpServer interface{} `field:"optional" json:"mcpServer" yaml:"mcpServer"`
	// Configuration for Datadog MCP server integration.
	//
	// Specifies the server name, endpoint URL, optional description, and webhook settings to enable the Agent Space to query metrics, traces, and logs from Datadog.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-mcpserverdatadog
	//
	McpServerDatadog interface{} `field:"optional" json:"mcpServerDatadog" yaml:"mcpServerDatadog"`
	// Configuration for New Relic MCP server integration.
	//
	// Specifies the New Relic account ID and MCP endpoint URL to enable the Agent Space to query metrics, traces, and logs from New Relic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-mcpservernewrelic
	//
	McpServerNewRelic interface{} `field:"optional" json:"mcpServerNewRelic" yaml:"mcpServerNewRelic"`
	// Configuration for Splunk MCP server integration.
	//
	// Specifies the server name, endpoint URL, optional description, and webhook settings to enable the Agent Space to query logs, metrics, and events from Splunk.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-mcpserversplunk
	//
	McpServerSplunk interface{} `field:"optional" json:"mcpServerSplunk" yaml:"mcpServerSplunk"`
	// Configuration for ServiceNow instance integration.
	//
	// Specifies the instance URL, instance ID, and webhook settings to enable the Agent Space to create, update, and manage ServiceNow incidents and change requests.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// Configuration for Slack workspace integration.
	//
	// Specifies the workspace ID, workspace name, and transmission targets to enable the Agent Space to send notifications to designated Slack channels.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-slack
	//
	Slack interface{} `field:"optional" json:"slack" yaml:"slack"`
	// Configuration for AWS source account integration.
	//
	// Specifies the account ID, assumable role ARN, and resources to be monitored in the source account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-sourceaws
	//
	SourceAws interface{} `field:"optional" json:"sourceAws" yaml:"sourceAws"`
}


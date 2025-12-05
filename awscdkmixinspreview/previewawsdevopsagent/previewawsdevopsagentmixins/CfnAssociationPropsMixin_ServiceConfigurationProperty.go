package previewawsdevopsagentmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var resourceMetadata interface{}
//
//   serviceConfigurationProperty := &ServiceConfigurationProperty{
//   	Aws: &AWSConfigurationProperty{
//   		AccountId: jsii.String("accountId"),
//   		AccountType: jsii.String("accountType"),
//   		AssumableRoleArn: jsii.String("assumableRoleArn"),
//   		Resources: []interface{}{
//   			&AWSResourceProperty{
//   				ResourceArn: jsii.String("resourceArn"),
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
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   		EnvId: jsii.String("envId"),
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
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   		InstanceIdentifier: jsii.String("instanceIdentifier"),
//   		ProjectId: jsii.String("projectId"),
//   		ProjectPath: jsii.String("projectPath"),
//   	},
//   	McpServer: &MCPServerConfigurationProperty{
//   		Description: jsii.String("description"),
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   		Tools: []*string{
//   			jsii.String("tools"),
//   		},
//   	},
//   	McpServerDatadog: &MCPServerDatadogConfigurationProperty{
//   		Description: jsii.String("description"),
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   	},
//   	McpServerNewRelic: &MCPServerNewRelicConfigurationProperty{
//   		AccountId: jsii.String("accountId"),
//   		Endpoint: jsii.String("endpoint"),
//   	},
//   	McpServerSplunk: &MCPServerSplunkConfigurationProperty{
//   		Description: jsii.String("description"),
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   		Endpoint: jsii.String("endpoint"),
//   		Name: jsii.String("name"),
//   	},
//   	ServiceNow: &ServiceNowConfigurationProperty{
//   		EnableWebhookUpdates: jsii.Boolean(false),
//   		InstanceId: jsii.String("instanceId"),
//   	},
//   	Slack: &SlackConfigurationProperty{
//   		TransmissionTarget: &SlackTransmissionTargetProperty{
//   			IncidentResponseTarget: &SlackChannelProperty{
//   				ChannelId: jsii.String("channelId"),
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
//   		Resources: []interface{}{
//   			&AWSResourceProperty{
//   				ResourceArn: jsii.String("resourceArn"),
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
type CfnAssociationPropsMixin_ServiceConfigurationProperty struct {
	// AWS association for 'monitor' account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-aws
	//
	Aws interface{} `field:"optional" json:"aws" yaml:"aws"`
	// Dynatrace monitoring configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-dynatrace
	//
	Dynatrace interface{} `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// EventChannelconfiguration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-eventchannel
	//
	EventChannel interface{} `field:"optional" json:"eventChannel" yaml:"eventChannel"`
	// GitHub repository integration configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-github
	//
	GitHub interface{} `field:"optional" json:"gitHub" yaml:"gitHub"`
	// GitLab project integration configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-gitlab
	//
	GitLab interface{} `field:"optional" json:"gitLab" yaml:"gitLab"`
	// MCP server configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-mcpserver
	//
	McpServer interface{} `field:"optional" json:"mcpServer" yaml:"mcpServer"`
	// Datadog MCP server configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-mcpserverdatadog
	//
	McpServerDatadog interface{} `field:"optional" json:"mcpServerDatadog" yaml:"mcpServerDatadog"`
	// NewRelic MCP server configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-mcpservernewrelic
	//
	McpServerNewRelic interface{} `field:"optional" json:"mcpServerNewRelic" yaml:"mcpServerNewRelic"`
	// Splunk MCP server configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-mcpserversplunk
	//
	McpServerSplunk interface{} `field:"optional" json:"mcpServerSplunk" yaml:"mcpServerSplunk"`
	// ServiceNow integration configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-servicenow
	//
	ServiceNow interface{} `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// Slack workspace integration configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-slack
	//
	Slack interface{} `field:"optional" json:"slack" yaml:"slack"`
	// AWS association for 'source' account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-serviceconfiguration.html#cfn-devopsagent-association-serviceconfiguration-sourceaws
	//
	SourceAws interface{} `field:"optional" json:"sourceAws" yaml:"sourceAws"`
}


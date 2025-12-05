package previewawsdevopsagentmixins


// Properties for CfnAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var resourceMetadata interface{}
//
//   cfnAssociationMixinProps := &CfnAssociationMixinProps{
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	Configuration: &ServiceConfigurationProperty{
//   		Aws: &AWSConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			AccountType: jsii.String("accountType"),
//   			AssumableRoleArn: jsii.String("assumableRoleArn"),
//   			Resources: []interface{}{
//   				&AWSResourceProperty{
//   					ResourceArn: jsii.String("resourceArn"),
//   					ResourceMetadata: resourceMetadata,
//   					ResourceType: jsii.String("resourceType"),
//   				},
//   			},
//   			Tags: []KeyValuePairProperty{
//   				&KeyValuePairProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   		Dynatrace: &DynatraceConfigurationProperty{
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			EnvId: jsii.String("envId"),
//   			Resources: []*string{
//   				jsii.String("resources"),
//   			},
//   		},
//   		EventChannel: &EventChannelConfigurationProperty{
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   		},
//   		GitHub: &GitHubConfigurationProperty{
//   			Owner: jsii.String("owner"),
//   			OwnerType: jsii.String("ownerType"),
//   			RepoId: jsii.String("repoId"),
//   			RepoName: jsii.String("repoName"),
//   		},
//   		GitLab: &GitLabConfigurationProperty{
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			InstanceIdentifier: jsii.String("instanceIdentifier"),
//   			ProjectId: jsii.String("projectId"),
//   			ProjectPath: jsii.String("projectPath"),
//   		},
//   		McpServer: &MCPServerConfigurationProperty{
//   			Description: jsii.String("description"),
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   			Tools: []*string{
//   				jsii.String("tools"),
//   			},
//   		},
//   		McpServerDatadog: &MCPServerDatadogConfigurationProperty{
//   			Description: jsii.String("description"),
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   		},
//   		McpServerNewRelic: &MCPServerNewRelicConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   		McpServerSplunk: &MCPServerSplunkConfigurationProperty{
//   			Description: jsii.String("description"),
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   		},
//   		ServiceNow: &ServiceNowConfigurationProperty{
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			InstanceId: jsii.String("instanceId"),
//   		},
//   		Slack: &SlackConfigurationProperty{
//   			TransmissionTarget: &SlackTransmissionTargetProperty{
//   				IncidentResponseTarget: &SlackChannelProperty{
//   					ChannelId: jsii.String("channelId"),
//   					ChannelName: jsii.String("channelName"),
//   				},
//   			},
//   			WorkspaceId: jsii.String("workspaceId"),
//   			WorkspaceName: jsii.String("workspaceName"),
//   		},
//   		SourceAws: &SourceAwsConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			AccountType: jsii.String("accountType"),
//   			AssumableRoleArn: jsii.String("assumableRoleArn"),
//   			Resources: []interface{}{
//   				&AWSResourceProperty{
//   					ResourceArn: jsii.String("resourceArn"),
//   					ResourceMetadata: resourceMetadata,
//   					ResourceType: jsii.String("resourceType"),
//   				},
//   			},
//   			Tags: []KeyValuePairProperty{
//   				&KeyValuePairProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	ServiceId: jsii.String("serviceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html
//
type CfnAssociationMixinProps struct {
	// The unique identifier of the AgentSpace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-agentspaceid
	//
	AgentSpaceId *string `field:"optional" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The configuration that directs how AgentSpace interacts with the given service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// The identifier for the associated service.
	//
	// For SourceAws and Aws configurations, this must be 'aws'. For all other service types, this is a UUID generated from the RegisterService command
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-serviceid
	//
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
}


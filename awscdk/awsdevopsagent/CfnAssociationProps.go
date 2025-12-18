package awsdevopsagent


// Properties for defining a `CfnAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var resourceMetadata interface{}
//
//   cfnAssociationProps := &CfnAssociationProps{
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	Configuration: &ServiceConfigurationProperty{
//   		Aws: &AWSConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			AccountType: jsii.String("accountType"),
//   			AssumableRoleArn: jsii.String("assumableRoleArn"),
//
//   			// the properties below are optional
//   			Resources: []interface{}{
//   				&AWSResourceProperty{
//   					ResourceArn: jsii.String("resourceArn"),
//
//   					// the properties below are optional
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
//   			EnvId: jsii.String("envId"),
//
//   			// the properties below are optional
//   			EnableWebhookUpdates: jsii.Boolean(false),
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
//   			ProjectId: jsii.String("projectId"),
//   			ProjectPath: jsii.String("projectPath"),
//
//   			// the properties below are optional
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			InstanceIdentifier: jsii.String("instanceIdentifier"),
//   		},
//   		McpServer: &MCPServerConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   			Tools: []*string{
//   				jsii.String("tools"),
//   			},
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   		},
//   		McpServerDatadog: &MCPServerDatadogConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   		},
//   		McpServerNewRelic: &MCPServerNewRelicConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   		McpServerSplunk: &MCPServerSplunkConfigurationProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   		},
//   		ServiceNow: &ServiceNowConfigurationProperty{
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			InstanceId: jsii.String("instanceId"),
//   		},
//   		Slack: &SlackConfigurationProperty{
//   			TransmissionTarget: &SlackTransmissionTargetProperty{
//   				IncidentResponseTarget: &SlackChannelProperty{
//   					ChannelId: jsii.String("channelId"),
//
//   					// the properties below are optional
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
//
//   			// the properties below are optional
//   			Resources: []interface{}{
//   				&AWSResourceProperty{
//   					ResourceArn: jsii.String("resourceArn"),
//
//   					// the properties below are optional
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
//
//   	// the properties below are optional
//   	LinkedAssociationIds: []*string{
//   		jsii.String("linkedAssociationIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html
//
type CfnAssociationProps struct {
	// The unique identifier of the Agent Space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-agentspaceid
	//
	AgentSpaceId *string `field:"required" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The configuration that directs how the Agent Space interacts with the given service.
	//
	// You can specify only one configuration type per association.
	//
	// *Allowed Values* : `SourceAws` | `Aws` | `GitHub` | `GitLab` | `Slack` | `Dynatrace` | `ServiceNow` | `MCPServer` | `MCPServerNewRelic` | `MCPServerDatadog` | `MCPServerSplunk` | `EventChannel`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The identifier for the associated service.
	//
	// For `SourceAws` and `Aws` configurations, this must be `aws` . For all other service types, this is a UUID generated from the RegisterService command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-serviceid
	//
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// Set of linked association IDs for parent-child relationships.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-linkedassociationids
	//
	LinkedAssociationIds *[]*string `field:"optional" json:"linkedAssociationIds" yaml:"linkedAssociationIds"`
}


package awsdevopsagent


// Properties for CfnAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var resourceMetadata interface{}
//
//   cfnAssociationMixinProps := &CfnAssociationMixinProps{
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	Capabilities: &CapabilitiesProperty{
//   		CodeChangeRiskAnalysis: &CapabilityConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			TriggerFilterGroups: []interface{}{
//   				&TriggerFilterGroupProperty{
//   					Events: []*string{
//   						jsii.String("events"),
//   					},
//   					TargetBranches: &PatternFilterProperty{
//   						Patterns: []*string{
//   							jsii.String("patterns"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		CodeChangeRiskAnalysisAutomatedTesting: &AutomatedTestingCapabilityConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	Configuration: &ServiceConfigurationProperty{
//   		Aws: &AWSConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			AccountType: jsii.String("accountType"),
//   			AgentElevatedRoleArn: jsii.String("agentElevatedRoleArn"),
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
//   		Azure: &AzureConfigurationProperty{
//   			SubscriptionId: jsii.String("subscriptionId"),
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
//   			ReleaseManagementAssociationId: jsii.String("releaseManagementAssociationId"),
//   			RepoId: jsii.String("repoId"),
//   			RepoName: jsii.String("repoName"),
//   		},
//   		GitLab: &GitLabConfigurationProperty{
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			InstanceIdentifier: jsii.String("instanceIdentifier"),
//   			ProjectId: jsii.String("projectId"),
//   			ProjectPath: jsii.String("projectPath"),
//   			ReleaseManagementAssociationId: jsii.String("releaseManagementAssociationId"),
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
//   			EnabledElevatedTools: []*string{
//   				jsii.String("enabledElevatedTools"),
//   			},
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   		},
//   		McpServerGrafana: &MCPServerGrafanaConfigurationProperty{
//   			EnabledElevatedTools: []*string{
//   				jsii.String("enabledElevatedTools"),
//   			},
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			Endpoint: jsii.String("endpoint"),
//   			Tools: []*string{
//   				jsii.String("tools"),
//   			},
//   		},
//   		McpServerNewRelic: &MCPServerNewRelicConfigurationProperty{
//   			AccountId: jsii.String("accountId"),
//   			Endpoint: jsii.String("endpoint"),
//   		},
//   		McpServerSigV4: &MCPServerSigV4ConfigurationProperty{
//   			Tools: []*string{
//   				jsii.String("tools"),
//   			},
//   		},
//   		McpServerSplunk: &MCPServerSplunkConfigurationProperty{
//   			Description: jsii.String("description"),
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			Endpoint: jsii.String("endpoint"),
//   			Name: jsii.String("name"),
//   		},
//   		PagerDuty: &PagerDutyConfigurationProperty{
//   			CustomerEmail: jsii.String("customerEmail"),
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			Services: []*string{
//   				jsii.String("services"),
//   			},
//   		},
//   		ReleaseManagement: &ReleaseManagementConfigurationProperty{
//   			Name: jsii.String("name"),
//   			NetworkAccess: &NetworkAccessConfigurationProperty{
//   				PrivateAccess: &PrivateNetworkAccessProperty{
//   					PrivateConnectionName: jsii.String("privateConnectionName"),
//   					RuntimeRoleArn: jsii.String("runtimeRoleArn"),
//   				},
//   			},
//   		},
//   		ServiceNow: &ServiceNowConfigurationProperty{
//   			EnableWebhookUpdates: jsii.Boolean(false),
//   			InstanceId: jsii.String("instanceId"),
//   		},
//   		Slack: &SlackConfigurationProperty{
//   			Bidirectional: &SlackBidirectionalConfigurationProperty{
//   				Enabled: jsii.Boolean(false),
//   				RoleArn: jsii.String("roleArn"),
//   			},
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
//   			AgentElevatedRoleArn: jsii.String("agentElevatedRoleArn"),
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
//   	LinkedAssociationIds: []*string{
//   		jsii.String("linkedAssociationIds"),
//   	},
//   	ServiceId: jsii.String("serviceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html
//
type CfnAssociationMixinProps struct {
	// The unique identifier of the Agent Space.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-agentspaceid
	//
	AgentSpaceId *string `field:"optional" json:"agentSpaceId" yaml:"agentSpaceId"`
	// Automated AWS DevOps Agent capabilities for the repository in this association.
	//
	// Code change risk analysis analyzes pull requests and merge requests for deployment risk, and automated testing extends that analysis by running tests against the change. Capabilities are supported for GitHub and GitLab associations. When this property is set, it fully defines both capabilities below: a capability omitted here is removed from the association. Removing this property from a template that previously set it disables both capabilities. Capabilities configured outside CloudFormation are left unchanged only while this property has never been set. On rollback, both capabilities are restored to the previous template's values, not to settings made outside CloudFormation. Other capabilities on the association are not affected. Updates read and rewrite the association's full capability set, so a capability change made outside CloudFormation while a stack update is in progress may be overwritten.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-capabilities
	//
	Capabilities interface{} `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The configuration that directs how the Agent Space interacts with the given service.
	//
	// You can specify only one configuration type per association.
	//
	// *Allowed Values* : `SourceAws` | `Aws` | `GitHub` | `GitLab` | `Slack` | `Dynatrace` | `ServiceNow` | `MCPServer` | `MCPServerNewRelic` | `MCPServerDatadog` | `MCPServerSplunk` | `EventChannel`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Set of linked association IDs for parent-child relationships.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-linkedassociationids
	//
	LinkedAssociationIds *[]*string `field:"optional" json:"linkedAssociationIds" yaml:"linkedAssociationIds"`
	// The identifier for the associated service.
	//
	// For `SourceAws` and `Aws` configurations, this must be `aws` . For all other service types, this is a UUID generated from the RegisterService command.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html#cfn-devopsagent-association-serviceid
	//
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
}


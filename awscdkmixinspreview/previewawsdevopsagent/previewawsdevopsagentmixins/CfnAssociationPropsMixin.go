package previewawsdevopsagentmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdevopsagent/previewawsdevopsagentmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::DevOpsAgent::Association` resource specifies an association between an Agent Space and a service, defining how the Agent Space interacts with external services like GitHub, Slack, AWS accounts, and others.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var resourceMetadata interface{}
//
//   cfnAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnAssociationPropsMixin(&CfnAssociationMixinProps{
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
//   	LinkedAssociationIds: []*string{
//   		jsii.String("linkedAssociationIds"),
//   	},
//   	ServiceId: jsii.String("serviceId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-association.html
//
type CfnAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAssociationPropsMixin
type jsiiProxy_CfnAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAssociationPropsMixin) Props() *CfnAssociationMixinProps {
	var returns *CfnAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DevOpsAgent::Association`.
func NewCfnAssociationPropsMixin(props *CfnAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DevOpsAgent::Association`.
func NewCfnAssociationPropsMixin_Override(c CfnAssociationPropsMixin, props *CfnAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}


package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::BedrockAgentCore::Harness - a managed agentic loop service that provides a turnkey solution for running stateful, tool-equipped AI agents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var awsIam interface{}
//   var disabled interface{}
//   var inputSchema interface{}
//   var mergeStrategy IMergeStrategy
//   var none interface{}
//
//   cfnHarnessPropsMixin := awscdkcfnpropertymixins.Aws_bedrockagentcore.NewCfnHarnessPropsMixin(&CfnHarnessMixinProps{
//   	AllowedTools: []*string{
//   		jsii.String("allowedTools"),
//   	},
//   	AuthorizerConfiguration: &AuthorizerConfigurationProperty{
//   		CustomJwtAuthorizer: &CustomJWTAuthorizerConfigurationProperty{
//   			AllowedAudience: []*string{
//   				jsii.String("allowedAudience"),
//   			},
//   			AllowedClients: []*string{
//   				jsii.String("allowedClients"),
//   			},
//   			AllowedScopes: []*string{
//   				jsii.String("allowedScopes"),
//   			},
//   			CustomClaims: []interface{}{
//   				&CustomClaimValidationTypeProperty{
//   					AuthorizingClaimMatchValue: &AuthorizingClaimMatchValueTypeProperty{
//   						ClaimMatchOperator: jsii.String("claimMatchOperator"),
//   						ClaimMatchValue: &ClaimMatchValueTypeProperty{
//   							MatchValueString: jsii.String("matchValueString"),
//   							MatchValueStringList: []*string{
//   								jsii.String("matchValueStringList"),
//   							},
//   						},
//   					},
//   					InboundTokenClaimName: jsii.String("inboundTokenClaimName"),
//   					InboundTokenClaimValueType: jsii.String("inboundTokenClaimValueType"),
//   				},
//   			},
//   			DiscoveryUrl: jsii.String("discoveryUrl"),
//   			PrivateEndpoint: &PrivateEndpointProperty{
//   				ManagedVpcResource: &ManagedVpcResourceProperty{
//   					EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   					RoutingDomain: jsii.String("routingDomain"),
//   					SecurityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					SubnetIds: []*string{
//   						jsii.String("subnetIds"),
//   					},
//   					Tags: map[string]*string{
//   						"tagsKey": jsii.String("tags"),
//   					},
//   					VpcIdentifier: jsii.String("vpcIdentifier"),
//   				},
//   				SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   					ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   				},
//   			},
//   			PrivateEndpointOverrides: []interface{}{
//   				&PrivateEndpointOverrideProperty{
//   					Domain: jsii.String("domain"),
//   					PrivateEndpoint: &PrivateEndpointProperty{
//   						ManagedVpcResource: &ManagedVpcResourceProperty{
//   							EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   							RoutingDomain: jsii.String("routingDomain"),
//   							SecurityGroupIds: []*string{
//   								jsii.String("securityGroupIds"),
//   							},
//   							SubnetIds: []*string{
//   								jsii.String("subnetIds"),
//   							},
//   							Tags: map[string]*string{
//   								"tagsKey": jsii.String("tags"),
//   							},
//   							VpcIdentifier: jsii.String("vpcIdentifier"),
//   						},
//   						SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   							ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Environment: &HarnessEnvironmentProviderProperty{
//   		AgentCoreRuntimeEnvironment: &HarnessAgentCoreRuntimeEnvironmentProperty{
//   			AgentRuntimeArn: jsii.String("agentRuntimeArn"),
//   			AgentRuntimeId: jsii.String("agentRuntimeId"),
//   			AgentRuntimeName: jsii.String("agentRuntimeName"),
//   			FilesystemConfigurations: []interface{}{
//   				&FilesystemConfigurationProperty{
//   					EfsAccessPoint: &EfsAccessPointConfigurationProperty{
//   						AccessPointArn: jsii.String("accessPointArn"),
//   						MountPath: jsii.String("mountPath"),
//   					},
//   					S3FilesAccessPoint: &S3FilesAccessPointConfigurationProperty{
//   						AccessPointArn: jsii.String("accessPointArn"),
//   						MountPath: jsii.String("mountPath"),
//   					},
//   					SessionStorage: &SessionStorageConfigurationProperty{
//   						MountPath: jsii.String("mountPath"),
//   					},
//   				},
//   			},
//   			LifecycleConfiguration: &LifecycleConfigurationProperty{
//   				IdleRuntimeSessionTimeout: jsii.Number(123),
//   				MaxLifetime: jsii.Number(123),
//   			},
//   			NetworkConfiguration: &NetworkConfigurationProperty{
//   				NetworkMode: jsii.String("networkMode"),
//   				NetworkModeConfig: &VpcConfigProperty{
//   					SecurityGroups: []*string{
//   						jsii.String("securityGroups"),
//   					},
//   					Subnets: []*string{
//   						jsii.String("subnets"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	EnvironmentArtifact: &HarnessEnvironmentArtifactProperty{
//   		ContainerConfiguration: &ContainerConfigurationProperty{
//   			ContainerUri: jsii.String("containerUri"),
//   		},
//   	},
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	HarnessName: jsii.String("harnessName"),
//   	MaxIterations: jsii.Number(123),
//   	MaxTokens: jsii.Number(123),
//   	Memory: &HarnessMemoryConfigurationProperty{
//   		AgentCoreMemoryConfiguration: &HarnessAgentCoreMemoryConfigurationProperty{
//   			ActorId: jsii.String("actorId"),
//   			Arn: jsii.String("arn"),
//   			MessagesCount: jsii.Number(123),
//   			RetrievalConfig: map[string]interface{}{
//   				"retrievalConfigKey": &HarnessAgentCoreMemoryRetrievalConfigProperty{
//   					"relevanceScore": jsii.String("relevanceScore"),
//   					"strategyId": jsii.String("strategyId"),
//   					"topK": jsii.String("topK"),
//   				},
//   			},
//   		},
//   		Disabled: disabled,
//   		ManagedMemoryConfiguration: &HarnessManagedMemoryConfigurationProperty{
//   			Arn: jsii.String("arn"),
//   			EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   			EventExpiryDuration: jsii.Number(123),
//   			Strategies: []*string{
//   				jsii.String("strategies"),
//   			},
//   		},
//   	},
//   	Model: &HarnessModelConfigurationProperty{
//   		BedrockModelConfig: &HarnessBedrockModelConfigProperty{
//   			ApiFormat: jsii.String("apiFormat"),
//   			MaxTokens: jsii.Number(123),
//   			ModelId: jsii.String("modelId"),
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   		GeminiModelConfig: &HarnessGeminiModelConfigProperty{
//   			ApiKeyArn: jsii.String("apiKeyArn"),
//   			MaxTokens: jsii.Number(123),
//   			ModelId: jsii.String("modelId"),
//   			Temperature: jsii.Number(123),
//   			TopK: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   		LiteLlmModelConfig: &HarnessLiteLlmModelConfigProperty{
//   			ApiBase: jsii.String("apiBase"),
//   			ApiKeyArn: jsii.String("apiKeyArn"),
//   			MaxTokens: jsii.Number(123),
//   			ModelId: jsii.String("modelId"),
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   		OpenAiModelConfig: &HarnessOpenAiModelConfigProperty{
//   			ApiFormat: jsii.String("apiFormat"),
//   			ApiKeyArn: jsii.String("apiKeyArn"),
//   			MaxTokens: jsii.Number(123),
//   			ModelId: jsii.String("modelId"),
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   	},
//   	Skills: []interface{}{
//   		&HarnessSkillProperty{
//   			AwsSkills: &HarnessSkillAwsSkillsSourceProperty{
//   				Paths: []*string{
//   					jsii.String("paths"),
//   				},
//   			},
//   			Git: &HarnessSkillGitSourceProperty{
//   				Auth: &HarnessSkillGitAuthProperty{
//   					CredentialArn: jsii.String("credentialArn"),
//   					Username: jsii.String("username"),
//   				},
//   				Path: jsii.String("path"),
//   				Url: jsii.String("url"),
//   			},
//   			Path: jsii.String("path"),
//   			S3: &HarnessSkillS3SourceProperty{
//   				Uri: jsii.String("uri"),
//   			},
//   		},
//   	},
//   	SystemPrompt: []interface{}{
//   		&HarnessSystemContentBlockProperty{
//   			Text: jsii.String("text"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeoutSeconds: jsii.Number(123),
//   	Tools: []interface{}{
//   		&HarnessToolProperty{
//   			Config: &HarnessToolConfigurationProperty{
//   				AgentCoreBrowser: &HarnessAgentCoreBrowserConfigProperty{
//   					BrowserArn: jsii.String("browserArn"),
//   				},
//   				AgentCoreCodeInterpreter: &HarnessAgentCoreCodeInterpreterConfigProperty{
//   					CodeInterpreterArn: jsii.String("codeInterpreterArn"),
//   				},
//   				AgentCoreGateway: &HarnessAgentCoreGatewayConfigProperty{
//   					GatewayArn: jsii.String("gatewayArn"),
//   					OutboundAuth: &HarnessGatewayOutboundAuthProperty{
//   						AwsIam: awsIam,
//   						None: none,
//   						Oauth: &OAuthCredentialProviderProperty{
//   							CustomParameters: map[string]*string{
//   								"customParametersKey": jsii.String("customParameters"),
//   							},
//   							DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   							GrantType: jsii.String("grantType"),
//   							ProviderArn: jsii.String("providerArn"),
//   							Scopes: []*string{
//   								jsii.String("scopes"),
//   							},
//   						},
//   					},
//   				},
//   				InlineFunction: &HarnessInlineFunctionConfigProperty{
//   					Description: jsii.String("description"),
//   					InputSchema: inputSchema,
//   				},
//   				RemoteMcp: &HarnessRemoteMcpConfigProperty{
//   					Headers: map[string]*string{
//   						"headersKey": jsii.String("headers"),
//   					},
//   					Url: jsii.String("url"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Truncation: &HarnessTruncationConfigurationProperty{
//   		Config: &HarnessTruncationStrategyConfigurationProperty{
//   			SlidingWindow: &HarnessSlidingWindowConfigurationProperty{
//   				MessagesCount: jsii.Number(123),
//   			},
//   			Summarization: &HarnessSummarizationConfigurationProperty{
//   				PreserveRecentMessages: jsii.Number(123),
//   				SummarizationSystemPrompt: jsii.String("summarizationSystemPrompt"),
//   				SummaryRatio: jsii.Number(123),
//   			},
//   		},
//   		Strategy: jsii.String("strategy"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html
//
type CfnHarnessPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnHarnessMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnHarnessPropsMixin
type jsiiProxy_CfnHarnessPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnHarnessPropsMixin) Props() *CfnHarnessMixinProps {
	var returns *CfnHarnessMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarnessPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::BedrockAgentCore::Harness`.
func NewCfnHarnessPropsMixin(props *CfnHarnessMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnHarnessPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnHarnessPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHarnessPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnHarnessPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::BedrockAgentCore::Harness`.
func NewCfnHarnessPropsMixin_Override(c CfnHarnessPropsMixin, props *CfnHarnessMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnHarnessPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnHarnessPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHarnessPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnHarnessPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHarnessPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_bedrockagentcore.CfnHarnessPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHarnessPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnHarnessPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


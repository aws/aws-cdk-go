package awsbedrockagentcore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnHarnessPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var awsIam interface{}
//   var disabled interface{}
//   var inputSchema interface{}
//   var none interface{}
//
//   cfnHarnessMixinProps := &CfnHarnessMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html
//
type CfnHarnessMixinProps struct {
	// The tools that the agent is allowed to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-allowedtools
	//
	AllowedTools *[]*string `field:"optional" json:"allowedTools" yaml:"allowedTools"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-authorizerconfiguration
	//
	AuthorizerConfiguration interface{} `field:"optional" json:"authorizerConfiguration" yaml:"authorizerConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-environmentartifact
	//
	EnvironmentArtifact interface{} `field:"optional" json:"environmentArtifact" yaml:"environmentArtifact"`
	// Environment variables to set in the harness runtime environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The ARN of the IAM role that the harness assumes when running.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The name of the harness.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-harnessname
	//
	HarnessName *string `field:"optional" json:"harnessName" yaml:"harnessName"`
	// The maximum number of iterations the agent loop can execute per invocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-maxiterations
	//
	MaxIterations *float64 `field:"optional" json:"maxIterations" yaml:"maxIterations"`
	// The maximum number of tokens the agent can generate per iteration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-maxtokens
	//
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-memory
	//
	Memory interface{} `field:"optional" json:"memory" yaml:"memory"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-model
	//
	Model interface{} `field:"optional" json:"model" yaml:"model"`
	// The skills available to the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-skills
	//
	Skills interface{} `field:"optional" json:"skills" yaml:"skills"`
	// The system prompt that defines the agent's behavior.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-systemprompt
	//
	SystemPrompt interface{} `field:"optional" json:"systemPrompt" yaml:"systemPrompt"`
	// Tags to apply to the harness resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The maximum duration in seconds for the agent loop execution per invocation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-timeoutseconds
	//
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The tools available to the agent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-tools
	//
	Tools interface{} `field:"optional" json:"tools" yaml:"tools"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html#cfn-bedrockagentcore-harness-truncation
	//
	Truncation interface{} `field:"optional" json:"truncation" yaml:"truncation"`
}


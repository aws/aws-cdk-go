package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrockagentcore"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::BedrockAgentCore::Harness - a managed agentic loop service that provides a turnkey solution for running stateful, tool-equipped AI agents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var awsIam interface{}
//   var disabled interface{}
//   var inputSchema interface{}
//   var none interface{}
//
//   cfnHarness := awscdk.Aws_bedrockagentcore.NewCfnHarness(this, jsii.String("MyCfnHarness"), &CfnHarnessProps{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	HarnessName: jsii.String("harnessName"),
//   	Model: &HarnessModelConfigurationProperty{
//   		BedrockModelConfig: &HarnessBedrockModelConfigProperty{
//   			ModelId: jsii.String("modelId"),
//
//   			// the properties below are optional
//   			ApiFormat: jsii.String("apiFormat"),
//   			MaxTokens: jsii.Number(123),
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   		GeminiModelConfig: &HarnessGeminiModelConfigProperty{
//   			ApiKeyArn: jsii.String("apiKeyArn"),
//   			ModelId: jsii.String("modelId"),
//
//   			// the properties below are optional
//   			MaxTokens: jsii.Number(123),
//   			Temperature: jsii.Number(123),
//   			TopK: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   		LiteLlmModelConfig: &HarnessLiteLlmModelConfigProperty{
//   			ModelId: jsii.String("modelId"),
//
//   			// the properties below are optional
//   			ApiBase: jsii.String("apiBase"),
//   			ApiKeyArn: jsii.String("apiKeyArn"),
//   			MaxTokens: jsii.Number(123),
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   		OpenAiModelConfig: &HarnessOpenAiModelConfigProperty{
//   			ApiKeyArn: jsii.String("apiKeyArn"),
//   			ModelId: jsii.String("modelId"),
//
//   			// the properties below are optional
//   			ApiFormat: jsii.String("apiFormat"),
//   			MaxTokens: jsii.Number(123),
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AllowedTools: []*string{
//   		jsii.String("allowedTools"),
//   	},
//   	AuthorizerConfiguration: &AuthorizerConfigurationProperty{
//   		CustomJwtAuthorizer: &CustomJWTAuthorizerConfigurationProperty{
//   			DiscoveryUrl: jsii.String("discoveryUrl"),
//
//   			// the properties below are optional
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
//   			PrivateEndpoint: &PrivateEndpointProperty{
//   				ManagedVpcResource: &ManagedVpcResourceProperty{
//   					EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   					SubnetIds: []*string{
//   						jsii.String("subnetIds"),
//   					},
//   					VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   					// the properties below are optional
//   					RoutingDomain: jsii.String("routingDomain"),
//   					SecurityGroupIds: []*string{
//   						jsii.String("securityGroupIds"),
//   					},
//   					Tags: map[string]*string{
//   						"tagsKey": jsii.String("tags"),
//   					},
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
//   							SubnetIds: []*string{
//   								jsii.String("subnetIds"),
//   							},
//   							VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   							// the properties below are optional
//   							RoutingDomain: jsii.String("routingDomain"),
//   							SecurityGroupIds: []*string{
//   								jsii.String("securityGroupIds"),
//   							},
//   							Tags: map[string]*string{
//   								"tagsKey": jsii.String("tags"),
//   							},
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
//
//   				// the properties below are optional
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
//   	MaxIterations: jsii.Number(123),
//   	MaxTokens: jsii.Number(123),
//   	Memory: &HarnessMemoryConfigurationProperty{
//   		AgentCoreMemoryConfiguration: &HarnessAgentCoreMemoryConfigurationProperty{
//   			Arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			ActorId: jsii.String("actorId"),
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
//   	Skills: []interface{}{
//   		&HarnessSkillProperty{
//   			AwsSkills: &HarnessSkillAwsSkillsSourceProperty{
//   				Paths: []*string{
//   					jsii.String("paths"),
//   				},
//   			},
//   			Git: &HarnessSkillGitSourceProperty{
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				Auth: &HarnessSkillGitAuthProperty{
//   					CredentialArn: jsii.String("credentialArn"),
//
//   					// the properties below are optional
//   					Username: jsii.String("username"),
//   				},
//   				Path: jsii.String("path"),
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
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			Config: &HarnessToolConfigurationProperty{
//   				AgentCoreBrowser: &HarnessAgentCoreBrowserConfigProperty{
//   					BrowserArn: jsii.String("browserArn"),
//   				},
//   				AgentCoreCodeInterpreter: &HarnessAgentCoreCodeInterpreterConfigProperty{
//   					CodeInterpreterArn: jsii.String("codeInterpreterArn"),
//   				},
//   				AgentCoreGateway: &HarnessAgentCoreGatewayConfigProperty{
//   					GatewayArn: jsii.String("gatewayArn"),
//
//   					// the properties below are optional
//   					OutboundAuth: &HarnessGatewayOutboundAuthProperty{
//   						AwsIam: awsIam,
//   						None: none,
//   						Oauth: &OAuthCredentialProviderProperty{
//   							ProviderArn: jsii.String("providerArn"),
//   							Scopes: []*string{
//   								jsii.String("scopes"),
//   							},
//
//   							// the properties below are optional
//   							CustomParameters: map[string]*string{
//   								"customParametersKey": jsii.String("customParameters"),
//   							},
//   							DefaultReturnUrl: jsii.String("defaultReturnUrl"),
//   							GrantType: jsii.String("grantType"),
//   						},
//   					},
//   				},
//   				InlineFunction: &HarnessInlineFunctionConfigProperty{
//   					Description: jsii.String("description"),
//   					InputSchema: inputSchema,
//   				},
//   				RemoteMcp: &HarnessRemoteMcpConfigProperty{
//   					Url: jsii.String("url"),
//
//   					// the properties below are optional
//   					Headers: map[string]*string{
//   						"headersKey": jsii.String("headers"),
//   					},
//   				},
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Truncation: &HarnessTruncationConfigurationProperty{
//   		Strategy: jsii.String("strategy"),
//
//   		// the properties below are optional
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
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-harness.html
//
type CfnHarness interface {
	awscdk.CfnResource
	awscdk.IInspectable
	interfacesawsbedrockagentcore.IHarnessRef
	awscdk.ITaggableV2
	// The tools that the agent is allowed to use.
	AllowedTools() *[]*string
	SetAllowedTools(val *[]*string)
	// The Amazon Resource Name (ARN) of the harness.
	AttrArn() *string
	// The timestamp when the harness was created.
	AttrCreatedAt() *string
	// The ARN of the underlying AgentCore Runtime.
	AttrEnvironmentAgentCoreRuntimeEnvironmentAgentRuntimeArn() *string
	// The ID of the underlying AgentCore Runtime.
	AttrEnvironmentAgentCoreRuntimeEnvironmentAgentRuntimeId() *string
	// The name of the underlying AgentCore Runtime.
	AttrEnvironmentAgentCoreRuntimeEnvironmentAgentRuntimeName() *string
	// The unique identifier of the harness.
	AttrHarnessId() *string
	// The ARN of the managed memory resource.
	//
	// Read-only, populated by the service.
	AttrMemoryManagedMemoryConfigurationArn() *string
	AttrStatus() *string
	// The timestamp when the harness was last updated.
	AttrUpdatedAt() *string
	// The version of the harness.
	//
	// Incremented on every successful update.
	AttrVersion() *string
	AuthorizerConfiguration() interface{}
	SetAuthorizerConfiguration(val interface{})
	// Tag Manager which manages the tags for this resource.
	CdkTagManager() awscdk.TagManager
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	CfnPropertyNames() *map[string]*string
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	Env() *interfaces.ResourceEnvironment
	Environment() interface{}
	SetEnvironment(val interface{})
	EnvironmentArtifact() interface{}
	SetEnvironmentArtifact(val interface{})
	// Environment variables to set in the harness runtime environment.
	EnvironmentVariables() interface{}
	SetEnvironmentVariables(val interface{})
	// The ARN of the IAM role that the harness assumes when running.
	ExecutionRoleArn() *string
	SetExecutionRoleArn(val *string)
	// The name of the harness.
	HarnessName() *string
	SetHarnessName(val *string)
	// A reference to a Harness resource.
	HarnessRef() *interfacesawsbedrockagentcore.HarnessReference
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The maximum number of iterations the agent loop can execute per invocation.
	MaxIterations() *float64
	SetMaxIterations(val *float64)
	// The maximum number of tokens the agent can generate per iteration.
	MaxTokens() *float64
	SetMaxTokens(val *float64)
	Memory() interface{}
	SetMemory(val interface{})
	Model() interface{}
	SetModel(val interface{})
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The skills available to the agent.
	Skills() interface{}
	SetSkills(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// The system prompt that defines the agent's behavior.
	SystemPrompt() interface{}
	SetSystemPrompt(val interface{})
	// Tags to apply to the harness resource.
	Tags() *[]*awscdk.CfnTag
	SetTags(val *[]*awscdk.CfnTag)
	// The maximum duration in seconds for the agent loop execution per invocation.
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	// The tools available to the agent.
	Tools() interface{}
	SetTools(val interface{})
	Truncation() interface{}
	SetTruncation(val interface{})
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This method has been renamed to `addResourceDependency` to more clearly
	// set it apart from `construct.node.addDependency`. See the documentation
	// of that function for more details.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method has been renamed to `addResourceDependency`, which makes it
	// more clear that this method operates at a different level from the
	// construct-level `construct.node.addDependency()` mechanism.
	// Deprecated: Use `addResourceDependency` instead.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	//
	// This method only adds dependencies between L1 resources. If you are
	// looking for a generic construct-to-construct dependency mechanism that works
	// for all constructs including L2s, use `construct.node.addDependency` instead.
	AddResourceDependency(target awscdk.CfnResource, reason *string)
	// Sets the cross-stack reference strength for this resource.
	//
	// When set, any cross-stack reference to this resource will use the specified
	// strength instead of the global default from the consuming stack's context.
	ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength)
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	CfnPropertyName(cdkPropertyName *string) *string
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources and stacks this resource depends on.
	//
	// For resources depended on directly, returns the `CfnResource` object. For
	// dependencies on other stacks, returns the `Stack` object. The order of the
	// array is not guaranteed.
	ObtainDependencies() *[]interface{}
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	// Deprecated: Use `removeResourceDependency` instead.
	RemoveDependency(target awscdk.CfnResource)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveResourceDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CfnHarness
type jsiiProxy_CfnHarness struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
	internal.Type__interfacesawsbedrockagentcoreIHarnessRef
	internal.Type__awscdkITaggableV2
}

func (j *jsiiProxy_CfnHarness) AllowedTools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedTools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AttrCreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AttrEnvironmentAgentCoreRuntimeEnvironmentAgentRuntimeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentAgentCoreRuntimeEnvironmentAgentRuntimeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AttrEnvironmentAgentCoreRuntimeEnvironmentAgentRuntimeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentAgentCoreRuntimeEnvironmentAgentRuntimeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AttrEnvironmentAgentCoreRuntimeEnvironmentAgentRuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrEnvironmentAgentCoreRuntimeEnvironmentAgentRuntimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AttrHarnessId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrHarnessId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AttrMemoryManagedMemoryConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrMemoryManagedMemoryConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AttrStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AttrUpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AttrVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) AuthorizerConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) CdkTagManager() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"cdkTagManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) CfnPropertyNames() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"cfnPropertyNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Environment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) EnvironmentArtifact() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentArtifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) EnvironmentVariables() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) HarnessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"harnessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) HarnessRef() *interfacesawsbedrockagentcore.HarnessReference {
	var returns *interfacesawsbedrockagentcore.HarnessReference
	_jsii_.Get(
		j,
		"harnessRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) MaxIterations() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIterations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) MaxTokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Memory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Model() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Skills() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skills",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) SystemPrompt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Tags() *[]*awscdk.CfnTag {
	var returns *[]*awscdk.CfnTag
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Tools() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) Truncation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"truncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnHarness) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::BedrockAgentCore::Harness`.
func NewCfnHarness(scope constructs.Construct, id *string, props *CfnHarnessProps) CfnHarness {
	_init_.Initialize()

	if err := validateNewCfnHarnessParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnHarness{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.CfnHarness",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::BedrockAgentCore::Harness`.
func NewCfnHarness_Override(c CfnHarness, scope constructs.Construct, id *string, props *CfnHarnessProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.CfnHarness",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnHarness)SetAllowedTools(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedTools",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetAuthorizerConfiguration(val interface{}) {
	if err := j.validateSetAuthorizerConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizerConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetEnvironment(val interface{}) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetEnvironmentArtifact(val interface{}) {
	if err := j.validateSetEnvironmentArtifactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentArtifact",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetEnvironmentVariables(val interface{}) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetHarnessName(val *string) {
	if err := j.validateSetHarnessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"harnessName",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetMaxIterations(val *float64) {
	_jsii_.Set(
		j,
		"maxIterations",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetMaxTokens(val *float64) {
	_jsii_.Set(
		j,
		"maxTokens",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetMemory(val interface{}) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetModel(val interface{}) {
	if err := j.validateSetModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"model",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetSkills(val interface{}) {
	if err := j.validateSetSkillsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skills",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetSystemPrompt(val interface{}) {
	if err := j.validateSetSystemPromptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemPrompt",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetTags(val *[]*awscdk.CfnTag) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetTimeoutSeconds(val *float64) {
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetTools(val interface{}) {
	if err := j.validateSetToolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tools",
		val,
	)
}

func (j *jsiiProxy_CfnHarness)SetTruncation(val interface{}) {
	if err := j.validateSetTruncationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"truncation",
		val,
	)
}

func CfnHarness_ArnForHarness(resource interfacesawsbedrockagentcore.IHarnessRef) *string {
	_init_.Initialize()

	if err := validateCfnHarness_ArnForHarnessParameters(resource); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.CfnHarness",
		"arnForHarness",
		[]interface{}{resource},
		&returns,
	)

	return returns
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnHarness_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHarness_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.CfnHarness",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks whether the given object is a CfnHarness.
func CfnHarness_IsCfnHarness(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHarness_IsCfnHarnessParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.CfnHarness",
		"isCfnHarness",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given object is a CfnResource.
func CfnHarness_IsCfnResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHarness_IsCfnResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.CfnHarness",
		"isCfnResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnHarness_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnHarness_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.CfnHarness",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnHarness_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_bedrockagentcore.CfnHarness",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnHarness) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnHarness) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnHarness) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnHarness) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnHarness) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnHarness) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnHarness) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnHarness) AddResourceDependency(target awscdk.CfnResource, reason *string) {
	if err := c.validateAddResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addResourceDependency",
		[]interface{}{target, reason},
	)
}

func (c *jsiiProxy_CfnHarness) ApplyCrossStackReferenceStrength(strength awscdk.ReferenceStrength) {
	if err := c.validateApplyCrossStackReferenceStrengthParameters(strength); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyCrossStackReferenceStrength",
		[]interface{}{strength},
	)
}

func (c *jsiiProxy_CfnHarness) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnHarness) CfnPropertyName(cdkPropertyName *string) *string {
	if err := c.validateCfnPropertyNameParameters(cdkPropertyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"cfnPropertyName",
		[]interface{}{cdkPropertyName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHarness) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHarness) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHarness) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnHarness) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHarness) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnHarness) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnHarness) RemoveResourceDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveResourceDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeResourceDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnHarness) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHarness) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnHarness) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHarness) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnHarness) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

func (c *jsiiProxy_CfnHarness) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}


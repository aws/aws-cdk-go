# Amazon Bedrock Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

| **Language**                                                                                   | **Package**                             |
| :--------------------------------------------------------------------------------------------- | --------------------------------------- |
| ![Typescript Logo](https://docs.aws.amazon.com/cdk/api/latest/img/typescript32.png) TypeScript | `@aws-cdk/aws-bedrock-alpha` |

[Amazon Bedrock](https://aws.amazon.com/bedrock/) is a fully managed service that offers a choice of high-performing foundation models (FMs) from leading AI companies and Amazon through a single API, along with a broad set of capabilities you need to build generative AI applications with security, privacy, and responsible AI.

This construct library facilitates the deployment of Bedrock Agents, enabling you to create sophisticated AI applications that can interact with your systems and data sources.

## Table of contents

* [Agents](#agents)

  * [Create an Agent](#create-an-agent)
  * [Action groups](#action-groups)
  * [Prepare the Agent](#prepare-the-agent)
  * [Prompt Override Configuration](#prompt-override-configuration)
  * [Memory Configuration](#memory-configuration)
  * [Agent Collaboration](#agent-collaboration)
  * [Custom Orchestration](#custom-orchestration)
  * [Agent Alias](#agent-alias)
* [Guardrails](#guardrails)

  * [Guardrail Properties](#guardrail-properties)
  * [Filter Types](#filter-types)

    * [Content Filters](#content-filters)
    * [Denied Topics](#denied-topics)
    * [Word Filters](#word-filters)
    * [PII Filters](#pii-filters)
    * [Regex Filters](#regex-filters)
    * [Contextual Grounding Filters](#contextual-grounding-filters)
  * [Guardrail Methods](#guardrail-methods)
  * [Guardrail Permissions](#guardrail-permissions)
  * [Guardrail Metrics](#guardrail-metrics)
  * [Importing Guardrails](#importing-guardrails)
  * [Guardrail Versioning](#guardrail-versioning)
* [Prompts](#prompts)

  * [Prompt Variants](#prompt-variants)
  * [Basic Text Prompt](#basic-text-prompt)
  * [Chat Prompt](#chat-prompt)
  * [Agent Prompt](#agent-prompt)
  * [Prompt Properties](#prompt-properties)
  * [Prompt Version](#prompt-version)
  * [Import Methods](#import-methods)
* [Inference Profiles](#inference-profiles)

  * [Using Inference Profiles](#using-inference-profiles)
  * [Types of Inference Profiles](#types-of-inference-profiles)
  * [Prompt Routers](#prompt-routers)
  * [Inference Profile Permissions](#inference-profile-permissions)
  * [Inference Profiles Import Methods](#inference-profiles-import-methods)

## Agents

Amazon Bedrock Agents allow generative AI applications to automate complex, multistep tasks by seamlessly integrating with your company's systems, APIs, and data sources. It uses the reasoning of foundation models (FMs), APIs, and data to break down user requests, gather relevant information, and efficiently complete tasks.

### Create an Agent

Building an agent is straightforward and fast.
The following example creates an Agent with a simple instruction and default prompts:

```go
agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
})
```

You can also create an agent with a guardrail:

```go
// Create a guardrail to filter inappropriate content
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
	Description: jsii.String("Legal ethical guardrails."),
})

guardrail.AddContentFilter(&ContentFilter{
	Type: bedrock.ContentFilterType_SEXUAL,
	InputStrength: bedrock.ContentFilterStrength_HIGH,
	OutputStrength: bedrock.ContentFilterStrength_MEDIUM,
})

// Create an agent with the guardrail
agentWithGuardrail := bedrock.NewAgent(this, jsii.String("AgentWithGuardrail"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
	Guardrail: guardrail,
})
```

### Agent Properties

The Bedrock Agent class supports the following properties.

| Name | Type | Required | Description |
|---|---|---|---|
| name | string | No | The name of the agent. Defaults to a name generated by CDK |
| instruction | string | Yes | The instruction used by the agent that determines how it will perform its task. Must have a minimum of 40 characters |
| foundationModel | IBedrockInvokable | Yes | The foundation model used for orchestration by the agent |
| existingRole | iam.IRole | No | The existing IAM Role for the agent to use. Must have a trust policy allowing Bedrock service to assume the role. Defaults to a new created role |
| shouldPrepareAgent | boolean | No | Specifies whether to automatically update the `DRAFT` version of the agent after making changes. Defaults to false |
| idleSessionTTL | Duration | No | How long sessions should be kept open for the agent. Session expires if no conversation occurs during this time. Defaults to 1 hour |
| kmsKey | kms.IKey | No | The KMS key of the agent if custom encryption is configured. Defaults to AWS managed key |
| description | string | No | A description of the agent. Defaults to no description |
| actionGroups | AgentActionGroup[] | No | The Action Groups associated with the agent |
| guardrail | IGuardrail | No | The guardrail that will be associated with the agent. Defaults to no guardrail |
| memory | Memory | No | The type and configuration of the memory to maintain context across multiple sessions and recall past interactions. Defaults to no memory |
| promptOverrideConfiguration | PromptOverrideConfiguration | No | Overrides some prompt templates in different parts of an agent sequence configuration |
| userInputEnabled | boolean | No | Select whether the agent can prompt additional information from the user when it lacks enough information. Defaults to false |
| codeInterpreterEnabled | boolean | No | Select whether the agent can generate, run, and troubleshoot code when trying to complete a task. Defaults to false |
| forceDelete | boolean | No | Whether to delete the resource even if it's in use. Defaults to true |
| agentCollaboration | AgentCollaboration | No | Configuration for agent collaboration settings, including type and collaborators. This property allows you to define how the agent collaborates with other agents and what collaborators it can work with. Defaults to no agent collaboration configuration |
| customOrchestrationExecutor | CustomOrchestrationExecutor | No | The Lambda function to use for custom orchestration. If provided, orchestrationType is set to CUSTOM_ORCHESTRATION. If not provided, orchestrationType defaults to DEFAULT. Defaults to default orchestration |

### Action Groups

An action group defines functions your agent can call. The functions are Lambda functions. The action group uses an OpenAPI schema to tell the agent what your functions do and how to call them.

#### Action Group Properties

The AgentActionGroup class supports the following properties.

| Name | Type | Required | Description |
|---|---|---|---|
| name | string | No | The name of the action group. Defaults to a name generated in the format 'action_group_quick_start_UUID' |
| description | string | No | A description of the action group |
| apiSchema | ApiSchema | No | The OpenAPI schema that defines the functions in the action group |
| executor | ActionGroupExecutor | No | The Lambda function that executes the actions in the group |
| enabled | boolean | No | Whether the action group is enabled. Defaults to true |
| forceDelete | boolean | No | Whether to delete the resource even if it's in use. Defaults to false |
| functionSchema | FunctionSchema | No | Defines functions that each define parameters that the agent needs to invoke from the user |
| parentActionGroupSignature | ParentActionGroupSignature | No | The AWS Defined signature for enabling certain capabilities in your agent |

There are three ways to provide an API schema for your action group:

From a local asset file (requires binding to scope):

```go
actionGroupFunction := lambda.NewFunction(this, jsii.String("ActionGroupFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_12(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("../lambda/action-group"))),
})

// When using ApiSchema.fromLocalAsset, you must bind the schema to a scope
schema := bedrock.ApiSchema_FromLocalAsset(path.join(__dirname, jsii.String("action-group.yaml")))
schema.Bind(this)

actionGroup := bedrock.NewAgentActionGroup(&AgentActionGroupProps{
	Name: jsii.String("query-library"),
	Description: jsii.String("Use these functions to get information about the books in the library."),
	Executor: bedrock.ActionGroupExecutor_FromLambda(actionGroupFunction),
	Enabled: jsii.Boolean(true),
	ApiSchema: schema,
})

agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
})

agent.AddActionGroup(actionGroup)
```

From an inline OpenAPI schema:

```go
inlineSchema := bedrock.ApiSchema_FromInline(jsii.String(`
openapi: 3.0.3
info:
  title: Library API
  version: 1.0.0
paths:
  /search:
    get:
      summary: Search for books
      operationId: searchBooks
      parameters:
        - name: query
          in: query
          required: true
          schema:
            type: string
`))

actionGroupFunction := lambda.NewFunction(this, jsii.String("ActionGroupFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_12(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("../lambda/action-group"))),
})

actionGroup := bedrock.NewAgentActionGroup(&AgentActionGroupProps{
	Name: jsii.String("query-library"),
	Description: jsii.String("Use these functions to get information about the books in the library."),
	Executor: bedrock.ActionGroupExecutor_FromLambda(actionGroupFunction),
	Enabled: jsii.Boolean(true),
	ApiSchema: inlineSchema,
})

agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
})

agent.AddActionGroup(actionGroup)
```

From an existing S3 file:

```go
bucket := s3.Bucket_FromBucketName(this, jsii.String("ExistingBucket"), jsii.String("my-schema-bucket"))
s3Schema := bedrock.ApiSchema_FromS3File(bucket, jsii.String("schemas/action-group.yaml"))

actionGroupFunction := lambda.NewFunction(this, jsii.String("ActionGroupFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_12(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("../lambda/action-group"))),
})

actionGroup := bedrock.NewAgentActionGroup(&AgentActionGroupProps{
	Name: jsii.String("query-library"),
	Description: jsii.String("Use these functions to get information about the books in the library."),
	Executor: bedrock.ActionGroupExecutor_FromLambda(actionGroupFunction),
	Enabled: jsii.Boolean(true),
	ApiSchema: s3Schema,
})

agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
})

agent.AddActionGroup(actionGroup)
```

### Using FunctionSchema with Action Groups

As an alternative to using OpenAPI schemas, you can define functions directly using the `FunctionSchema` class. This approach provides a more structured way to define the functions that your agent can call.

```go
actionGroupFunction := lambda.NewFunction(this, jsii.String("ActionGroupFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_12(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("../lambda/action-group"))),
})

// Define a function schema with parameters
functionSchema := bedrock.NewFunctionSchema(&FunctionSchemaProps{
	Functions: []functionProps{
		&functionProps{
			Name: jsii.String("searchBooks"),
			Description: jsii.String("Search for books in the library catalog"),
			Parameters: map[string]functionParameterProps{
				"query": &functionParameterProps{
					"type": bedrock.ParameterType_STRING,
					"required": jsii.Boolean(true),
					"description": jsii.String("The search query string"),
				},
				"maxResults": &functionParameterProps{
					"type": bedrock.ParameterType_INTEGER,
					"required": jsii.Boolean(false),
					"description": jsii.String("Maximum number of results to return"),
				},
				"includeOutOfPrint": &functionParameterProps{
					"type": bedrock.ParameterType_BOOLEAN,
					"required": jsii.Boolean(false),
					"description": jsii.String("Whether to include out-of-print books"),
				},
			},
			RequireConfirmation: bedrock.RequireConfirmation_DISABLED,
		},
		&functionProps{
			Name: jsii.String("getBookDetails"),
			Description: jsii.String("Get detailed information about a specific book"),
			Parameters: map[string]*functionParameterProps{
				"bookId": &functionParameterProps{
					"type": bedrock.ParameterType_STRING,
					"required": jsii.Boolean(true),
					"description": jsii.String("The unique identifier of the book"),
				},
			},
			RequireConfirmation: bedrock.RequireConfirmation_ENABLED,
		},
	},
})

// Create an action group using the function schema
actionGroup := bedrock.NewAgentActionGroup(&AgentActionGroupProps{
	Name: jsii.String("library-functions"),
	Description: jsii.String("Functions for interacting with the library catalog"),
	Executor: bedrock.ActionGroupExecutor_FromLambda(actionGroupFunction),
	FunctionSchema: functionSchema,
	Enabled: jsii.Boolean(true),
})

agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
	ActionGroups: []agentActionGroup{
		actionGroup,
	},
})
```

The `FunctionSchema` approach offers several advantages:

* Type-safe definition of functions and parameters
* Built-in validation of parameter names, descriptions, and other properties
* Clear structure that maps directly to the AWS Bedrock API
* Support for parameter types including string, number, integer, boolean, array, and object
* Option to require user confirmation before executing specific functions

If you chose to load your schema file from S3, the construct will provide the necessary permissions to your agent's execution role to access the schema file from the specific bucket. Similar to performing the operation through the console, the agent execution role will get a permission like:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "AmazonBedrockAgentS3PolicyProd",
            "Effect": "Allow",
            "Action": [
                "s3:GetObject"
            ],
            "Resource": [
                "arn:aws:s3:::<BUCKET_NAME>/<OBJECT_KEY>"
            ],
            "Condition": {
                "StringEquals": {
                    "aws:ResourceAccount": "ACCOUNT_NUMBER"
                }
            }
        }
    ]
}
```

```go
// create a bucket containing the input schema
schemaBucket := s3.NewBucket(this, jsii.String("SchemaBucket"), &BucketProps{
	EnforceSSL: jsii.Boolean(true),
	Versioned: jsii.Boolean(true),
	PublicReadAccess: jsii.Boolean(false),
	BlockPublicAccess: s3.BlockPublicAccess_BLOCK_ALL(),
	Encryption: s3.BucketEncryption_S3_MANAGED,
	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	AutoDeleteObjects: jsii.Boolean(true),
})

// deploy the local schema file to S3
deployement := awscdk.Aws_s3_deployment.NewBucketDeployment(this, jsii.String("DeployWebsite"), &BucketDeploymentProps{
	Sources: []iSource{
		awscdk.*Aws_s3_deployment.Source_Asset(path.join(__dirname, jsii.String("../inputschema"))),
	},
	DestinationBucket: schemaBucket,
	DestinationKeyPrefix: jsii.String("inputschema"),
})

// create the agent
agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
	UserInputEnabled: jsii.Boolean(true),
	ShouldPrepareAgent: jsii.Boolean(true),
})

// create a lambda function
actionGroupFunction := lambda.NewFunction(this, jsii.String("ActionGroupFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_12(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("../lambda/action-group"))),
})

// create an action group and read the schema file from S3
actionGroup := bedrock.NewAgentActionGroup(&AgentActionGroupProps{
	Name: jsii.String("query-library"),
	Description: jsii.String("Use these functions to get information about the books in the library."),
	Executor: bedrock.ActionGroupExecutor_FromLambda(actionGroupFunction),
	Enabled: jsii.Boolean(true),
	ApiSchema: bedrock.ApiSchema_FromS3File(schemaBucket, jsii.String("inputschema/action-group.yaml")),
})

// add the action group to the agent
agent.AddActionGroup(actionGroup)

// add dependency for the agent on the s3 deployment
agent.Node.AddDependency(deployement)
```

### Prepare the Agent

The `Agent` constructs take an optional parameter `shouldPrepareAgent` to indicate that the Agent should be prepared after any updates to an agent or action group. This may increase the time to create and update those resources. By default, this value is false.

#### Prepare Agent Properties

| Name | Type | Required | Description |
|---|---|---|---|
| shouldPrepareAgent | boolean | No | Whether to automatically update the DRAFT version of the agent after making changes. Defaults to false |

Creating an agent alias will not prepare the agent, so if you create an alias using the `AgentAlias` resource then you should set `shouldPrepareAgent` to ***true***.

### Prompt Override Configuration

Bedrock Agents allows you to customize the prompts and LLM configuration for different steps in the agent sequence. The implementation provides type-safe configurations for each step type, ensuring correct usage at compile time.

#### Prompt Override Configuration Properties

| Name | Type | Required | Description |
|---|---|---|---|
| steps | PromptStepConfiguration[] | Yes | Array of step configurations for different parts of the agent sequence |
| parser | lambda.IFunction | No | Lambda function for custom parsing of agent responses |

#### Prompt Step Configuration Properties

Each step in the `steps` array supports the following properties:

| Name | Type | Required | Description |
|---|---|---|---|
| stepType | AgentStepType | Yes | The type of step being configured (PRE_PROCESSING, ORCHESTRATION, POST_PROCESSING, ROUTING_CLASSIFIER, MEMORY_SUMMARIZATION, KNOWLEDGE_BASE_RESPONSE_GENERATION) |
| stepEnabled | boolean | No | Whether this step is enabled. Defaults to true |
| customPromptTemplate | string | No | Custom prompt template to use for this step |
| inferenceConfig | InferenceConfiguration | No | Configuration for model inference parameters |
| foundationModel | BedrockFoundationModel | No | Alternative foundation model to use for this step (only valid for ROUTING_CLASSIFIER step) |
| useCustomParser | boolean | No | Whether to use a custom parser for this step. Requires parser to be provided in PromptOverrideConfiguration |

#### Inference Configuration Properties

When providing `inferenceConfig`, the following properties are supported:

| Name | Type | Required | Description |
|---|---|---|---|
| temperature | number | No | Controls randomness in the model's output (0.0-1.0) |
| topP | number | No | Controls diversity via nucleus sampling (0.0-1.0) |
| topK | number | No | Controls diversity by limiting the cumulative probability |
| maximumLength | number | No | Maximum length of generated text |
| stopSequences | string[] | No | Sequences where the model should stop generating |

The following steps can be configured:

* PRE_PROCESSING: Prepares the user input for orchestration
* ORCHESTRATION: Main step that determines the agent's actions
* POST_PROCESSING: Refines the agent's response
* ROUTING_CLASSIFIER: Classifies and routes requests to appropriate collaborators
* MEMORY_SUMMARIZATION: Summarizes conversation history for memory retention
* KNOWLEDGE_BASE_RESPONSE_GENERATION: Generates responses using knowledge base content

Example with pre-processing configuration:

```go
agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
	Instruction: jsii.String("You are a helpful assistant."),
	PromptOverrideConfiguration: bedrock.PromptOverrideConfiguration_FromSteps([]promptStepConfigBase{
		&promptStepConfigBase{
			StepType: bedrock.AgentStepType_PRE_PROCESSING,
			StepEnabled: jsii.Boolean(true),
			CustomPromptTemplate: jsii.String("Your custom prompt template here"),
			InferenceConfig: &InferenceConfiguration{
				Temperature: jsii.Number(0),
				TopP: jsii.Number(1),
				TopK: jsii.Number(250),
				MaximumLength: jsii.Number(1),
				StopSequences: []*string{
					jsii.String("\n\nHuman:"),
				},
			},
		},
	}),
})
```

Example with routing classifier and foundation model:

```go
agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
	Instruction: jsii.String("You are a helpful assistant."),
	PromptOverrideConfiguration: bedrock.PromptOverrideConfiguration_FromSteps([]promptStepConfigBase{
		&PromptRoutingClassifierConfigCustomParser{
			StepType: bedrock.AgentStepType_ROUTING_CLASSIFIER,
			StepEnabled: jsii.Boolean(true),
			CustomPromptTemplate: jsii.String("Your routing template here"),
			FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_V2(),
		}.(promptRoutingClassifierConfigCustomParser),
	}),
})
```

Using a custom Lambda parser:

```go
parserFunction := lambda.NewFunction(this, jsii.String("ParserFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_10(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(jsii.String("lambda")),
})

agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
	Instruction: jsii.String("You are a helpful assistant."),
	PromptOverrideConfiguration: bedrock.PromptOverrideConfiguration_WithCustomParser(&CustomParserProps{
		Parser: parserFunction,
		PreProcessingStep: &PromptPreProcessingConfigCustomParser{
			StepType: bedrock.AgentStepType_PRE_PROCESSING,
			UseCustomParser: jsii.Boolean(true),
		},
	}),
})
```

Foundation models can only be specified for the ROUTING_CLASSIFIER step.

### Memory Configuration

Agents can maintain context across multiple sessions and recall past interactions using memory. This feature is useful for creating a more coherent conversational experience.

#### Memory Configuration Properties

| Name | Type | Required | Description |
|---|---|---|---|
| maxRecentSessions | number | No | Maximum number of recent session summaries to retain |
| memoryDuration | Duration | No | How long to retain session summaries |

Example:

```go
agent := bedrock.NewAgent(this, jsii.String("MyAgent"), &AgentProps{
	AgentName: jsii.String("MyAgent"),
	Instruction: jsii.String("Your instruction here"),
	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
	Memory: awsbedrockalpha.Memory_SessionSummary(&SessionSummaryMemoryProps{
		MaxRecentSessions: jsii.Number(10),
		 // Keep the last 10 session summaries
		MemoryDuration: awscdk.Duration_Days(jsii.Number(20)),
	}),
})
```

### Agent Collaboration

Agent Collaboration enables multiple Bedrock Agents to work together on complex tasks. This feature allows agents to specialize in different areas and collaborate to provide more comprehensive responses to user queries.

#### Agent Collaboration Properties

| Name | Type | Required | Description |
|---|---|---|---|
| type | AgentCollaboratorType | Yes | Type of collaboration (SUPERVISOR or PEER) |
| collaborators | AgentCollaborator[] | Yes | List of agent collaborators |

#### Agent Collaborator Properties

| Name | Type | Required | Description |
|---|---|---|---|
| agentAlias | AgentAlias | Yes | The agent alias to collaborate with |
| collaborationInstruction | string | Yes | Instructions for how to collaborate with this agent |
| collaboratorName | string | Yes | Name of the collaborator |
| relayConversationHistory | boolean | No | Whether to relay conversation history to the collaborator. Defaults to false |

Example:

```go
// Create a specialized agent
customerSupportAgent := bedrock.NewAgent(this, jsii.String("CustomerSupportAgent"), &AgentProps{
	Instruction: jsii.String("You specialize in answering customer support questions."),
	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
})

// Create an agent alias
customerSupportAlias := bedrock.NewAgentAlias(this, jsii.String("CustomerSupportAlias"), &AgentAliasProps{
	Agent: customerSupportAgent,
	AgentAliasName: jsii.String("production"),
})

// Create a main agent that collaborates with the specialized agent
mainAgent := bedrock.NewAgent(this, jsii.String("MainAgent"), &AgentProps{
	Instruction: jsii.String("You route specialized questions to other agents."),
	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
	AgentCollaboration: map[string]interface{}{
		"type": bedrock.AgentCollaboratorType_SUPERVISOR,
		"collaborators": []AgentCollaborator{
			bedrock.NewAgentCollaborator(&AgentCollaboratorProps{
				"agentAlias": customerSupportAlias,
				"collaborationInstruction": jsii.String("Route customer support questions to this agent."),
				"collaboratorName": jsii.String("CustomerSupport"),
				"relayConversationHistory": jsii.Boolean(true),
			}),
		},
	},
})
```

### Custom Orchestration

Custom Orchestration allows you to override the default agent orchestration flow with your own Lambda function. This enables more control over how the agent processes user inputs and invokes action groups.

When you provide a customOrchestrationExecutor, the agent's orchestrationType is automatically set to CUSTOM_ORCHESTRATION. If no customOrchestrationExecutor is provided, the orchestrationType defaults to DEFAULT, using Amazon Bedrock's built-in orchestration.

#### Custom Orchestration Properties

| Name | Type | Required | Description |
|---|---|---|---|
| function | lambda.IFunction | Yes | The Lambda function that implements the custom orchestration logic |

Example:

```go
orchestrationFunction := lambda.NewFunction(this, jsii.String("OrchestrationFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_10(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(jsii.String("lambda/orchestration")),
})

agent := bedrock.NewAgent(this, jsii.String("CustomOrchestrationAgent"), &AgentProps{
	Instruction: jsii.String("You are a helpful assistant with custom orchestration logic."),
	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
	CustomOrchestrationExecutor: bedrock.CustomOrchestrationExecutor_FromLambda(orchestrationFunction),
})
```

### Agent Alias

After you have sufficiently iterated on your working draft and are satisfied with the behavior of your agent, you can set it up for deployment and integration into your application by creating aliases.

To deploy your agent, you need to create an alias. During alias creation, Amazon Bedrock automatically creates a version of your agent. The alias points to this newly created version. You can point the alias to a previously created version if necessary. You then configure your application to make API calls to that alias.

By default, the Agent resource creates a test alias named 'AgentTestAlias' that points to the 'DRAFT' version. This test alias is accessible via the `testAlias` property of the agent. You can also create additional aliases for different environments using the AgentAlias construct.

#### Agent Alias Properties

| Name | Type | Required | Description |
|---|---|---|---|
| agent | Agent | Yes | The agent to create an alias for |
| agentAliasName | string | No | The name of the agent alias. Defaults to a name generated by CDK |
| description | string | No | A description of the agent alias. Defaults to no description |
| routingConfiguration | AgentAliasRoutingConfiguration | No | Configuration for routing traffic between agent versions |
| agentVersion | string | No | The version of the agent to use. If not specified, a new version is created |

When redeploying an agent with changes, you must ensure the agent version is updated to avoid deployment failures with "agent already exists" errors. The recommended way to handle this is to include the `lastUpdated` property in the agent's description, which automatically updates whenever the agent is modified. This ensures a new version is created on each deployment.

Example:

```go
agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
})

agentAlias := bedrock.NewAgentAlias(this, jsii.String("myAlias"), &AgentAliasProps{
	AgentAliasName: jsii.String("production"),
	Agent: agent,
	Description: fmt.Sprintf("Production version of my agent. Created at %v", agent.LastUpdated),
})
```

## Guardrails

Amazon Bedrock's Guardrails feature enables you to implement robust governance and control mechanisms for your generative AI applications, ensuring alignment with your specific use cases and responsible AI policies. Guardrails empowers you to create multiple tailored policy configurations, each designed to address the unique requirements and constraints of different use cases. These policy configurations can then be seamlessly applied across multiple foundation models (FMs) and Agents, ensuring a consistent user experience and standardizing safety, security, and privacy controls throughout your generative AI ecosystem.

With Guardrails, you can define and enforce granular, customizable policies to precisely govern the behavior of your generative AI applications. You can configure the following policies in a guardrail to avoid undesirable and harmful content and remove sensitive information for privacy protection.

Content filters – Adjust filter strengths to block input prompts or model responses containing harmful content.
Denied topics – Define a set of topics that are undesirable in the context of your application. These topics will be blocked if detected in user queries or model responses.
Word filters – Configure filters to block undesirable words, phrases, and profanity. Such words can include offensive terms, competitor names etc.
Sensitive information filters – Block or mask sensitive information such as personally identifiable information (PII) or custom regex in user inputs and model responses.
You can create a Guardrail with a minimum blockedInputMessaging, blockedOutputsMessaging and default content filter policy.

### Basic Guardrail Creation

#### TypeScript

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
	Description: jsii.String("Legal ethical guardrails."),
})
```

### Guardrail Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| guardrailName | string | Yes | The name of the guardrail |
| description | string | No | The description of the guardrail |
| blockedInputMessaging | string | No | The message to return when the guardrail blocks a prompt. Default: "Sorry, your query violates our usage policy." |
| blockedOutputsMessaging | string | No | The message to return when the guardrail blocks a model response. Default: "Sorry, I am unable to answer your question because of our usage policy." |
| kmsKey | IKey | No | A custom KMS key to use for encrypting data. Default: Your data is encrypted by default with a key that AWS owns and manages for you. |
| crossRegionConfig | GuardrailCrossRegionConfigProperty | No | The cross-region configuration for the guardrail. This enables cross-region inference for enhanced language support and filtering capabilities. Default: No cross-region configuration |
| contentFilters | ContentFilter[] | No | The content filters to apply to the guardrail |
| contentFiltersTierConfig | TierConfig | No | The tier configuration to apply to content filters. Default: TierConfig.CLASSIC |
| deniedTopics | Topic[] | No | Up to 30 denied topics to block user inputs or model responses associated with the topic |
| topicsTierConfig | TierConfig | No | The tier configuration to apply to topic filters. Default: TierConfig.CLASSIC |
| wordFilters | string[] | No | The word filters to apply to the guardrail |
| managedWordListFilters | ManagedWordFilterType[] | No | The managed word filters to apply to the guardrail |
| piiFilters | PIIFilter[] | No | The PII filters to apply to the guardrail |
| regexFilters | RegexFilter[] | No | The regular expression (regex) filters to apply to the guardrail |
| contextualGroundingFilters | ContextualGroundingFilter[] | No | The contextual grounding filters to apply to the guardrail |

### Filter Types

#### Content Filters

Content filters allow you to block input prompts or model responses containing harmful content. You can adjust the filter strength and configure separate actions for input and output.

##### Content Filter Configuration

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
	// Configure tier for content filters (optional)
	ContentFiltersTierConfig: bedrock.TierConfig_STANDARD,
})

guardrail.AddContentFilter(&ContentFilter{
	Type: bedrock.ContentFilterType_SEXUAL,
	InputStrength: bedrock.ContentFilterStrength_HIGH,
	OutputStrength: bedrock.ContentFilterStrength_MEDIUM,
	// props below are optional
	InputAction: bedrock.GuardrailAction_BLOCK,
	InputEnabled: jsii.Boolean(true),
	OutputAction: bedrock.GuardrailAction_NONE,
	OutputEnabled: jsii.Boolean(true),
	InputModalities: []modalityType{
		bedrock.*modalityType_TEXT,
		bedrock.*modalityType_IMAGE,
	},
	OutputModalities: []*modalityType{
		bedrock.*modalityType_TEXT,
	},
})
```

Available content filter types:

* `SEXUAL`: Describes input prompts and model responses that indicates sexual interest, activity, or arousal
* `VIOLENCE`: Describes input prompts and model responses that includes glorification of or threats to inflict physical pain
* `HATE`: Describes input prompts and model responses that discriminate, criticize, insult, denounce, or dehumanize a person or group
* `INSULTS`: Describes input prompts and model responses that includes demeaning, humiliating, mocking, insulting, or belittling language
* `MISCONDUCT`: Describes input prompts and model responses that seeks or provides information about engaging in misconduct activity
* `PROMPT_ATTACK`: Enable to detect and block user inputs attempting to override system instructions

Available content filter strengths:

* `NONE`: No filtering
* `LOW`: Light filtering
* `MEDIUM`: Moderate filtering
* `HIGH`: Strict filtering

Available guardrail actions:

* `BLOCK`: Blocks the content from being processed
* `ANONYMIZE`: Masks the content with an identifier tag
* `NONE`: Takes no action

> Warning: the ANONYMIZE action is not available in all configurations. Please refer to the documentation of each filter to see which ones
> support

Available modality types:

* `TEXT`: Text modality for content filters
* `IMAGE`: Image modality for content filters

#### Tier Configuration

Guardrails support tier configurations that determine the level of language support and robustness for content and topic filters. You can configure separate tier settings for content filters and topic filters.

##### Tier Configuration Options

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
	// Configure tier for content filters
	ContentFiltersTierConfig: bedrock.TierConfig_STANDARD,
	// Configure tier for topic filters
	TopicsTierConfig: bedrock.TierConfig_CLASSIC,
})
```

Available tier configurations:

* `CLASSIC`: Provides established guardrails functionality supporting English, French, and Spanish languages
* `STANDARD`: Provides a more robust solution than the CLASSIC tier and has more comprehensive language support. This tier requires that your guardrail use cross-Region inference

> Note: The STANDARD tier provides enhanced language support and more comprehensive filtering capabilities, but requires cross-Region inference to be enabled for your guardrail.

#### Cross-Region Configuration

You can configure a system-defined guardrail profile to use with your guardrail. Guardrail profiles define the destination AWS Regions where guardrail inference requests can be automatically routed. Using guardrail profiles helps maintain guardrail performance and reliability when demand increases.

##### Cross-Region Configuration Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| guardrailProfileArn | string | Yes | The ARN of the system-defined guardrail profile that defines the destination AWS Regions where guardrail inference requests can be automatically routed |

##### Cross-Region Configuration Example

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
	Description: jsii.String("Guardrail with cross-region configuration for enhanced language support"),
	CrossRegionConfig: &GuardrailCrossRegionConfigProperty{
		GuardrailProfileArn: jsii.String("arn:aws:bedrock:us-east-1:123456789012:guardrail-profile/my-profile"),
	},
	// Use STANDARD tier for enhanced capabilities
	ContentFiltersTierConfig: bedrock.TierConfig_STANDARD,
	TopicsTierConfig: bedrock.TierConfig_STANDARD,
})
```

> Note: Cross-region configuration is required when using the STANDARD tier for content and topic filters. It helps maintain guardrail performance and reliability when demand increases by automatically routing inference requests to appropriate regions.

You will need to provide the necessary permissions for cross region: https://docs.aws.amazon.com/bedrock/latest/userguide/guardrail-profiles-permissions.html .

#### Denied Topics

Denied topics allow you to define a set of topics that are undesirable in the context of your application. These topics will be blocked if detected in user queries or model responses. You can configure separate actions for input and output.

##### Denied Topic Configuration

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
	// Configure tier for topic filters (optional)
	TopicsTierConfig: bedrock.TierConfig_STANDARD,
})

// Use a predefined topic
guardrail.AddDeniedTopicFilter(bedrock.Topic_FINANCIAL_ADVICE())

// Create a custom topic with input/output actions
guardrail.AddDeniedTopicFilter(bedrock.Topic_Custom(&CustomTopicProps{
	Name: jsii.String("Legal_Advice"),
	Definition: jsii.String("Offering guidance or suggestions on legal matters, legal actions, interpretation of laws, or legal rights and responsibilities."),
	Examples: []*string{
		jsii.String("Can I sue someone for this?"),
		jsii.String("What are my legal rights in this situation?"),
		jsii.String("Is this action against the law?"),
		jsii.String("What should I do to file a legal complaint?"),
		jsii.String("Can you explain this law to me?"),
	},
	// props below are optional
	InputAction: bedrock.GuardrailAction_BLOCK,
	InputEnabled: jsii.Boolean(true),
	OutputAction: bedrock.GuardrailAction_NONE,
	OutputEnabled: jsii.Boolean(true),
}))
```

#### Word Filters

Word filters allow you to block specific words, phrases, or profanity in user inputs and model responses. You can configure separate actions for input and output.

##### Word Filter Configuration

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
})

// Add managed word list with input/output actions
guardrail.AddManagedWordListFilter(&ManagedWordFilter{
	Type: bedrock.ManagedWordFilterType_PROFANITY,
	InputAction: bedrock.GuardrailAction_BLOCK,
	InputEnabled: jsii.Boolean(true),
	OutputAction: bedrock.GuardrailAction_NONE,
	OutputEnabled: jsii.Boolean(true),
})

// Add individual words
guardrail.AddWordFilter(&WordFilter{
	Text: jsii.String("drugs"),
})
guardrail.AddWordFilter(&WordFilter{
	Text: jsii.String("competitor"),
})

// Add words from a file
guardrail.AddWordFilterFromFile(jsii.String("./scripts/wordsPolicy.csv"))
```

#### PII Filters

PII filters allow you to detect and handle personally identifiable information in user inputs and model responses. You can configure separate actions for input and output.

The PII types are organized into enum-like classes for better type safety and transpilation compatibility:

* **GeneralPIIType**: General PII types like addresses, emails, names, phone numbers
* **FinancePIIType**: Financial information like credit card numbers, PINs, SWIFT codes
* **InformationTechnologyPIIType**: IT-related data like URLs, IP addresses, AWS keys
* **USASpecificPIIType**: US-specific identifiers like SSNs, passport numbers
* **CanadaSpecificPIIType**: Canada-specific identifiers like health numbers, SINs
* **UKSpecificPIIType**: UK-specific identifiers like NHS numbers, NI numbers

##### PII Filter Configuration

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
})

// Add PII filter for addresses with input/output actions
guardrail.AddPIIFilter(&PIIFilter{
	Type: bedrock.GeneralPIIType_ADDRESS(),
	Action: bedrock.GuardrailAction_BLOCK,
	// below props are optional
	InputAction: bedrock.GuardrailAction_BLOCK,
	InputEnabled: jsii.Boolean(true),
	OutputAction: bedrock.GuardrailAction_ANONYMIZE,
	OutputEnabled: jsii.Boolean(true),
})

// Add PII filter for credit card numbers with input/output actions
guardrail.AddPIIFilter(&PIIFilter{
	Type: bedrock.FinancePIIType_CREDIT_DEBIT_CARD_NUMBER(),
	Action: bedrock.GuardrailAction_BLOCK,
	// below props are optional
	InputAction: bedrock.GuardrailAction_BLOCK,
	InputEnabled: jsii.Boolean(true),
	OutputAction: bedrock.GuardrailAction_ANONYMIZE,
	OutputEnabled: jsii.Boolean(true),
})

// Add PII filter for email addresses
guardrail.AddPIIFilter(&PIIFilter{
	Type: bedrock.GeneralPIIType_EMAIL(),
	Action: bedrock.GuardrailAction_ANONYMIZE,
})

// Add PII filter for US Social Security Numbers
guardrail.AddPIIFilter(&PIIFilter{
	Type: bedrock.USASpecificPIIType_US_SOCIAL_SECURITY_NUMBER(),
	Action: bedrock.GuardrailAction_BLOCK,
})

// Add PII filter for IP addresses
guardrail.AddPIIFilter(&PIIFilter{
	Type: bedrock.InformationTechnologyPIIType_IP_ADDRESS(),
	Action: bedrock.GuardrailAction_ANONYMIZE,
})
```

##### Available PII Types

**GeneralPIIType:**

* `ADDRESS`: Physical addresses
* `AGE`: Individual's age
* `DRIVER_ID`: Driver's license numbers
* `EMAIL`: Email addresses
* `LICENSE_PLATE`: Vehicle license plates
* `NAME`: Individual names
* `PASSWORD`: Passwords
* `PHONE`: Phone numbers
* `USERNAME`: User account names
* `VEHICLE_IDENTIFICATION_NUMBER`: Vehicle VINs

**FinancePIIType:**

* `CREDIT_DEBIT_CARD_CVV`: Card verification codes
* `CREDIT_DEBIT_CARD_EXPIRY`: Card expiration dates
* `CREDIT_DEBIT_CARD_NUMBER`: Credit/debit card numbers
* `PIN`: Personal identification numbers
* `SWIFT_CODE`: Bank SWIFT codes
* `INTERNATIONAL_BANK_ACCOUNT_NUMBER`: IBAN numbers

**InformationTechnologyPIIType:**

* `URL`: Web addresses
* `IP_ADDRESS`: IPv4 addresses
* `MAC_ADDRESS`: Network interface MAC addresses
* `AWS_ACCESS_KEY`: AWS access key IDs
* `AWS_SECRET_KEY`: AWS secret access keys

**USASpecificPIIType:**

* `US_BANK_ACCOUNT_NUMBER`: US bank account numbers
* `US_BANK_ROUTING_NUMBER`: US bank routing numbers
* `US_INDIVIDUAL_TAX_IDENTIFICATION_NUMBER`: US ITINs
* `US_PASSPORT_NUMBER`: US passport numbers
* `US_SOCIAL_SECURITY_NUMBER`: US Social Security Numbers

**CanadaSpecificPIIType:**

* `CA_HEALTH_NUMBER`: Canadian Health Service Numbers
* `CA_SOCIAL_INSURANCE_NUMBER`: Canadian Social Insurance Numbers

**UKSpecificPIIType:**

* `UK_NATIONAL_HEALTH_SERVICE_NUMBER`: UK NHS numbers
* `UK_NATIONAL_INSURANCE_NUMBER`: UK National Insurance numbers
* `UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER`: UK UTR numbers

#### Regex Filters

Regex filters allow you to detect and handle custom patterns in user inputs and model responses. You can configure separate actions for input and output.

##### Regex Filter Configuration

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
})
// Add regex filter with input/output actions
guardrail.AddRegexFilter(&RegexFilter{
	Name: jsii.String("TestRegexFilter"),
	Pattern: jsii.String("test-pattern"),
	Action: bedrock.GuardrailAction_ANONYMIZE,
	// below props are optional
	Description: jsii.String("This is a test regex filter"),
	InputAction: bedrock.GuardrailAction_BLOCK,
	InputEnabled: jsii.Boolean(true),
	OutputAction: bedrock.GuardrailAction_ANONYMIZE,
	OutputEnabled: jsii.Boolean(true),
})
```

#### Contextual Grounding Filters

Contextual grounding filters allow you to ensure that model responses are factually correct and relevant to the user's query. You can configure the action and enable/disable the filter.

##### Contextual Grounding Filter Configuration

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
})
// Add contextual grounding filter with action and enabled flag
guardrail.AddContextualGroundingFilter(&ContextualGroundingFilter{
	Type: bedrock.ContextualGroundingFilterType_GROUNDING,
	Threshold: jsii.Number(0.8),
	// the properties below are optional
	Action: bedrock.GuardrailAction_BLOCK,
	Enabled: jsii.Boolean(true),
})
```

### Guardrail Methods

| Method | Description |
|--------|-------------|
| `addContentFilter()` | Adds a content filter to the guardrail |
| `addDeniedTopicFilter()` | Adds a denied topic filter to the guardrail |
| `addWordFilter()` | Adds a word filter to the guardrail |
| `addManagedWordListFilter()` | Adds a managed word list filter to the guardrail |
| `addWordFilterFromFile()` | Adds word filters from a file to the guardrail |
| `addPIIFilter()` | Adds a PII filter to the guardrail |
| `addRegexFilter()` | Adds a regex filter to the guardrail |
| `addContextualGroundingFilter()` | Adds a contextual grounding filter to the guardrail |
| `createVersion()` | Creates a new version of the guardrail |

### Guardrail Permissions

Guardrails provide methods to grant permissions to other resources that need to interact with the guardrail.

#### Permission Methods

| Method | Description | Parameters |
|--------|-------------|------------|
| `grant(grantee, ...actions)` | Grants the given principal identity permissions to perform actions on this guardrail | `grantee`: The principal to grant permissions to<br>`actions`: The actions to grant (e.g., `bedrock:GetGuardrail`, `bedrock:ListGuardrails`) |
| `grantApply(grantee)` | Grants the given identity permissions to apply the guardrail | `grantee`: The principal to grant permissions to |

#### Permission Examples

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
})

lambdaFunction := lambda.NewFunction(this, jsii.String("testLambda"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_12(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("../lambda/my-code"))),
})

// Grant specific permissions to a Lambda function
guardrail.Grant(lambdaFunction, jsii.String("bedrock:GetGuardrail"), jsii.String("bedrock:ListGuardrails"))

// Grant permissions to apply the guardrail
guardrail.GrantApply(lambdaFunction)
```

### Guardrail Metrics

Amazon Bedrock provides metrics for your guardrails, allowing you to monitor their effectiveness and usage. These metrics are available in CloudWatch and can be used to create dashboards and alarms.

#### Metrics Examples

```go
import cloudwatch "github.com/aws/aws-cdk-go/awscdk"


guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
})
// Get a specific metric for this guardrail
invocationsMetric := guardrail.MetricInvocations(&MetricOptions{
	Statistic: jsii.String("Sum"),
	Period: awscdk.Duration_Minutes(jsii.Number(5)),
})

// Create a CloudWatch alarm for high invocation latency
// Create a CloudWatch alarm for high invocation latency
cloudwatch.NewAlarm(this, jsii.String("HighLatencyAlarm"), &AlarmProps{
	Metric: guardrail.MetricInvocationLatency(),
	Threshold: jsii.Number(1000),
	 // 1 second
	EvaluationPeriods: jsii.Number(3),
})

// Get metrics for all guardrails
allInvocationsMetric := bedrock.Guardrail_MetricAllInvocations()
```

### Importing Guardrails

You can import existing guardrails using the `fromGuardrailAttributes` or `fromCfnGuardrail` methods.

#### Import Configuration

```go
var stack stack

cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
})
// Import an existing guardrail by ARN
importedGuardrail := bedrock.Guardrail_FromGuardrailAttributes(stack, jsii.String("TestGuardrail"), &GuardrailAttributes{
	GuardrailArn: jsii.String("arn:aws:bedrock:us-east-1:123456789012:guardrail/oygh3o8g7rtl"),
	GuardrailVersion: jsii.String("1"),
	 //optional
	KmsKey: cmk,
})
```

```go
import bedrockl1 "github.com/aws/aws-cdk-go/awscdk"

// Import a guardrail created through the L1 CDK CfnGuardrail construct
l1guardrail := bedrockl1.NewCfnGuardrail(this, jsii.String("MyCfnGuardrail"), &CfnGuardrailProps{
	BlockedInputMessaging: jsii.String("blockedInputMessaging"),
	BlockedOutputsMessaging: jsii.String("blockedOutputsMessaging"),
	Name: jsii.String("namemycfnguardrails"),
	WordPolicyConfig: &WordPolicyConfigProperty{
		WordsConfig: []interface{}{
			&WordConfigProperty{
				Text: jsii.String("drugs"),
			},
		},
	},
})

importedGuardrail := bedrock.Guardrail_FromCfnGuardrail(l1guardrail)
```

### Guardrail Versioning

Guardrails support versioning, allowing you to track changes and maintain multiple versions of your guardrail configurations.

#### Version Configuration

```go
guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
	GuardrailName: jsii.String("my-BedrockGuardrails"),
})
// Create a new version of the guardrail
guardrail.CreateVersion(jsii.String("testversion"))
```

## Prompts

Amazon Bedrock provides the ability to create and save prompts using Prompt management so that you can save time by applying the same prompt to different workflows. You can include variables in the prompt so that you can adjust the prompt for different use case.

The `Prompt` resource allows you to create a new prompt.

### Prompt Variants

Prompt variants in the context of Amazon Bedrock refer to alternative configurations of a prompt, including its message or the model and inference configurations used. Prompt variants are the building blocks of prompts - you must create at least one prompt variant to create a prompt. Prompt variants allow you to create different versions of a prompt, test them, and save the variant that works best for your use case.

There are three types of prompt variants:

* **Basic Text Prompt**: Simple text-based prompts for straightforward interactions
* **Chat variant**: Conversational prompts that support system messages, user/assistant message history, and tools
* **Agent variant**: Prompts designed to work with Bedrock Agents

### Basic Text Prompt

Text prompts are the simplest form of prompts, consisting of plain text instructions with optional variables. They are ideal for straightforward tasks like summarization, content generation, or question answering where you need a direct text-based interaction with the model.

```go
cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
})
claudeModel := bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_SONNET_V1_0()

variant1 := bedrock.PromptVariant_Text(&TextPromptVariantProps{
	VariantName: jsii.String("variant1"),
	Model: claudeModel,
	PromptVariables: []*string{
		jsii.String("topic"),
	},
	PromptText: jsii.String("This is my first text prompt. Please summarize our conversation on: {{topic}}."),
	InferenceConfiguration: bedrock.PromptInferenceConfiguration_Text(&PromptInferenceConfigurationProps{
		Temperature: jsii.Number(1),
		TopP: jsii.Number(0.999),
		MaxTokens: jsii.Number(2000),
	}),
})

prompt1 := bedrock.NewPrompt(this, jsii.String("prompt1"), &PromptProps{
	PromptName: jsii.String("prompt1"),
	Description: jsii.String("my first prompt"),
	DefaultVariant: variant1,
	Variants: []iPromptVariant{
		variant1,
	},
	KmsKey: cmk,
})
```

### Chat Prompt

Use this template type when the model supports the Converse API or the Anthropic Claude Messages API. This allows you to include a System prompt and previous User messages and Assistant messages for context.

```go
cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
})

variantChat := bedrock.PromptVariant_Chat(&ChatPromptVariantProps{
	VariantName: jsii.String("variant1"),
	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
	Messages: []chatMessage{
		bedrock.*chatMessage_User(jsii.String("From now on, you speak Japanese!")),
		bedrock.*chatMessage_Assistant(jsii.String("Konnichiwa!")),
		bedrock.*chatMessage_*User(jsii.String("From now on, you speak {{language}}!")),
	},
	System: jsii.String("You are a helpful assistant that only speaks the language you`re told."),
	PromptVariables: []*string{
		jsii.String("language"),
	},
	ToolConfiguration: &ToolConfiguration{
		ToolChoice: bedrock.ToolChoice_AUTO(),
		Tools: []tool{
			bedrock.*tool_Function(&FunctionToolProps{
				Name: jsii.String("top_song"),
				Description: jsii.String("Get the most popular song played on a radio station."),
				InputSchema: map[string]interface{}{
					"type": jsii.String("object"),
					"properties": map[string]map[string]*string{
						"sign": map[string]*string{
							"type": jsii.String("string"),
							"description": jsii.String("The call sign for the radio station for which you want the most popular song. Example calls signs are WZPZ and WKR."),
						},
					},
					"required": []*string{
						jsii.String("sign"),
					},
				},
			}),
		},
	},
})

bedrock.NewPrompt(this, jsii.String("prompt1"), &PromptProps{
	PromptName: jsii.String("prompt-chat"),
	Description: jsii.String("my first chat prompt"),
	DefaultVariant: variantChat,
	Variants: []iPromptVariant{
		variantChat,
	},
	KmsKey: cmk,
})
```

### Agent Prompt

Agent prompts are designed to work with Bedrock Agents, allowing you to create prompts that can be used by agents to perform specific tasks. Agent prompts use text prompts as their foundation and can reference agent aliases and include custom instructions for how the agent should behave.

```go
cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
})

// Assuming you have an existing agent and alias
agent := bedrock.Agent_FromAgentAttributes(this, jsii.String("ImportedAgent"), &AgentAttributes{
	AgentArn: jsii.String("arn:aws:bedrock:region:account:agent/agent-id"),
	RoleArn: jsii.String("arn:aws:iam::account:role/agent-role"),
})

agentAlias := bedrock.AgentAlias_FromAttributes(this, jsii.String("ImportedAlias"), &AgentAliasAttributes{
	AliasId: jsii.String("alias-id"),
	AliasName: jsii.String("my-alias"),
	AgentVersion: jsii.String("1"),
	Agent: agent,
})

agentVariant := bedrock.PromptVariant_Agent(&AgentPromptVariantProps{
	VariantName: jsii.String("agent-variant"),
	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
	AgentAlias: agentAlias,
	PromptText: jsii.String("Use the agent to help with: {{task}}. Please be thorough and provide detailed explanations."),
	PromptVariables: []*string{
		jsii.String("task"),
	},
})

bedrock.NewPrompt(this, jsii.String("agentPrompt"), &PromptProps{
	PromptName: jsii.String("agent-prompt"),
	Description: jsii.String("Prompt for agent interactions"),
	DefaultVariant: agentVariant,
	Variants: []iPromptVariant{
		agentVariant,
	},
	KmsKey: cmk,
})
```

### Prompt Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| promptName | string | Yes | The name of the prompt |
| description | string | No | A description of the prompt |
| defaultVariant | PromptVariant | Yes | The default variant to use for the prompt |
| variants | PromptVariant[] | No | Additional variants for the prompt |
| kmsKey | kms.IKey | No | The KMS key to use for encrypting the prompt. Defaults to AWS managed key |
| tags | Record<string, string> | No | Tags to apply to the prompt |

### Prompt Version

A prompt version is a snapshot of a prompt at a specific point in time that you create when you are satisfied with a set of configurations. Versions allow you to deploy your prompt and easily switch between different configurations for your prompt and update your application with the most appropriate version for your use-case.

You can create a Prompt version by using the PromptVersion class or by using the .createVersion(..) on a Prompt object. It is recommended to use the .createVersion(..) method. It uses a hash based mechanism to update the version whenever a certain configuration property changes.

```go
cmk := kms.NewKey(this, jsii.String("cmk"), &KeyProps{
})
claudeModel := bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_SONNET_V1_0()

variant1 := bedrock.PromptVariant_Text(&TextPromptVariantProps{
	VariantName: jsii.String("variant1"),
	Model: claudeModel,
	PromptVariables: []*string{
		jsii.String("topic"),
	},
	PromptText: jsii.String("This is my first text prompt. Please summarize our conversation on: {{topic}}."),
	InferenceConfiguration: bedrock.PromptInferenceConfiguration_Text(&PromptInferenceConfigurationProps{
		Temperature: jsii.Number(1),
		TopP: jsii.Number(0.999),
		MaxTokens: jsii.Number(2000),
	}),
})

prompt1 := bedrock.NewPrompt(this, jsii.String("prompt1"), &PromptProps{
	PromptName: jsii.String("prompt1"),
	Description: jsii.String("my first prompt"),
	DefaultVariant: variant1,
	Variants: []iPromptVariant{
		variant1,
	},
	KmsKey: cmk,
})

promptVersion := bedrock.NewPromptVersion(this, jsii.String("MyPromptVersion"), &PromptVersionProps{
	Prompt: prompt1,
	Description: jsii.String("my first version"),
})
//or alternatively:
// const promptVersion = prompt1.createVersion('my first version');
versionString := promptVersion.Version
```

### Import Methods

You can use the `fromPromptAttributes` method to import an existing Bedrock Prompt into your CDK application.

```go
// Import an existing prompt by ARN
importedPrompt := bedrock.Prompt_FromPromptAttributes(this, jsii.String("ImportedPrompt"), &PromptAttributes{
	PromptArn: jsii.String("arn:aws:bedrock:region:account:prompt/prompt-id"),
	KmsKey: kms.Key_FromKeyArn(this, jsii.String("ImportedKey"), jsii.String("arn:aws:kms:region:account:key/key-id")),
	 // optional
	PromptVersion: jsii.String("1"),
})
```

## Inference Profiles

Amazon Bedrock Inference Profiles provide a way to manage and optimize inference configurations for your foundation models. They allow you to define reusable configurations that can be applied across different prompts and agents.

### Using Inference Profiles

Inference profiles can be used with prompts and agents to maintain consistent inference configurations across your application.

#### With Agents

```go
// Create a cross-region inference profile
crossRegionProfile := bedrock.CrossRegionInferenceProfile_FromConfig(&CrossRegionInferenceProfileProps{
	GeoRegion: bedrock.CrossRegionInferenceProfileRegion_US,
	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
})

// Use the cross-region profile with an agent
agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
	FoundationModel: crossRegionProfile,
	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about agriculture."),
})
```

#### With Prompts

```go
// Create a prompt router for intelligent model selection
promptRouter := bedrock.PromptRouter_FromDefaultId(bedrock.DefaultPromptRouterIdentifier_ANTHROPIC_CLAUDE_V1(), jsii.String("us-east-1"))

// Use the prompt router with a prompt variant
variant := bedrock.PromptVariant_Text(&TextPromptVariantProps{
	VariantName: jsii.String("variant1"),
	PromptText: jsii.String("What is the capital of France?"),
	Model: promptRouter,
})

bedrock.NewPrompt(this, jsii.String("Prompt"), &PromptProps{
	PromptName: jsii.String("prompt-router-test"),
	Variants: []iPromptVariant{
		variant,
	},
})
```

### Types of Inference Profiles

Amazon Bedrock offers two types of inference profiles:

#### Application Inference Profiles

Application inference profiles are user-defined profiles that help you track costs and model usage. They can be created for a single region or for multiple regions using a cross-region inference profile.

##### Single Region Application Profile

```go
// Create an application inference profile for one Region
appProfile := bedrock.NewApplicationInferenceProfile(this, jsii.String("MyApplicationProfile"), &ApplicationInferenceProfileProps{
	ApplicationInferenceProfileName: jsii.String("claude-3-sonnet-v1"),
	ModelSource: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_SONNET_V1_0(),
	Description: jsii.String("Application profile for cost tracking"),
	Tags: map[string]*string{
		"Environment": jsii.String("Production"),
	},
})
```

##### Multi-Region Application Profile

```go
// Create a cross-region inference profile
crossRegionProfile := bedrock.CrossRegionInferenceProfile_FromConfig(&CrossRegionInferenceProfileProps{
	GeoRegion: bedrock.CrossRegionInferenceProfileRegion_US,
	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V2_0(),
})

// Create an application inference profile across regions
appProfile := bedrock.NewApplicationInferenceProfile(this, jsii.String("MyMultiRegionProfile"), &ApplicationInferenceProfileProps{
	ApplicationInferenceProfileName: jsii.String("claude-35-sonnet-v2-multi-region"),
	ModelSource: crossRegionProfile,
	Description: jsii.String("Multi-region application profile for cost tracking"),
})
```

#### System Defined Inference Profiles

Cross-region inference enables you to seamlessly manage unplanned traffic bursts by utilizing compute across different AWS Regions. With cross-region inference, you can distribute traffic across multiple AWS Regions, enabling higher throughput and enhanced resilience during periods of peak demands.

Before using a CrossRegionInferenceProfile, ensure that you have access to the models and regions defined in the inference profiles. For instance, if you use the system defined inference profile "us.anthropic.claude-3-5-sonnet-20241022-v2:0", inference requests will be routed to US East (Virginia) us-east-1, US East (Ohio) us-east-2 and US West (Oregon) us-west-2. Thus, you need to have model access enabled in those regions for the model anthropic.claude-3-5-sonnet-20241022-v2:0.

##### System Defined Profile Configuration

```go
crossRegionProfile := bedrock.CrossRegionInferenceProfile_FromConfig(&CrossRegionInferenceProfileProps{
	GeoRegion: bedrock.CrossRegionInferenceProfileRegion_US,
	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V2_0(),
})
```

### Prompt Routers

Amazon Bedrock intelligent prompt routing provides a single serverless endpoint for efficiently routing requests between different foundational models within the same model family. It can help you optimize for response quality and cost. They offer a comprehensive solution for managing multiple AI models through a single serverless endpoint, simplifying the process for you. Intelligent prompt routing predicts the performance of each model for each request, and dynamically routes each request to the model that it predicts is most likely to give the desired response at the lowest cost.

#### Default and Custom Prompt Routers

```go
// Use a default prompt router
variant := bedrock.PromptVariant_Text(&TextPromptVariantProps{
	VariantName: jsii.String("variant1"),
	PromptText: jsii.String("What is the capital of France?"),
	Model: bedrock.PromptRouter_FromDefaultId(bedrock.DefaultPromptRouterIdentifier_ANTHROPIC_CLAUDE_V1(), jsii.String("us-east-1")),
})

bedrock.NewPrompt(this, jsii.String("Prompt"), &PromptProps{
	PromptName: jsii.String("prompt-router-test"),
	Variants: []iPromptVariant{
		variant,
	},
})
```

### Inference Profile Permissions

Use the `grantProfileUsage` method to grant appropriate permissions to resources that need to use the inference profile.

#### Granting Profile Usage Permissions

```go
// Create an application inference profile
profile := bedrock.NewApplicationInferenceProfile(this, jsii.String("MyProfile"), &ApplicationInferenceProfileProps{
	ApplicationInferenceProfileName: jsii.String("my-profile"),
	ModelSource: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
})

// Create a Lambda function
lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_11(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("def handler(event, context): return \"Hello\"")),
})

// Grant the Lambda function permission to use the inference profile
profile.GrantProfileUsage(lambdaFunction)

// Use a system defined inference profile
crossRegionProfile := bedrock.CrossRegionInferenceProfile_FromConfig(&CrossRegionInferenceProfileProps{
	GeoRegion: bedrock.CrossRegionInferenceProfileRegion_US,
	Model: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0(),
})

// Grant permissions to use the cross-region inference profile
crossRegionProfile.GrantProfileUsage(lambdaFunction)
```

The `grantProfileUsage` method adds the necessary IAM permissions to the resource, allowing it to use the inference profile. This includes permissions to call `bedrock:GetInferenceProfile` and `bedrock:ListInferenceProfiles` actions on the inference profile resource.

### Inference Profiles Import Methods

You can import existing application inference profiles using the following methods:

```go
// Import an inference profile through attributes
importedProfile := bedrock.ApplicationInferenceProfile_FromApplicationInferenceProfileAttributes(this, jsii.String("ImportedProfile"), &ApplicationInferenceProfileAttributes{
	InferenceProfileArn: jsii.String("arn:aws:bedrock:us-east-1:123456789012:application-inference-profile/my-profile-id"),
	InferenceProfileIdentifier: jsii.String("my-profile-id"),
})
```

You can also import an application inference profile from an existing L1 CloudFormation construct:

```go
// Create or reference an existing L1 CfnApplicationInferenceProfile
cfnProfile := awscdk.Aws_bedrock.NewCfnApplicationInferenceProfile(this, jsii.String("CfnProfile"), &CfnApplicationInferenceProfileProps{
	InferenceProfileName: jsii.String("my-cfn-profile"),
	ModelSource: &InferenceProfileModelSourceProperty{
		CopyFrom: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_3_5_SONNET_V1_0().InvokableArn,
	},
	Description: jsii.String("Profile created via L1 construct"),
})

// Import the L1 construct as an L2 ApplicationInferenceProfile
importedFromCfn := bedrock.ApplicationInferenceProfile_FromCfnApplicationInferenceProfile(cfnProfile)

// Grant permissions to use the imported profile
lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
	Runtime: lambda.Runtime_PYTHON_3_11(),
	Handler: jsii.String("index.handler"),
	Code: lambda.Code_FromInline(jsii.String("def handler(event, context): return \"Hello\"")),
})

importedFromCfn.GrantProfileUsage(lambdaFunction)
```

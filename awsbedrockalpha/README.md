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
* [Prompts](#prompts)

  * [Prompt Variants](#prompt-variants)
  * [Basic Text Prompt](#basic-text-prompt)
  * [Chat Prompt](#chat-prompt)
  * [Agent Prompt](#agent-prompt)
  * [Prompt Properties](#prompt-properties)
  * [Prompt Version](#prompt-version)
  * [Import Methods](#import-methods)

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

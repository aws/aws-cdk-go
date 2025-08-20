package awsbedrockalpha


// Properties for a function schema.
//
// Example:
//   actionGroupFunction := lambda.NewFunction(this, jsii.String("ActionGroupFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_12(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("../lambda/action-group"))),
//   })
//
//   // Define a function schema with parameters
//   functionSchema := bedrock.NewFunctionSchema(&FunctionSchemaProps{
//   	Functions: []functionProps{
//   		&functionProps{
//   			Name: jsii.String("searchBooks"),
//   			Description: jsii.String("Search for books in the library catalog"),
//   			Parameters: map[string]functionParameterProps{
//   				"query": &functionParameterProps{
//   					"type": bedrock.ParameterType_STRING,
//   					"required": jsii.Boolean(true),
//   					"description": jsii.String("The search query string"),
//   				},
//   				"maxResults": &functionParameterProps{
//   					"type": bedrock.ParameterType_INTEGER,
//   					"required": jsii.Boolean(false),
//   					"description": jsii.String("Maximum number of results to return"),
//   				},
//   				"includeOutOfPrint": &functionParameterProps{
//   					"type": bedrock.ParameterType_BOOLEAN,
//   					"required": jsii.Boolean(false),
//   					"description": jsii.String("Whether to include out-of-print books"),
//   				},
//   			},
//   			RequireConfirmation: bedrock.RequireConfirmation_DISABLED,
//   		},
//   		&functionProps{
//   			Name: jsii.String("getBookDetails"),
//   			Description: jsii.String("Get detailed information about a specific book"),
//   			Parameters: map[string]*functionParameterProps{
//   				"bookId": &functionParameterProps{
//   					"type": bedrock.ParameterType_STRING,
//   					"required": jsii.Boolean(true),
//   					"description": jsii.String("The unique identifier of the book"),
//   				},
//   			},
//   			RequireConfirmation: bedrock.RequireConfirmation_ENABLED,
//   		},
//   	},
//   })
//
//   // Create an action group using the function schema
//   actionGroup := bedrock.NewAgentActionGroup(&AgentActionGroupProps{
//   	Name: jsii.String("library-functions"),
//   	Description: jsii.String("Functions for interacting with the library catalog"),
//   	Executor: bedrock.ActionGroupExecutor_FromLambda(actionGroupFunction),
//   	FunctionSchema: functionSchema,
//   	Enabled: jsii.Boolean(true),
//   })
//
//   agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
//   	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
//   	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
//   	ActionGroups: []agentActionGroup{
//   		actionGroup,
//   	},
//   })
//
// Experimental.
type FunctionSchemaProps struct {
	// Functions defined in the schema.
	// Experimental.
	Functions *[]*FunctionProps `field:"required" json:"functions" yaml:"functions"`
}


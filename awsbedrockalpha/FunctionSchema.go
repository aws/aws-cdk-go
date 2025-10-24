package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a function schema for a Bedrock Agent Action Group.
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
//   	Functions: []FunctionProps{
//   		&FunctionProps{
//   			Name: jsii.String("searchBooks"),
//   			Description: jsii.String("Search for books in the library catalog"),
//   			Parameters: map[string]FunctionParameterProps{
//   				"query": &FunctionParameterProps{
//   					"type": bedrock.ParameterType_STRING,
//   					"required": jsii.Boolean(true),
//   					"description": jsii.String("The search query string"),
//   				},
//   				"maxResults": &FunctionParameterProps{
//   					"type": bedrock.ParameterType_INTEGER,
//   					"required": jsii.Boolean(false),
//   					"description": jsii.String("Maximum number of results to return"),
//   				},
//   				"includeOutOfPrint": &FunctionParameterProps{
//   					"type": bedrock.ParameterType_BOOLEAN,
//   					"required": jsii.Boolean(false),
//   					"description": jsii.String("Whether to include out-of-print books"),
//   				},
//   			},
//   			RequireConfirmation: bedrock.RequireConfirmation_DISABLED,
//   		},
//   		&FunctionProps{
//   			Name: jsii.String("getBookDetails"),
//   			Description: jsii.String("Get detailed information about a specific book"),
//   			Parameters: map[string]FunctionParameterProps{
//   				"bookId": &FunctionParameterProps{
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
//   	ActionGroups: []AgentActionGroup{
//   		actionGroup,
//   	},
//   })
//
// Experimental.
type FunctionSchema interface {
	// The functions defined in the schema.
	// Experimental.
	Functions() *[]Function
}

// The jsii proxy struct for FunctionSchema
type jsiiProxy_FunctionSchema struct {
	_ byte // padding
}

func (j *jsiiProxy_FunctionSchema) Functions() *[]Function {
	var returns *[]Function
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}


// Experimental.
func NewFunctionSchema(props *FunctionSchemaProps) FunctionSchema {
	_init_.Initialize()

	if err := validateNewFunctionSchemaParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionSchema{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.FunctionSchema",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewFunctionSchema_Override(f FunctionSchema, props *FunctionSchemaProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.FunctionSchema",
		[]interface{}{props},
		f,
	)
}


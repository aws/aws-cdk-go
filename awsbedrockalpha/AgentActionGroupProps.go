package awsbedrockalpha


// ****************************************************************************                        PROPS - Action Group Class ***************************************************************************.
//
// Example:
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("ExistingBucket"), jsii.String("my-schema-bucket"))
//   s3Schema := bedrock.ApiSchema_FromS3File(bucket, jsii.String("schemas/action-group.yaml"))
//
//   actionGroupFunction := lambda.NewFunction(this, jsii.String("ActionGroupFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_12(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("../lambda/action-group"))),
//   })
//
//   actionGroup := bedrock.NewAgentActionGroup(&AgentActionGroupProps{
//   	Name: jsii.String("query-library"),
//   	Description: jsii.String("Use these functions to get information about the books in the library."),
//   	Executor: bedrock.ActionGroupExecutor_FromLambda(actionGroupFunction),
//   	Enabled: jsii.Boolean(true),
//   	ApiSchema: s3Schema,
//   })
//
//   agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
//   	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
//   	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
//   })
//
//   agent.AddActionGroup(actionGroup)
//
// Experimental.
type AgentActionGroupProps struct {
	// The API Schema defining the functions available to the agent.
	// Default: undefined - No API Schema is provided.
	//
	// Experimental.
	ApiSchema ApiSchema `field:"optional" json:"apiSchema" yaml:"apiSchema"`
	// A description of the action group.
	// Default: undefined - No description is provided.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the action group is available for the agent to invoke or not when sending an InvokeAgent request.
	// Default: true - The action group is enabled.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The action group executor that implements the API functions.
	// Default: undefined - No executor is provided.
	//
	// Experimental.
	Executor ActionGroupExecutor `field:"optional" json:"executor" yaml:"executor"`
	// Specifies whether to delete the resource even if it's in use.
	// Default: false - The resource will not be deleted if it's in use.
	//
	// Experimental.
	ForceDelete *bool `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Defines functions that each define parameters that the agent needs to invoke from the user.
	//
	// NO L2 yet as this doesn't make much sense IMHO.
	// Default: undefined - No function schema is provided.
	//
	// Experimental.
	FunctionSchema FunctionSchema `field:"optional" json:"functionSchema" yaml:"functionSchema"`
	// The name of the action group.
	// Default: - A unique name is generated in the format 'action_group_quick_start_UUID'.
	//
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The AWS Defined signature for enabling certain capabilities in your agent.
	//
	// When this property is specified, you must leave the description, apiSchema,
	// and actionGroupExecutor fields blank for this action group.
	// Default: undefined - No parent action group signature is provided.
	//
	// Experimental.
	ParentActionGroupSignature ParentActionGroupSignature `field:"optional" json:"parentActionGroupSignature" yaml:"parentActionGroupSignature"`
}


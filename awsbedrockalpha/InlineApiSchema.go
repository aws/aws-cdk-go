package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Class to define an API Schema from an inline string.
//
// The schema can be provided directly as a string in either JSON or YAML format.
//
// Example:
//   inlineSchema := bedrock.ApiSchema_FromInline(jsii.String(`
//   openapi: 3.0.3
//   info:
//     title: Library API
//     version: 1.0.0
//   paths:
//     /search:
//       get:
//         summary: Search for books
//         operationId: searchBooks
//         parameters:
//           - name: query
//             in: query
//             required: true
//             schema:
//               type: string
//   `))
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
//   	ApiSchema: inlineSchema,
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
type InlineApiSchema interface {
	ApiSchema
	// The inline OpenAPI schema definition as a string, if using an inline schema.
	//
	// Can be in JSON or YAML format.
	// Experimental.
	InlineSchema() *string
	// The S3 location of the API schema file, if using an S3-based schema.
	//
	// Contains the bucket name and object key information.
	// Experimental.
	S3File() *awss3.Location
}

// The jsii proxy struct for InlineApiSchema
type jsiiProxy_InlineApiSchema struct {
	jsiiProxy_ApiSchema
}

func (j *jsiiProxy_InlineApiSchema) InlineSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InlineApiSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


// Experimental.
func NewInlineApiSchema(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateNewInlineApiSchemaParameters(schema); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineApiSchema{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.InlineApiSchema",
		[]interface{}{schema},
		&j,
	)

	return &j
}

// Experimental.
func NewInlineApiSchema_Override(i InlineApiSchema, schema *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.InlineApiSchema",
		[]interface{}{schema},
		i,
	)
}

// Creates an API Schema from an inline string.
// Experimental.
func InlineApiSchema_FromInline(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateInlineApiSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.InlineApiSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates an API Schema from a local file.
// Experimental.
func InlineApiSchema_FromLocalAsset(path *string) AssetApiSchema {
	_init_.Initialize()

	if err := validateInlineApiSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns AssetApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.InlineApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
// Experimental.
func InlineApiSchema_FromS3File(bucket awss3.IBucket, objectKey *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateInlineApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.InlineApiSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey},
		&returns,
	)

	return returns
}


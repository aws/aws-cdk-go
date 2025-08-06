package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// API Schema from a local asset.
//
// The asset is uploaded to an S3 staging bucket, then moved to its final location
// by CloudFormation during deployment.
//
// Example:
//   actionGroupFunction := lambda.NewFunction(this, jsii.String("ActionGroupFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_12(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("../lambda/action-group"))),
//   })
//
//   // When using ApiSchema.fromLocalAsset, you must bind the schema to a scope
//   schema := bedrock.ApiSchema_FromLocalAsset(path.join(__dirname, jsii.String("action-group.yaml")))
//   schema.Bind(this)
//
//   actionGroup := bedrock.NewAgentActionGroup(&AgentActionGroupProps{
//   	Name: jsii.String("query-library"),
//   	Description: jsii.String("Use these functions to get information about the books in the library."),
//   	Executor: bedrock.ActionGroupExecutor_FromLambda(actionGroupFunction),
//   	Enabled: jsii.Boolean(true),
//   	ApiSchema: schema,
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
type AssetApiSchema interface {
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
	// Binds this API schema to a construct scope.
	//
	// This method initializes the S3 asset if it hasn't been initialized yet.
	// Must be called before rendering the schema as CFN properties.
	// Experimental.
	Bind(scope constructs.Construct)
}

// The jsii proxy struct for AssetApiSchema
type jsiiProxy_AssetApiSchema struct {
	jsiiProxy_ApiSchema
}

func (j *jsiiProxy_AssetApiSchema) InlineSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetApiSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


// Experimental.
func NewAssetApiSchema(path *string, options *awss3assets.AssetOptions) AssetApiSchema {
	_init_.Initialize()

	if err := validateNewAssetApiSchemaParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetApiSchema{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.AssetApiSchema",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetApiSchema_Override(a AssetApiSchema, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.AssetApiSchema",
		[]interface{}{path, options},
		a,
	)
}

// Creates an API Schema from an inline string.
// Experimental.
func AssetApiSchema_FromInline(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateAssetApiSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.AssetApiSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates an API Schema from a local file.
// Experimental.
func AssetApiSchema_FromLocalAsset(path *string) AssetApiSchema {
	_init_.Initialize()

	if err := validateAssetApiSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns AssetApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.AssetApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
// Experimental.
func AssetApiSchema_FromS3File(bucket awss3.IBucket, objectKey *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateAssetApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.AssetApiSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetApiSchema) Bind(scope constructs.Construct) {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{scope},
	)
}


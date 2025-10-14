package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Represents the concept of an API Schema for a Bedrock Agent Action Group.
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
type ApiSchema interface {
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

// The jsii proxy struct for ApiSchema
type jsiiProxy_ApiSchema struct {
	_ byte // padding
}

func (j *jsiiProxy_ApiSchema) InlineSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiSchema_Override(a ApiSchema, s3File *awss3.Location, inlineSchema *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.ApiSchema",
		[]interface{}{s3File, inlineSchema},
		a,
	)
}

// Creates an API Schema from an inline string.
// Experimental.
func ApiSchema_FromInline(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateApiSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.ApiSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates an API Schema from a local file.
// Experimental.
func ApiSchema_FromLocalAsset(path *string) AssetApiSchema {
	_init_.Initialize()

	if err := validateApiSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns AssetApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.ApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
// Experimental.
func ApiSchema_FromS3File(bucket awss3.IBucketRef, objectKey *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.ApiSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey},
		&returns,
	)

	return returns
}


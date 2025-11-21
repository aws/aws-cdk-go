package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Class to define an API Schema from an inline string.
//
// The schema can be provided directly as a string.
// Validation is performed at the target configuration level where the schema type is known.
//
// Example:
//   inlineSchema := agentcore.ApiSchema_FromInline(jsii.String(`
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
// Experimental.
type InlineApiSchema interface {
	ApiSchema
	// The account ID of the S3 bucket owner for cross-account access.
	// Experimental.
	BucketOwnerAccountId() *string
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
	// Bind the schema to a construct.
	// Experimental.
	Bind(scope constructs.Construct)
	// Grant permissions to the role.
	// Experimental.
	GrantPermissionsToRole(role awsiam.IRole)
}

// The jsii proxy struct for InlineApiSchema
type jsiiProxy_InlineApiSchema struct {
	jsiiProxy_ApiSchema
}

func (j *jsiiProxy_InlineApiSchema) BucketOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccountId",
		&returns,
	)
	return returns
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
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineApiSchema",
		[]interface{}{schema},
		&j,
	)

	return &j
}

// Experimental.
func NewInlineApiSchema_Override(i InlineApiSchema, schema *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineApiSchema",
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
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineApiSchema",
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
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
// Experimental.
func InlineApiSchema_FromS3File(bucket awss3.IBucket, objectKey *string, bucketOwnerAccountId *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateInlineApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineApiSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey, bucketOwnerAccountId},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineApiSchema) Bind(scope constructs.Construct) {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"bind",
		[]interface{}{scope},
	)
}

func (i *jsiiProxy_InlineApiSchema) GrantPermissionsToRole(role awsiam.IRole) {
	if err := i.validateGrantPermissionsToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"grantPermissionsToRole",
		[]interface{}{role},
	)
}


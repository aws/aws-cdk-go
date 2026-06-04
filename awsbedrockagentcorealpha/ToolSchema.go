package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Defines the schema for tools available to a gateway target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   toolSchema := bedrock_agentcore_alpha.ToolSchema_FromLocalAsset(jsii.String("path"))
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type ToolSchema interface {
	// The account ID of the S3 bucket owner for cross-account access.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	BucketOwnerAccountId() *string
	// The inline tool schema definition as a string, if using an inline schema.
	//
	// Can be in JSON or YAML format.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	InlineSchema() *[]*ToolDefinition
	// The S3 location of the tool schema file, if using an S3-based schema.
	//
	// Contains the bucket name and object key information.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	S3File() *awss3.Location
	// Bind the schema to a construct.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Bind(scope constructs.Construct)
	// Grant permissions to the role.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	GrantPermissionsToRole(role awsiam.IRole)
}

// The jsii proxy struct for ToolSchema
type jsiiProxy_ToolSchema struct {
	_ byte // padding
}

func (j *jsiiProxy_ToolSchema) BucketOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolSchema) InlineSchema() *[]*ToolDefinition {
	var returns *[]*ToolDefinition
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ToolSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func NewToolSchema_Override(t ToolSchema, s3File *awss3.Location, bucketOwnerAccountId *string, inlineSchema *[]*ToolDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ToolSchema",
		[]interface{}{s3File, bucketOwnerAccountId, inlineSchema},
		t,
	)
}

// Creates a Tool Schema from an inline string.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func ToolSchema_FromInline(schema *[]*ToolDefinition) InlineToolSchema {
	_init_.Initialize()

	if err := validateToolSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineToolSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ToolSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates a tool Schema from a local file.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func ToolSchema_FromLocalAsset(path *string) ToolSchema {
	_init_.Initialize()

	if err := validateToolSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns ToolSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ToolSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates a Tool Schema from an S3 File.
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func ToolSchema_FromS3File(bucket awss3.IBucket, objectKey *string, bucketOwnerAccountId *string) S3ToolSchema {
	_init_.Initialize()

	if err := validateToolSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ToolSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ToolSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey, bucketOwnerAccountId},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_ToolSchema) Bind(scope constructs.Construct) {
	if err := t.validateBindParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"bind",
		[]interface{}{scope},
	)
}

func (t *jsiiProxy_ToolSchema) GrantPermissionsToRole(role awsiam.IRole) {
	if err := t.validateGrantPermissionsToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"grantPermissionsToRole",
		[]interface{}{role},
	)
}


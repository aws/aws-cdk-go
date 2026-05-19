package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Class to define a Tool Schema from an inline string.
//
// The schema can be provided directly as a string in either JSON or YAML format.
//
// Example:
//   toolSchema := agentcore.ToolSchema_FromInline([]ToolDefinition{
//   	&ToolDefinition{
//   		Name: jsii.String("hello_world"),
//   		Description: jsii.String("A simple hello world tool"),
//   		InputSchema: &SchemaDefinition{
//   			Type: agentcore.SchemaDefinitionType_OBJECT(),
//   			Properties: map[string]SchemaDefinition{
//   				"name": &SchemaDefinition{
//   					"type": agentcore.SchemaDefinitionType_STRING(),
//   					"description": jsii.String("The name to greet"),
//   				},
//   			},
//   			Required: []*string{
//   				jsii.String("name"),
//   			},
//   		},
//   	},
//   })
//
type InlineToolSchema interface {
	ToolSchema
	// The account ID of the S3 bucket owner for cross-account access.
	BucketOwnerAccountId() *string
	// The inline tool schema definition as a string, if using an inline schema.
	//
	// Can be in JSON or YAML format.
	InlineSchema() *[]*ToolDefinition
	// The S3 location of the tool schema file, if using an S3-based schema.
	//
	// Contains the bucket name and object key information.
	S3File() *awss3.Location
	// Bind the schema to a construct.
	Bind(scope constructs.Construct)
	// Grant permissions to the role.
	GrantPermissionsToRole(role awsiam.IRole)
}

// The jsii proxy struct for InlineToolSchema
type jsiiProxy_InlineToolSchema struct {
	jsiiProxy_ToolSchema
}

func (j *jsiiProxy_InlineToolSchema) BucketOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InlineToolSchema) InlineSchema() *[]*ToolDefinition {
	var returns *[]*ToolDefinition
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InlineToolSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


func NewInlineToolSchema(schema *[]*ToolDefinition) InlineToolSchema {
	_init_.Initialize()

	if err := validateNewInlineToolSchemaParameters(schema); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineToolSchema{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.InlineToolSchema",
		[]interface{}{schema},
		&j,
	)

	return &j
}

func NewInlineToolSchema_Override(i InlineToolSchema, schema *[]*ToolDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.InlineToolSchema",
		[]interface{}{schema},
		i,
	)
}

// Creates a Tool Schema from an inline string.
func InlineToolSchema_FromInline(schema *[]*ToolDefinition) InlineToolSchema {
	_init_.Initialize()

	if err := validateInlineToolSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineToolSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.InlineToolSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates a tool Schema from a local file.
func InlineToolSchema_FromLocalAsset(path *string) ToolSchema {
	_init_.Initialize()

	if err := validateInlineToolSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns ToolSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.InlineToolSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates a Tool Schema from an S3 File.
func InlineToolSchema_FromS3File(bucket awss3.IBucket, objectKey *string, bucketOwnerAccountId *string) S3ToolSchema {
	_init_.Initialize()

	if err := validateInlineToolSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ToolSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.InlineToolSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey, bucketOwnerAccountId},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InlineToolSchema) Bind(scope constructs.Construct) {
	if err := i.validateBindParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"bind",
		[]interface{}{scope},
	)
}

func (i *jsiiProxy_InlineToolSchema) GrantPermissionsToRole(role awsiam.IRole) {
	if err := i.validateGrantPermissionsToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"grantPermissionsToRole",
		[]interface{}{role},
	)
}


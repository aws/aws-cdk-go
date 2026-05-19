package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
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
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var schemaDefinition_ SchemaDefinition
//
//   inlineToolSchema := bedrock_agentcore_alpha.NewInlineToolSchema([]ToolDefinition{
//   	&ToolDefinition{
//   		Description: jsii.String("description"),
//   		InputSchema: &SchemaDefinition{
//   			Type: bedrock_agentcore_alpha.SchemaDefinitionType_STRING,
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			Items: schemaDefinition_,
//   			Properties: map[string]SchemaDefinition{
//   				"propertiesKey": schemaDefinition_,
//   			},
//   			Required: []*string{
//   				jsii.String("required"),
//   			},
//   		},
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		OutputSchema: &SchemaDefinition{
//   			Type: bedrock_agentcore_alpha.SchemaDefinitionType_STRING,
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			Items: schemaDefinition_,
//   			Properties: map[string]SchemaDefinition{
//   				"propertiesKey": schemaDefinition_,
//   			},
//   			Required: []*string{
//   				jsii.String("required"),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type InlineToolSchema interface {
	ToolSchema
	// The account ID of the S3 bucket owner for cross-account access.
	// Experimental.
	BucketOwnerAccountId() *string
	// The inline tool schema definition as a string, if using an inline schema.
	//
	// Can be in JSON or YAML format.
	// Experimental.
	InlineSchema() *[]*ToolDefinition
	// The S3 location of the tool schema file, if using an S3-based schema.
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


// Experimental.
func NewInlineToolSchema(schema *[]*ToolDefinition) InlineToolSchema {
	_init_.Initialize()

	if err := validateNewInlineToolSchemaParameters(schema); err != nil {
		panic(err)
	}
	j := jsiiProxy_InlineToolSchema{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineToolSchema",
		[]interface{}{schema},
		&j,
	)

	return &j
}

// Experimental.
func NewInlineToolSchema_Override(i InlineToolSchema, schema *[]*ToolDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineToolSchema",
		[]interface{}{schema},
		i,
	)
}

// Creates a Tool Schema from an inline string.
// Experimental.
func InlineToolSchema_FromInline(schema *[]*ToolDefinition) InlineToolSchema {
	_init_.Initialize()

	if err := validateInlineToolSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineToolSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineToolSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates a tool Schema from a local file.
// Experimental.
func InlineToolSchema_FromLocalAsset(path *string) ToolSchema {
	_init_.Initialize()

	if err := validateInlineToolSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns ToolSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineToolSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates a Tool Schema from an S3 File.
// Experimental.
func InlineToolSchema_FromS3File(bucket awss3.IBucket, objectKey *string, bucketOwnerAccountId *string) S3ToolSchema {
	_init_.Initialize()

	if err := validateInlineToolSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ToolSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.InlineToolSchema",
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


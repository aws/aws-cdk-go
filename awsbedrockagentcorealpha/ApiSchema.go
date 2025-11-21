package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the concept of an API Schema for a Gateway Target.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   smithySchema := agentcore.ApiSchema_FromLocalAsset(path.join(__dirname, jsii.String("models"), jsii.String("smithy-model.json")))
//   smithySchema.Bind(this)
//
//   // Create a gateway target with Smithy Model and OAuth
//   target := agentcore.GatewayTarget_ForSmithy(this, jsii.String("MySmithyTarget"), &GatewayTargetSmithyProps{
//   	GatewayTargetName: jsii.String("my-smithy-target"),
//   	Description: jsii.String("Target for Smithy model integration"),
//   	Gateway: gateway,
//   	SmithyModel: smithySchema,
//   })
//
// Experimental.
type ApiSchema interface {
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

// The jsii proxy struct for ApiSchema
type jsiiProxy_ApiSchema struct {
	_ byte // padding
}

func (j *jsiiProxy_ApiSchema) BucketOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccountId",
		&returns,
	)
	return returns
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
func NewApiSchema_Override(a ApiSchema, s3File *awss3.Location, bucketOwnerAccountId *string, inlineSchema *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiSchema",
		[]interface{}{s3File, bucketOwnerAccountId, inlineSchema},
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
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiSchema",
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
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
// Experimental.
func ApiSchema_FromS3File(bucket awss3.IBucket, objectKey *string, bucketOwnerAccountId *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.ApiSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey, bucketOwnerAccountId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiSchema) Bind(scope constructs.Construct) {
	if err := a.validateBindParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bind",
		[]interface{}{scope},
	)
}

func (a *jsiiProxy_ApiSchema) GrantPermissionsToRole(role awsiam.IRole) {
	if err := a.validateGrantPermissionsToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"grantPermissionsToRole",
		[]interface{}{role},
	)
}


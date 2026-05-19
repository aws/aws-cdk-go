package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
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
type AssetApiSchema interface {
	ApiSchema
	// The account ID of the S3 bucket owner for cross-account access.
	BucketOwnerAccountId() *string
	// The inline OpenAPI schema definition as a string, if using an inline schema.
	//
	// Can be in JSON or YAML format.
	InlineSchema() *string
	// The S3 location of the API schema file, if using an S3-based schema.
	//
	// Contains the bucket name and object key information.
	S3File() *awss3.Location
	// Binds this API schema to a construct scope.
	//
	// This method initializes the S3 asset if it hasn't been initialized yet.
	// Must be called before rendering the schema as CFN properties.
	Bind(scope constructs.Construct)
	// Grant permissions to the role.
	GrantPermissionsToRole(role awsiam.IRole)
}

// The jsii proxy struct for AssetApiSchema
type jsiiProxy_AssetApiSchema struct {
	jsiiProxy_ApiSchema
}

func (j *jsiiProxy_AssetApiSchema) BucketOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccountId",
		&returns,
	)
	return returns
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


func NewAssetApiSchema(path *string, options *awss3assets.AssetOptions) AssetApiSchema {
	_init_.Initialize()

	if err := validateNewAssetApiSchemaParameters(path, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetApiSchema{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.AssetApiSchema",
		[]interface{}{path, options},
		&j,
	)

	return &j
}

func NewAssetApiSchema_Override(a AssetApiSchema, path *string, options *awss3assets.AssetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.AssetApiSchema",
		[]interface{}{path, options},
		a,
	)
}

// Creates an API Schema from an inline string.
func AssetApiSchema_FromInline(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateAssetApiSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.AssetApiSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates an API Schema from a local file.
func AssetApiSchema_FromLocalAsset(path *string) AssetApiSchema {
	_init_.Initialize()

	if err := validateAssetApiSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns AssetApiSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.AssetApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
func AssetApiSchema_FromS3File(bucket awss3.IBucket, objectKey *string, bucketOwnerAccountId *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateAssetApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.AssetApiSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey, bucketOwnerAccountId},
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

func (a *jsiiProxy_AssetApiSchema) GrantPermissionsToRole(role awsiam.IRole) {
	if err := a.validateGrantPermissionsToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"grantPermissionsToRole",
		[]interface{}{role},
	)
}


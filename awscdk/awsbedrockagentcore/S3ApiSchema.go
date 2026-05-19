package awsbedrockagentcore

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Class to define an API Schema from an S3 object.
//
// Example:
//   gateway := agentcore.NewGateway(this, jsii.String("MyGateway"), &GatewayProps{
//   	GatewayName: jsii.String("my-gateway"),
//   })
//
//   // Create an API key credential provider in Token Vault
//   apiKeyProvider := agentcore.NewApiKeyCredentialProvider(this, jsii.String("MyApiKeyProvider"), &ApiKeyCredentialProviderProps{
//   	ApiKeyCredentialProviderName: jsii.String("my-apikey"),
//   })
//
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("ExistingBucket"), jsii.String("my-schema-bucket"))
//   s3mySchema := agentcore.ApiSchema_FromS3File(bucket, jsii.String("schemas/myschema.yaml"))
//
//   // Add an OpenAPI target using the L2 construct directly
//   target := gateway.AddOpenApiTarget(jsii.String("MyTarget"), &AddOpenApiTargetOptions{
//   	GatewayTargetName: jsii.String("my-api-target"),
//   	Description: jsii.String("Target for external API integration"),
//   	ApiSchema: s3mySchema,
//   	CredentialProviderConfigurations: []ICredentialProviderConfig{
//   		agentcore.GatewayCredentialProvider_FromApiKeyIdentity(apiKeyProvider, &FromApiKeyIdentityOptions{
//   			CredentialLocation: agentcore.ApiKeyCredentialLocation_Header(&ApiKeyAdditionalConfiguration{
//   				CredentialParameterName: jsii.String("X-API-Key"),
//   			}),
//   		}),
//   	},
//   })
//
//   // This makes sure your s3 bucket is available before target
//   target.Node.AddDependency(bucket)
//
type S3ApiSchema interface {
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
	// Bind the schema to a construct.
	Bind(scope constructs.Construct)
	// Grant permissions to the role.
	GrantPermissionsToRole(role awsiam.IRole)
}

// The jsii proxy struct for S3ApiSchema
type jsiiProxy_S3ApiSchema struct {
	jsiiProxy_ApiSchema
}

func (j *jsiiProxy_S3ApiSchema) BucketOwnerAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketOwnerAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ApiSchema) InlineSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inlineSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ApiSchema) S3File() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3File",
		&returns,
	)
	return returns
}


func NewS3ApiSchema(location *awss3.Location, bucketOwnerAccountId *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateNewS3ApiSchemaParameters(location); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ApiSchema{}

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.S3ApiSchema",
		[]interface{}{location, bucketOwnerAccountId},
		&j,
	)

	return &j
}

func NewS3ApiSchema_Override(s S3ApiSchema, location *awss3.Location, bucketOwnerAccountId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_bedrockagentcore.S3ApiSchema",
		[]interface{}{location, bucketOwnerAccountId},
		s,
	)
}

// Creates an API Schema from an inline string.
func S3ApiSchema_FromInline(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.S3ApiSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates an API Schema from a local file.
func S3ApiSchema_FromLocalAsset(path *string) AssetApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns AssetApiSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.S3ApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
func S3ApiSchema_FromS3File(bucket awss3.IBucket, objectKey *string, bucketOwnerAccountId *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_bedrockagentcore.S3ApiSchema",
		"fromS3File",
		[]interface{}{bucket, objectKey, bucketOwnerAccountId},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ApiSchema) Bind(scope constructs.Construct) {
	if err := s.validateBindParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bind",
		[]interface{}{scope},
	)
}

func (s *jsiiProxy_S3ApiSchema) GrantPermissionsToRole(role awsiam.IRole) {
	if err := s.validateGrantPermissionsToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"grantPermissionsToRole",
		[]interface{}{role},
	)
}


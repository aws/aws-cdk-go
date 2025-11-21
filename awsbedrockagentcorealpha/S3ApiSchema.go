package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
)

// Class to define an API Schema from an S3 object.
//
// Example:
//   bucket := s3.Bucket_FromBucketName(this, jsii.String("ExistingBucket"), jsii.String("my-schema-bucket"))
//   s3Schema := agentcore.ApiSchema_FromS3File(bucket, jsii.String("schemas/action-group.yaml"))
//
// Experimental.
type S3ApiSchema interface {
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


// Experimental.
func NewS3ApiSchema(location *awss3.Location, bucketOwnerAccountId *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateNewS3ApiSchemaParameters(location); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ApiSchema{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.S3ApiSchema",
		[]interface{}{location, bucketOwnerAccountId},
		&j,
	)

	return &j
}

// Experimental.
func NewS3ApiSchema_Override(s S3ApiSchema, location *awss3.Location, bucketOwnerAccountId *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-agentcore-alpha.S3ApiSchema",
		[]interface{}{location, bucketOwnerAccountId},
		s,
	)
}

// Creates an API Schema from an inline string.
// Experimental.
func S3ApiSchema_FromInline(schema *string) InlineApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromInlineParameters(schema); err != nil {
		panic(err)
	}
	var returns InlineApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.S3ApiSchema",
		"fromInline",
		[]interface{}{schema},
		&returns,
	)

	return returns
}

// Creates an API Schema from a local file.
// Experimental.
func S3ApiSchema_FromLocalAsset(path *string) AssetApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromLocalAssetParameters(path); err != nil {
		panic(err)
	}
	var returns AssetApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.S3ApiSchema",
		"fromLocalAsset",
		[]interface{}{path},
		&returns,
	)

	return returns
}

// Creates an API Schema from an S3 File.
// Experimental.
func S3ApiSchema_FromS3File(bucket awss3.IBucket, objectKey *string, bucketOwnerAccountId *string) S3ApiSchema {
	_init_.Initialize()

	if err := validateS3ApiSchema_FromS3FileParameters(bucket, objectKey); err != nil {
		panic(err)
	}
	var returns S3ApiSchema

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.S3ApiSchema",
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


package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/constructs-go/constructs/v10"
)

// OpenAPI specification from an S3 archive.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket bucket
//
//   s3ApiDefinition := awscdk.Aws_apigateway.NewS3ApiDefinition(bucket, jsii.String("key"), jsii.String("objectVersion"))
//
type S3ApiDefinition interface {
	ApiDefinition
	// Called when the specification is initialized to allow this object to bind to the stack, add resources and have fun.
	Bind(_scope constructs.Construct) *ApiDefinitionConfig
	// Called after the CFN RestApi resource has been created to allow the Api Definition to bind to it.
	//
	// Specifically it's required to allow assets to add
	// metadata for tooling like SAM CLI to be able to find their origins.
	BindAfterCreate(_scope constructs.Construct, _restApi IRestApi)
}

// The jsii proxy struct for S3ApiDefinition
type jsiiProxy_S3ApiDefinition struct {
	jsiiProxy_ApiDefinition
}

func NewS3ApiDefinition(bucket awss3.IBucket, key *string, objectVersion *string) S3ApiDefinition {
	_init_.Initialize()

	if err := validateNewS3ApiDefinitionParameters(bucket, key); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ApiDefinition{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.S3ApiDefinition",
		[]interface{}{bucket, key, objectVersion},
		&j,
	)

	return &j
}

func NewS3ApiDefinition_Override(s S3ApiDefinition, bucket awss3.IBucket, key *string, objectVersion *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.S3ApiDefinition",
		[]interface{}{bucket, key, objectVersion},
		s,
	)
}

// Loads the API specification from a local disk asset.
func S3ApiDefinition_FromAsset(file *string, options *awss3assets.AssetOptions) AssetApiDefinition {
	_init_.Initialize()

	if err := validateS3ApiDefinition_FromAssetParameters(file, options); err != nil {
		panic(err)
	}
	var returns AssetApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.S3ApiDefinition",
		"fromAsset",
		[]interface{}{file, options},
		&returns,
	)

	return returns
}

// Creates an API definition from a specification file in an S3 bucket.
func S3ApiDefinition_FromBucket(bucket awss3.IBucket, key *string, objectVersion *string) S3ApiDefinition {
	_init_.Initialize()

	if err := validateS3ApiDefinition_FromBucketParameters(bucket, key); err != nil {
		panic(err)
	}
	var returns S3ApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.S3ApiDefinition",
		"fromBucket",
		[]interface{}{bucket, key, objectVersion},
		&returns,
	)

	return returns
}

// Create an API definition from an inline object.
//
// The inline object must follow the
// schema of OpenAPI 2.0 or OpenAPI 3.0
//
// Example:
//     apigateway.ApiDefinition.fromInline({
//       openapi: '3.0.2',
//       paths: {
//         '/pets': {
//           get: {
//             'responses': {
//               200: {
//                 content: {
//                   'application/json': {
//                     schema: {
//                       $ref: '#/components/schemas/Empty',
//                     },
//                   },
//                 },
//               },
//             },
//             'x-amazon-apigateway-integration': {
//               responses: {
//                 default: {
//                   statusCode: '200',
//                 },
//               },
//               requestTemplates: {
//                 'application/json': '{"statusCode": 200}',
//               },
//               passthroughBehavior: 'when_no_match',
//               type: 'mock',
//             },
//           },
//         },
//       },
//       components: {
//         schemas: {
//           Empty: {
//             title: 'Empty Schema',
//             type: 'object',
//           },
//         },
//       },
//     });
//
func S3ApiDefinition_FromInline(definition interface{}) InlineApiDefinition {
	_init_.Initialize()

	if err := validateS3ApiDefinition_FromInlineParameters(definition); err != nil {
		panic(err)
	}
	var returns InlineApiDefinition

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_apigateway.S3ApiDefinition",
		"fromInline",
		[]interface{}{definition},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ApiDefinition) Bind(_scope constructs.Construct) *ApiDefinitionConfig {
	if err := s.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *ApiDefinitionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ApiDefinition) BindAfterCreate(_scope constructs.Construct, _restApi IRestApi) {
	if err := s.validateBindAfterCreateParameters(_scope, _restApi); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bindAfterCreate",
		[]interface{}{_scope, _restApi},
	)
}

